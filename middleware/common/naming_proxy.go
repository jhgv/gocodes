package common

import "github.com/jhgv/gocodes/middleware/core/proxy"

// RemoteObject : TODO: make object generic for any object
type RemoteObject struct {
	serviceName string
	object      *proxy.TextHelperProxy
}

type NamingProxy struct {
	Host          string
	Port          int
	remoteObjects []RemoteObject
}

func (np *NamingProxy) Bind(serviceName string, object *proxy.TextHelperProxy) {
	// TODO: avoid registering duplicated objects
	object.SetHost(np.Host)
	object.SetPort(np.Port)
	remoteObj := RemoteObject{serviceName: serviceName, object: object}
	updatedObjectBind := append(np.remoteObjects, remoteObj)
	np.remoteObjects = updatedObjectBind
}

func (np *NamingProxy) Lookup(serviceName string) (*proxy.TextHelperProxy, error) {
	for i := 0; i < len(np.remoteObjects); i++ {
		if np.remoteObjects[i].serviceName == serviceName {
			return np.remoteObjects[i].object, nil
		}
	}
	// TODO: return error if object is not found
	return nil, nil
}
