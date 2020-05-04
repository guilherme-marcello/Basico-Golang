package main
import "fmt"
// float64 é designado por double em C/C++/Java
// float32 normalmente é designado apenas por float ou single
// Em Go, exite float32 e float64.

// float32 -> ocupa 32 bits de memória; é mais rápido
// float64 -> ocupa 64 bits de memória; é mais preciso

func main(){
	var valor_flutuante float64 = 1.2
	nome := "Exemplo de Float:"
	fmt.Println(nome,valor_flutuante)
}

// Ler: https://en.wikipedia.org/wiki/IEEE_754#Basic_and_interchange_formats
