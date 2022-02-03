package main

import (
	"calculator/cmd/add/numberLoader"
	sum2 "calculator/sum"
	"fmt"
	"io"
	"os"
)

func main(){
	run(os.Args[1:], os.Stdout)
}

func run(args []string, output io.Writer){
	loader := numberLoader.NewNumberLoader(os.DirFS("./"))

	nums, err := loader.Load(args)
	if err != nil {
		fmt.Fprint(output,err)
	}
	sum := sum2.Sum(nums...)

	fmt.Fprint(output, sum)
}
