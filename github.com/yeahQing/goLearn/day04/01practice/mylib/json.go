package mylib

//Teacher 教师
type Teacher struct {
	ID     int
	Gender string
	Name   string
}

//Department 部门
type Department struct {
	Title    string
	Teachers []*Teacher
}
