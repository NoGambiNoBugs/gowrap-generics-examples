// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package decorators

import (
	"context"

	"github.com/NoGambiNoBugs/gowrap-generics-examples/entity"
	"github.com/NoGambiNoBugs/gowrap-generics-examples/port"
)

// CustomerQueryTarget is an alias for port.CustomerQuery
type CustomerQueryTarget port.CustomerQuery

// CustomerQueryExample implements CustomerQueryTarget
type CustomerQueryExample struct {
	base CustomerQueryTarget
}

// Get implements port.CustomerQuery
func (d CustomerQueryExample) Get(ctx context.Context, input int) (output entity.Customer, err error) {
	return d.base.Get(ctx, input)
}

// GetByEmail implements port.CustomerQuery
func (d CustomerQueryExample) GetByEmail(ctx context.Context, email string) (output entity.Customer, err error) {
	return d.base.GetByEmail(ctx, email)
}

// NewCustomerQueryExample instruments an implementation of the CustomerQueryTarget
func NewCustomerQueryExample(base CustomerQueryTarget) CustomerQueryExample {
	return CustomerQueryExample{
		base: base,
	}
}