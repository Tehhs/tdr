package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TodoModel struct {
	FileName string
}

func (m TodoModel) View() string {
	s := `*****************************************
*
*		THIS IS THE TODO VIEW (WIP)
*
*****************************************

`

	return s
}

func (btm TodoModel) Init() tea.Cmd {
	return nil
}

func (m TodoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func NewTodoModel() TodoModel {
	return TodoModel{}
}
