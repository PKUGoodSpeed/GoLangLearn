package main

import "fmt"

func main(){
	for i:=0; i<10000000000; i += 1 {
		if(i%500000000 == 0){
			fmt.Printf("Hello world %d \n", i)
		}
	}
}
