package main

import "time"

func main() {
	go helloWorld() //making a go routine

	//Required to wait for goroutine to complete execution
	time.Sleep(time.Second) //This is used to create a time delay to wait for the goroutine to complete execution.
}

func helloWorld() {
	println("Hello World")
}

/*Launching anonymous functions as Goroutines*/ //Uncomment the below code and try executing it
/*
func main(){
	go func(){
		println("Hello World")
	}()
	time.Sleep(time.Second)
}
*/

/*Passing data to anonymous function*/
/*
func main() {
	go func (msg string){
		println(msg)
	}("Hello World")
	time.Sleep(time.Second)
}
*/

/*Defining a afunction within the scope of main function and using it to print messages concurrently*/
/*
func main(){
	messagePrinter := func(msg string){
		println(msg)
	}
	go messagePrinter("Hello World")
	go messagePrinter("Hello Goroutine")
	time.Sleep(time.Second)
}
*/
