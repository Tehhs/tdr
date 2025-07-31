package comments

import (
	"strings"
	"testing"
)

func TestBasicJavascriptComments(t *testing.T) { 
	
	var content string = `


	;(()=>{
		const a = "test"
		//todo(test): this comment should be picked up
	})();


	`
	
	parseResult, err := Parse(&content, "js")

	if err != nil { 
		t.Errorf("Failed to parse %v", err)
	}

	if parseResult == nil {
		t.Errorf("Parse result returned nil")
	}

	if len(parseResult.Comments) != 1 { 
		t.Errorf("Did not parse the correct amount of comments; Got %d", len(parseResult.Comments))
	}

}



func TestMultipleJavascriptComments(t *testing.T) { 
	
	var content string = `

function f(){ 
	//todo: this is another comment that should be picked up
	const number = 2
	//todo(test): this comment should be picked up
}

	`
	
	parseResult, err := Parse(&content, "go")

	if err != nil { 
		t.Errorf("Failed to parse %v", err)
	}

	if parseResult == nil {
		t.Errorf("Parse result returned nil")
	}

	if len(parseResult.Comments) != 2 { 
		t.Errorf("Did not parse the correct amount of comments; Got %d", len(parseResult.Comments))
	}

}


func TestSequentialLineJavascriptComments(t *testing.T) { 
	
	var content string = `

function SomeFunction() { 
	const number = 2
	//todo(test): this comment should be picked up
	//and this should also be in the same comment block TESTFORTHISSTRING
}

	`
	
	parseResult, err := Parse(&content, "go")

	if err != nil { 
		t.Errorf("Failed to parse %v", err)
	}

	if parseResult == nil {
		t.Errorf("Parse result returned nil")
	}

	if len(parseResult.Comments) == 0 { 
		t.Errorf("Did not pick up on any errors")
	}

	commentBlockStr := parseResult.Comments[0].String()
	if !strings.Contains(commentBlockStr, "TESTFORTHISSTRING") { 
		t.Error("Did not pick up on second line of comment string")
	}

}

