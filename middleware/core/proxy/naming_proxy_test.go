package proxy_test

import (
	"testing"
)

func TestNamingProxy(t *testing.T) {
	namingProxy := &NamingProxy{}
	var expectedObj = &TextHelperProxy{}
	namingProxy.Bind("Test", expectedObj)
	actualObject, _ := namingProxy.Lookup("Test")
	if actualObject == nil {
		t.Error("NamingProxy was expected to find an object, none was found")
	}

	if actualObject != expectedObj {
		t.Error("NamingProxy returned wrong object")
	}
}
