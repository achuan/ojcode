package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	var retArray [][]int
	sort.Ints(candidates)

	var dfs func(preTravel, candidates []int, index, target int)

	dfs = func(preTravel, candidates []int, begin, target int) {
		for index := begin;  index < len(candidates); index++{
			if candidates[index] > target{
				break
			}
			if index > begin && candidates[index] == candidates[index-1]{
				continue
			}

			if candidates[index] == target{
				value := append(append([]int{}, preTravel...), candidates[index])
				retArray = append(retArray, value)
			}
			preTravel = append(preTravel, candidates[index])
			dfs(preTravel, candidates, index+1, target - candidates[index])
			preTravel = preTravel[0:len(preTravel)-1]
		}
	}
	dfs([]int{}, candidates, 0, target)

	return retArray
}

func main(){
	//candidates := []int{2,3,6,7}
	//target := 7

	candidates := []int{2,5,2,1,2}
	target := 5
	//candidates := []int{10,1,2,7,6,1,5}
	//target := 8
	fmt.Printf("candidates:%v, target:%v, result:%v\n", candidates, target,
		combinationSum2(candidates, target))

}
