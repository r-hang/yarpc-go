// Code generated by thriftrw-plugin-yarpc
// @generated

// Copyright (c) 2023 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package onewayserver

import (
	context "context"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	oneway "go.uber.org/yarpc/internal/crossdock/thrift/oneway"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the Oneway service.
type Interface interface {
	Echo(
		ctx context.Context,
		Token *string,
	) error
}

// New prepares an implementation of the Oneway service for
// registration.
//
// 	handler := OnewayHandler{}
// 	dispatcher.Register(onewayserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Oneway",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "echo",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.Echo),
					NoWire: echo_NoWireHandler{impl},
				},
				Signature:    "Echo(Token *string)",
				ThriftModule: oneway.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type yarpcErrorNamer interface{ YARPCErrorName() string }

type yarpcErrorCoder interface{ YARPCErrorCode() *yarpcerrors.Code }

func (h handler) Echo(ctx context.Context, body wire.Value) error {
	var args oneway.Oneway_Echo_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.Echo(ctx, args.Token)
}

type echo_NoWireHandler struct{ impl Interface }

func (h echo_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args oneway.Oneway_Echo_Args

		err error
	)

	if _, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args); err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'Oneway' procedure 'Echo': %w", err)
	}

	return thrift.NoWireResponse{}, h.impl.Echo(ctx, args.Token)

}
