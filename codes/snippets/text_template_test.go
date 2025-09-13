// text/template example.
package snippets

// Github
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	IssuesURL = "https://api.github.com/search/issues"
)

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown格式
}

// User
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssue
func SearchIssue(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	queryUrl := IssuesURL + "?q=" + q
	fmt.Println(queryUrl)
	resp, err := http.Get(queryUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

import (
	"os"
	"testing"
	"text/template"
	"time"

	// "github.com/zhoujiagen/examples/github"
)

// 展示了模板语言: 动作, 循环, 函数
const textTemplate = `{{.TotalCount}} issues:
{{range .Items}}------------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var textReport = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(textTemplate))

func TestTemplate(t *testing.T) {
	terms := []string{"HBase"}
	// result, err := github.SearchIssue(terms)
	result, err := SearchIssue(terms)
	if err != nil {
		t.Error(err)
	}

	if err := textReport.Execute(os.Stdout, result); err != nil {
		t.Error(err)
	}
}