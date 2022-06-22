package main

func ArrayGroup(arr []string) [][]string {
	var result [][]string
	var temp string
	var index int

	for _, v := range arr {
		if len(result) == 0 {
			result = append(result, []string{v})
			temp = v
		} else if v != temp {
			result = append(result, []string{v})
			temp = v
			index++
		} else if v == temp {
			result[index] = append(result[index], v)
		}
	}

	return result
}
