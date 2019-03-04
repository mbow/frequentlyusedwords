package frequentlyusedwords_test

import (
	"frequentlyUsedWords/internal/frequentlyusedwords"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_matchesGoldenFile(t *testing.T) {
	file, err := os.Open("test_fixtures/golden")
	require.NoError(t, err, "test did not fail helper did - open")
	expected, err := ioutil.ReadAll(file)
	require.NoError(t, err, "test did not fail helper did - readall")

	output, err := frequentlyusedwords.ReadFile("test_fixtures/mobydick.txt")
	require.NoError(t, err, "test failed")
	require.Equal(t, expected, output, "output does not match golden file")
}

func Test_HandlesBinaryFiles(t *testing.T) {
	_, err := frequentlyusedwords.ReadFile("test_fixtures/notTextBinary")
	require.EqualError(t, err, "file was not charset=utf-8, was application/octet-stream, exit")
}
