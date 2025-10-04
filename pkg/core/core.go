package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Tehhs/tdr/pkg/comments"
	"github.com/Tehhs/tdr/pkg/tdrl"
	"github.com/Tehhs/tdr/pkg/util"
)

type TDRCore struct {
}

func NewTDRCore() TDRCore {
	return TDRCore{}
}

// todo(refactor): not sure if this should be here
type TodoBlock struct {
	CommentBlock *comments.CommentBlock
	RawContent   *ContentType
}

type ContentType struct {

	//The actual content to process
	Content *[]byte

	//Name of the content. This might be the filename, filepath+filename,
	//or just any random name for the content in the case where there are
	//no files
	ContentName *string

	//In the case of files, this might be the extension name (example: "go", "js", "jsx").
	//This is required to work out how to extract comments from the language. If there is
	//no file, you have to work out what the extension is and supply it here.
	ContentType *string
}

func (c TDRCore) getTodosFromContent(content ContentType) (*[]TodoBlock, error) {
	if content.Content == nil { 
		return nil, fmt.Errorf("cannot process todo content with no actual content")
	}

	if content.ContentType == nil {
		return nil, fmt.Errorf("cannot process todo content with no content type")
	}

	if content.ContentName == nil {
		return nil, fmt.Errorf("cannot process todo content with no content name")
	}

	//todo(refactor): Comments layer should return comment layer errors, and
	//have a special file extension not supported error to check for here
	//instead of just ingoring.
	parseResult, err := comments.Parse(util.Ptr(string(*content.Content)), *content.ContentType)

	if err != nil || parseResult == nil {
		return nil, fmt.Errorf("could not parse comments out of %s", *content.ContentName)
	}

	var todoCommentBlocks []TodoBlock = []TodoBlock{}

	for _, commentBlock := range parseResult.Comments {
		hasTodo := strings.Contains(strings.ToLower(commentBlock.String()), "todo")
		if hasTodo {
			todoCommentBlocks = append(todoCommentBlocks, TodoBlock{
				CommentBlock: commentBlock,
				RawContent:   &content,
			})
		}
	}

	return &todoCommentBlocks, nil
}

func (c TDRCore) getTodosCommentsFromFileOrDirectory(fileOrFolder string) (*[]TodoBlock, error) {
	fileOrFolderInfo, err := os.Stat(fileOrFolder)
	if err != nil {
		return nil, fmt.Errorf("could not get stats of file or folder %s", fileOrFolder)
	}
	if fileOrFolderInfo.IsDir() {
		panic("unimplemented") //todo(unimplemented): unimplemented
	}

	if !fileOrFolderInfo.IsDir() {
		contents, err := os.ReadFile(fileOrFolder)
		if err != nil {
			return nil, fmt.Errorf("could not read file %s", fileOrFolder)
		}
		
		if !strings.Contains(fileOrFolder, ".") { 
			return nil, fmt.Errorf("file %s does not have a type", fileOrFolder)
		}
		filenameSegments := strings.Split(fileOrFolder, ".")
		extension := filenameSegments[len(filenameSegments)-1]
		extension = strings.ToLower(extension)
		extension = strings.TrimSpace(extension)

		return c.getTodosFromContent(ContentType{
			Content:     &contents,
			ContentName: util.Ptr(fileOrFolderInfo.Name()),
			ContentType: &extension,
		})
	}

	panic("unimplemented")
}

type ProcessedNamedContent struct { 
	Name *string 
	Processed []tdrl.TDRLTodo
}

type ProcessOutput struct { 
	Todos []ProcessedNamedContent
}

func (c TDRCore) Process(fileOrFolder string) (ProcessOutput, error) { 
	tdrlParser := tdrl.NewParser()
	
	todoCommentBlocks, err := c.getTodosCommentsFromFileOrDirectory(fileOrFolder)
	if err != nil || todoCommentBlocks == nil { 
		return ProcessOutput{}, err
	}

	processedTodoOutput := []ProcessedNamedContent{}

	for _, todoCommentBlock := range *todoCommentBlocks { 
		
		if todoCommentBlock.CommentBlock == nil {
			//todo(refactor): Not sure if this should happen. Need to check and possibly refactor
			log.Print("Warning: comment block is nil\n")
			continue 
		}
		
		todoLines := []string{}
		for _, line := range todoCommentBlock.CommentBlock.Lines {
			todoLines = append(todoLines, line.String())
		}
		todoContent := strings.Join(todoLines, " ")
		todoContent = strings.ToLower(todoContent)

		processedTodos := tdrlParser.ProcessTodo(todoContent)
		processedTodoOutput = append(processedTodoOutput, ProcessedNamedContent{
			Name: todoCommentBlock.RawContent.ContentName, //todo(ptr): possible nil ptr deref
			Processed: processedTodos,
		})
		
	}

	return ProcessOutput{
		Todos: processedTodoOutput,
	}, nil 
}
