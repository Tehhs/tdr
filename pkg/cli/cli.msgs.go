package cli

import "github.com/Tehhs/tdr/pkg/core"

type ScannedMsg struct {
	HasScanError bool
	ScanErrorMsg string

	NewTodos *[]core.ProcessedNamedContent
}
