package factory

import (
	"fmt"
	"reflect"
	"testing"
)

type FactorySpecification struct {
	Name string

	factory *factory

	t *testing.T

	given         []interface{}
	commands      []Command
	event         interface{}
	expectedError string
	lastError     error
}

func NewFactorySpecification(name string, t *testing.T) *FactorySpecification {
	return &FactorySpecification{
		Name:     name,
		factory:  FromEvents(make([]interface{}, 0)),
		commands: make([]Command, 0),
		t:        t,
	}
}

func (spec *FactorySpecification) Given(events ...interface{}) {
	spec.given = events
	spec.factory = FromEvents(events)
}

func (spec *FactorySpecification) When(cmd Command) {
	spec.commands = append(spec.commands, cmd)
}

func (spec *FactorySpecification) Then(event interface{}) {
	spec.event = event
}

func (spec *FactorySpecification) ThenError(message string) {
	spec.expectedError = message
}

func (spec *FactorySpecification) executeGiven() {
	fmt.Println("Given")
	for _, event := range spec.given {
		fmt.Println("    ", event)
	}
}

func (spec *FactorySpecification) executeCommands() {
	fmt.Println("When")
	for _, action := range spec.commands {
		spec.lastError = action.Execute(spec.factory)
		fmt.Println("    ", action)
	}
}

func (spec *FactorySpecification) executeExpectations() {
	if spec.event == nil {
		return
	}

	if len(spec.factory.events) == 0 {
		spec.t.Fail()
		return
	}

	lastEvent := spec.factory.events[len(spec.factory.events)-1]

	fmt.Println("Then")
	fmt.Println("    ", spec.event)

	if !reflect.DeepEqual(lastEvent, spec.event) {
		spec.t.Fail()
		return
	}
}

func (spec *FactorySpecification) executeErrorExpectations() {
	if spec.expectedError == "" {
		return
	}

	fmt.Println("Then error")
	fmt.Println("    ", spec.expectedError)

	if spec.lastError == nil {
		spec.t.Fail()
		return
	}

	if spec.expectedError != spec.lastError.Error() {
		spec.t.Fail()
		return
	}
}

func (spec *FactorySpecification) Run(t *testing.T) {
	fmt.Println(spec.Name)

	spec.executeGiven()
	spec.executeCommands()
	spec.executeExpectations()
	spec.executeErrorExpectations()

	fmt.Println("")
}
