/* WaitGroups comes in the synchronization package to help us synchronize multiple Goroutines. We have to wait for one Goroutine to finish,
we add 1 to the group, and once all of them are added, we ask the group to wait. When the Goroutine finishes, it says done and the
WaitGroup will take one from the group. */
package main

import (
	"fmt"
	"sync"
)

func funcname() { //change funcname to main when executing
	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		fmt.Println("Hello World")
		wait.Done()
	}()
	wait.Wait()

}

/*
func main(){
	var wait sync.WaitGroup

	goRoutines := 5
	wait.Add(goRoutines)

	for i:= 0;i<goRoutines;i++{
		gofunc(goRoutineID int) {
			fmt.Printf("ID: %d :Hello from goroutines!\n",goRoutineID)
			wait.Done()
		}(i)
	}
	wait.Wait()
}
*/
