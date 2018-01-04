package actions

import (
	"errors"
	"testing"
)

func Test_Action_Execute_executes_the_given_function(t *testing.T) {
	test := "Before Test"
	f := func() { test = "After Test" }
	a := NewAction("test", "this is a test", f)

	a.Execute()
	actual := test
	expected := "After Test"
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func Test_Choice_AddAction_returns_an_error_when_an_action_with_the_same_name_already_exists(t *testing.T) {
	a1 := NewAction("test1", "test1 description", func() {})
	a2 := NewAction("test1", "test1 description", func() {})
	c := NewChoice()

	c.AddAction(a1)
	actual := c.AddAction(a2)

	expected := errors.New("An action with this name already exists")
	if actual.Error() != expected.Error() {
		t.Error("Expected", expected, "but got", actual)
	}
}

func Test_choiceIsValid_returns_false_when_no_action_exists(t *testing.T) {
	c := NewChoice()

	actual := c.(*ChoiceImpl).choiceIsValid("")
	expected := false

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func Test_choiceIsValid_returns_false_when_no_action_with_the_name_matching_the_choice_exists(t *testing.T) {
	a := NewAction("test", "desc", func() {})
	c := NewChoice()
	c.AddAction(a)

	actual := c.(*ChoiceImpl).choiceIsValid("")
	expected := false

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func Test_choiceIsValid_returns_true_when_an_action_with_the_name_matching_the_choice_exists(t *testing.T) {
	a := NewAction("test", "desc", func() {})
	c := NewChoice()
	c.AddAction(a)

	actual := c.(*ChoiceImpl).choiceIsValid("test")
	expected := true

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
