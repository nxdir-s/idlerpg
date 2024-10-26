package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "Hello from events consumer\n")
}
