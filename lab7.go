package main

import (
	"fmt"
	"math/rand"
	"time"
)

func exec(max int) int{
	rand.Seed(42)
	v := rand.Intn(max)
	time.Sleep(time.Duration(v) * time.Millisecond)
	return v
}

func aux( max_sleep_ms int ) <- chan int {

	saida := make (chan int)
	for i:=0 ; i<1000; i++{
		go func(){
			saida <- exec(max_sleep_ms)
		}()
	} 
	return saida

}

func main(){

	canal1 := aux(20)
	canal2 := aux(200)

	soma := 0

	for i:= 0; i <500; i++{
		soma = soma + <-canal1
		fmt.Println("canal1", soma)
		soma = soma + <-canal2
		fmt.Println("canal2", soma)
	}
	fmt.Println(soma)
}