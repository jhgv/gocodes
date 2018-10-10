package proxy

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/jhgv/gocodes/middleware/patterns/client/requestor"
)

type TextfyProxy struct {
	ClientProxy
}

func (tp *TextfyProxy) UpperText(message string) (string, error) {
	var req = requestor.Requestor{}
	var inv requestor.Invocation
	var ter requestor.Termination
	params := []string{message}

	inv.SetHost(tp.getHost())
	inv.SetPort(tp.getPort())
	inv.SetParameters(params)
	inv.SetOperationName(nameOf(tp.UpperText))

	ter, _ = req.Invoke(inv)

	return ter.GetResult(), nil
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
