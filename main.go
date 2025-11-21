//codigo feito em 21/11/2025,  por isto, os valores de dolar e euro podem estar errados

package main

import "fmt"

//aqui começa praticamente o programa para o user
func programa() {
	euro := 6.22
	dolar := 5.41

	var escolha int
	var quantidade float64
	var resultado float64

	for {
		fmt.Println("=== ESCOLHA ===")
		fmt.Println("1 - dolar")
		fmt.Println("2 - euro")
		fmt.Println("0  - sair")

		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			fmt.Println("quantos dolares quer converter pra real: ")
			fmt.Scan(&quantidade)
			resultado = quantidade * dolar
			fmt.Println("lembre-se: 1 dolar custa 5.41")
			fmt.Printf("resultado: R$%.2f\n", resultado)
		case 2:
			fmt.Println("quantos euros quer converter pra real: ")
			fmt.Scan(&quantidade)
			resultado = quantidade * euro
			fmt.Println("lembre-se: 1 euro custa 6.22")
			fmt.Printf("resultado: R$%.2f\n", resultado)
		case 0:
			fmt.Println("saindo...")
			return
		default:
			fmt.Println("erro: opcao invalidal")
		}
	}
}

func main() {

	var calcular string
	fmt.Print("deseja calcular(S/N): ")
	fmt.Scanln(&calcular)

	if calcular == "S" || calcular == "s" {
		programa()
	} else {
		fmt.Println("saindo...")
		return
	}
}
