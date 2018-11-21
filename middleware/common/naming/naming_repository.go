package naming

import (
	"fmt"
	"github.com/jhgv/gocodes/middleware/core/proxy"
)

type NamingRepository struct {
	namingRecords []NamingRecord
}

type namingRecordError struct {
	arg  string
	message string
}

func (e *namingRecordError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.message)
}

func (nr *NamingRepository) addRecord(serviceName string, clientProxy proxy.FileConverterProxy) {
	namingRecord := &NamingRecord{}
	namingRecord.SetServiceName(serviceName)
	namingRecord.SetClientProxy(clientProxy)
	nr.namingRecords = append(nr.namingRecords, *namingRecord)
}

func (nr *NamingRepository) getClientProxyByServiceName(serviceName string) (proxy.FileConverterProxy, error) {
	for _, record := range nr.namingRecords {
		if record.serviceName == serviceName {
			return record.clientProxy, nil
		}
	}
	var f proxy.FileConverterProxy
	return f, &namingRecordError{serviceName, "Can't find service with this name!"}
}
