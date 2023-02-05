package main

import "fmt"

type DataSet struct {
	data map[int]int
}

func (dataSet *DataSet) add(key int) {
	value, valPresent := dataSet.data[key]
	if !valPresent {
		dataSet.data[key] = 1
	} else {
		dataSet.data[key] = value + 1
	}
}

func (dataSet *DataSet) checkFrequency() {
	for _, v := range dataSet.data {
		if v != 1 {
			fmt.Println("Contains Duplicate")
			return
		}
	}
	fmt.Println("No Duplicates Found")
}

func main() {
	mySlice := []int{1, 1}
	tempMap := make(map[int]int)
	hs := &DataSet{
		data: tempMap,
	}
	for _, num := range mySlice {
		hs.add(num)
	}
	hs.checkFrequency()
}
