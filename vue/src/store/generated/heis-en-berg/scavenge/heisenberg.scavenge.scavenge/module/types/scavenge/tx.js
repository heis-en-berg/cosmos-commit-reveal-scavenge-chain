/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'heisenberg.scavenge.scavenge';
const baseMsgSubmitScavenge = { creator: '', solutionHash: '', description: '', reward: '' };
export const MsgSubmitScavenge = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.solutionHash !== '') {
            writer.uint32(18).string(message.solutionHash);
        }
        if (message.description !== '') {
            writer.uint32(26).string(message.description);
        }
        if (message.reward !== '') {
            writer.uint32(34).string(message.reward);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSubmitScavenge };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.solutionHash = reader.string();
                    break;
                case 3:
                    message.description = reader.string();
                    break;
                case 4:
                    message.reward = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSubmitScavenge };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.solutionHash !== undefined && object.solutionHash !== null) {
            message.solutionHash = String(object.solutionHash);
        }
        else {
            message.solutionHash = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = '';
        }
        if (object.reward !== undefined && object.reward !== null) {
            message.reward = String(object.reward);
        }
        else {
            message.reward = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.solutionHash !== undefined && (obj.solutionHash = message.solutionHash);
        message.description !== undefined && (obj.description = message.description);
        message.reward !== undefined && (obj.reward = message.reward);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSubmitScavenge };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.solutionHash !== undefined && object.solutionHash !== null) {
            message.solutionHash = object.solutionHash;
        }
        else {
            message.solutionHash = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = '';
        }
        if (object.reward !== undefined && object.reward !== null) {
            message.reward = object.reward;
        }
        else {
            message.reward = '';
        }
        return message;
    }
};
const baseMsgSubmitScavengeResponse = {};
export const MsgSubmitScavengeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSubmitScavengeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgSubmitScavengeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSubmitScavengeResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    SubmitScavenge(request) {
        const data = MsgSubmitScavenge.encode(request).finish();
        const promise = this.rpc.request('heisenberg.scavenge.scavenge.Msg', 'SubmitScavenge', data);
        return promise.then((data) => MsgSubmitScavengeResponse.decode(new Reader(data)));
    }
}