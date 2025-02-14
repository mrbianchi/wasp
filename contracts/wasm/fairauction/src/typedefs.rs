// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use wasmlib::*;
use wasmlib::host::*;

use crate::structs::*;

pub struct ArrayOfImmutableAgentID {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableAgentID {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_agent_id(&self, index: i32) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.obj_id, Key32(index))
    }
}

pub type ImmutableBidderList = ArrayOfImmutableAgentID;

pub struct ArrayOfMutableAgentID {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableAgentID {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_agent_id(&self, index: i32) -> ScMutableAgentID {
        ScMutableAgentID::new(self.obj_id, Key32(index))
    }
}

pub type MutableBidderList = ArrayOfMutableAgentID;

pub struct MapAgentIDToImmutableBid {
    pub(crate) obj_id: i32,
}

impl MapAgentIDToImmutableBid {
    pub fn get_bid(&self, key: &ScAgentID) -> ImmutableBid {
        ImmutableBid { obj_id: self.obj_id, key_id: key.get_key_id() }
    }
}

pub type ImmutableBids = MapAgentIDToImmutableBid;

pub struct MapAgentIDToMutableBid {
    pub(crate) obj_id: i32,
}

impl MapAgentIDToMutableBid {
    pub fn clear(&self) {
        clear(self.obj_id)
    }

    pub fn get_bid(&self, key: &ScAgentID) -> MutableBid {
        MutableBid { obj_id: self.obj_id, key_id: key.get_key_id() }
    }
}

pub type MutableBids = MapAgentIDToMutableBid;

// @formatter:on
