package main

import "fmt"

//Neste caso, o For funciona também como o conhecido "While", já que apenas é concedida uma condição e ENQUANTO aquilo for verdade, a execução será realizada!

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//Logo, se você omiir a condição do laço, ele se tornara INFINITO.
