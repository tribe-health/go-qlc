package sync

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/sync/pb"
)

type FrontierReq struct {
	StartAddress types.Address
	Age          uint32
	Count        uint32
}

func NewFrontierReq(addr types.Address, Age, Count uint32) (packet *FrontierReq) {
	return &FrontierReq{
		StartAddress: addr,
		Age:          Age,
		Count:        Count,
	}
}

// ToProto converts domain frontier into proto frontier
func FrontierReqToProto(fr *FrontierReq) ([]byte, error) {
	//pb := new(pb.Frontier)
	address := fr.StartAddress.Bytes()
	frpb := &pb.FrontierReq{
		Address: address,
		Age:     fr.Age,
		Count:   fr.Count,
	}
	data, err := proto.Marshal(frpb)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/// FrontierReqFromProto parse the data into frontier message
func FrontierReqFromProto(data []byte) (*FrontierReq, error) {
	fr := new(pb.FrontierReq)

	if err := proto.Unmarshal(data, fr); err != nil {
		fmt.Println("Failed to unmarshal FrontierReqPacket message.")
		return nil, err
	}
	address, err := types.BytesToAddress(fr.Address)
	if err != nil {
		fmt.Println("address error")
	}
	pb := &FrontierReq{
		StartAddress: address,
		Age:          fr.Age,
		Count:        fr.Count,
	}
	return pb, nil
}

type FrontierResponse struct {
	Frontier *types.Frontier
}

func NewFrontierRsp(fr *types.Frontier) (packet *FrontierResponse) {
	return &FrontierResponse{
		Frontier: fr,
	}
}

// ToProto converts domain FrontierResponse into proto FrontierResponse
func FrontierResponseToProto(fr *FrontierResponse) ([]byte, error) {
	pi := &pb.FrontierRsp{
		HeaderBlock: fr.Frontier.HeaderBlock[:],
		OpenBlock:   fr.Frontier.OpenBlock[:],
	}
	data, err := proto.Marshal(pi)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/// FrontierResponseFromProto parse the data into frontier message
func FrontierResponseFromProto(data []byte) (*FrontierResponse, error) {
	fr := new(pb.FrontierRsp)
	frp := new(types.Frontier)
	if err := proto.Unmarshal(data, fr); err != nil {
		fmt.Println("Failed to unmarshal FrontierRspPacket message.")
		return nil, err
	}
	err := frp.HeaderBlock.UnmarshalBinary(fr.HeaderBlock[:])
	if err != nil {
		return nil, err
	}
	err = frp.OpenBlock.UnmarshalBinary(fr.OpenBlock[:])
	if err != nil {
		return nil, err
	}
	frps := &FrontierResponse{
		Frontier: frp,
	}
	return frps, nil
}
