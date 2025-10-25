package cli

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
)

func (m BubbleTeaModel) update_scanRequest(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, ok := msg.(ScanRequest)
	if !ok {
		slog.Error("invalid scan request")
		return m, nil
	}

	return m, scanCmd(m.TDRCore, "/")

}

func (m BubbleTeaModel) update_scannedMessage(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.IsLoading = false

	scannedMessage, ok := msg.(ScannedMsg)

	if !ok {
		slog.Error("invalid scanned message")
		return m, nil
	}

	if scannedMessage.HasScanError {
		slog.Error(scannedMessage.ScanErrorMsg)
		//Maybe display this error
		return m, nil
	}

	m.ProcessedNamedContenet = scannedMessage.NewTodos

	return m, nil

}

func (m BubbleTeaModel) update_keyMsg(msg tea.Msg) (tea.Model, tea.Cmd) {
	var keyMsg *tea.KeyMsg = nil
	if km, ok := msg.(tea.KeyMsg); ok {
		keyMsg = &km
	} else {
		slog.Error("call to update_keyMsg with non tea.KeyMsg type arg")
		return m, nil
	}

	switch keyMsg.String() {

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

	return m, nil
}
