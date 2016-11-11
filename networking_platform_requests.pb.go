// Code generated by protoc-gen-go.
// source: networking_platform_requests.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type BuyItemAndroidRequest struct {
	BuyItemIntent string `protobuf:"bytes,1,opt,name=buy_item_intent" json:"buy_item_intent,omitempty"`
}

func (m *BuyItemAndroidRequest) Reset()         { *m = BuyItemAndroidRequest{} }
func (m *BuyItemAndroidRequest) String() string { return proto.CompactTextString(m) }
func (*BuyItemAndroidRequest) ProtoMessage()    {}

type BuyItemPokeCoinsRequest struct {
	ItemId string `protobuf:"bytes,1,opt,name=item_id" json:"item_id,omitempty"`
}

func (m *BuyItemPokeCoinsRequest) Reset()         { *m = BuyItemPokeCoinsRequest{} }
func (m *BuyItemPokeCoinsRequest) String() string { return proto.CompactTextString(m) }
func (*BuyItemPokeCoinsRequest) ProtoMessage()    {}

type SendEncryptedSignatureRequest struct {
	EncryptedSignature []byte `protobuf:"bytes,1,opt,name=encrypted_signature,proto3" json:"encrypted_signature,omitempty"`
}

func (m *SendEncryptedSignatureRequest) Reset()         { *m = SendEncryptedSignatureRequest{} }
func (m *SendEncryptedSignatureRequest) String() string { return proto.CompactTextString(m) }
func (*SendEncryptedSignatureRequest) ProtoMessage()    {}

func init() {
}
