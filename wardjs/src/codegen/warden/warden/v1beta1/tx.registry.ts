//@ts-nocheck
import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgUpdateParams, MsgNewSpace, MsgAddSpaceOwner, MsgRemoveSpaceOwner, MsgNewKeychain, MsgAddKeychainParty, MsgUpdateSpace, MsgUpdateKeychain, MsgNewKeyRequest, MsgUpdateKeyRequest, MsgNewSignatureRequest, MsgFulfilSignatureRequest, MsgNewSignTransactionRequest } from "./tx.js";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/warden.warden.v1beta1.MsgUpdateParams", MsgUpdateParams], ["/warden.warden.v1beta1.MsgNewSpace", MsgNewSpace], ["/warden.warden.v1beta1.MsgAddSpaceOwner", MsgAddSpaceOwner], ["/warden.warden.v1beta1.MsgRemoveSpaceOwner", MsgRemoveSpaceOwner], ["/warden.warden.v1beta1.MsgNewKeychain", MsgNewKeychain], ["/warden.warden.v1beta1.MsgAddKeychainParty", MsgAddKeychainParty], ["/warden.warden.v1beta1.MsgUpdateSpace", MsgUpdateSpace], ["/warden.warden.v1beta1.MsgUpdateKeychain", MsgUpdateKeychain], ["/warden.warden.v1beta1.MsgNewKeyRequest", MsgNewKeyRequest], ["/warden.warden.v1beta1.MsgUpdateKeyRequest", MsgUpdateKeyRequest], ["/warden.warden.v1beta1.MsgNewSignatureRequest", MsgNewSignatureRequest], ["/warden.warden.v1beta1.MsgFulfilSignatureRequest", MsgFulfilSignatureRequest], ["/warden.warden.v1beta1.MsgNewSignTransactionRequest", MsgNewSignTransactionRequest]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    updateParams(value: MsgUpdateParams) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateParams",
        value: MsgUpdateParams.encode(value).finish()
      };
    },
    newSpace(value: MsgNewSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSpace",
        value: MsgNewSpace.encode(value).finish()
      };
    },
    addSpaceOwner(value: MsgAddSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddSpaceOwner",
        value: MsgAddSpaceOwner.encode(value).finish()
      };
    },
    removeSpaceOwner(value: MsgRemoveSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgRemoveSpaceOwner",
        value: MsgRemoveSpaceOwner.encode(value).finish()
      };
    },
    newKeychain(value: MsgNewKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeychain",
        value: MsgNewKeychain.encode(value).finish()
      };
    },
    addKeychainParty(value: MsgAddKeychainParty) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddKeychainParty",
        value: MsgAddKeychainParty.encode(value).finish()
      };
    },
    updateSpace(value: MsgUpdateSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateSpace",
        value: MsgUpdateSpace.encode(value).finish()
      };
    },
    updateKeychain(value: MsgUpdateKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeychain",
        value: MsgUpdateKeychain.encode(value).finish()
      };
    },
    newKeyRequest(value: MsgNewKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeyRequest",
        value: MsgNewKeyRequest.encode(value).finish()
      };
    },
    updateKeyRequest(value: MsgUpdateKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeyRequest",
        value: MsgUpdateKeyRequest.encode(value).finish()
      };
    },
    newSignatureRequest(value: MsgNewSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignatureRequest",
        value: MsgNewSignatureRequest.encode(value).finish()
      };
    },
    fulfilSignatureRequest(value: MsgFulfilSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgFulfilSignatureRequest",
        value: MsgFulfilSignatureRequest.encode(value).finish()
      };
    },
    newSignTransactionRequest(value: MsgNewSignTransactionRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignTransactionRequest",
        value: MsgNewSignTransactionRequest.encode(value).finish()
      };
    }
  },
  withTypeUrl: {
    updateParams(value: MsgUpdateParams) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateParams",
        value
      };
    },
    newSpace(value: MsgNewSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSpace",
        value
      };
    },
    addSpaceOwner(value: MsgAddSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddSpaceOwner",
        value
      };
    },
    removeSpaceOwner(value: MsgRemoveSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgRemoveSpaceOwner",
        value
      };
    },
    newKeychain(value: MsgNewKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeychain",
        value
      };
    },
    addKeychainParty(value: MsgAddKeychainParty) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddKeychainParty",
        value
      };
    },
    updateSpace(value: MsgUpdateSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateSpace",
        value
      };
    },
    updateKeychain(value: MsgUpdateKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeychain",
        value
      };
    },
    newKeyRequest(value: MsgNewKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeyRequest",
        value
      };
    },
    updateKeyRequest(value: MsgUpdateKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeyRequest",
        value
      };
    },
    newSignatureRequest(value: MsgNewSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignatureRequest",
        value
      };
    },
    fulfilSignatureRequest(value: MsgFulfilSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgFulfilSignatureRequest",
        value
      };
    },
    newSignTransactionRequest(value: MsgNewSignTransactionRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignTransactionRequest",
        value
      };
    }
  },
  toJSON: {
    updateParams(value: MsgUpdateParams) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateParams",
        value: MsgUpdateParams.toJSON(value)
      };
    },
    newSpace(value: MsgNewSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSpace",
        value: MsgNewSpace.toJSON(value)
      };
    },
    addSpaceOwner(value: MsgAddSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddSpaceOwner",
        value: MsgAddSpaceOwner.toJSON(value)
      };
    },
    removeSpaceOwner(value: MsgRemoveSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgRemoveSpaceOwner",
        value: MsgRemoveSpaceOwner.toJSON(value)
      };
    },
    newKeychain(value: MsgNewKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeychain",
        value: MsgNewKeychain.toJSON(value)
      };
    },
    addKeychainParty(value: MsgAddKeychainParty) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddKeychainParty",
        value: MsgAddKeychainParty.toJSON(value)
      };
    },
    updateSpace(value: MsgUpdateSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateSpace",
        value: MsgUpdateSpace.toJSON(value)
      };
    },
    updateKeychain(value: MsgUpdateKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeychain",
        value: MsgUpdateKeychain.toJSON(value)
      };
    },
    newKeyRequest(value: MsgNewKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeyRequest",
        value: MsgNewKeyRequest.toJSON(value)
      };
    },
    updateKeyRequest(value: MsgUpdateKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeyRequest",
        value: MsgUpdateKeyRequest.toJSON(value)
      };
    },
    newSignatureRequest(value: MsgNewSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignatureRequest",
        value: MsgNewSignatureRequest.toJSON(value)
      };
    },
    fulfilSignatureRequest(value: MsgFulfilSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgFulfilSignatureRequest",
        value: MsgFulfilSignatureRequest.toJSON(value)
      };
    },
    newSignTransactionRequest(value: MsgNewSignTransactionRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignTransactionRequest",
        value: MsgNewSignTransactionRequest.toJSON(value)
      };
    }
  },
  fromJSON: {
    updateParams(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateParams",
        value: MsgUpdateParams.fromJSON(value)
      };
    },
    newSpace(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSpace",
        value: MsgNewSpace.fromJSON(value)
      };
    },
    addSpaceOwner(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddSpaceOwner",
        value: MsgAddSpaceOwner.fromJSON(value)
      };
    },
    removeSpaceOwner(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgRemoveSpaceOwner",
        value: MsgRemoveSpaceOwner.fromJSON(value)
      };
    },
    newKeychain(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeychain",
        value: MsgNewKeychain.fromJSON(value)
      };
    },
    addKeychainParty(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddKeychainParty",
        value: MsgAddKeychainParty.fromJSON(value)
      };
    },
    updateSpace(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateSpace",
        value: MsgUpdateSpace.fromJSON(value)
      };
    },
    updateKeychain(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeychain",
        value: MsgUpdateKeychain.fromJSON(value)
      };
    },
    newKeyRequest(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeyRequest",
        value: MsgNewKeyRequest.fromJSON(value)
      };
    },
    updateKeyRequest(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeyRequest",
        value: MsgUpdateKeyRequest.fromJSON(value)
      };
    },
    newSignatureRequest(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignatureRequest",
        value: MsgNewSignatureRequest.fromJSON(value)
      };
    },
    fulfilSignatureRequest(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgFulfilSignatureRequest",
        value: MsgFulfilSignatureRequest.fromJSON(value)
      };
    },
    newSignTransactionRequest(value: any) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignTransactionRequest",
        value: MsgNewSignTransactionRequest.fromJSON(value)
      };
    }
  },
  fromPartial: {
    updateParams(value: MsgUpdateParams) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateParams",
        value: MsgUpdateParams.fromPartial(value)
      };
    },
    newSpace(value: MsgNewSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSpace",
        value: MsgNewSpace.fromPartial(value)
      };
    },
    addSpaceOwner(value: MsgAddSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddSpaceOwner",
        value: MsgAddSpaceOwner.fromPartial(value)
      };
    },
    removeSpaceOwner(value: MsgRemoveSpaceOwner) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgRemoveSpaceOwner",
        value: MsgRemoveSpaceOwner.fromPartial(value)
      };
    },
    newKeychain(value: MsgNewKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeychain",
        value: MsgNewKeychain.fromPartial(value)
      };
    },
    addKeychainParty(value: MsgAddKeychainParty) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgAddKeychainParty",
        value: MsgAddKeychainParty.fromPartial(value)
      };
    },
    updateSpace(value: MsgUpdateSpace) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateSpace",
        value: MsgUpdateSpace.fromPartial(value)
      };
    },
    updateKeychain(value: MsgUpdateKeychain) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeychain",
        value: MsgUpdateKeychain.fromPartial(value)
      };
    },
    newKeyRequest(value: MsgNewKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewKeyRequest",
        value: MsgNewKeyRequest.fromPartial(value)
      };
    },
    updateKeyRequest(value: MsgUpdateKeyRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgUpdateKeyRequest",
        value: MsgUpdateKeyRequest.fromPartial(value)
      };
    },
    newSignatureRequest(value: MsgNewSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignatureRequest",
        value: MsgNewSignatureRequest.fromPartial(value)
      };
    },
    fulfilSignatureRequest(value: MsgFulfilSignatureRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgFulfilSignatureRequest",
        value: MsgFulfilSignatureRequest.fromPartial(value)
      };
    },
    newSignTransactionRequest(value: MsgNewSignTransactionRequest) {
      return {
        typeUrl: "/warden.warden.v1beta1.MsgNewSignTransactionRequest",
        value: MsgNewSignTransactionRequest.fromPartial(value)
      };
    }
  }
};