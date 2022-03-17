package list

type List[T any] interface {
	Append(value T)
	Prepend(value T)
	Delete(value T) bool
	Exist(value T) bool
	Length() int
}
