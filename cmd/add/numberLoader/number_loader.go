package numberLoader

import (
	"bufio"
	"fmt"
	"io/fs"
	"strconv"
)

type NumberLoader struct {
	numbers []int
	history map[int]int
	fs fs.FS
}

func NewNumberLoader(fs fs.FS) *NumberLoader {
	return &NumberLoader{fs: fs, numbers: []int{}, history: map[int]int{}}
}

func (n* NumberLoader) Load(args []string) ([]int, error){
	if len(args) == 0{
		file, err := n.fs.Open("input.txt")
		if err != nil {
			return nil, err
		}

		n.readFromTxt(file)
		file.Close()
	}
	var files []string
	if len(args) > 0 && args[0] == "--input-file"{
			if len(args) >=2 && len(args)%2==0{
				fmt.Println("true")
				for i := 0; i < len(args); i++ {
					if args[i] == "--input-file" {
						files = append(files,args[i+1])
						i++
					}
				}
			}
			fmt.Println(files)
		for _, fileName := range files {
			file, err := n.fs.Open(fileName)
			if err != nil {
				return nil, err
			}

			n.readFromTxt(file)
			file.Close()
		}
	}
	return n.numbers, nil
}

func (n* NumberLoader) readFromTxt(file fs.File){
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n.extractNumber(scanner.Text())
	}
}

func (n* NumberLoader) extractNumber(input string) {
	i, err := strconv.Atoi(input)
	if err == nil {
		if n.history[i] == 0 {
			n.numbers = append(n.numbers, i)
		}
		n.history[i]=1
	}
}