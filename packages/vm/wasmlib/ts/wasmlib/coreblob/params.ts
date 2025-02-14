// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib"
import * as sc from "./index";

export class MapStringToImmutableBytes {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getBytes(key: string): wasmlib.ScImmutableBytes {
        return new wasmlib.ScImmutableBytes(this.objID, wasmlib.Key32.fromString(key).getKeyID());
    }
}

export class ImmutableStoreBlobParams extends wasmlib.ScMapID {

    blobs(): sc.MapStringToImmutableBytes {
        return new sc.MapStringToImmutableBytes(this.mapID);
    }
}

export class MapStringToMutableBytes {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID)
    }

    getBytes(key: string): wasmlib.ScMutableBytes {
        return new wasmlib.ScMutableBytes(this.objID, wasmlib.Key32.fromString(key).getKeyID());
    }
}

export class MutableStoreBlobParams extends wasmlib.ScMapID {

    blobs(): sc.MapStringToMutableBytes {
        return new sc.MapStringToMutableBytes(this.mapID);
    }
}

export class ImmutableGetBlobFieldParams extends wasmlib.ScMapID {

    field(): wasmlib.ScImmutableString {
        return new wasmlib.ScImmutableString(this.mapID, wasmlib.Key32.fromString(sc.ParamField));
    }

    hash(): wasmlib.ScImmutableHash {
        return new wasmlib.ScImmutableHash(this.mapID, wasmlib.Key32.fromString(sc.ParamHash));
    }
}

export class MutableGetBlobFieldParams extends wasmlib.ScMapID {

    field(): wasmlib.ScMutableString {
        return new wasmlib.ScMutableString(this.mapID, wasmlib.Key32.fromString(sc.ParamField));
    }

    hash(): wasmlib.ScMutableHash {
        return new wasmlib.ScMutableHash(this.mapID, wasmlib.Key32.fromString(sc.ParamHash));
    }
}

export class ImmutableGetBlobInfoParams extends wasmlib.ScMapID {

    hash(): wasmlib.ScImmutableHash {
        return new wasmlib.ScImmutableHash(this.mapID, wasmlib.Key32.fromString(sc.ParamHash));
    }
}

export class MutableGetBlobInfoParams extends wasmlib.ScMapID {

    hash(): wasmlib.ScMutableHash {
        return new wasmlib.ScMutableHash(this.mapID, wasmlib.Key32.fromString(sc.ParamHash));
    }
}
