package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	type useCase struct {
		name           string
		itemList       map[int]item
		emailList      []string
		expectedOutput map[string]int
	}
	casesTable := []useCase{
		{
			name: "Normal usage, multiple people pay less, multiple people pay more",
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
			name: "Normal usage, everyone pays the same",
			itemList: map[int]item{
				1: {
					name:   "bala",
					amount: 1,
					price:  100,
				},
				2: {
					name:   "alho",
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

		{
			name: "Normal usage, multiple people pay less, one person pays more",
			itemList: map[int]item{
				1: {
					name:   "chiclete",
					amount: 1,
					price:  151,
				},
				2: {
					name:   "alho",
					amount: 10,
					price:  15,
				},
			},
			emailList: []string{
				"a@mail.com",
				"b@mail.com",
				"c@mail.com",
			},
			expectedOutput: map[string]int{"a@mail.com": 101, "b@mail.com": 100, "c@mail.com": 100},
		},

		{
			name: "No price in some item",
			itemList: map[int]item{
				1: {
					name:   "bala",
					amount: 1,
					price:  00,
				},
				2: {
					name:   "alho",
					amount: 10,
					price:  10,
				},
			},
			emailList: []string{
				"a@mail.com",
				"b@mail.com",
			},
			expectedOutput: map[string]int{"Algo de errado com a lista de itens": 1},
		},

		{
			name:     "Empty items list",
			itemList: map[int]item{},
			emailList: []string{
				"a@mail.com",
				"b@mail.com",
			},
			expectedOutput: map[string]int{"Lista de itens vazia": 1},
		},

		{
			name: "Empty e-mail list",
			itemList: map[int]item{
				1: {
					name:   "bala",
					amount: 1,
					price:  10,
				},
				2: {
					name:   "alho",
					amount: 10,
					price:  15,
				},
			},
			emailList:      []string{},
			expectedOutput: map[string]int{"Lista de e-mails vazia": 1},
		},

		{
			name: "One or more blank e-mails",
			itemList: map[int]item{
				1: {
					name:   "bala",
					amount: 1,
					price:  10,
				},
				2: {
					name:   "alho",
					amount: 10,
					price:  15,
				},
			},
			emailList: []string{
				"",
				"b@mail.com",
			},
			expectedOutput: map[string]int{"Algum e-mail em branco na lista de e-mails": 1},
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
