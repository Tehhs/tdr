package components

import (

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/lipgloss/table"
)

type Tag struct {
	TagName string
	Amt     int
}

type TagsModel struct {
	FileName string
	Tags     []Tag
	TagIndex int
}

var (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")

	cellStyle    = lipgloss.NewStyle().Padding(0, 1).Padding(1, 0)
	oddRowStyle  = cellStyle.Foreground(gray)
	evenRowStyle = cellStyle.Foreground(lightGray)
)

func (m TagsModel) View() string {
	s := ""

	
	// amtPerRow := 2 
	// rows := [][]string{}
	// for i, item := range m.Tags { 
	// 	rowIndex := i / amtPerRow
	// 	colIndex := i % amtPerRow
	// 	rowIndex[rowIndex][colIndex] = item.TagName

	// }

	

	// t := table.New().
	// 	Border(lipgloss.HiddenBorder()).
	// 	Width(80).
	// 	StyleFunc(func(row, col int) lipgloss.Style {
	// 		switch {
	// 		case row%2 == 0:
	// 			return evenRowStyle
	// 		default:
	// 			return oddRowStyle
	// 		}
	// 	})
		// Rows(rows...)

	// s += t.Render()

	return s
}

func (btm TagsModel) Init() tea.Cmd {
	return nil
}

func (m TagsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func NewTagsModel() TagsModel {

	tagsModel := TagsModel{}

	tagsModel.Tags = []Tag{
		{TagName: "Tag1"},
		{TagName: "Tag2"},
		{TagName: "Tag3"},
	}
	return tagsModel
}
