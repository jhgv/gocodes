package requestor

type Termination struct {
	result string
}

func (t *Termination) SetResult(result string) {
	t.result = result
}

func (t *Termination) GetResult() string {
	return t.result
}
