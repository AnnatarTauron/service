package main

import "fmt"

func main() {
	checker := NewChecker()

	client1 := NewGoMetrClient("1")
	client2 := NewGoMetrClient("2")
	client3 := NewGoMetrClient("3")

	checker.Add(client1)
	checker.Add(client2)
	checker.Add(client3)

	fmt.Println(checker)
}
