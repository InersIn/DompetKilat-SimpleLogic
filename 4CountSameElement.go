package main

type CountElement struct {
	Count int
	Char  string
}

func CountSameElements(arr []string) []CountElement {
	var result []CountElement
	grouped := ArrayGroup(arr)

	for _, element := range grouped {
		result = append(result, CountElement{Count: len(element), Char: element[0]})
	}

	return []CountElement{}
}
