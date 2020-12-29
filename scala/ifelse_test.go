package scala

import "testing"

func TestIfElseBasic(t *testing.T) {
	assertCode(t, If(true).Then(Code("true code")).Else(Code("false code")), "true code")
	assertCode(t, If(false).Then(Code("true code")).Else(Code("false code")), "false code")
}

func TestIfElseOmit(t *testing.T) {
	assertCode(t, If(true), "")
	assertCode(t, If(false), "")
}

func TestOnlyIf(t *testing.T) {
	assertCode(t, Only(Code("some code")).If(true), "some code")
	assertCode(t, Only(Code("some code")).If(false), "")
}

func TestIfElseLazy(t *testing.T) {
	codeTrue := If(true).Then(Lazy(func() Writable { return Code("true code") }))
	assertCode(t, codeTrue, "true code")
	codeFalse := If(false).Else(Lazy(func() Writable { return Code("false code") }))
	assertCode(t, codeFalse, "false code")
}