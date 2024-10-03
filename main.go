package main

import "fmt"

func main(){
	start();
}

func start() {
	var option string;
	running := true

	for (running){
		printMenu();
		fmt.Scanln(&option);
		switch(option) {
			//Sum
			case "1":
				fmt.Println("O resultado é: ", sum(readNumbers()));
				break;
			//Subtract
			case "2":
				fmt.Println("O resultado é: ", subtract(readNumbers()));
				break;
			//Multiply
			case "3":
				fmt.Println("O resultado é: ", multiply(readNumbers()));
				break;
			//Divide
			case "4":
				fmt.Println("O resultado é: ", divide(readNumbers()));
				break;
			//Exit
			case "5":
				fmt.Println("Programa finalizado.");
				running = false;
				break;
			default:
				fmt.Println("Opção inválida, escolha um dos números válidos");
		}
		fmt.Println("Aperte [Enter] para continuar.")
		fmt.Scanln();
	}
}

func printMenu(){
	fmt.Println("====== CALCULADORA ======");
	fmt.Println("[1] Soma");
	fmt.Println("[2] Subtração");
	fmt.Println("[3] Multiplicação");
	fmt.Println("[4] Divisão");
	fmt.Println("[5] Sair");
}

func readNumbers() []float32 {
	var n1 float32;
	var n2 float32;

	fmt.Println("Digite o primeiro número:");
	fmt.Scanln(&n1);
	fmt.Println("Digite o segundo número:");
	fmt.Scanln(&n2);

	result := []float32{n1, n2}
	return result;
}

func sum(n []float32) float32 {
	return (n[0] + n[1])
}

func subtract(n []float32) float32 {
	return (n[0] - n[1])
}

func multiply(n []float32) float32 {
	return (n[0] * n[1])
}

func divide(n []float32) float32 {
	return (n[0] / n[1])
}
