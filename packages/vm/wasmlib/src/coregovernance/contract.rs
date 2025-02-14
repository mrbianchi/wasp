// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use std::ptr;

use crate::*;
use crate::coregovernance::*;

pub struct AddAllowedStateControllerAddressCall {
    pub func:   ScFunc,
    pub params: MutableAddAllowedStateControllerAddressParams,
}

pub struct ClaimChainOwnershipCall {
    pub func: ScFunc,
}

pub struct DelegateChainOwnershipCall {
    pub func:   ScFunc,
    pub params: MutableDelegateChainOwnershipParams,
}

pub struct RemoveAllowedStateControllerAddressCall {
    pub func:   ScFunc,
    pub params: MutableRemoveAllowedStateControllerAddressParams,
}

pub struct RotateStateControllerCall {
    pub func:   ScFunc,
    pub params: MutableRotateStateControllerParams,
}

pub struct SetChainInfoCall {
    pub func:   ScFunc,
    pub params: MutableSetChainInfoParams,
}

pub struct SetContractFeeCall {
    pub func:   ScFunc,
    pub params: MutableSetContractFeeParams,
}

pub struct SetDefaultFeeCall {
    pub func:   ScFunc,
    pub params: MutableSetDefaultFeeParams,
}

pub struct GetAllowedStateControllerAddressesCall {
    pub func:    ScView,
    pub results: ImmutableGetAllowedStateControllerAddressesResults,
}

pub struct GetChainInfoCall {
    pub func:    ScView,
    pub results: ImmutableGetChainInfoResults,
}

pub struct GetFeeInfoCall {
    pub func:    ScView,
    pub params:  MutableGetFeeInfoParams,
    pub results: ImmutableGetFeeInfoResults,
}

pub struct GetMaxBlobSizeCall {
    pub func:    ScView,
    pub results: ImmutableGetMaxBlobSizeResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn add_allowed_state_controller_address(_ctx: & dyn ScFuncCallContext) -> AddAllowedStateControllerAddressCall {
        let mut f = AddAllowedStateControllerAddressCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_ADD_ALLOWED_STATE_CONTROLLER_ADDRESS),
            params: MutableAddAllowedStateControllerAddressParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn claim_chain_ownership(_ctx: & dyn ScFuncCallContext) -> ClaimChainOwnershipCall {
        ClaimChainOwnershipCall {
            func: ScFunc::new(HSC_NAME, HFUNC_CLAIM_CHAIN_OWNERSHIP),
        }
    }
    pub fn delegate_chain_ownership(_ctx: & dyn ScFuncCallContext) -> DelegateChainOwnershipCall {
        let mut f = DelegateChainOwnershipCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_DELEGATE_CHAIN_OWNERSHIP),
            params: MutableDelegateChainOwnershipParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn remove_allowed_state_controller_address(_ctx: & dyn ScFuncCallContext) -> RemoveAllowedStateControllerAddressCall {
        let mut f = RemoveAllowedStateControllerAddressCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_REMOVE_ALLOWED_STATE_CONTROLLER_ADDRESS),
            params: MutableRemoveAllowedStateControllerAddressParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn rotate_state_controller(_ctx: & dyn ScFuncCallContext) -> RotateStateControllerCall {
        let mut f = RotateStateControllerCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_ROTATE_STATE_CONTROLLER),
            params: MutableRotateStateControllerParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn set_chain_info(_ctx: & dyn ScFuncCallContext) -> SetChainInfoCall {
        let mut f = SetChainInfoCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_CHAIN_INFO),
            params: MutableSetChainInfoParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn set_contract_fee(_ctx: & dyn ScFuncCallContext) -> SetContractFeeCall {
        let mut f = SetContractFeeCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_CONTRACT_FEE),
            params: MutableSetContractFeeParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn set_default_fee(_ctx: & dyn ScFuncCallContext) -> SetDefaultFeeCall {
        let mut f = SetDefaultFeeCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_DEFAULT_FEE),
            params: MutableSetDefaultFeeParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn get_allowed_state_controller_addresses(_ctx: & dyn ScViewCallContext) -> GetAllowedStateControllerAddressesCall {
        let mut f = GetAllowedStateControllerAddressesCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_ALLOWED_STATE_CONTROLLER_ADDRESSES),
            results: ImmutableGetAllowedStateControllerAddressesResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn get_chain_info(_ctx: & dyn ScViewCallContext) -> GetChainInfoCall {
        let mut f = GetChainInfoCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_CHAIN_INFO),
            results: ImmutableGetChainInfoResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn get_fee_info(_ctx: & dyn ScViewCallContext) -> GetFeeInfoCall {
        let mut f = GetFeeInfoCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_FEE_INFO),
            params:  MutableGetFeeInfoParams { id: 0 },
            results: ImmutableGetFeeInfoResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_max_blob_size(_ctx: & dyn ScViewCallContext) -> GetMaxBlobSizeCall {
        let mut f = GetMaxBlobSizeCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_MAX_BLOB_SIZE),
            results: ImmutableGetMaxBlobSizeResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
}

// @formatter:on
