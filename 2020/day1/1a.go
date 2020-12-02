package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer elapsed("Search")() //checks how long the entire process runs and prints out result to console

	nums := make([]int, 0)
	set := make(map[int]struct{}) //stores second value
	var void struct{}
	var result, rem int

	//load variables from text file
	file, err := ioutil.ReadFile("nums.txt")

	if err == nil {
		scanner := bufio.NewScanner(
			strings.NewReader(string(file)))
		for scanner.Scan() {
			val, _ := strconv.Atoi(scanner.Text())
			nums = append(nums, val)

		}
	} else {
		fmt.Println("File not found")
	}
	sort.Ints(nums)

out:
	for i := range nums {

		rem = 2020 - nums[i]
		if _, ok := set[rem]; ok {
			result = rem * nums[i]
			fmt.Printf("%v * %v = %v\n", rem, nums[i], result)
			break out
		} else {
			set[rem] = void
		}

	}

}
func elapsed(event string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v\n",
			event,
			time.Since(start))
	}
}
