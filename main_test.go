package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	type useCase struct {
		itemList       map[int]item
		emailList      []string
		expectedOutput map[string]int
	}
	casesTable := []useCase{
		{
			itemList: map[int]item{
				1: {
					name:   "batata",
					amount: 1,
					price:  100,
				},
				2: {
					name:   "tomate",
					amount: 10,
					price:  15,
				},
			},
			emailList: []string{
				"mary@mail.com",
				"john@mail.com",
				"rose@mail.com",
				"sunny@mail.com",
				"julia@mail.com",
				"debra@mail.com",
				"adam@mail.com",
			},
			expectedOutput: map[string]int{"adam@mail.com": 35, "debra@mail.com": 35, "john@mail.com": 36, "julia@mail.com": 36, "mary@mail.com": 36, "rose@mail.com": 36, "sunny@mail.com": 36},
		},
		{
			itemList: map[int]item{
				1: {
					name:   "batata",
					amount: 1,
					price:  100,
				},
				2: {
					name:   "tomate",
					amount: 10,
					price:  15,
				},
			},
			emailList: []string{
				"a@mail.com",
				"b@mail.com",
			},
			expectedOutput: map[string]int{"a@mail.com": 125, "b@mail.com": 125},
		},
	}

	for _, useCase := range casesTable {
		//fmt.Println("INSIDE THE TEST", useCase.itemList)

		resultado, error := DivideTab(useCase.itemList, useCase.emailList)
		if error != nil {
			fmt.Println("Error:", error)
		}
		esperado := useCase.expectedOutput
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Resultado: '%v', esperado: '%v'", resultado, esperado)
		}

	}
}
