package grpcclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"google.golang.org/grpc"
)

func CallRPCJSON(
	server string,
	svc *Service,
	method *Method,
	payload map[string]interface{},
	protoset *ProtoSet,
) (map[string]interface{}, error) {

	if svc == nil || method == nil || method.Desc == nil {
		return nil, fmt.Errorf("service, method, or descriptor is nil")
	}

	// 1️⃣ Connect to gRPC server
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stub := grpcdynamic.NewStub(conn)

	// 2️⃣ Build dynamic request message
	reqMsg := dynamic.NewMessage(method.Desc.GetInputType())

	// Convert payload map to JSON bytes
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Unmarshal JSON into dynamic message
	if err := reqMsg.UnmarshalJSON(jsonBytes); err != nil {
		return nil, fmt.Errorf("failed to populate request: %v", err)
	}

	// 3️⃣ Call RPC
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := stub.InvokeRpc(ctx, method.Desc, reqMsg)
	if err != nil {
		return nil, fmt.Errorf("RPC call failed: %v", err)
	}

	// 4️⃣ Convert response to map
	var out map[string]interface{}
	switch m := resp.(type) {
	case *dynamic.Message:
		b, err := m.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %v", err)
		}
		if err := json.Unmarshal(b, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response JSON: %v", err)
		}
	default:
		// fallback: marshal proto.Message to JSON then unmarshal
		b, err := json.Marshal(resp)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal proto.Message: %v", err)
		}
		if err := json.Unmarshal(b, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response JSON: %v", err)
		}
	}

	return out, nil
}
