// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamAddress     = 0
	IdxParamAgentID     = 1
	IdxParamBlockIndex  = 2
	IdxParamBytes       = 3
	IdxParamChainID     = 4
	IdxParamColor       = 5
	IdxParamHash        = 6
	IdxParamHname       = 7
	IdxParamIndex       = 8
	IdxParamInt16       = 9
	IdxParamInt32       = 10
	IdxParamInt64       = 11
	IdxParamName        = 12
	IdxParamRecordIndex = 13
	IdxParamRequestID   = 14
	IdxParamString      = 15
	IdxParamValue       = 16
	IdxResultCount      = 17
	IdxResultIotas      = 18
	IdxResultLength     = 19
	IdxResultRecord     = 20
	IdxResultValue      = 21
	IdxStateArrays      = 22
)

const keyMapLen = 23

var keyMap = [keyMapLen]wasmlib.Key{
	ParamAddress,
	ParamAgentID,
	ParamBlockIndex,
	ParamBytes,
	ParamChainID,
	ParamColor,
	ParamHash,
	ParamHname,
	ParamIndex,
	ParamInt16,
	ParamInt32,
	ParamInt64,
	ParamName,
	ParamRecordIndex,
	ParamRequestID,
	ParamString,
	ParamValue,
	ResultCount,
	ResultIotas,
	ResultLength,
	ResultRecord,
	ResultValue,
	StateArrays,
}

var idxMap [keyMapLen]wasmlib.Key32
