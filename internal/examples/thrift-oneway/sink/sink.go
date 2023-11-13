// Code generated by thriftrw v1.29.2. DO NOT EDIT.
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

package sink

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type SinkRequest struct {
	Message string `json:"message,required"`
}

// ToWire translates a SinkRequest struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *SinkRequest) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SinkRequest struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SinkRequest struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SinkRequest
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SinkRequest) FromWire(w wire.Value) error {
	var err error

	messageIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}

	if !messageIsSet {
		return errors.New("field Message of SinkRequest is required")
	}

	return nil
}

// Encode serializes a SinkRequest struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a SinkRequest struct could not be encoded.
func (v *SinkRequest) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Message); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a SinkRequest struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a SinkRequest struct could not be generated from the wire
// representation.
func (v *SinkRequest) Decode(sr stream.Reader) error {

	messageIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Message, err = sr.ReadString()
			if err != nil {
				return err
			}
			messageIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !messageIsSet {
		return errors.New("field Message of SinkRequest is required")
	}

	return nil
}

// String returns a readable string representation of a SinkRequest
// struct.
func (v *SinkRequest) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++

	return fmt.Sprintf("SinkRequest{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SinkRequest match the
// provided SinkRequest.
//
// This function performs a deep comparison.
func (v *SinkRequest) Equals(rhs *SinkRequest) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message == rhs.Message) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SinkRequest.
func (v *SinkRequest) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message", v.Message)
	return err
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *SinkRequest) GetMessage() (o string) {
	if v != nil {
		o = v.Message
	}
	return
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "sink",
	Package:  "go.uber.org/yarpc/internal/examples/thrift-oneway/sink",
	FilePath: "sink.thrift",
	SHA1:     "f723dc578c8ae4251c75f873a5a24b6c1a9df61f",
	Raw:      rawIDL,
}

const rawIDL = "service Hello {\n    oneway void sink(1:SinkRequest snk)\n}\n\nstruct SinkRequest {\n    1: required string message;\n}\n"

// Hello_Sink_Args represents the arguments for the Hello.sink function.
//
// The arguments for sink are sent and received over the wire as this struct.
type Hello_Sink_Args struct {
	Snk *SinkRequest `json:"snk,omitempty"`
}

// ToWire translates a Hello_Sink_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Hello_Sink_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Snk != nil {
		w, err = v.Snk.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _SinkRequest_Read(w wire.Value) (*SinkRequest, error) {
	var v SinkRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Hello_Sink_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Hello_Sink_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Hello_Sink_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Hello_Sink_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Snk, err = _SinkRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a Hello_Sink_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Hello_Sink_Args struct could not be encoded.
func (v *Hello_Sink_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Snk != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.Snk.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

func _SinkRequest_Decode(sr stream.Reader) (*SinkRequest, error) {
	var v SinkRequest
	err := v.Decode(sr)
	return &v, err
}

// Decode deserializes a Hello_Sink_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Hello_Sink_Args struct could not be generated from the wire
// representation.
func (v *Hello_Sink_Args) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TStruct:
			v.Snk, err = _SinkRequest_Decode(sr)
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a Hello_Sink_Args
// struct.
func (v *Hello_Sink_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Snk != nil {
		fields[i] = fmt.Sprintf("Snk: %v", v.Snk)
		i++
	}

	return fmt.Sprintf("Hello_Sink_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Hello_Sink_Args match the
// provided Hello_Sink_Args.
//
// This function performs a deep comparison.
func (v *Hello_Sink_Args) Equals(rhs *Hello_Sink_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Snk == nil && rhs.Snk == nil) || (v.Snk != nil && rhs.Snk != nil && v.Snk.Equals(rhs.Snk))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Hello_Sink_Args.
func (v *Hello_Sink_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Snk != nil {
		err = multierr.Append(err, enc.AddObject("snk", v.Snk))
	}
	return err
}

// GetSnk returns the value of Snk if it is set or its
// zero value if it is unset.
func (v *Hello_Sink_Args) GetSnk() (o *SinkRequest) {
	if v != nil && v.Snk != nil {
		return v.Snk
	}

	return
}

// IsSetSnk returns true if Snk is not nil.
func (v *Hello_Sink_Args) IsSetSnk() bool {
	return v != nil && v.Snk != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "sink" for this struct.
func (v *Hello_Sink_Args) MethodName() string {
	return "sink"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be OneWay for this struct.
func (v *Hello_Sink_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

// Hello_Sink_Helper provides functions that aid in handling the
// parameters and return values of the Hello.sink
// function.
var Hello_Sink_Helper = struct {
	// Args accepts the parameters of sink in-order and returns
	// the arguments struct for the function.
	Args func(
		snk *SinkRequest,
	) *Hello_Sink_Args
}{}

func init() {
	Hello_Sink_Helper.Args = func(
		snk *SinkRequest,
	) *Hello_Sink_Args {
		return &Hello_Sink_Args{
			Snk: snk,
		}
	}

}
