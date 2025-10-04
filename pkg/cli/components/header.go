package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type HeaderModel struct {
	FileName string
}

func (m HeaderModel) View() string {
	s := fmt.Sprintf("File Name: %s\n\n", m.FileName)

	return s
}

func (btm HeaderModel) Init() tea.Cmd {
	return nil
}

func (m HeaderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func NewHeaderModel() HeaderModel {
	return HeaderModel{}
}
