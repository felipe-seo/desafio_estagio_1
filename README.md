# desafio_estagio_1
Desafio 1 da trilha de estágio em back end go da Stone

Neste desafio as entradas são uma lista de items e uma lista de e-mails.

Criei um tipo item que tem os campos name, amount e price.

Criei as listas de items e de e-mails.

Passei as listas como argumentos para a função divideTab que calcula a conta de cada cliente.

A função divideTab recebe como parâmetros:
  um map onde as chaves são inteiros e os valores são items(lista de items);
  um slice de string(lista de emails)

Dentro da função divideTab verifico se as listas estão vazias

Se elas passarem pela verificação:
  calculo o valor total da conta, somando os resultados da multiplicação do preço pela quantidade;
  divido o total pelo número de e-mails;
  caso haja sobra, por exemplo 4 pessoas e uma conta de 43 reais, divido os 3 reais que sobraram entre os e-mails;
  retorno um map contendo as informações.


Referência:
```
Imagine uma lista de compras. Ela possui:

Itens
Quantidade de cada item
Preço por unidade/peso/pacote de cada item
Desenvolva uma função (ou método) que irá receber uma lista de compras (como a detalhada acima) e uma lista de e-mails. Aqui, cada e-mail representa uma pessoa.

A função deve:

Calcular a soma dos valores, ou seja, multiplicar o preço de cada item por sua quantidade e somar todos os itens
Dividir o valor de forma igual entre a quantidade de e-mails
Retornar um mapa/dicionário onde a chave será o e-mail e o valor será quanto ele deve pagar nessa conta
⚠️ IMPORTANTE

Quando fizer a divisão, é importante que não falte nenhum centavo! Cuidado quando tiver, por exemplo, um valor total de R$ 1,00 para ser dividido entre 3 e-mails. Isso daria uma dízima infinita. No entanto, o correto aqui é que duas pessoas fiquem com o valor 0,33 e a terceira fique com 0,34.
Para facilitar e como boa prática, não trabalhe com valores com decimais. Ou seja, ponto flutuante ou float. Use inteiros para representar os valores! Como a menor unidade na nossa moeda é o centavo, use valores inteiros em centavos. Assim, um real será representado por 100 (cem centavos é igual a um real).
```
