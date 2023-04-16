package string_utils

import (
	"fmt"
	"math/big"
	"strings"
)

func StringifyArray(arr []*big.Int) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, v.String())
	}
	return strings.Join(strArr, ",")
}

// write a method given a comma seperated value of string numbers, get the big Ints back
func DestringifyArray(arr string) ([]*big.Int, error) {
	var intArr []*big.Int
	for _, v := range strings.Split(arr, ",") {
		i, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse string to big int: %s", v)
		}
		intArr = append(intArr, i)
	}
	return intArr, nil
}
