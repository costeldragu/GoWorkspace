package loops

import "fmt"

func ClasicFor(lenght int) {
	for x := 0; x < lenght; x++ {
		fmt.Println(x)
	}
}

func WhileLoop(length int) {
	x := 0
	for x < length {
		fmt.Println(x)
		x++
	}
}

func LoopForever(length int)  {
	x := 0
	for {
		fmt.Println(x)
		if(x > length) {
			break
		}
		x++
	}
}

func LoopWithContinue(length int)  {
	x := 0
	for {
		if(x%2 == 0) {
			x++
			continue
		}
		fmt.Println(x)
		if(x > length) {
			break
		}
		x++
	}
}