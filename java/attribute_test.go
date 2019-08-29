package java

import (
	"gopoetry/util"
	"testing"
)

func TestAttributeSimple(t *testing.T) {
	util.AssertCode(t, Attribute("MyAttribute(a, b, c = c)"), `@MyAttribute(a, b, c = c)`)
}