package modal

type Student struct {
	Id    []byte
	Name  string
	Classes [][]byte
}

func (Student) TableName() string {
	return "students"
}
