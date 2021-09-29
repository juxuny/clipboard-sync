package lib

func UniqueID(intSlice []ID) []ID {
	keys := make(map[ID]bool)
	list := []ID{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueUint64(intSlice []uint64) []uint64 {
	keys := make(map[uint64]bool)
	list := []uint64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveUint64(intSlice []uint64, target uint64) []uint64 {
	list := []uint64{}
	for _, v := range intSlice {
		if v != target {
			list = append(list, v)
		}
	}
	return list
}

func RemoveID(intSlice []ID, target ID) []ID {
	list := []ID{}
	for _, v := range intSlice {
		if v != target {
			list = append(list, v)
		}
	}
	return list
}

// 数组分页
func PageUint64(intSlice []uint64, pageNum, pageSize int) []uint64 {
	start := (pageNum - 1) * pageSize
	end := start + pageSize
	if start >= len(intSlice) {
		return []uint64{}
	}
	if end > len(intSlice) {
		end = len(intSlice)
	}
	return intSlice[start:end]
}

// 数组分页
func PageID(intSlice []ID, pageNum, pageSize int) []ID {
	start := (pageNum - 1) * pageSize
	end := start + pageSize
	if start >= len(intSlice) {
		return []ID{}
	}
	if end > len(intSlice) {
		end = len(intSlice)
	}
	return intSlice[start:end]
}
