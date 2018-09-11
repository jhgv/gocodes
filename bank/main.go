package main

import (
	"github.com/jhgv/gocodes/bank/models"
)

func main() {
	conta := models.Conta{}
	conta.Creditar(100.0)
	conta.PrintStatus()
}
