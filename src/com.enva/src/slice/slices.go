package slice

import "fmt"

func CreateSlice() {

	 s := []int{
		0,1,2,3,
	}

    a := []int {
    	5,6,7,8,9,
	}

	s = append(s,a...)

	for _ ,entry := range s {
		fmt.Println(entry)
	}


}
