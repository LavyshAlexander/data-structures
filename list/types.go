package list

type List interface {
	Append(value int)
	Prepend(value int)
	Delete(value int) bool
	Exist(value int) bool
	Length() int
}
