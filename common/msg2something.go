package main

import (
    "bytes"
    "errors"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/ptypes/any"

    "github.com/golang/protobuf/jsonpb"
    _struct "github.com/golang/protobuf/ptypes/struct"
)

func MessageToStruct(msg proto.Message) (*_struct.Struct, error) {
    if msg == nil {
        return nil, errors.New("nil message")
    }

    buf := &bytes.Buffer{}
    if err := (&jsonpb.Marshaler{OrigName: true}).Marshal(buf, msg); err != nil {
        return nil, err
    }

    pbs := &_struct.Struct{}
    if err := jsonpb.Unmarshal(buf, pbs); err != nil {
        return nil, err
    }

    return pbs, nil
}

func MessageToAny(msg proto.Message) (*any.Any ,error) {
	if msg == nil {
        return nil, errors.New("nil message")
    }

    buffer := proto.NewBuffer(nil)
    buffer.SetDeterministic(true)
    
    if err := buffer.Marshal(msg); err != nil {
            return nil,err
    }

    return &any.Any{
        TypeUrl: "type.googleapis.com/" + proto.MessageName(msg),
        Value:   buffer.Bytes(),
    }, nil
}