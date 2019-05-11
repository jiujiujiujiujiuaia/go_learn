package common

type count int

func New(v int) count{
	return count(v)
}
