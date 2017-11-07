package loops

import "fmt"

func GetRangesArray() {

}

func GetRangeSlice() {
	slice := []string {"one","two","three"}
	for index , input := range slice {
		fmt.Print(index)
		fmt.Println( "=>" + input);
	}

}

func GetRangeMap() {

}
