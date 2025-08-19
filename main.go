package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// model stored the application state, can be of any type

type model struct {
	choices []string //items on the to-do list
	cursor int //which to-do list item the cursor is pointing at
	selected map[int]struct{} //which to-do items are selected
}

//initialization state

func initialModel() model {
	return model {
		choices: []string{"Bring carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{})
	}
}


func (m model) Init() tea.Cmd {
	return nil //no IO right now
}


//Update method is the "something happened!" function, it update the model based on the event that happened
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) { // ????
		case tea.KeyMsg:
			
			switch msg.String() {
				case "ctrl+c", "q":
					return m, tea.Quit

		}
	}
}
