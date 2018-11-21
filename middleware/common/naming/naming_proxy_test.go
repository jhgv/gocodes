package naming

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Data struct {
	Params []interface{}
}

type MyObj struct {
	Name string
	Last MyObj2

}

type MyObj2 struct {
	LastName string
}

func (mo *MyObj) Hello() {
	fmt.Printf("Hello %s", mo.Name)
}

func TestNamingProxy(t *testing.T) {

	myObj := MyObj{Name: "Joao", Last: MyObj2{"Veras"}}

	params := make([]interface{}, 2)
	params[0] = "Teste"
	params[1] = myObj

	data := Data {
		Params: params,
	}

	fmt.Print(data)
	var msgUnmarshalled = &Data{}

	msgMarshalled, _ := json.Marshal(data)

	json.Unmarshal(msgMarshalled, msgUnmarshalled)

	fmt.Print(msgUnmarshalled.Params)


}