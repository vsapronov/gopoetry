package scala

type StatementsDeclaration struct {
	statements []Writable
	block      bool
	scope      bool
}

func (self *StatementsDeclaration) Add(code Writable) *StatementsDeclaration {
	self.statements = append(self.statements, code)
	return self
}

func (self *StatementsDeclaration) Scope() *StatementsDeclaration {
	body := Scope()
	self.Add(body)
	return body
}

func (self *StatementsDeclaration) Def(name string) *MethodDeclaration {
	method := Method(name)
	self.Add(method)
	return method
}

func (self *StatementsDeclaration) Val(name string, type_ string) *FieldDeclaration {
	field := Val(name, type_)
	self.Add(field)
	return field
}

func (self *StatementsDeclaration) Var(name string, type_ string) *FieldDeclaration {
	field := Var(name, type_)
	self.Add(field)
	return field
}

func Statements(statements ...Writable) *StatementsDeclaration {
	return &StatementsDeclaration{statements: statements, block: false, scope: false}
}

func Block(statements ...Writable) *StatementsDeclaration {
	return &StatementsDeclaration{statements: statements, block: true, scope: false}
}

func Scope(statements ...Writable) *StatementsDeclaration {
	return &StatementsDeclaration{statements: statements, block: true, scope: true}
}

func ScopeInline(statements ...Writable) *StatementsDeclaration {
	return &StatementsDeclaration{statements: statements, block: false, scope: true}
}

func (self *StatementsDeclaration) WriteCode(writer CodeWriter) {
	if self.scope {
		writer.Write("{")
		if self.block {
			writer.Eol()
		} else {
			writer.Write(" ")
		}
	}
	if self.block {
		writer.Indent()
	}
	for _, statement := range self.statements {
		statement.WriteCode(writer)
	}
	if self.block {
		writer.UnIndent()
	}
	if self.scope {
		if !self.block {
			writer.Write(" ")
		}
		writer.Write("}")
		writer.Eol()
	}
}
