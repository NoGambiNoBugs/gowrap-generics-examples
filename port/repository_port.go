package port

import (
	"context"

	"github.com/NoGambiNoBugs/gowrap-generics-examples/entity"
)

////go:generate gowrap gen -g -i Inserter -t ../templates/template.go.tmpl -o ../decorators/inserter_decorator.go
type Inserter[I, O any] interface {
	Insert(ctx context.Context, input I) (output O, err error)
}

////go:generate gowrap gen -g -i Updater -t ../templates/template.go.tmpl -o ../decorators/updater_decorator.go
type Updater[I, O any] interface {
	Update(ctx context.Context, input I) (output O, err error)
}

////go:generate gowrap gen -g -i Getter -t ../templates/template.go.tmpl -o ../decorators/getter_decorator.go
type Getter[I, O any] interface {
	Get(ctx context.Context, input I) (output O, err error)
}

////go:generate gowrap gen -g -i Deleter -t ../templates/template.go.tmpl -o ../decorators/deleter_decorator.go
type Deleter[I, O any] interface {
	Delete(ctx context.Context, input I) (output O, err error)
}

////go:generate gowrap gen -g -i CustomerCommand -t ../templates/template.go.tmpl -o ../decorators/customer_command_decorator.go
type CustomerCommand interface {
	Inserter[entity.Customer, entity.Customer]
	Updater[entity.Customer, entity.Customer]
	Deleter[int, entity.Customer]
}

////go:generate gowrap gen -g -i CustomerQuery -t ../templates/template.go.tmpl -o ../decorators/customer_query_decorator.go
type CustomerQuery interface {
	Getter[int, entity.Customer]
	GetByEmail(ctx context.Context, email string) (output entity.Customer, err error)
}

////go:generate gowrap gen -g -i PartnerCommand -t ../templates/template.go.tmpl -o ../decorators/partner_command_decorator.go
type PartnerCommand interface {
	Inserter[entity.Partner, entity.Partner]
	Updater[entity.Partner, entity.Partner]
	Deleter[int, entity.Partner]
}

////go:generate gowrap gen -g -i PartnerQuery -t ../templates/template.go.tmpl -o ../decorators/partner_query_decorator.go
type PartnerQuery interface {
	Getter[int, entity.Partner]
	GetByReference(ctx context.Context, reference string) (output entity.Partner, err error)
}

////go:generate gowrap gen -g -i OrderCommand -t ../templates/template.go.tmpl -o ../decorators/order_command_decorator.go
type OrderCommand[T any] interface {
	Inserter[entity.Order[T], entity.Order[T]]
	Updater[entity.Order[T], entity.Order[T]]
	Deleter[int, entity.Order[T]]
}

////go:generate gowrap gen -g -i CustomerOrderQuery -t ../templates/template.go.tmpl -o ../decorators/customer_order_query_decorator.go
type CustomerOrderQuery interface {
	Getter[int, entity.Order[entity.Customer]]
}

////go:generate gowrap gen -g -i PartnerOrderQuery -t ../templates/template.go.tmpl -o ../decorators/partner_order_query_decorator.go
type PartnerOrderQuery interface {
	Getter[int, entity.Order[entity.Partner]]
}
