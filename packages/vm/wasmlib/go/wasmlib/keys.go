// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

type MapKey interface {
	KeyID() Key32
}

type Key string

func (key Key) KeyID() Key32 {
	return GetKeyIDFromString(string(key))
}

type Key32 int32

func (key Key32) KeyID() Key32 {
	return key
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

const (
	KeyAccountID       = Key32(-1)
	KeyAddress         = Key32(-2)
	KeyBalances        = Key32(-3)
	KeyBase58Decode    = Key32(-4)
	KeyBase58Encode    = Key32(-5)
	KeyBlsAddress      = Key32(-6)
	KeyBlsAggregate    = Key32(-7)
	KeyBlsValid        = Key32(-8)
	KeyCall            = Key32(-9)
	KeyCaller          = Key32(-10)
	KeyChainID         = Key32(-11)
	KeyChainOwnerID    = Key32(-12)
	KeyColor           = Key32(-13)
	KeyContract        = Key32(-14)
	KeyContractCreator = Key32(-15)
	KeyDeploy          = Key32(-16)
	KeyEd25519Address  = Key32(-17)
	KeyEd25519Valid    = Key32(-18)
	KeyEvent           = Key32(-19)
	KeyExports         = Key32(-20)
	KeyHashBlake2b     = Key32(-21)
	KeyHashSha3        = Key32(-22)
	KeyHname           = Key32(-23)
	KeyIncoming        = Key32(-24)
	KeyLength          = Key32(-25)
	KeyLog             = Key32(-26)
	KeyMaps            = Key32(-27)
	KeyMinted          = Key32(-28)
	KeyPanic           = Key32(-29)
	KeyParams          = Key32(-30)
	KeyPost            = Key32(-31)
	KeyRandom          = Key32(-32)
	KeyRequestID       = Key32(-33)
	KeyResults         = Key32(-34)
	KeyReturn          = Key32(-35)
	KeyState           = Key32(-36)
	KeyTimestamp       = Key32(-37)
	KeyTrace           = Key32(-38)
	KeyTransfers       = Key32(-39)
	KeyUtility         = Key32(-40)
	KeyZzzzzzz         = Key32(-41)
)
