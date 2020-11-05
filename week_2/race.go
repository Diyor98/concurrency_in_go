package main

import (
	"fmt"
)

func plusTwo(x *int) {
	*x += 2
}
func plusOne(x *int) {
	*x++
}

func main() {

	var myint int = 1

	for i := 0; i < 10; i++ {
		go plusTwo(&myint)
		fmt.Println(myint)
		go plusOne(&myint)
		fmt.Println(myint)
	}
}

/*
Race condition is a situation when a program runs two or more threads
and the result of the program depends on the order of execution of these
threads as the program does not provide any mechanism to control the order.
In short, every time the program is run it may output a different result.

These problem may occur when two or more threads are given access to a shared
part of memory (for example, a variable) without any synchronization, and each
thread tries to change the data contained in this part of memory in accordance
with its instructions. Since the program was not given any syncronization mechanism
and the processor schedules processes inside it depending on their number and priority,
the threads may be executed in different order, and, as result, we may get different
output each time the program is run.

*/
