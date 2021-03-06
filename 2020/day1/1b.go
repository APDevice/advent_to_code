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
		scanner := bufio.NewScanner(
			strings.NewReader(string(file)))
		for scanner.Scan() {
			val, _ := strconv.Atoi(scanner.Text())
			nums = append(nums, val)
			set[val] = void

		}
	} else {
		fmt.Println("File not found")
	}
	//sort.Ints(nums)
	jMax := len(nums) - 1
out:
	for i := range nums {
		for j := jMax; j > -1; j-- {
			if nums[i]+nums[j] >= 2020 {
				jMax--
				continue
			}
			rem = 2020 - nums[i] - nums[j]
			if _, ok := set[rem]; ok {
				result = rem * nums[i] * nums[j]
				fmt.Printf("%v * %v * %v = %v\n", rem, nums[i], nums[j], result)
				break out
			}

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
