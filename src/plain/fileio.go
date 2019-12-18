package plain

import (
	"bufio"
	"io"
	"os"
)

func GetFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		testr := []string{""}
		return testr, err
	}
	var result []string
	bf := bufio.NewReader(file)
	for true {
		line, _, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		result = append(result, string(line))
	}
	return result, nil
}
