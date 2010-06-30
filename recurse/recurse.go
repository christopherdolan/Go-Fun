package main

import "fmt"

func main() {
	foo(10);
}

func foo(count int) {
	if (count > 0) {
		fmt.Printf("%d\n", count);
		foo(count - 1);
	}
}
