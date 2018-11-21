package security

import (
	"fmt"
	"testing"
)

func TestEncrypter(t *testing.T) {
	e := Encrypter{}
	palavra := "Ola meu nome é joão"
	key := []byte("1234567890123456")


	r := e.Encrypt(key, []byte(palavra))

	resultado := fmt.Sprintf("%s", r)

	if resultado == palavra {
		t.Error("Falha na criptografia")
	}

	dec := e.Decrypt(key, r)
	decResultado := fmt.Sprintf("%s", dec)

	if decResultado != palavra {
		t.Error("Falha na descriptografando")
	}

}


