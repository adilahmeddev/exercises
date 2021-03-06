package main

import (
	"calculator/cmd/add/numberLoader"
	"calculator/formatter"
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
		return
	}
	sum := sum2.Sum(nums...)

	formatter := formatter.Formatter{}
	fmt.Fprint(output, formatter.FormatNumber(sum))
}
