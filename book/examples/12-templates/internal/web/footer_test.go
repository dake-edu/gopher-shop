package web

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"testing"
	"time"
)

func TestFooterUsesRenderTimeYearWithoutPageData(t *testing.T) {
	page := pageTemplate("home")
	year := time.Now().Year()
	var output bytes.Buffer
	if err := page.ExecuteTemplate(&output, "footer", nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), fmt.Sprintf("&copy; %d eGopher", year)) {
		t.Fatalf("missing copyright: %s", output.String())
	}

	// Keep one parsed template and simulate the clock crossing New Year.
	page = page.Funcs(template.FuncMap{"year": func() int { return year }})
	for _, next := range []int{2026, 2027} {
		year = next
		output.Reset()
		if err := page.ExecuteTemplate(&output, "footer", nil); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(output.String(), fmt.Sprintf("&copy; %d ", next)) {
			t.Fatalf("year was frozen before rendering: %s", output.String())
		}
	}
}
