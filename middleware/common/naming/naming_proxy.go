package naming

import (
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/core/requestor"
	"reflect"
	"runtime"
	"strings"
)

type NamingProxy struct {
	proxy.ClientProxy
}

func NewNamingProxy(host string, port int) *NamingProxy {
	np := new(NamingProxy)
	np.SetHost(host)
	np.SetPort(port)
	return np
}

func (np *NamingProxy) Bind(serviceName string, clientProxy proxy.FileConverterProxy) {
	var req = &requestor.Requestor{}
	var inv = &requestor.Invocation{}

	inv.SetHost(np.GetHost())
	inv.SetPort(np.GetPort())
	inv.SetOperationName(nameOf((*NamingProxy).Bind))
	params := make([]interface{}, 2)
	params[0] = serviceName
	params[1] = clientProxy
	inv.SetParams(params)

	req.InvokeN(inv)

}

func (np *NamingProxy) Lookup(serviceName string) (proxy.FileConverterProxy, error) {

	var req = &requestor.Requestor{}
	var inv = &requestor.Invocation{}
	var ter *requestor.Termination

	inv.SetHost(np.GetHost())
	inv.SetPort(np.GetPort())
	inv.SetOperationName(nameOf((*NamingProxy).Lookup))
	params := make([]interface{}, 1)
	params[0] = serviceName
	inv.SetParams(params)

	ter, _ = req.InvokeN(inv)

	clientProxyParams := ter.GetResult().(map[string]interface{})
	clientProxy := proxy.FileConverterProxy{}
	clientProxy.SetHost(clientProxyParams["Host"].(string))
	clientProxy.SetPort(int(clientProxyParams["Port"].(float64)))
	clientProxy.SetID(int(clientProxyParams["ID"].(float64)))

	return clientProxy, nil
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
