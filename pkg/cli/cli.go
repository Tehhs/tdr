package cli

import (
	"fmt"
	"os"

	components "github.com/Tehhs/tdr/pkg/cli/components"
	"github.com/Tehhs/tdr/pkg/core"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/lipgloss"
)

type TdrCli struct {
	TDRClient *core.TDRCore
}

type NewCLIParams struct {
}

type ScreenState int

const (
	screenTags ScreenState = iota
	screenTodos
)

type ScanState int

const (
	ScanInProgress ScanState = iota
	ScanFinished
	ScanErrored
	ScanHasntScanned
)

type BubbleTeaModel struct {
	FileOrFolder *string
	HeaderModel  components.HeaderModel
	TagsModel    components.TagsModel
	TodoModel    components.TodoModel

	ViewState              ScreenState
	TDRCore                *core.TDRCore
	DisplayLoadingPage     bool
	IsLoading              bool
	ProcessedNamedContenet *[]core.ProcessedNamedContent
}

func initialModel() BubbleTeaModel {

	coreInstance := core.NewTDRCore()
	// processOutput, err := coreInstance.Process("./main.go")

	return BubbleTeaModel{
		TagsModel:          components.NewTagsModel(),
		TodoModel:          components.NewTodoModel(),
		HeaderModel:        components.NewHeaderModel(),
		TDRCore:            coreInstance,
		DisplayLoadingPage: true,
		IsLoading:          true,
	}
}

func (btm BubbleTeaModel) Init() tea.Cmd {
	return scanCmd(btm.TDRCore, "/")
}

func (m BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m.update_keyMsg(msg)

	case ScanRequest:
		return m.update_scanRequest(msg)

	case ScannedMsg:
		return m.update_scannedMessage(msg)

	// default:
		// there can be other bubble tea events so dont emit errors here
		// slog.Error("unknown update message")
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
