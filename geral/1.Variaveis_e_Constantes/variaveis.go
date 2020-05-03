package main
import "fmt"

func main() {
// As variáveis podem ser declaradas de várias formas
//1. "var NOME DataType"

	var v1 int = 3
	fmt.Println(v1)

//2. "NOME := VALOR"

	v2 := ":)"
	fmt.Println(v2)

//3. 
//"var(
// 	NOME = VALOR
//)"

	var(
		v3 = "Hello"
		v4 = 2020
	)
	fmt.Println(v3,v4)

}
