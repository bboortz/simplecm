package main

import (
	"strconv"
)

func convertStringArrayToIntArray(arr []string) []int {
    var result = []int{}

    for _, i := range arr {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        result = append(result, j)
    }

	return result

}
