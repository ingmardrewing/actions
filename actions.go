package actions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Action interface {
	GetName() string
	GetDescription() string
	Execute()
}

/* Actions */

func NewAction(name, description string, exec func()) Action {
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

func (a *ActionImpl) Execute() {
	if a.exec != nil {
		a.exec()
	}
}

/* Choice */

type Choice interface {
	AddAction(action Action) error
	AskUser() Action
}

func NewChoice() Choice {
	return new(ChoiceImpl)
}

type ChoiceImpl struct {
	actions []Action
}

func (c *ChoiceImpl) AddAction(a Action) error {
	if c.getActionByName(a.GetName()) == nil {
		c.actions = append(c.actions, a)
		return nil
	}
	return errors.New("An action with this name already exists")
}

func (c *ChoiceImpl) displayActions() {
	for _, a := range c.actions {
		fmt.Printf("%s: %s\n", a.GetName(), a.GetDescription())
	}
}

func (c *ChoiceImpl) choiceIsValid(choice string) bool {
	a := c.getActionByName(choice)
	return a != nil
}

func (c *ChoiceImpl) AskUser() Action {
	choice := ""
	for !c.choiceIsValid(choice) {
		fmt.Println("Your options:")
		c.displayActions()
		choice = c.GetUsersChoice()
	}
	return c.getActionByName(choice)
}

func (c *ChoiceImpl) getActionByName(name string) Action {
	for _, a := range c.actions {
		if a.GetName() == name {
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
