// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package decorators

import (
	"context"

	"github.com/NoGambiNoBugs/gowrap-generics-examples/entity"
	"github.com/NoGambiNoBugs/gowrap-generics-examples/port"
)

// CustomerOrderQueryTarget is an alias for port.CustomerOrderQuery
type CustomerOrderQueryTarget port.CustomerOrderQuery

// CustomerOrderQueryExample implements CustomerOrderQueryTarget
type CustomerOrderQueryExample struct {
	base CustomerOrderQueryTarget
}

// Get implements port.CustomerOrderQuery
func (d CustomerOrderQueryExample) Get(ctx context.Context, input int) (output entity.Order[entity.Customer], err error) {
	return d.base.Get(ctx, input)
}

// NewCustomerOrderQueryExample instruments an implementation of the CustomerOrderQueryTarget
func NewCustomerOrderQueryExample(base CustomerOrderQueryTarget) CustomerOrderQueryExample {
	return CustomerOrderQueryExample{
		base: base,
	}
}