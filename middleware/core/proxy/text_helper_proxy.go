package proxy

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/jhgv/gocodes/middleware/core/requestor"
)

type TextHelperProxy struct {
	ClientProxy
}

func (tp *TextHelperProxy) UpperText(message string) string {
	return sendInvocation(tp, nameOf((*TextHelperProxy).UpperText), message)
}

func (tp *TextHelperProxy) LowerText(message string) string {
	return sendInvocation(tp, nameOf((*TextHelperProxy).LowerText), message)
}

func (tp *TextHelperProxy) InvertText(message string) string {
	return sendInvocation(tp, nameOf((*TextHelperProxy).InvertText), message)
}

func sendInvocation(tp *TextHelperProxy, methodName string, message string) string {
	var req = &requestor.Requestor{}
	var inv = &requestor.Invocation{}
	var ter *requestor.Termination
	params := []string{message}

	inv.SetHost(tp.GetHost())
	inv.SetPort(tp.GetPort())
	inv.SetParameters(params)
	inv.SetOperationName(methodName)
	inv.SetObjectID(tp.GetID())

	ter, _ = req.Invoke(inv)

	return ter.GetResult()

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
