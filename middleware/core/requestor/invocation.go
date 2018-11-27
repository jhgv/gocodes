package requestor

type Invocation struct {
	objectID      int
	host          string
	port          int
	operationName string
	EnabledCompression bool
	File    []byte
	Params    []interface{}
}

func (inv *Invocation) SetObjectID(objectID int) {
	inv.objectID = objectID
}

func (inv *Invocation) GetObjectID() int {
	return inv.objectID
}

func (inv *Invocation) SetHost(host string) {
	inv.host = host
}

func (inv *Invocation) GetHost() string {
	return inv.host
}

func (inv *Invocation) SetPort(port int) {
	inv.port = port
}

func (inv *Invocation) GetPort() int {
	return inv.port
}

func (inv *Invocation) SetOperationName(operationName string) {
	inv.operationName = operationName
}

func (inv *Invocation) GetOperationName() string {
	return inv.operationName
}

func (inv *Invocation) SetFile(file []byte) {
	inv.File = file
}

func (inv *Invocation) GetFile() []byte {
	return inv.File
}

func (inv *Invocation) SetParams(params []interface{}) {
	inv.Params = params
}

func (inv *Invocation) GetParams() []interface{} {
	return inv.Params
}
