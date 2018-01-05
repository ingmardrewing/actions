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

type action interface {
	getName() string
	getDescription() string
	execute()
}

/* Actions */

func newAction(name, description string, exec func()) action {
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

func (a *ActionImpl) getDescription() string {
	return a.description
}

func (a *ActionImpl) getName() string {
	return a.name
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
}

// Implementation of the Choice interface, holding a splice of actions
type ChoiceImpl struct {
	actions []action
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
		fmt.Printf(tmpl, a.getName(), a.getDescription())
	}
}

func (c *ChoiceImpl) findLongestActionName() string {
	l := 0
	for _, a := range c.actions {
		if len(a.getName()) > l {
			l = len(a.getName())
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
		choice = c.GetUsersChoice()
	}
	a := c.getActionByName(choice)
	a.execute()
}

func (c *ChoiceImpl) getActionByName(name string) action {
	for _, a := range c.actions {
		if a.getName() == name {
			return a
		}
	}
	return nil
}

func (c *ChoiceImpl) GetUsersChoice() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
