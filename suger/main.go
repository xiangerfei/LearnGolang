package main

import "fmt"

func main() {
	suger("a", 1, 3, 4, 6)
	fmt.Println("vim-go")
}


func suger(values string, values2... int){
	for k, v := range values2{
		fmt.Println("--->:", k, v)
	}
}

