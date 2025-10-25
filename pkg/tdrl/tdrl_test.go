package tdrl

import (
	"testing"

	"github.com/Tehhs/tdr/pkg/util"
)

func Test_BasicTodos(t *testing.T) {
	tdrlParser := NewParser()

	todos := tdrlParser.ProcessTodo("todo: this is a basic todo")

	if len(todos) != 1 {
		t.Error("could not process todos")
	}

}

func Test_Tags(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []string
	}{
		{
			name:    "Simple Tags",
			content: "todo(tag1, tag2, tag3): this is a todo",
			want:    []string{"tag1", "tag2", "tag3"},
		},
		//More here when we work out what exactly we like in tags
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			got := p.ProcessTodo(tt.content)

			if len(got) != 1 {
				t.Errorf("Got %d amount of todos returned from parsing '%s'; Should be 1", len(got), tt.content)
			}

			todo := got[0]

			if !util.ArraysEqual(todo.Tags, tt.want) {
				t.Errorf("Tags for '%s' = %+v, want %+v", tt.content, todo.Tags, tt.want)
			}
		})
	}
}
