package main

import (
	"fmt"

	"github.com/yeochinyi/tut-anatomy-mod/calc"

	calcNew "github.com/yeochinyi/tut-anatomy-mod/v2/calc"
)

func main() {
	results := calc.Add(1, 2)
	fmt.Println(results)

	fmt.Println(calcNew.Add(1, 2))
}
