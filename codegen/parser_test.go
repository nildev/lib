package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"

	. "gopkg.in/check.v1"
)

type StructSuite struct{}

var _ = Suite(&StructSuite{})

func (s *StructSuite) TestIfCorrectFieldsAreBeingGeneratedForStruct(c *C) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "fixtures/funcs.go", nil, parser.AllErrors)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == "CheckPrimitive" {
				fields := makeFields(x.Type.Params)

				c.Assert("string", Equals, fields["msg"].Type)
				c.Assert("int", Equals, fields["num"].Type)
				c.Assert("int8", Equals, fields["num8"].Type)
				c.Assert("int16", Equals, fields["num16"].Type)
				c.Assert("int32", Equals, fields["num32"].Type)
				c.Assert("int64", Equals, fields["num64"].Type)
				c.Assert("float32", Equals, fields["flt32"].Type)
				c.Assert("float64", Equals, fields["flt64"].Type)

				ofields := makeFields(x.Type.Results)

				c.Assert("string", Equals, ofields["rmsg"].Type)
				c.Assert("int", Equals, ofields["rnum"].Type)
				c.Assert("int8", Equals, ofields["rnum8"].Type)
				c.Assert("int16", Equals, ofields["rnum16"].Type)
				c.Assert("int32", Equals, ofields["rnum32"].Type)
				c.Assert("int64", Equals, ofields["rnum64"].Type)
				c.Assert("float32", Equals, ofields["rflt32"].Type)
				c.Assert("float64", Equals, ofields["rflt64"].Type)
			}

			if x.Name.Name == "CheckNestedMap" {
				fields := makeFields(x.Type.Params)

				c.Assert(fields["mapSS"].Type, Equals, "map[string]string")
				c.Assert(fields["mapSMI"].Type, Equals, "map[string]map[string]int")
				c.Assert(fields["mapASAMSAS"].Type, Equals, "map[int][]map[string][]string")
				c.Assert(fields["mapPSSI"].Type, Equals, "map[*string]sub.MyInterface")
				c.Assert(fields["mapMPSI"].Type, Equals, "map[*MyStruct]int")
				c.Assert(fields["mapMSI"].Type, Equals, "map[MyStruct]int")
				c.Assert(fields["mapMSII"].Type, Equals, "map[sub.MyInterface]int")
				c.Assert(fields["mapMSIISI"].Type, Equals, "map[struct{}]interface{}")

				fields = makeFields(x.Type.Results)

				c.Assert(fields["rmapSS"].Type, Equals, "map[string]string")
				c.Assert(fields["rmapSMI"].Type, Equals, "map[string]map[string]int")
				c.Assert(fields["rmapASAMSAS"].Type, Equals, "map[int][]map[string][]string")
				c.Assert(fields["rmapPSSI"].Type, Equals, "map[*string]sub.MyInterface")
				c.Assert(fields["rmapMPSI"].Type, Equals, "map[*MyStruct]int")
				c.Assert(fields["rmapMSI"].Type, Equals, "map[MyStruct]int")
				c.Assert(fields["rmapMSII"].Type, Equals, "map[sub.MyInterface]int")
				c.Assert(fields["rmapMSIISI"].Type, Equals, "map[struct{}]interface{}")
			}

			// for variadic funcs we will make only 1 field
			if x.Name.Name == "CheckVariadicParam" {
				fields := makeFields(x.Type.Params)

				c.Assert(fields["x"].Type, Equals, "map[sub.MyInterface]int")
			}

			if x.Name.Name == "CheckComplex" {
				fields := makeFields(x.Type.Params)

				c.Assert(fields["str1"].Type, Equals, "MyStruct")
				c.Assert(fields["str2"].Type, Equals, "MyOtherStruct")
				c.Assert(fields["inteM"].Type, Equals, "sub.MyInterface")
				c.Assert(fields["inte"].Type, Equals, "interface{}")
				c.Assert(fields["st"].Type, Equals, "struct{}")
				c.Assert(fields["ai"].Type, Equals, "alias.MyInterface")

				fields = makeFields(x.Type.Results)

				c.Assert(fields["rstr1"].Type, Equals, "MyStruct")
				c.Assert(fields["rstr2"].Type, Equals, "MyOtherStruct")
				c.Assert(fields["rinteM"].Type, Equals, "sub.MyInterface")
				c.Assert(fields["rinte"].Type, Equals, "interface{}")
				c.Assert(fields["rst"].Type, Equals, "struct{}")
				c.Assert(fields["rai"].Type, Equals, "alias.MyInterface")
			}
		}
		return true
	})

}

func (s *StructSuite) TestIfCorrectStructsAreBeingGenerated(c *C) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "fixtures/funcs.go", nil, parser.ParseComments)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == "CheckPrimitive" {
				i, o := MakeInputOutputStructs(x, f.Imports, f.Comments)

				c.Assert(i.Name, Equals, "InputCheckPrimitive")
				c.Assert(o.Name, Equals, "OutputCheckPrimitive")
				c.Assert(i.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Alias, Equals, "")
				c.Assert(i.Imports["alias"].Alias, Equals, "alias")
				c.Assert(i.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/sub")
				c.Assert(i.Imports["alias"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/suby")
				c.Assert(o.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Alias, Equals, "")
				c.Assert(o.Imports["alias"].Alias, Equals, "alias")
				c.Assert(o.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/sub")
				c.Assert(o.Imports["alias"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/suby")
			}

			if x.Name.Name == "CheckNestedMap" {
				i, o := MakeInputOutputStructs(x, f.Imports, f.Comments)

				c.Assert(i.Name, Equals, "InputCheckNestedMap")
				c.Assert(o.Name, Equals, "OutputCheckNestedMap")
				c.Assert(i.Imports["alias"].Alias, Equals, "alias")
				c.Assert(i.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/sub")
				c.Assert(i.Imports["alias"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/suby")
				c.Assert(o.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Alias, Equals, "")
				c.Assert(o.Imports["alias"].Alias, Equals, "alias")
				c.Assert(o.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/sub")
				c.Assert(o.Imports["alias"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/suby")
			}

			// for variadic funcs we will make only 1 field
			if x.Name.Name == "CheckVariadicParam" {
				i, o := MakeInputOutputStructs(x, f.Imports, f.Comments)

				c.Assert(i.Name, Equals, "InputCheckVariadicParam")
				c.Assert(o, IsNil)

				c.Assert(i.Imports["alias"].Alias, Equals, "alias")
				c.Assert(i.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/sub")
				c.Assert(i.Imports["alias"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/suby")
			}

			if x.Name.Name == "CheckComplex" {
				i, o := MakeInputOutputStructs(x, f.Imports, f.Comments)

				c.Assert(i.Name, Equals, "InputCheckComplex")
				c.Assert(o.Name, Equals, "OutputCheckComplex")
				c.Assert(i.Imports["alias"].Alias, Equals, "alias")
				c.Assert(i.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/sub")
				c.Assert(i.Imports["alias"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/suby")
				c.Assert(o.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Alias, Equals, "")
				c.Assert(o.Imports["alias"].Alias, Equals, "alias")
				c.Assert(o.Imports["bitbucket.org/nildev/lib/codegen/fixtures/sub"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/sub")
				c.Assert(o.Imports["alias"].Path, Equals, "bitbucket.org/nildev/lib/codegen/fixtures/suby")
			}
		}
		return true
	})
}

func (s *StructSuite) TestIfCorrectFuncsAreBeingGenerated(c *C) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "fixtures/funcs.go", nil, parser.ParseComments)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == "CheckPrimitive" {
				fn := MakeFunc(x, f.Imports, f.Comments)

				c.Assert(fn.GetHandlerName(), Equals, "CheckPrimitiveHandler")
				c.Assert(fn.GetMethod(), Equals, "POST")
				c.Assert(fn.GetFullName(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures:CheckPrimitive")
				c.Assert(fn.GetPattern(), Equals, "/CheckPrimitive")
				c.Assert(fn.GetPkgPath(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures")
				c.Assert(fn.GetOnlyPkgName(), Equals, "fixtures")
			}

			if x.Name.Name == "CheckNestedMap" {
				fn := MakeFunc(x, f.Imports, f.Comments)

				c.Assert(fn.GetHandlerName(), Equals, "CheckNestedMapHandler")
				c.Assert(fn.GetMethod(), Equals, "POST")
				c.Assert(fn.GetFullName(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures:CheckNestedMap")
				c.Assert(fn.GetPattern(), Equals, "/CheckNestedMap")
				c.Assert(fn.GetPkgPath(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures")
				c.Assert(fn.GetOnlyPkgName(), Equals, "fixtures")
			}

			if x.Name.Name == "CheckVariadicParam" {
				fn := MakeFunc(x, f.Imports, f.Comments)
				c.Assert(fn.GetHandlerName(), Equals, "CheckVariadicParamHandler")
				c.Assert(fn.GetMethod(), Equals, "POST")
				c.Assert(fn.GetFullName(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures:CheckVariadicParam")
				c.Assert(fn.GetPattern(), Equals, "/CheckVariadicParam")
				c.Assert(fn.GetPkgPath(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures")
				c.Assert(fn.GetOnlyPkgName(), Equals, "fixtures")
			}

			if x.Name.Name == "CheckComplex" {
				fn := MakeFunc(x, f.Imports, f.Comments)
				c.Assert(fn.GetHandlerName(), Equals, "CheckComplexHandler")
				c.Assert(fn.GetMethod(), Equals, "POST")
				c.Assert(fn.GetFullName(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures:CheckComplex")
				c.Assert(fn.GetPattern(), Equals, "/CheckComplex")
				c.Assert(fn.GetPkgPath(), Equals, "bitbucket.org/nildev/lib/codegen/fixtures")
				c.Assert(fn.GetOnlyPkgName(), Equals, "fixtures")
			}
		}
		return true
	})
}
