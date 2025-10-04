package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TagsModel struct {
	FileName string
}

func (m TagsModel) View() string {
	s := `

*****************************************
*
*		THIS IS THE TAGS VIEW (WIP)
*
*****************************************

`

	return s
}

func (btm TagsModel) Init() tea.Cmd {
	return nil
}

func (m TagsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func NewTagsModel() TagsModel {
	return TagsModel{}
}
