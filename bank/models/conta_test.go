package models

import (
	"testing"
	)

func TestDebitar(t *testing.T) {
	conta := Conta{}
	conta.Creditar(100)

	conta2 := Conta{}

	conta.Transferir(45, &conta2)

	expectedSaldoConta := 55.0
	expectedSaldoConta2 := 45.0

	if conta.GetSaldo() != expectedSaldoConta {
		t.Errorf("Expected: %f | Actual: %f", expectedSaldoConta, conta.GetSaldo())
	}

	if conta2.GetSaldo() != expectedSaldoConta2 {
		t.Errorf("Expected: %f | Actual: %f", expectedSaldoConta2, conta2.GetSaldo())
	}

}