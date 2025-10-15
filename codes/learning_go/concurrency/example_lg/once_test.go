package example_lg

import (
	"fmt"
	"sync"
	"testing"
)

// Example of sync.Once. - 'Learning Go': 12. Concurrency in Go, P.308

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

// use sync.OnceValue
var initParserCached func() SlowComplicatedParser = sync.OnceValue(initParser)

func Parse(dataToParse string) string {
	// do initialize once
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func ParseCached(dataToParse string) string {
	_parser := initParserCached()
	return _parser.Parse(dataToParse)
}

// initialize the parser: should be called only once.
func initParser() SlowComplicatedParser {
	fmt.Println("initialize parser")

	return SCPI{}
}

// mock implementantion
type SCPI struct {
}

func (s SCPI) Parse(in string) string {
	if len(in) > 1 {
		return in[0:1]
	}
	return ""
}

func TestOnce(t *testing.T) {
	result := Parse("hello")
	fmt.Println("result=", result)
	result = Parse("golang")
	fmt.Println("result=", result)

	fmt.Println()

	result = ParseCached("hello")
	fmt.Println("result=", result)
	result = ParseCached("golang")
	fmt.Println("result=", result)
}
