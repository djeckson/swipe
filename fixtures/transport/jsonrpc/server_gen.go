//+build !swipe

// Code generated by Swipe v1.20.0. DO NOT EDIT.

//go:generate swipe
package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/l-vitaly/go-kit/transport/http/jsonrpc"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/swipe-io/swipe/fixtures/service"
	"net/http"
	"strings"
)

func encodeResponseJSONRPCServiceInterface(_ context.Context, result interface{}) (json.RawMessage, error) {
	b, err := ffjson.Marshal(result)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func MakeServiceInterfaceEndpointCodecMap(ep EndpointSet, ns ...string) jsonrpc.EndpointCodecMap {
	var namespace = strings.Join(ns, ".")
	if len(ns) > 0 {
		namespace += "."
	}
	return jsonrpc.EndpointCodecMap{
		namespace + "create": jsonrpc.EndpointCodec{
			Endpoint: ep.CreateEndpoint,
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req createRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to createRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		namespace + "delete": jsonrpc.EndpointCodec{
			Endpoint: ep.DeleteEndpoint,
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req deleteRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to deleteRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		namespace + "get": jsonrpc.EndpointCodec{
			Endpoint: ep.GetEndpoint,
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req getRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to getRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		namespace + "getAll": jsonrpc.EndpointCodec{
			Endpoint: ep.GetAllEndpoint,
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				return nil, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
		namespace + "testMethod": jsonrpc.EndpointCodec{
			Endpoint: ep.TestMethodEndpoint,
			Decode: func(_ context.Context, msg json.RawMessage) (interface{}, error) {
				var req testMethodRequestServiceInterface
				err := ffjson.Unmarshal(msg, &req)
				if err != nil {
					return nil, fmt.Errorf("couldn't unmarshal body to testMethodRequestServiceInterface: %s", err)
				}
				return req, nil
			},
			Encode: encodeResponseJSONRPCServiceInterface,
		},
	}
}

// HTTP JSONRPC Transport
func MakeHandlerJSONRPCServiceInterface(s service.Interface, opts ...ServiceInterfaceServerOption) (http.Handler, error) {
	sopt := &serverServiceInterfaceOpts{}
	for _, o := range opts {
		o(sopt)
	}
	ep := MakeEndpointSet(s)
	ep.CreateEndpoint = middlewareChain(append(sopt.genericEndpointMiddleware, sopt.createEndpointMiddleware...))(ep.CreateEndpoint)
	ep.DeleteEndpoint = middlewareChain(append(sopt.genericEndpointMiddleware, sopt.deleteEndpointMiddleware...))(ep.DeleteEndpoint)
	ep.GetEndpoint = middlewareChain(append(sopt.genericEndpointMiddleware, sopt.getEndpointMiddleware...))(ep.GetEndpoint)
	ep.GetAllEndpoint = middlewareChain(append(sopt.genericEndpointMiddleware, sopt.getAllEndpointMiddleware...))(ep.GetAllEndpoint)
	ep.TestMethodEndpoint = middlewareChain(append(sopt.genericEndpointMiddleware, sopt.testMethodEndpointMiddleware...))(ep.TestMethodEndpoint)
	r := mux.NewRouter()
	handler := jsonrpc.NewServer(MakeServiceInterfaceEndpointCodecMap(ep), sopt.genericServerOption...)
	r.Methods("POST").Path("/rpc/{method}").Handler(handler)
	return r, nil
}