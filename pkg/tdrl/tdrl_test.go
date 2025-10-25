package tdrl

import (
	"fmt"
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

func Test_ProcessedText(t *testing.T) {
	p := NewParser()

	//Simple 

	text := "this should work"
	todos := p.ProcessTodo(fmt.Sprintf("todo: %s", text))

	if len(todos) != 1 {
		t.Error("returned no todos or invalid amount of todos")
	}

	todo := todos[0]

	if todo.ProcessedContent != text {
		t.Errorf("Failed to process content. Got '%s', wanted '%s'", todo.ProcessedContent, text)
	}

	//Probably want to add more test for if there's no leading space like "todo:this should work" returning "this should work"

	//todo(test): Need to make sure processed text works with tags too 
}

func Test_WithoutTags(t *testing.T) {
	p := NewParser()
	text := "this should work"
	todos := p.ProcessTodo(fmt.Sprintf("todo: %s", text))

	if len(todos) != 1 {
		t.Error("returned no todos or invalid amount of todos")
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
