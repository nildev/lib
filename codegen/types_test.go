package codegen

import . "gopkg.in/check.v1"

type TypesSuite struct{}

var _ = Suite(&TypesSuite{})

func (s *TypesSuite) TestIfCorrectTypesAreReturned(c *C) {
	nf := &Field{Name: "notFound", Type: "string"}

	str := &Field{Name: "string", Type: "string"}
	strPtr := &Field{Name: "stringPtr", Type: "*string"}
	inti := &Field{Name: "int", Type: "int"}
	intiPtr := &Field{Name: "intPtr", Type: "*int"}
	inti8 := &Field{Name: "int8", Type: "int8"}
	intiPtr8 := &Field{Name: "intPtr8", Type: "*int8"}
	inti16 := &Field{Name: "int16", Type: "int16"}
	intiPtr16 := &Field{Name: "intPtr16", Type: "*int16"}
	inti32 := &Field{Name: "int32", Type: "int32"}
	intiPtr32 := &Field{Name: "intPtr32", Type: "*int32"}
	inti64 := &Field{Name: "int64", Type: "int64"}
	intiPtr64 := &Field{Name: "intPtr64", Type: "*int64"}
	flt32 := &Field{Name: "flt32", Type: "float32"}
	fltPtr32 := &Field{Name: "fltPtr32", Type: "*float32"}
	flt64 := &Field{Name: "flt64", Type: "float64"}
	fltPtr64 := &Field{Name: "fltPtr64", Type: "*float64"}
	booly := &Field{Name: "booly", Type: "bool"}
	boolyPtr := &Field{Name: "boolyPtr", Type: "*bool"}

	data := map[string]string{
		"string":    "my value",
		"stringPtr": "my value x",
		"int":       "2",
		"intPtr":    "6",
		"int8":      "6",
		"intPtr8":   "6",
		"int16":     "6",
		"intPtr16":  "6",
		"int32":     "6",
		"intPtr32":  "6",
		"int64":     "6",
		"intPtr64":  "6",
		"flt32":     "6.77",
		"fltPtr32":  "6.77",
		"flt64":     "6.77",
		"fltPtr64":  "6.77",
		"booly":     "true",
		"boolyPtr":  "false",
	}

	type (
		Tmp struct {
			String    string
			StringPtr *string
			Int       int
			IntPtr    *int
			Int8      int8
			IntPtr8   *int8
			Int16     int16
			IntPtr16  *int16
			Int32     int32
			IntPtr32  *int32
			Int64     int64
			IntPtr64  *int64
			Flt32     float32
			FltPtr32  *float32
			Flt64     float64
			FltPtr64  *float64
			Booly     bool
			BoolyPtr  *bool
		}
	)

	xTmp := &Tmp{}

	nilval, e := nf.GetVarValue(data)
	c.Assert(e, IsNil)
	c.Assert(nilval, IsNil)

	strVal, e := str.GetVarValue(data)
	xTmp.String = strVal.(string)

	c.Assert(e, IsNil)
	c.Assert(xTmp.String, Equals, "my value")

	strValPtr, e := strPtr.GetVarValue(data)
	xTmp.StringPtr = strValPtr.(*string)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.StringPtr, Equals, "my value x")

	intVal, e := inti.GetVarValue(data)
	xTmp.Int = intVal.(int)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Int, Equals, 2)

	intValPtr, e := intiPtr.GetVarValue(data)
	xTmp.IntPtr = intValPtr.(*int)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.IntPtr, Equals, 6)

	intVal8, e := inti8.GetVarValue(data)
	xTmp.Int8 = intVal8.(int8)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Int8, Equals, int8(6))

	intValPtr8, e := intiPtr8.GetVarValue(data)
	xTmp.IntPtr8 = intValPtr8.(*int8)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.IntPtr8, Equals, int8(6))

	intVal16, e := inti16.GetVarValue(data)
	xTmp.Int16 = intVal16.(int16)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Int16, Equals, int16(6))

	intValPtr16, e := intiPtr16.GetVarValue(data)
	xTmp.IntPtr16 = intValPtr16.(*int16)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.IntPtr16, Equals, int16(6))

	intVal32, e := inti32.GetVarValue(data)
	xTmp.Int32 = intVal32.(int32)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Int32, Equals, int32(6))

	intValPtr32, e := intiPtr32.GetVarValue(data)
	xTmp.IntPtr32 = intValPtr32.(*int32)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.IntPtr32, Equals, int32(6))

	intVal64, e := inti64.GetVarValue(data)
	xTmp.Int64 = intVal64.(int64)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Int64, Equals, int64(6))

	intValPtr64, e := intiPtr64.GetVarValue(data)
	xTmp.IntPtr64 = intValPtr64.(*int64)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.IntPtr64, Equals, int64(6))

	fltVal32, e := flt32.GetVarValue(data)
	xTmp.Flt32 = fltVal32.(float32)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Flt32, Equals, float32(6.77))

	fltValPtr32, e := fltPtr32.GetVarValue(data)
	xTmp.FltPtr32 = fltValPtr32.(*float32)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.FltPtr32, Equals, float32(6.77))

	fltVal64, e := flt64.GetVarValue(data)
	xTmp.Flt64 = fltVal64.(float64)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Flt64, Equals, float64(6.77))

	fltValPtr64, e := fltPtr64.GetVarValue(data)
	xTmp.FltPtr64 = fltValPtr64.(*float64)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.FltPtr64, Equals, float64(6.77))

	boolyVal, e := booly.GetVarValue(data)
	xTmp.Booly = boolyVal.(bool)
	c.Assert(e, IsNil)
	c.Assert(xTmp.Booly, Equals, true)

	boolyValPtr, e := boolyPtr.GetVarValue(data)
	xTmp.BoolyPtr = boolyValPtr.(*bool)
	c.Assert(e, IsNil)
	c.Assert(*xTmp.BoolyPtr, Equals, false)
}
