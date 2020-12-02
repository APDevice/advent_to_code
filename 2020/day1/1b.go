package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer elapsed("Search")() //checks how long the entire process runs and prints out result to console

	nums := make([]int, 0)
	set := make(map[int]struct{}) //stores third value
	var void struct{}
	var result, rem int

	//load variables from text file
	file, err := ioutil.ReadFile("nums.txt")
	if err == nil {
		scanner := bufio.NewScanner(strings.NewReader(string(file)))
		for scanner.Scan() {
			val, _ := strconv.Atoi(scanner.Text())
			nums = append(nums, val)

		}
	} else {
		fmt.Println("File not found")
	}

out:
	for i := range nums {
		for j := range nums {
			if nums[i]+nums[j] >= 2020 {
				continue
			}
			rem = 2020 - nums[i] - nums[j]
			if _, ok := set[rem]; ok {
				result = rem * nums[i] * nums[j]
				fmt.Println(result)
				break out
			} else {
				set[nums[i]+nums[j]] = void
			}

		}
	}

}
func elapsed(event string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", event, time.Since(start))
	}
}
