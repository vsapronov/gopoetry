package java

import (
	"testing"
)

func TestStatementsSimple(t *testing.T) {
	expected := `
line1()
line2()
`
	AssertCode(t, Statements().Lines("line1()", "line2()"), expected)
}

func TestStatementsWithBlock(t *testing.T) {
	expected := `
line1()
line2 {
  nextedLine1()
  nextedLine2()
}
`
	statements := Statements()
	statements.Line("line1()")
	statements.Append("line2")
	statements.Block().Lines(
		"nextedLine1()",
		"nextedLine2()",
	)
	AssertCode(t, statements, expected)
}
