// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package decorators

import (
	"context"

	"github.com/NoGambiNoBugs/gowrap-generics-examples/entity"
	"github.com/NoGambiNoBugs/gowrap-generics-examples/port"
)

// PartnerCommandTarget is an alias for port.PartnerCommand
type PartnerCommandTarget port.PartnerCommand

// PartnerCommandExample implements PartnerCommandTarget
type PartnerCommandExample struct {
	base PartnerCommandTarget
}

// Delete implements port.PartnerCommand
func (d PartnerCommandExample) Delete(ctx context.Context, input int) (output entity.Partner, err error) {
	return d.base.Delete(ctx, input)
}

// Insert implements port.PartnerCommand
func (d PartnerCommandExample) Insert(ctx context.Context, input entity.Partner) (output entity.Partner, err error) {
	return d.base.Insert(ctx, input)
}

// Update implements port.PartnerCommand
func (d PartnerCommandExample) Update(ctx context.Context, input entity.Partner) (output entity.Partner, err error) {
	return d.base.Update(ctx, input)
}

// NewPartnerCommandExample instruments an implementation of the PartnerCommandTarget
func NewPartnerCommandExample(base PartnerCommandTarget) PartnerCommandExample {
	return PartnerCommandExample{
		base: base,
	}
}