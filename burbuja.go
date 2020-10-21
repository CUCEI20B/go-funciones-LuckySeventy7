package main

import "fmt"

func Burbuja(listaDeNumeros []int64){
	var aux int64

	for i := 0; i < len(listaDeNumeros); i++ {
		for j := 0; j < len(listaDeNumeros); j++ {
			if listaDeNumeros[i] < listaDeNumeros[j] {
				aux = listaDeNumeros[i]
				listaDeNumeros[i] = listaDeNumeros[j]
				listaDeNumeros[j] = aux
			}
		}
	}
	//return listaDeNumeros
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

	//fmt.Println(burbuja(s)) 
}
