package types

type Page[T any] struct {
	Items      *[]T
	TotalItems int
	TotalPages int
	Page       int
	PageSize   int
}
