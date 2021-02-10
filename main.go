package main

import (
	"fmt"
	"github.com/epsilon-akshay/pkg_example/internal/random"
)
import "github.com/epsilon-akshay/pkg_example/pkg/string"

func main() {
	fmt.Print("hi")
	fmt.Print(string.Hash())
	fmt.Print(random.Random())
}