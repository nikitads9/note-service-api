package utils

import "errors"

func SplitSlice(numbers []int64, batchSize int64) ([][]int64, error) {
	var begin, end int64

	if batchSize < 0 || len(numbers) == 0 {
		return nil, errors.New("invalid batch size or empty slice")
	}

	quantity := len(numbers) / int(batchSize)
	if len(numbers)%int(batchSize) != 0 {
		quantity += 1
	}
	end = batchSize
	batches := make([][]int64, quantity)

	for i := 0; i < quantity; i++ {
		if end > int64(len(numbers)) {
			batches[i] = numbers[begin:]
			break
		}
		batches[i] = numbers[begin:end]
		begin = end
		end += batchSize
	}
	return batches, nil
}

func SwapKeyValue(myMap map[int64]string) (map[string]int64, error) {
	if len(myMap) == 0 {
		return nil, errors.New("null map received")
	}

	resMap := make(map[string]int64, len(myMap))
	for key, value := range myMap {
		resMap[value] = key
	}
	return resMap, nil
}

func SliceFilter(base []int64, filter []int64) ([]int64, error) { //при [1, 2, 3] [1, 2] выведет [3]
	if len(filter) == 0 || len(base) == 0 {
		return nil, errors.New("null slice received")
	}

	fltr, err := SliceToMap(filter)

	if err != nil {
		return nil, errors.New("was unable to convert this slice into map")
	}

	for _, val := range filter {
		fltr[val] = struct{}{}
	}

	for i := 0; i < len(base); i++ {
		if _, found := fltr[base[i]]; found {
			base = append(base[:i], base[i+1:]...)
			i -= 1
		}
	}
	return base, nil
}

func SliceToMap(inSlice []int64) (map[int64]struct{}, error) {
	if len(inSlice) == 0 {
		return nil, errors.New("null slice received")
	}
	outMap := make(map[int64]struct{}, len(inSlice))

	for _, val := range inSlice {
		outMap[val] = struct{}{}
	}
	return outMap, nil
}

//Написать тесты ко всем
