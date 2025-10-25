package cli

import (
	"github.com/Tehhs/tdr/pkg/core"
	tea "github.com/charmbracelet/bubbletea"
)

func scanCmd(tdrc *core.TDRCore, filepath string) tea.Cmd {

	if tdrc == nil {
		panic("tdr is not initialized")
	}

	return func() tea.Msg {

		output, err := tdrc.Process(filepath)

		if err != nil {
			return ScannedMsg{
				HasScanError: true,
				ScanErrorMsg: "Could not scan for todos",
			}
		}

		return ScannedMsg{
			NewTodos: &output.Todos,
		}
	}
}
