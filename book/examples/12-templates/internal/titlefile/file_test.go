package titlefile

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadTitle(t *testing.T) {
	path := filepath.Join(t.TempDir(), "title.txt")
	if err := os.WriteFile(path, []byte("  Шаңырақ Go  "), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := Load(path)
	if err != nil || got != "Шаңырақ Go" {
		t.Fatalf("got (%q, %v)", got, err)
	}
}

func TestMissingTitleKeepsCause(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.txt")
	got, err := Load(path)
	if got != "" || !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("got (%q, %v), want missing-file cause", got, err)
	}
}

type trackedReader struct {
	reader   io.Reader
	closeErr error
	closes   int
}

func (r *trackedReader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *trackedReader) Close() error {
	r.closes++
	return r.closeErr
}

func TestTitleFileBoundaries(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{"valid", "  Go  ", "Go", nil},
		{"empty", "", "", ErrInvalidTitle},
		{"multiline", "Go\nShop", "", ErrInvalidTitle},
		{"invalid UTF-8", string([]byte{0xff}), "", ErrInvalidTitle},
		{"exact byte limit", strings.Repeat(" ", 1022) + "Go", "Go", nil},
		{"above byte limit", strings.Repeat(" ", 1023) + "Go", "", ErrTitleTooLarge},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reader := &trackedReader{reader: strings.NewReader(tc.input)}
			got, err := readTitleAndClose(reader)
			if got != tc.want || !errors.Is(err, tc.wantErr) || reader.closes != 1 {
				t.Fatalf("got (%q, %v), closes=%d", got, err, reader.closes)
			}
		})
	}
}

type failingReader struct{}

func (failingReader) Read(p []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestReadAndCloseErrorsBothSurvive(t *testing.T) {
	closeErr := errors.New("close failed")
	reader := &trackedReader{reader: failingReader{}, closeErr: closeErr}
	got, err := readTitleAndClose(reader)
	if got != "" || reader.closes != 1 || !errors.Is(err, io.ErrUnexpectedEOF) || !errors.Is(err, closeErr) {
		t.Fatalf("got (%q, %v), closes=%d", got, err, reader.closes)
	}
}

func TestCloseErrorDiscardsTitle(t *testing.T) {
	closeErr := errors.New("close failed")
	reader := &trackedReader{reader: strings.NewReader("Go"), closeErr: closeErr}
	got, err := readTitleAndClose(reader)
	if got != "" || !errors.Is(err, closeErr) || reader.closes != 1 {
		t.Fatalf("got (%q, %v), closes=%d", got, err, reader.closes)
	}
}
