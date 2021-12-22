package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {

	item1 := item{
		name:   "batata",
		amount: 1,
		price:  100,
	}
	item2 := item{
		name:   "tomate",
		amount: 1,
		price:  109,
	}
	//crio a lista de items
	itemList := map[int]item{
		1: item1,
		2: item2,
	}

	emailList := []string{
		"mary@mail.com",
		"john@mail.com",
		"rose@mail.com",
		"sunny@mail.com",
		"julia@mail.com",
		"debra@mail.com",
		"adam@mail.com",
	}

	t.Run("Test tab divider", func(t *testing.T) {
		resultado, err := DivideTab(itemList, emailList)
		esperado := map[string]int{"adam@mail.com": 29, "debra@mail.com": 30, "john@mail.com": 30, "julia@mail.com": 30, "mary@mail.com": 30, "rose@mail.com": 30, "sunny@mail.com": 30}

		if err != nil {
			t.Errorf("Erro: '%s'", err)
		}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Resultado: '%v', esperado: '%v'", resultado, esperado)
		}

	})

}
