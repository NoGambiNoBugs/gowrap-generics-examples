package entity

type Order[T any] struct {
	Id     int
	Number string
	Target T
}
