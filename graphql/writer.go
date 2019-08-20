package graphql

import (
	"strings"
)

type CodeWriter interface {
	Indent()
	Unindent()
	Begin()
	End()
	Eol()
	Write(code string)
}

func prefix(indentation int) string {
	tab := strings.Repeat(" ", 2)
	return strings.Repeat(tab, indentation)
}

type codeWriter struct {
	builder     strings.Builder
	indentation int
	newLine     bool
}

func (self *codeWriter) Indent() {
	self.indentation += 1
}

func (self *codeWriter) Unindent() {
	self.indentation -= 1
}

func (self *codeWriter) Begin() {
	self.Write("{")
	self.Eol()
	self.indentation += 1
}

func (self *codeWriter) End() {
	self.indentation -= 1
	self.Write("}")
	self.Eol()
}

func (self *codeWriter) Eol() {
	self.Write("\n")
	self.newLine = true
}

func (self *codeWriter) Write(code string) {
	if self.newLine {
		self.builder.WriteString(prefix(self.indentation))
		self.newLine = false
	}
	self.builder.WriteString(code)
}

func (self *codeWriter) Code() string {
	return self.builder.String()
}

func CreateWriter() codeWriter {
	return codeWriter{builder: strings.Builder{}, indentation: 0, newLine: true}
}