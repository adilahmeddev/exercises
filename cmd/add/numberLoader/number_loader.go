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
		err := n.extractFromDefault()
		if err != nil {
			return nil, err
		}
	}
	var files []string
	if len(args) > 0 && args[0] == "--input-file" {
		err := n.extractFromInputFile(args, files)
		if err != nil {
			return nil, err
		}
	} else {
		err := n.extractNumberFromArgs(args)
		if err != nil {
			return nil, err
		}
	}
	return n.numbers, nil
}

func (n *NumberLoader) extractNumberFromArgs(args []string) error {
	for _, arg := range args {
		err := n.extractNumber(arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *NumberLoader) extractFromInputFile(args []string, files []string) error {
	if len(args) >= 2 && len(args)%2 == 0 {
		for i := 0; i < len(args); i++ {
			if args[i] == "--input-file" {
				files = append(files, args[i+1])
				i++
			}
		}
	}
	for _, fileName := range files {
		file, err := n.fs.Open(fileName)
		if err != nil {
			return err
		}
		err = n.readFromTxt(file)
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *NumberLoader) extractFromDefault() error {
	file, err := n.fs.Open("input.txt")
	if err != nil {
		return err
	}
	err = n.readFromTxt(file)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (n* NumberLoader) readFromTxt(file fs.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err := n.extractNumber(scanner.Text())
		if err != nil {
			return err
		}
	}
	return nil
}

func (n* NumberLoader) extractNumber(input string) error {
	i, err := strconv.Atoi(input)
	if err == nil {
		if n.history[i] == 0 {
			n.numbers = append(n.numbers, i)
		}
		n.history[i]=1
	} else {
		return fmt.Errorf("invalid input")
	}
	return nil
}