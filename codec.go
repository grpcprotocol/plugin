package plugin

import (
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func MarshalExecReq(req proto.Message) (*ExecReq, error) {
	var (
		err     error
		payload *any.Any
	)
	if payload, err = anypb.New(req); err != nil {
		return nil, err
	}

	return &ExecReq{Payload: payload}, nil
}

func MarshalExecRsp(rsp proto.Message) (*ExecRsp, error) {
	var (
		err   error
		reply *any.Any
	)
	if reply, err = anypb.New(rsp); err != nil {
		return nil, err
	}

	return &ExecRsp{Reply: reply}, nil
}

func UnmarshalExecReq(req *ExecReq, msg proto.Message) error {
	return anypb.UnmarshalTo(req.GetPayload(), msg, proto.UnmarshalOptions{})
}

func UnmarshalExecRsp(rsp *ExecRsp, msg proto.Message) error {
	return anypb.UnmarshalTo(rsp.GetReply(), msg, proto.UnmarshalOptions{})
}
