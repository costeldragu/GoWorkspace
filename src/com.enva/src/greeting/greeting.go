package greeting

import "fmt"

type Printer func(string)

type Salutation struct {
	Name     string;
	Greeting string;
}

func Greet(salutation Salutation, do Printer, isFomal bool) {
	message, alternative := CreateMessage("Costel", "Salut")
	if prefix := GetPrefix(salutation.Name); isFomal {
		do(prefix + message)
	} else {
		do(alternative)
	}
}
func Print(s string) {
	fmt.Println(s)
}

func CreateMessage(name, greetig string) (message string, alternate string) {
	message = greetig + " " + name
	alternate = "Hey" + " " + name
	return
}

func GetPrefix(name string) (prefix string) {
	switch name {
	case "Bob":
	fallthrough
	case "Joe" ,"Costel":
		prefix = "Dr "
	case "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "

	}
	return
}

func GetPrefixSecond(name string) (prefix string)  {
	switch  {
	case name == "Bob":
		prefix = "Hello bob"
	default:
		prefix = "unknown"
	}

	return
}

/**
 * switch between types
 */
func TestSwitchType(x interface{})  {
	switch x.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("unknown")
	}

}