package main

import "fmt"

//Diferencas entre Map e Struct - fazer nota
func main() {

	//Tipos de declarações

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	printMap(colors)

	//var nomeSobrenome map[string]string
	//fmt.Println(nomeSobrenome)

	//cores := make(map[int]string)
	//cores[0] = "#ffffff"
	//fmt.Println(cores)

	////Removendo
	//delete(cores, 0)
	//fmt.Println(cores)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
