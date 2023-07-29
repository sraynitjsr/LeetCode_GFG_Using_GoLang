package blind75

impotr "fmt"

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

func (dataSet *DataSet) checkFrequency() bool {
	for _, v := range dataSet.data {
		if v != 1 {
			return true
		}
	}
	return false
}

func ContainsDuplicate() bool {
  mySlice := []int{1, 2, 3, 2} 
	tempMap := make(map[int]int)
	hs := &DataSet{
		data: tempMap,
	}
	for _, num := range mySlice {
		hs.add(num)
	}
	return hs.checkFrequency()
}
