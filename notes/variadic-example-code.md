```go

package main

import "fmt"


//Exemplo de parâmetros variádicos
func main() {
		fmt.Println(Demo(1,2,3))
		slice_n := []int{3,6,9}
		fmt.Println(Demo(slice_n...))

}


func Demo(numbers ...int) int{
	sum := 0
	for key,value := range numbers {
		fmt.Printf("A chave de valor = %d",key)	
		sum += value
		fmt.Println()
	}
	return sum 

}
```
