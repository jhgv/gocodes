package requestor

type Termination struct {
	result interface{}
}

func (t *Termination) SetResult(result interface{}) {
	t.result = result
}

func (t *Termination) GetResult() interface{} {
	return t.result
}
