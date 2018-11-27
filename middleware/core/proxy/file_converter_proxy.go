package proxy

import (
	"github.com/jhgv/gocodes/middleware/core/requestor"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type FileConverterProxy struct {
	ClientProxy
}

func (fp *FileConverterProxy) ConvertFile(file *os.File) []byte {
	return sendInvocation(fp, nameOf((*FileConverterProxy).ConvertFile), file)
}

func sendInvocation(fp *FileConverterProxy, methodName string, file *os.File) []byte {
	var req = &requestor.Requestor{}
	var inv = &requestor.Invocation{}
	var ter *requestor.Termination

	inv.EnabledCompression = true
	inv.SetHost(fp.GetHost())
	inv.SetPort(fp.GetPort())
	fileInfo, _ := file.Stat()
	b := make([]byte, fileInfo.Size())
	file.Read(b)
	inv.SetFile(b)
	inv.SetOperationName(methodName)
	inv.SetObjectID(fp.GetID())

	ter, _ = req.Invoke(inv)

	return ter.GetResult().([]byte)

}

func nameOf(f interface{}) string {
	var methodName string
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			methodName = rf.Name()
		}
	}
	if tokenizedName := strings.Split(methodName, "."); len(tokenizedName) > 0 {
		methodName = tokenizedName[len(tokenizedName)-1]
	}
	return methodName
}
