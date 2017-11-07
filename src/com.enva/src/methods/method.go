package methods

import "fmt"

type Salutation struct {
	Name string
	Greeting string
}

type Salutations []Salutation

func (salutations Salutations) Greet()  {
	for _,s := range salutations{
		fmt.Println(s.Greeting+" "+s.Name)
	}
}

func (salutation *Salutation) Rename(newName string)  {
	salutation.Name = newName
}

func  (salutation *Salutation) Update(newName string)  {
	salutation.Greeting = newName
}

func (salutation *Salutation) Write (p []byte) (n int, err error)  {
	// transform bytes to string
	s := string(p)
	salutation.Rename(s);
	n = len(s)
	err = nil
	return
}