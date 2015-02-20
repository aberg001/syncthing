// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package xdr_test

import (
	"bytes"
	"io"

	"github.com/calmh/xdr"
)

/*

XDRBenchStruct Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                         I1 (64 bits)                          +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                              I2                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|            0x0000             |              I3               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             uint8                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of Bs0                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     Bs0 (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of Bs1                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     Bs1 (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of S0                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     S0 (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of S1                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     S1 (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct XDRBenchStruct {
	unsigned hyper I1;
	unsigned int I2;
	unsigned int I3;
	uint8 I4;
	opaque Bs0<128>;
	opaque Bs1<>;
	string S0<128>;
	string S1<>;
}

*/

func (o XDRBenchStruct) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o XDRBenchStruct) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o XDRBenchStruct) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o XDRBenchStruct) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o XDRBenchStruct) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteUint64(o.I1)
	xw.WriteUint32(o.I2)
	xw.WriteUint16(o.I3)
	xw.WriteUint8(o.I4)
	if l := len(o.Bs0); l > 128 {
		return xw.Tot(), xdr.ElementSizeExceeded("Bs0", l, 128)
	}
	xw.WriteBytes(o.Bs0)
	xw.WriteBytes(o.Bs1)
	if l := len(o.S0); l > 128 {
		return xw.Tot(), xdr.ElementSizeExceeded("S0", l, 128)
	}
	xw.WriteString(o.S0)
	xw.WriteString(o.S1)
	return xw.Tot(), xw.Error()
}

func (o *XDRBenchStruct) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *XDRBenchStruct) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *XDRBenchStruct) decodeXDR(xr *xdr.Reader) error {
	o.I1 = xr.ReadUint64()
	o.I2 = xr.ReadUint32()
	o.I3 = xr.ReadUint16()
	o.I4 = xr.ReadUint8()
	o.Bs0 = xr.ReadBytesMax(128)
	o.Bs1 = xr.ReadBytes()
	o.S0 = xr.ReadStringMax(128)
	o.S1 = xr.ReadString()
	return xr.Error()
}

/*

repeatReader Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of data                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    data (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct repeatReader {
	opaque data<>;
}

*/

func (o repeatReader) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o repeatReader) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o repeatReader) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o repeatReader) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o repeatReader) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteBytes(o.data)
	return xw.Tot(), xw.Error()
}

func (o *repeatReader) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *repeatReader) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *repeatReader) decodeXDR(xr *xdr.Reader) error {
	o.data = xr.ReadBytes()
	return xr.Error()
}
