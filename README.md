[![Build Status](https://travis-ci.org/ingmardrewing/actions.svg?branch=master)](https://travis-ci.org/ingmardrewing/actions)

# Actions

A simple package making interactive console sessions possible

## Usage

I am using the package to have a simple recurring menu on the commandline.
Integration in a project is - imho - straight forward:
  - create a new Choice
  - add actions the user can execute
  - perpertuously call Choice.AskUser()

An action consist of a short name, a short description and a function that will be executed when the user selects the action.
A simple setup may look like this:

```
	c := actions.NewChoice()
	c.AddAction("exit", "Exits the Application", func() { os.Exit(0) })
	c.AddAction("make", "Generate website locally", generateSiteLocally)
```

Once you defined the actions, just put a call to the AskUser() method into a for loop somewhere in the main function, like this:

```
	for {
		c.AskUser()
	}
```

This way the user will be prompted to choose an action whenever the program starts or finishes an action.

