// Simple tool for interactive shell sessions
package actions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action interface {
	GetName() string
	GetDescription() string
	GetFunction() func()
	execute()
}

// Actions
func newAction(name, description string, exec func()) Action {
	a := new(ActionImpl)
	a.name = name
	a.description = description
	a.exec = exec
	return a
}

type ActionImpl struct {
	name        string
	description string
	exec        func()
}

func (a *ActionImpl) GetDescription() string {
	return a.description
}

func (a *ActionImpl) GetName() string {
	return a.name
}

func (a *ActionImpl) GetFunction() func() {
	return a.exec
}

func (a *ActionImpl) execute() {
	if a.exec != nil {
		a.exec()
	}
}

// Create a new Choice
func NewChoice() Choice {
	return new(ChoiceImpl)
}

// Choice interface
type Choice interface {
	AddAction(name, description string, callback func()) error
	AskUser()
	Actions() []Action
}

// Implementation of the Choice interface, holding a splice of actions
type ChoiceImpl struct {
	actions []Action
}

func (c *ChoiceImpl) Actions() []Action {
	return c.actions
}

// AddAction allows to add an action choosable for the user of the
// cli application using this package. The name should be short, for it
// must be typed by the user.
func (c *ChoiceImpl) AddAction(name, description string, callback func()) error {
	if c.getActionByName(name) == nil {
		a := newAction(name, description, callback)
		c.actions = append(c.actions, a)
		return nil
	}
	return errors.New("An action with this name already exists")
}

func (c *ChoiceImpl) displayActions() {
	l := c.findLongestActionName()
	tmpl := "%-" + l + "s: %s\n"
	for _, a := range c.actions {
		fmt.Printf(tmpl, a.GetName(), a.GetDescription())
	}
}

func (c *ChoiceImpl) findLongestActionName() string {
	l := 0
	for _, a := range c.actions {
		if len(a.GetName()) > l {
			l = len(a.GetName())
		}
	}
	return strconv.Itoa(l)
}

func (c *ChoiceImpl) choiceIsValid(choice string) bool {
	a := c.getActionByName(choice)
	return a != nil
}

// Prompting the user with the available actions.
// If put in a conditionless for loop this leads to
// a cli application displaying the actions again
// after any successful execution of an action.
func (c *ChoiceImpl) AskUser() {
	choice := ""
	for !c.choiceIsValid(choice) {
		fmt.Println("----- Available options:")
		c.displayActions()
		fmt.Println("----- Your choice:")
		choice = c.getUsersChoice()
	}
	a := c.getActionByName(choice)
	a.execute()
}

func (c *ChoiceImpl) getActionByName(name string) Action {
	for _, a := range c.actions {
		if a.GetName() == name {
			return a
		}
	}
	return nil
}

func (c *ChoiceImpl) getUsersChoice() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
