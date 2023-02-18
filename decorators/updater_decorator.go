// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package decorators

import (
	"context"

	"github.com/NoGambiNoBugs/gowrap-generics-examples/port"
)

// UpdaterTarget is an alias for port.Updater
type UpdaterTarget[I, O any] port.Updater[I, O]

// UpdaterExample implements UpdaterTarget
type UpdaterExample[I, O any] struct {
	base UpdaterTarget[I, O]
}

// Update implements port.Updater
func (d UpdaterExample[I, O]) Update(ctx context.Context, input I) (output O, err error) {
	return d.base.Update(ctx, input)
}

// NewUpdaterExample instruments an implementation of the UpdaterTarget
func NewUpdaterExample[I, O any](base UpdaterTarget[I, O]) UpdaterExample[I, O] {
	return UpdaterExample[I, O]{
		base: base,
	}
}