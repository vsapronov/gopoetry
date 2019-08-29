package java

import (
	"fmt"
	"gopoetry/util"
)

type UnitDeclaration struct {
	package_     string
	imports      []string
	declarations []util.Writable
}

func Unit(package_ string) *UnitDeclaration {
	return &UnitDeclaration{
		package_,
		[]string{},
		[]util.Writable{},
	}
}

func (self *UnitDeclaration) AddImports(imports ...string) *UnitDeclaration {
	self.imports = append(self.imports, imports...)
	return self
}

func (self *UnitDeclaration) Import(package_ string) *UnitDeclaration {
	self.AddImports(package_)
	return self
}

func (self *UnitDeclaration) AddDeclarations(declarations ...util.Writable) *UnitDeclaration {
	self.declarations = append(self.declarations, declarations...)
	return self
}

func (self *UnitDeclaration) Code() string {
	writer := util.CreateWriter(2)
	self.WriteCode(&writer)
	return writer.Code()
}

func (self *UnitDeclaration) WriteCode(writer util.CodeWriter) {
	writer.Write(fmt.Sprintf("package %s;", self.package_))
	writer.Eol()
	if len(self.imports) > 0 {
		writer.Eol()
		for _, import_ := range self.imports {
			writer.Write(fmt.Sprintf("import %s;", import_))
			writer.Eol()
		}
	}
	writer.Eol()
	for index, class := range self.declarations {
		if index > 0 {
			writer.Eol()
		}
		class.WriteCode(writer)
	}
}