package models

import "fmt"

// Conta : Objeto que representa uma conta
type Conta struct {
	saldo float64
}

// Creditar : recebe o valor a ser creditado na conta e atualiza o saldo
func (c *Conta) Creditar(valor float64) {
	c.saldo += valor
}

// Debitar : recebe o valor a ser debitado na conta e atualiza o saldo
func (c *Conta) Debitar(valor float64) {
	c.saldo -= valor
}

func (c* Conta) Transferir(valor float64, conta *Conta) {
	c.Debitar(valor)
	conta.Creditar(valor)
}

// GetSaldo : Retorna o saldo atual da conta
func (c *Conta) GetSaldo() float64 {
	return c.saldo
}

// PrintStatus : Imprime na tela o status atual da conta
func (c* Conta) PrintStatus() {
	fmt.Printf("Saldo atual da conta: %.2f", c.GetSaldo() )
}
