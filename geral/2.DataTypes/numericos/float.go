package main
import "fmt"
// float64 é designado double em C/C++/Java
// float32 normalmente é designado apenas por float
// Em Go, exite float32 e float64.

// float32 -> ocupa 32 bits de memória
// float64 -> ocupa 64 bits de memória

func main(){
	var valor_flutuante float64 = 1.2
	nome := "Exemplo de Float:"
	fmt.Println(nome,valor_flutuante)
}
