package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int{
	var retArray [][]int
	sort.Ints(candidates)
	var dfs func(preTravel, candidates []int, index, target int)
	dfs = func (preTravel, candidates []int, index, target int){
		for ; index < len(candidates); index++{
			if candidates[index] > target {
				break
			}
			if candidates[index] == target{
				value := append(append([]int{}, preTravel...), candidates[index])
				retArray = append(retArray, value)
				break
			}

			preTravel = append(preTravel, candidates[index])
			dfs(preTravel, candidates, index, target - candidates[index])
			preTravel = preTravel[0:len(preTravel)-1]
		}

		return
	}
	dfs([]int{}, candidates, 0, target)

	return retArray
}

func main(){
	//candidates := []int{2,3,6,7}
	//target := 7

	//candidates := []int{2,7,6,3,5,1}
	//target := 9
	candidates := []int{2,3,5}
	target := 8
	fmt.Printf("candidates:%v, target:%v, result:%v\n", candidates, target,
		combinationSum(candidates, target))

}



