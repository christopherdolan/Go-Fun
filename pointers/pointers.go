package main

import "fmt"

func main() {
	lstring := "i am a string";
	fmt.Printf("String literal: %s\n", lstring);
	astring := &lstring;
	fmt.Printf("String address: %s\n", astring);
	fmt.Printf("String pointer: %s\n", *astring);
	pstring2 := *"i am a string";
	fmt.Printf("String literal pointer: %s\n", &pstring2);
}
