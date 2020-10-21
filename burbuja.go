package main

import "fmt"

func Burbuja(lista []int64) {
	var aux int64
	for i := range lista{
		for j := range lista {
			if lista[i] < lista[j] {
				aux = lista[i]
				lista[i] = lista[j]
				lista[j] = aux
			}
		}
	}
	
}

func main(){
	var n int
	var numeros int64
	var s []int64 
	fmt.Scan(&n)
	for i :=0; i<n; i++  {
		fmt.Scan(&numeros)
		s = append(s, numeros)
		
	}

}
