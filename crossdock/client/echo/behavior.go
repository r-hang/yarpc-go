// Copyright (c) 2016 Uber Technologies, Inc.
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

package echo

import (
	"github.com/yarpc/yarpc-go/crossdock-go"
	"github.com/yarpc/yarpc-go/crossdock/client/params"
)

// echoEntry is an entry emitted by the echo behaviors.
type echoEntry struct {
	crossdock.Entry

	Transport string `json:"transport"`
	Encoding  string `json:"encoding"`
	Server    string `json:"server"`
}

// echoT wraps a sink to emit echoEntry entries instead.
type echoT struct {
	crossdock.T

	Transport string
	Encoding  string
	Server    string
}

func (t echoT) Put(e interface{}) {
	t.T.Put(echoEntry{
		Entry:     e.(crossdock.Entry),
		Transport: t.Transport,
		Encoding:  t.Encoding,
		Server:    t.Server,
	})
}

// createEchoT wraps a Sink to have transport, encoding, and server
// information.
func createEchoT(encoding string, t crossdock.T) crossdock.T {
	return echoT{
		T:         t,
		Transport: t.Param(params.Transport),
		Encoding:  encoding,
		Server:    t.Param(params.Server),
	}
}