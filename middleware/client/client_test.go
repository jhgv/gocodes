package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

type File struct {

}

type Body struct {
	File []byte
	Test string
}

type Message struct {
	Body Body
}

func TestFileMarshaller(t *testing.T) {
	absPath, _ := filepath.Abs("file.csv")
	//file, _ := ioutil.ReadFile(absPath)

	b, _ := ioutil.ReadFile(absPath)

	if b == nil {
		t.Error("Fail")
	}
}
