package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func importFile(fn string) []*passwordCheck {
	//load variables from text file
	pass := []*passwordCheck{}
	temp := make([]string, 3, 3)
	tempInts := make([]string, 2, 2)
	file, err := ioutil.ReadFile(fn)
	readError := 0

	if err == nil {
		scanner := bufio.NewScanner(
			strings.NewReader(string(file)))
		for scanner.Scan() {
			values := &passwordCheck{}
			temp = strings.Fields(scanner.Text())
			values.password = temp[2]
			values.char = temp[1][0]
			tempInts = strings.Split(temp[0], "-")
			values.min, err = strconv.Atoi(tempInts[0])
			values.max, err = strconv.Atoi(tempInts[1])
			pass = append(pass, values)
			if err != nil {
				readError++
			}
		}
	} else {
		fmt.Println("File not found")
	}
	return pass
}
