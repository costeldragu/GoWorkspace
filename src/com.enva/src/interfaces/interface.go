package interfaces

type Renamable interface {
	Rename(newName string)
	Update(newName string)
}

func RenameToFrog(r Renamable)  {
	r.Rename("Frog")
}