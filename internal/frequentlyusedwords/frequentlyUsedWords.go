package frequentlyusedwords

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"
)

// Filter function for a Scanner (based from bufio.ScanWords ) that returns each
// word modified to lower case and space-separated word of text a-z only, with
// surrounding spaces deleted. It will never return an empty string.
// The definition of space is set by unicode.IsSpace
func Filter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	data = bytes.ToLower(data)
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	// Scan until not a-z else take that as marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !(r >= 'a' && r <= 'z') {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func ReadFile(fileName string) ([]byte, error) {
	//open file and test content type so we know we can parse it
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "failed to read into buffer for content type check")
	}
	contentType := http.DetectContentType(buffer[:n])
	if !strings.Contains(contentType, "charset=utf-8") {
		return nil, fmt.Errorf("file was not charset=utf-8, was %v, exit", contentType)
	}
	if _, err := file.Seek(0, 0); err != nil {
		return nil, errors.Wrap(err, "failed to reset file after content check")
	}

	// scan the file and split words on a custom filter
	scanner := bufio.NewScanner(file)
	scanner.Split(Filter)

	//temp key map
	store := make(map[string]int)
	for scanner.Scan() {
		if scanner.Text() != "" {
			if val, ok := store[scanner.Text()]; ok {
				store[scanner.Text()] = val + 1
			} else {
				store[scanner.Text()] = 1
			}
		}
	}
	//false break loop check if EOF or error
	err = scanner.Err()
	if err == nil {
		log.Println("Scan completed and reached EOF")
	} else {
		return nil, err
	}

	//fmt.Println(store)
	type kv struct {
		Key   string
		Value int
	}

	ss := make([]kv, 0, 20)
	for k, v := range store {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 7, 0, 0, ' ', tabwriter.AlignRight)

	for i, kv := range ss {
		if i >= 20 {
			break
		}
		//this was a pita
		fmt.Fprintf(w, "%d\t %s\n", kv.Value, kv.Key)
	}
	w.Flush()
	return b.Bytes(), nil
}
