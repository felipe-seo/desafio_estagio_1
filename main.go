package main

import (
	"fmt"
)

type item struct {
	name   string
	amount int
	price  int
}

func main() {
	//crio os items
	item1 := item{
		name:   "batata",
		amount: 0,
		price:  100,
	}
	item2 := item{
		name:   "tomate",
		amount: 10,
		price:  15,
	}
	//crio a lista de items
	itemList := map[int]item{
		1: item1,
		2: item2,
	}
	//printo a lista
	//fmt.Println("Lista de items: ", itemList)

	//crio a lista de e-mails
	emailList := []string{
		"mary@mail.com",
		"john@mail.com",
		"rose@mail.com",
		"sunny@mail.com",
		"julia@mail.com",
		"debra@mail.com",
		"adam@mail.com",
	}
	//printo a lista
	//fmt.Println("Lista de emails: ", emailList)
	DivideTab(itemList, emailList)

}

//função que calcula a conta, recebe os itens e os clientes
func DivideTab(il map[int]item, el []string) (map[string]int, error) {
	total := 0
	tab := make(map[string]int, len(el))

	if verifyItems(il) != nil {
		return verifyItems(il), nil
	}
	if verifyEmails(el) != nil {
		return verifyEmails(el), nil
	}

	for _, element := range il {
		total += element.price * element.amount
	}

	remainder := int(total) % len(el)

	totalPerAccount := total / len(el)

	//distribuir a sobra, centavo por centavo, até não ter mais sobra
	for _, v := range el {
		if remainder > 0 {
			tab[v] = totalPerAccount + 1
			remainder--
		} else {
			tab[v] = totalPerAccount
		}
	}
	fmt.Println("Conta: ", tab)

	return tab, nil
}

func verifyItems(il map[int]item) map[string]int {
	if len(il) == 0 {
		//fmt.Println("Lista de itens vazia")
		return map[string]int{"Lista de itens vazia": 1}
	}
	for _, v := range il {
		if v.name == "" || v.amount == 0 || v.price == 0 {
			//fmt.Printf("Algo de errado com o item: %v", v)
			return map[string]int{"Algo de errado com a lista de itens": 1}
		}
	}
	return nil
}

func verifyEmails(el []string) map[string]int {
	if len(el) == 0 {
		//fmt.Println("Lista de e-mails vazia")
		return map[string]int{"Lista de e-mails vazia": 1}
	}
	//fazer algo para verificar se tem já tem o email na lista
	for _, v := range el {
		if v == "" {
			//fmt.Printf("Algo de errado com o email: %v", v)
			return map[string]int{"Algum e-mail em branco na lista de e-mails": 1}
		}
	}
	return nil
}

//avaliação
