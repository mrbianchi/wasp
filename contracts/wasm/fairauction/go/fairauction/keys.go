// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamColor          = 0
	IdxParamDescription    = 1
	IdxParamDuration       = 2
	IdxParamMinimumBid     = 3
	IdxParamOwnerMargin    = 4
	IdxResultBidders       = 5
	IdxResultColor         = 6
	IdxResultCreator       = 7
	IdxResultDeposit       = 8
	IdxResultDescription   = 9
	IdxResultDuration      = 10
	IdxResultHighestBid    = 11
	IdxResultHighestBidder = 12
	IdxResultMinimumBid    = 13
	IdxResultNumTokens     = 14
	IdxResultOwnerMargin   = 15
	IdxResultWhenStarted   = 16
	IdxStateAuctions       = 17
	IdxStateBidderList     = 18
	IdxStateBids           = 19
	IdxStateOwnerMargin    = 20
)

const keyMapLen = 21

var keyMap = [keyMapLen]wasmlib.Key{
	ParamColor,
	ParamDescription,
	ParamDuration,
	ParamMinimumBid,
	ParamOwnerMargin,
	ResultBidders,
	ResultColor,
	ResultCreator,
	ResultDeposit,
	ResultDescription,
	ResultDuration,
	ResultHighestBid,
	ResultHighestBidder,
	ResultMinimumBid,
	ResultNumTokens,
	ResultOwnerMargin,
	ResultWhenStarted,
	StateAuctions,
	StateBidderList,
	StateBids,
	StateOwnerMargin,
}

var idxMap [keyMapLen]wasmlib.Key32
