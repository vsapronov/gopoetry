package scala

import (
	"strings"
)

type CodeWriter interface {
	Eol()
	Indent()
	UnIndent()
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

func (self *codeWriter) UnIndent() {
	self.indentation -= 1
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
