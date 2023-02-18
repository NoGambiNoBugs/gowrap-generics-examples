// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package decorators

import (
	"context"

	"github.com/NoGambiNoBugs/gowrap-generics-examples/entity"
	"github.com/NoGambiNoBugs/gowrap-generics-examples/port"
)

// OrderCommandTarget is an alias for port.OrderCommand
type OrderCommandTarget[T any] port.OrderCommand[T]

// OrderCommandExample implements OrderCommandTarget
type OrderCommandExample[T any] struct {
	base OrderCommandTarget[T]
}

// Delete implements port.OrderCommand
func (d OrderCommandExample[T]) Delete(ctx context.Context, input int) (output entity.Order[T], err error) {
	return d.base.Delete(ctx, input)
}

// Insert implements port.OrderCommand
func (d OrderCommandExample[T]) Insert(ctx context.Context, input entity.Order[T]) (output entity.Order[T], err error) {
	return d.base.Insert(ctx, input)
}

// Update implements port.OrderCommand
func (d OrderCommandExample[T]) Update(ctx context.Context, input entity.Order[T]) (output entity.Order[T], err error) {
	return d.base.Update(ctx, input)
}

// NewOrderCommandExample instruments an implementation of the OrderCommandTarget
func NewOrderCommandExample[T any](base OrderCommandTarget[T]) OrderCommandExample[T] {
	return OrderCommandExample[T]{
		base: base,
	}
}
