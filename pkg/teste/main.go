package main

import "fmt"

func main() {
	event := []string{"teste", "teste2", "teste3", "teste4"}
	event = append(event[:0], event[1:]...)
	fmt.Println(event)

}
