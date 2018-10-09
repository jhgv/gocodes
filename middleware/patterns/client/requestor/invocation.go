package requestor

type Invocation struct {
	objectID      int
	host          string
	port          int
	operationName string
	parameters    []string
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

func (inv *Invocation) SetParameters(parameters []string) {
	inv.parameters = parameters
}

func (inv *Invocation) GetParameters() []string {
	return inv.parameters
}
