package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func sort(slice []int, wg *sync.WaitGroup, c chan []int) {
	defer wg.Done()
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}

	c <- slice
}

func main() {

	var wg sync.WaitGroup
	wg.Add(4)
	c := make(chan []int, 4)

	reader := bufio.NewReader(os.Stdin)

	mySlice := make([]int, 0)
	counter := 0

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "quit" && counter < 4 {
			fmt.Println("The number of integers should exceed three")
		} else if text == "quit" && counter >= 4 {
			break
		} else {
			i, _ := strconv.Atoi(text)
			mySlice = append(mySlice, i)
			counter++
		}

	}

	sliceLength := len(mySlice)
	go sort(mySlice[:sliceLength/4], &wg, c)
	go sort(mySlice[sliceLength/4:sliceLength/2], &wg, c)
	go sort(mySlice[sliceLength/2:3*sliceLength/4], &wg, c)
	go sort(mySlice[3*sliceLength/4:], &wg, c)
	wg.Wait()
	mySlice = make([]int, 0)
	mySlice = append(mySlice, <-c...)
	mySlice = append(mySlice, <-c...)
	mySlice = append(mySlice, <-c...)
	mySlice = append(mySlice, <-c...)

	for i := 0; i < len(mySlice); i++ {
		for j := 0; j < len(mySlice)-1; j++ {
			if mySlice[j] > mySlice[j+1] {
				temp := mySlice[j]
				mySlice[j] = mySlice[j+1]
				mySlice[j+1] = temp
			}
		}
	}
	fmt.Println(mySlice)

}
