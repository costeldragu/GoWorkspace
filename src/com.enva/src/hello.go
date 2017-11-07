package main

import "fmt"
import "./greeting"
import "./loops"
import (
	"./maps"
	"com.enva/src/slice"
	"com.enva/src/methods"
	"com.enva/src/interfaces"

)

func main() {
	//var message string = "Hello go"
	//var a , b ,c int = 1,2,3
	//
	//var greeting *string = &message

	//fmt.Println(greeting)

	 var s = greeting.Salutation{"Bob", "hello"};

	 message,alernate := greeting.CreateMessage(s.Name, s.Greeting)

	//var message Salutation = "hello"
	//fmt.Println(s.name,s.greeting)
	greeting.Greet(s,greeting.Print,true);
	fmt.Println(greeting.CreateMessage(s.Name, s.Greeting))
	fmt.Println(message)
	fmt.Println(alernate)


    loops.ClasicFor(10)
    loops.WhileLoop(15)
    loops.LoopForever(20)
    loops.LoopWithContinue(20)
    loops.GetRangeSlice()
    maps.RemoveMapValue("one")
    maps.CreateMap();

    slice.CreateSlice();

    salutations := methods.Salutations{
    	{"Peter","Hy"},
    	{"Mary","Hy"},
    	{"Bog","Hy"},
	}

	done := make(chan bool)

	salutations[2].Rename("nemmo")
	interfaces.RenameToFrog(&salutations[1])
	go func() {
		methods.Salutations.Greet(salutations)
		x := 1
		for {
			x++
			if(x > 100000000000) {
				fmt.Println(x)
				break
			}
		}
		done <- true
	}()

   fmt.Fprintf(&salutations[1],"The count is %d",10)
   <- done
   }

