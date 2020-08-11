package main

import (
	"fmt"
)

func main() {
	var a, b = 1, 5
	swap(&a, &b)
	fmt.Print(a, b)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
