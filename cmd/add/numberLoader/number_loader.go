package numberLoader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strconv"
	"strings"
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
		tee := bytes.Buffer{}
		reader := io.TeeReader(file, &tee)
		fileInfo, _ := file.Stat()

		output := make([]byte, fileInfo.Size())
		_, err = reader.Read(output)
		if err != nil {
			return err
		}
		isCSV := n.isCSV(tee.Bytes())
		if isCSV {
			err = n.readFromCSV(output)
			if err != nil {
				return err
			}
		} else {
			err = n.readFromTxt(output)
			if err != nil {
				return err
			}
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
	fileinfo, _ := file.Stat()
	fileContents := make([]byte, fileinfo.Size())
	_, err = file.Read(fileContents)
	if err != nil {
		return err
	}
	err = n.readFromTxt(fileContents)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (n* NumberLoader) readFromTxt(output []byte) error {
	reader := bytes.NewReader(output)
	scanner := bufio.NewScanner(reader)
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

func (n *NumberLoader) isCSV(content []byte) bool{

	for _, r := range content {
		if string(r) == ","{
			return true
		}
	}
	return false
}

func (n *NumberLoader) readFromCSV(content []byte) error {
	chars := strings.Split(string(content), ",")
	for _, char := range chars {
		err := n.extractNumber(char)
		if err != nil {
			return err
		}
	}
	return nil
}