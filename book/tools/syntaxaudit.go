// syntaxaudit records the first use of tokens in each registered code checkpoint.
package main

import (
	"encoding/json"
	"fmt"
	"go/scanner"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

type chapter struct {
	ID         string `json:"id"`
	Checkpoint string `json:"checkpoint"`
}

type occurrence struct {
	Chapter string `json:"chapter"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Token   string `json:"token"`
}

func main() {
	if err := audit(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func audit() error {
	data, err := os.ReadFile("book.json")
	if err != nil {
		return err
	}
	var metadata struct {
		Chapters []chapter `json:"chapters"`
	}
	if err := json.Unmarshal(data, &metadata); err != nil {
		return err
	}
	builtins := map[string]bool{
		"_": true, "nil": true, "true": true, "false": true,
		"len": true, "cap": true, "make": true, "append": true,
		"copy": true, "delete": true, "clear": true, "new": true,
		"close": true, "panic": true, "recover": true, "min": true,
		"max": true, "complex": true, "real": true, "imag": true,
		"print": true, "println": true,
	}
	var found []occurrence
	for _, ch := range metadata.Chapters {
		if ch.Checkpoint == "" {
			continue
		}
		var paths []string
		err := filepath.WalkDir(ch.Checkpoint, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if !entry.IsDir() && filepath.Ext(path) == ".go" {
				paths = append(paths, path)
			}
			return nil
		})
		if err != nil {
			return err
		}
		seen := make(map[string]bool)
		for _, path := range paths {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			set := token.NewFileSet()
			file := set.AddFile(path, -1, len(data))
			var lex scanner.Scanner
			lex.Init(file, data, nil, scanner.ScanComments)
			for {
				pos, tok, lit := lex.Scan()
				if tok == token.EOF {
					break
				}
				name := tok.String()
				if tok == token.IDENT {
					if !builtins[lit] {
						continue
					}
					name = lit
				}
				if !seen[name] {
					seen[name] = true
					found = append(found, occurrence{ch.ID, filepath.ToSlash(path), set.Position(pos).Line, name})
				}
			}
			if lex.ErrorCount != 0 {
				return fmt.Errorf("%s: lexical errors", path)
			}
		}
	}
	sort.Slice(found, func(i, j int) bool {
		if found[i].Chapter == found[j].Chapter {
			return found[i].Token < found[j].Token
		}
		return found[i].Chapter < found[j].Chapter
	})
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(found)
}
