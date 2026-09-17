package titlefile

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dake-edu/gopher-shop/book/examples/12-templates/internal/catalog"
)

const maxTitleFileBytes = 1024

var ErrTitleTooLarge = errors.New("файл названия слишком большой")
var ErrInvalidTitle = errors.New("название книги недопустимо")

func Load(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("открыть название %q: %w", path, err)
	}
	title, err := readTitleAndClose(file)
	if err != nil {
		return "", fmt.Errorf("прочитать название %q: %w", path, err)
	}
	return title, nil
}

// readTitleAndClose принимает на себя обязанность закрыть reader.
func readTitleAndClose(reader io.ReadCloser) (title string, err error) {
	defer func() {
		closeErr := reader.Close()
		if closeErr != nil {
			title = ""
			err = errors.Join(err, closeErr)
		}
	}()

	data, err := io.ReadAll(io.LimitReader(reader, maxTitleFileBytes+1))
	if err != nil {
		return "", err
	}
	if len(data) > maxTitleFileBytes {
		return "", ErrTitleTooLarge
	}
	title, ok := catalog.NormalizeTitle(string(data))
	if !ok {
		return "", ErrInvalidTitle
	}
	return title, nil
}
