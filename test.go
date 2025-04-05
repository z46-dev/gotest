package gotest

import (
	"fmt"

	"github.com/z46-dev/go-logger"
)

type Test struct {
	Name, Description string
	Caller            func(log *logger.Logger) (passFail bool)
	Blocking, passed  bool
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
	Blocking, passed  bool
}

func NewTestGroup(name string, initialTests ...*Test) *TestGroup {
	return &TestGroup{
		Name:  name,
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

type TestSuite struct {
	Name, Description string
	TestGroups        []*TestGroup
	Passed            bool
	Pass, Fail        int
}

func NewTestSuite(name string, initialGroups ...*TestGroup) *TestSuite {
	return &TestSuite{
		Name:       name,
		TestGroups: initialGroups,
	}
}

func (t *TestSuite) SetMeta(description string) *TestSuite {
	t.Description = description
	return t
}

func (t *TestSuite) AddGroups(groups ...*TestGroup) *TestSuite {
	t.TestGroups = append(t.TestGroups, groups...)
	return t
}

func (t *TestSuite) Run() {
	t.Passed = true
	var log *logger.Logger = logger.NewLogger().SetPrefix(fmt.Sprintf("[%s]", t.Name), logger.Bold)

	log.Statusf("Running test suite: %s\n", t.Name)
	if t.Description != "" {
		log.Basicf("Description: %s\n", t.Description)
	}

	for i, group := range t.TestGroups {
		group.passed = true

		log.Statusf("Running test group %d: %s\n", i+1, group.Name)
		if group.Description != "" {
			log.Basicf("Description: %s\n", group.Description)
		}

		for j, test := range group.Tests {
			log.Statusf("Running test %d: %s\n", j+1, test.Name)
			if test.Description != "" {
				log.Basicf("Description: %s\n", test.Description)
			}

			test.passed = test.Caller(log)

			if test.passed {
				log.Successf("Test %s passed.\n", test.Name)
				t.Pass++
			} else {
				log.Errorf("Test %s failed.\n", test.Name)
				t.Fail++
			}

			if !test.passed && test.Blocking {
				log.Errorf("Test %s failed and is blocking. Stopping test group %s.\n", test.Name, group.Name)
				group.passed = false
				break
			}
		}

		if group.passed {
			log.Successf("Test group %s completed.\n", group.Name)
		} else {
			log.Errorf("Test group %s failed.\n", group.Name)
			t.Passed = false
		}

		if group.Blocking && !group.passed {
			log.Errorf("Test group %s is blocking. Stopping test suite %s.\n", group.Name, t.Name)
			break
		}
	}

	if t.Passed {
		log.Successf("Test suite %s passed.\n", t.Name)
	} else {
		log.Errorf("Test suite %s failed.\n", t.Name)
	}

	log.Basicf("Test suite %s results: %d passed, %d failed.\n", t.Name, t.Pass, t.Fail)
}
