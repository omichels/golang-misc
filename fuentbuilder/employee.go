package fuentbuilder

type Employee struct {
	name   string
	works  string
	phone  string
	salary string
}

type handler func(e *Employee)

type EmployeeBuilder struct {
	actions []handler
}

func (e *EmployeeBuilder) Name(n string) *EmployeeBuilder {
	e.actions = append(e.actions, func(e *Employee) {
		e.name = n
	})
	return e
}

func (e *EmployeeBuilder) Works(n string) *EmployeeBuilder {
	e.actions = append(e.actions, func(e *Employee) {
		e.works = n
	})
	return e
}

func (e *EmployeeBuilder) Salary(n string) *EmployeeBuilder {
	e.actions = append(e.actions, func(e *Employee) {
		e.salary = n
	})
	return e
}

func (e *EmployeeBuilder) Phone(n string) *EmployeeBuilder {
	e.actions = append(e.actions, func(e *Employee) {
		e.phone = n
	})
	return e
}

func (e *EmployeeBuilder) Build() Employee {
	emp := Employee{}
	for _, action := range e.actions {
		action(&emp)
	}
	return emp
}

func NewEmployeeBuilder() *EmployeeBuilder {
	return &EmployeeBuilder{}
}
