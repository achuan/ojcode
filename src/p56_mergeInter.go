package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals,
		func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})

	var mergeInter [][]int
	var start, end int
	for i := 0; i < len(intervals); i++{
		if i == 0 {
			start = intervals[i][0]
			end = intervals[i][1]
			continue
		}
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
			continue
		}

		mergeInter = append(mergeInter, []int{start, end})
		start = intervals[i][0]
		end = intervals[i][1]
	}

	mergeInter = append(mergeInter, []int{start, end})

	return mergeInter
}

func main(){
	testArray := [][][]int{
		{{1,3}, {2,6},{8,10},{15,18}},
		{{1,4},{4,5}},
		{{1,4},{2,3}},
	}

	for _, item := range testArray{
		fmt.Printf("test array :%x merge interval is %x\n", item, merge(item))
	}
}


