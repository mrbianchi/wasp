// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package consensus

import (
	"sort"
	"time"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/identity"
	"github.com/iotaledger/hive.go/marshalutil"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/util"
	"go.dedis.ch/kyber/v3/sign/tbls"
	"golang.org/x/xerrors"
)

type BatchProposal struct {
	ValidatorIndex          uint16
	StateOutputID           ledgerstate.OutputID
	RequestIDs              []iscp.RequestID
	RequestHashes           [][32]byte
	Timestamp               time.Time
	ConsensusManaPledge     identity.ID
	AccessManaPledge        identity.ID
	FeeDestination          *iscp.AgentID
	SigShareOfStateOutputID tbls.SigShare
}

type consensusBatchParams struct {
	timestamp       time.Time // A preliminary timestamp. It can be adjusted based on timestamps of selected requests.
	accessPledge    identity.ID
	consensusPledge identity.ID
	feeDestination  *iscp.AgentID
	entropy         hashing.HashValue
}

func BatchProposalFromBytes(data []byte) (*BatchProposal, error) {
	return BatchProposalFromMarshalUtil(marshalutil.New(data))
}

const errFmt = "BatchProposalFromMarshalUtil: %w"

func BatchProposalFromMarshalUtil(mu *marshalutil.MarshalUtil) (*BatchProposal, error) {
	ret := &BatchProposal{}
	var err error
	ret.ValidatorIndex, err = mu.ReadUint16()
	if err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	ret.StateOutputID, err = ledgerstate.OutputIDFromMarshalUtil(mu)
	if err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	ret.AccessManaPledge, err = identity.IDFromMarshalUtil(mu)
	if err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	ret.ConsensusManaPledge, err = identity.IDFromMarshalUtil(mu)
	if err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	ret.FeeDestination, err = iscp.AgentIDFromMarshalUtil(mu)
	if err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	ret.Timestamp, err = mu.ReadTime()
	if err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	size, err := mu.ReadUint16()
	if err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	var sigShareSize byte
	if sigShareSize, err = mu.ReadByte(); err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	if ret.SigShareOfStateOutputID, err = mu.ReadBytes(int(sigShareSize)); err != nil {
		return nil, xerrors.Errorf(errFmt, err)
	}
	ret.RequestIDs = make([]iscp.RequestID, size)
	ret.RequestHashes = make([][32]byte, size)
	for i := range ret.RequestIDs {
		ret.RequestIDs[i], err = iscp.RequestIDFromMarshalUtil(mu)
		if err != nil {
			return nil, xerrors.Errorf(errFmt, err)
		}

		hashBytes, err := mu.ReadBytes(32)
		copy(ret.RequestHashes[i][:], hashBytes)
		if err != nil {
			return nil, xerrors.Errorf(errFmt, err)
		}
	}
	return ret, nil
}

func (b *BatchProposal) Bytes() []byte {
	mu := marshalutil.New()
	mu.WriteUint16(b.ValidatorIndex).
		Write(b.StateOutputID).
		Write(b.AccessManaPledge).
		Write(b.ConsensusManaPledge).
		Write(b.FeeDestination).
		WriteTime(b.Timestamp).
		WriteUint16(uint16(len(b.RequestIDs))).
		WriteByte(byte(len(b.SigShareOfStateOutputID))).
		WriteBytes(b.SigShareOfStateOutputID)
	for i := range b.RequestIDs {
		mu.Write(b.RequestIDs[i])
		mu.WriteBytes(b.RequestHashes[i][:])
	}
	return mu.Bytes()
}

// EnsureTimestampConsistent adjusts a batch timestamp, if it is not consistent with
// the requests in the BatchProposal and the previous transaction. The timestamp is consistent,
// if it is not bellow the timestamps of all the on-ledger requests and the previous transaction in the chain.
// This implement the "fixing" part described in IscpBatchTimestamp.tla.
func (b *BatchProposal) EnsureTimestampConsistent(requests []iscp.Request, stateTimestamp time.Time) error {
	maxReqTime := time.Time{}
	for i := range b.RequestIDs {
		if requests[i].Hash() != b.RequestHashes[i] {
			return xerrors.New("inconsistent requests in EnsureTimestampConsistent")
		}
		if maxReqTime.Before(requests[i].Timestamp()) {
			maxReqTime = requests[i].Timestamp()
		}
	}
	if b.Timestamp.Before(maxReqTime) {
		b.Timestamp = maxReqTime
	}
	if b.Timestamp.Before(stateTimestamp) {
		b.Timestamp = stateTimestamp
	}
	return nil
}

// calcBatchParameters from a given ACS deterministically calculates timestamp, access and consensus
// mana pledges and fee destination.
//
// Timestamp is calculated by taking maximal proposed timestamp excluding F highest proposals.
//
// TODO final version of pledges and fee destination
func (c *Consensus) calcBatchParameters(props []*BatchProposal) (*consensusBatchParams, error) {
	var retTS time.Time

	ts := make([]time.Time, len(props))
	for i := range ts {
		ts[i] = props[i].Timestamp
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].Before(ts[j])
	})
	proposalCount := len(props)                            // |acsProposals| >= N-F by ACS logic.
	maxFaulty := c.committee.Size() - c.committee.Quorum() // T = N-F ==> F = N-T
	retTS = ts[proposalCount-int(maxFaulty)-1]             // Max(|acsProposals|-F Lowest) ~= 66 percentile.

	indices := make([]uint16, len(props))
	for i := range indices {
		indices[i] = uint16(i)
	}
	// verify signatures calculate entropy
	sigSharesToAggregate := make([][]byte, len(props))
	for i, prop := range props {
		err := c.committee.DKShare().VerifySigShare(c.stateOutput.ID().Bytes(), prop.SigShareOfStateOutputID)
		if err != nil {
			return nil, xerrors.Errorf("INVALID SIGNATURE in ACS from peer #%d: %v", prop.ValidatorIndex, err)
		}
		sigSharesToAggregate[i] = prop.SigShareOfStateOutputID
	}
	// aggregate signatures for use as unpredictable entropy
	signatureWithPK, err := c.committee.DKShare().RecoverFullSignature(sigSharesToAggregate, c.stateOutput.ID().Bytes())
	if err != nil {
		return nil, xerrors.Errorf("recovering signature from ACS: %v", err)
	}

	// selects pseudo-random based on seed, the calculated timestamp
	selectedIndex := util.SelectDeterministicRandomUint16(indices, retTS.UnixNano())
	return &consensusBatchParams{
		timestamp:       retTS,
		accessPledge:    props[selectedIndex].AccessManaPledge,
		consensusPledge: props[selectedIndex].ConsensusManaPledge,
		feeDestination:  props[selectedIndex].FeeDestination,
		entropy:         hashing.HashData(signatureWithPK.Bytes()),
	}, nil
}

const keyLen = ledgerstate.OutputIDLength + 32

// calcIntersection a simple algorithm to calculate acceptable intersection. It simply takes all requests
// seen by 1/3+1 node. The assumptions is there can be at max 1/3 of bizantine nodes, so if something is reported
// by more that 1/3 of nodes it means it is correct
func calcIntersection(acs []*BatchProposal, n uint16) ([]iscp.RequestID, [][32]byte) {
	minNumberMentioned := n/3 + 1
	numMentioned := make(map[[keyLen]byte]uint16)

	maxLen := 0
	for _, prop := range acs {
		for i, reqid := range prop.RequestIDs {
			// save ID + Hash as key to avoid batching requests where different nodes have mismatching request content with the same ID
			hash := prop.RequestHashes[i]
			var key [keyLen]byte
			copy(key[:], append(reqid.Bytes(), hash[:]...))
			numMentioned[key]++
		}
		if len(prop.RequestIDs) > maxLen {
			maxLen = len(prop.RequestIDs)
		}
	}
	retIDs := make([]iscp.RequestID, 0, maxLen)
	retHashes := make([][32]byte, 0)
	for key, num := range numMentioned {
		if num < minNumberMentioned {
			continue
		}
		reqID, err := iscp.RequestIDFromBytes(key[:])
		if err != nil {
			continue
		}
		retIDs = append(retIDs, reqID)
		var hash [32]byte
		copy(hash[:], key[len(key)-32:])
		retHashes = append(retHashes, hash)
	}
	return retIDs, retHashes
}
