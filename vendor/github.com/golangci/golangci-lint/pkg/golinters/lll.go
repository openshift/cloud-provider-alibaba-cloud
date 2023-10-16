package golinters

import (
	"bufio"
	"fmt"
	"go/token"
	"os"
	"strings"
	"sync"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"github.com/golangci/golangci-lint/pkg/lint/linter"
	"github.com/golangci/golangci-lint/pkg/result"
)

const lllName = "lll"

const goCommentDirectivePrefix = "//go:"

//nolint:dupl
func NewLLL(settings *config.LllSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []goanalysis.Issue

	analyzer := &analysis.Analyzer{
		Name: lllName,
		Doc:  goanalysis.TheOnlyanalyzerDoc,
		Run: func(pass *analysis.Pass) (interface{}, error) {
			issues, err := runLll(pass, settings)
			if err != nil {
				return nil, err
			}

			if len(issues) == 0 {
				return nil, nil
			}

			mu.Lock()
			resIssues = append(resIssues, issues...)
			mu.Unlock()

			return nil, nil
		},
	}

	return goanalysis.NewLinter(
		lllName,
		"Reports long lines",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithIssuesReporter(func(*linter.Context) []goanalysis.Issue {
		return resIssues
	}).WithLoadMode(goanalysis.LoadModeSyntax)
}

func runLll(pass *analysis.Pass, settings *config.LllSettings) ([]goanalysis.Issue, error) {
	fileNames := getFileNames(pass)

	spaces := strings.Repeat(" ", settings.TabWidth)

	var issues []goanalysis.Issue
	for _, f := range fileNames {
		lintIssues, err := getLLLIssuesForFile(f, settings.LineLength, spaces)
		if err != nil {
			return nil, err
		}

		for i := range lintIssues {
			issues = append(issues, goanalysis.NewIssue(&lintIssues[i], pass))
		}
	}

	return issues, nil
}

func getLLLIssuesForFile(filename string, maxLineLen int, tabSpaces string) ([]result.Issue, error) {
	var res []result.Issue

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can't open file %s: %s", filename, err)
	}
	defer f.Close()

	lineNumber := 0
	multiImportEnabled := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()
		line = strings.ReplaceAll(line, "\t", tabSpaces)

		if strings.HasPrefix(line, goCommentDirectivePrefix) {
			continue
		}

		if strings.HasPrefix(line, "import") {
			multiImportEnabled = strings.HasSuffix(line, "(")
			continue
		}

		if multiImportEnabled {
			if line == ")" {
				multiImportEnabled = false
			}

			continue
		}

		lineLen := utf8.RuneCountInString(line)
		if lineLen > maxLineLen {
			res = append(res, result.Issue{
				Pos: token.Position{
					Filename: filename,
					Line:     lineNumber,
				},
				Text:       fmt.Sprintf("line is %d characters", lineLen),
				FromLinter: lllName,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		if err == bufio.ErrTooLong && maxLineLen < bufio.MaxScanTokenSize {
			// scanner.Scan() might fail if the line is longer than bufio.MaxScanTokenSize
			// In the case where the specified maxLineLen is smaller than bufio.MaxScanTokenSize
			// we can return this line as a long line instead of returning an error.
			// The reason for this change is that this case might happen with autogenerated files
			// The go-bindata tool for instance might generate a file with a very long line.
			// In this case, as it's an auto generated file, the warning returned by lll will
			// be ignored.
			// But if we return a linter error here, and this error happens for an autogenerated
			// file the error will be discarded (fine), but all the subsequent errors for lll will
			// be discarded for other files, and we'll miss legit error.
			res = append(res, result.Issue{
				Pos: token.Position{
					Filename: filename,
					Line:     lineNumber,
					Column:   1,
				},
				Text:       fmt.Sprintf("line is more than %d characters", bufio.MaxScanTokenSize),
				FromLinter: lllName,
			})
		} else {
			return nil, fmt.Errorf("can't scan file %s: %s", filename, err)
		}
	}

	return res, nil
}
