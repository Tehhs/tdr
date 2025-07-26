package core

import (
	"log"
	"os"
	"strings"

	"github.com/Tehhs/tdr/pkg/comments"
	"github.com/Tehhs/tdr/pkg/util"
)

type TDRCore struct {
}

func NewTDRCore() TDRCore {
	return TDRCore{}
}

//todo(refactor): not sure if this should be here
type TodoBlock struct {
	CommentBlock *comments.CommentBlock
	FileName     string
}

func (c TDRCore) ProcessFile(path string) []TodoBlock {
	var todoCommentBlocks []TodoBlock = []TodoBlock{}

	var extension *string = nil
	filePathParts := strings.Split(path, ".")
	extension = util.Ptr(filePathParts[len(filePathParts)-1])

	content, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("File '%s' could not be read.", path)
	}

	//todo(refactor): Comments layer should return comment layer errors, and
	//have a special file extension not supported error to check for here
	//instead of just ingoring.
	parseResult, err := comments.Parse(util.Ptr(string(content)), *extension)

	if err != nil {
		// log.Panic("Error parsing")
	}

	if parseResult == nil || parseResult.Comments == nil {
		return todoCommentBlocks
	}

	for _, commentBlock := range parseResult.Comments {
		hasTodo := strings.Contains(strings.ToLower(commentBlock.String()), "todo")
		if hasTodo {
			todoCommentBlocks = append(todoCommentBlocks, TodoBlock{
				CommentBlock: commentBlock,
				FileName:     path,
			})
		}
	}

	return todoCommentBlocks
}
