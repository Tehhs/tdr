package cli

import (
	"fmt"
	"os"

	components "github.com/Tehhs/tdr/pkg/cli/components"
	"github.com/Tehhs/tdr/pkg/core"
	"github.com/Tehhs/tdr/pkg/util"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/lipgloss"
)

type TdrCli struct {
	TDRClient *core.TDRCore
}

type NewCLIParams struct {
}

const (
	screenTags = iota
	screenTodos
)

type BubbleTeaModel struct {
	FileOrFolder *string
	HeaderModel  components.HeaderModel
	TagsModel    components.TagsModel
	TodoModel    components.TodoModel
	ViewState    int
}

func initialModel() BubbleTeaModel {
	return BubbleTeaModel{
		FileOrFolder: util.Ptr("examplefile.txt"),
		TagsModel:    components.NewTagsModel(),
		TodoModel:    components.NewTodoModel(),
		HeaderModel:  components.NewHeaderModel(),
	}
}

func (btm BubbleTeaModel) Init() tea.Cmd {
	return nil
}

func (m BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		//switch views
		case "up", "down", "s":
			if m.ViewState == screenTags {
				m.ViewState = screenTodos
			} else {
				m.ViewState = screenTags
			}
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m BubbleTeaModel) View() string {
	s := ""

	headerView := m.HeaderModel.View()
	s += fmt.Sprintf("%s", headerView)

	var mainView string = "no main view selected"
	switch m.ViewState {
	case screenTags:
		mainView = m.TagsModel.View()
	case screenTodos:
		mainView = m.TodoModel.View()
	}

	s += fmt.Sprintf("\n\n%s\n\n", mainView)

	// The footer
	footer := "\n\n[q - Quit] [s switch views]"


	finalView := lipgloss.JoinVertical(
		lipgloss.Left,
		headerView,
		mainView,
		footer,
	)

	// Send the UI for rendering
	return finalView
}

func New(args NewCLIParams) *TdrCli {

	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return nil
}
