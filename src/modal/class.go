package modal

type Class struct {
	Id    []byte
	Name  string
	Students [][]byte
}

func (Class) TableName() string {
	return "classes"
}