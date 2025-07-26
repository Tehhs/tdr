package comments

import (
	"strings"
	"testing"
)

func TestBasicComments(t *testing.T) { 
	
	var content string = `
package test

func SomeFunction(str *string) { 
	var number int = 2
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

	if len(parseResult.Comments) != 1 { 
		t.Errorf("Did not parse the correct amount of comments; Got %d", len(parseResult.Comments))
	}

}



func TestMultipleComments(t *testing.T) { 
	
	var content string = `
package test

func SomeFunction(str *string) { 
	//todo: this is another comment that should be picked up
	var number int = 2
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


func TestSequentialLineComments(t *testing.T) { 
	
	var content string = `
package test

func SomeFunction(str *string) { 
	var number int = 2
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

