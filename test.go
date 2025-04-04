package gotest

import "github.com/z46-dev/go-logger"

type Test struct {
	Name, Description string
	Caller            func(log *logger.Logger) (passFail bool)
	Blocking          bool
}

func NewTest(name string, caller func(log *logger.Logger) (passFail bool)) *Test {
	return &Test{
		Name:   name,
		Caller: caller,
	}
}

func (t *Test) SetMeta(description string, blocking bool) *Test {
	t.Description = description
	t.Blocking = blocking
	return t
}

type TestGroup struct {
	Name, Description string
	Tests             []*Test
	Blocking          bool
}

func NewTestGroup(name string, initialTests ...*Test) *TestGroup {
	return &TestGroup{
		Name: name,
		Tests: initialTests,
	}
}

func (t *TestGroup) SetMeta(description string, blocking bool) *TestGroup {
	t.Description = description
	t.Blocking = blocking
	return t
}


func (t *TestGroup) AddTests(tests ...*Test) *TestGroup {
	t.Tests = append(t.Tests, tests...)
	return t
}

type TestSuite struct{}
