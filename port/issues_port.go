package port

import "github.com/NoGambiNoBugs/gowrap-generics-examples/entity"

//go:generate gowrap gen -g -i Base -t ../templates/issue_65.go.tmpl -o ../decorators/issue_65.go
type Base interface {
	A()
	Nested
}

type Nested interface {
	B()
}

//go:generate gowrap gen -g -i Test -t ../templates/issue_62.go.tmpl -o ../decorators/issue_62.go
type Test interface {
	MyMethod(a int, b entity.Customer) (entity.Customer, error)
}
