package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

type TestCase struct {
	offset   int64
	limit    int64
	testFile string
}

func TestCopy(t *testing.T) {
	cases := []TestCase{
		{
			offset:   0,
			limit:    0,
			testFile: "out_offset0_limit0.txt",
		},
		{
			offset:   0,
			limit:    10,
			testFile: "out_offset0_limit10.txt",
		},
		{
			offset:   0,
			limit:    1000,
			testFile: "out_offset0_limit1000.txt",
		},
		{
			offset:   0,
			limit:    10000,
			testFile: "out_offset0_limit10000.txt",
		},
		{
			offset:   100,
			limit:    1000,
			testFile: "out_offset100_limit1000.txt",
		},
		{
			offset:   6000,
			limit:    1000,
			testFile: "out_offset6000_limit1000.txt",
		},
	}
	input := "testdata/input.txt"
	for _, f := range cases {
		file, err := ioutil.TempFile("", "")
		require.NoError(t, err)

		t.Run(f.testFile, func(t *testing.T) {
			err := Copy(input, file.Name(), f.offset, f.limit)
			require.NoError(t, err)

			copied, err := ioutil.ReadFile(file.Name())
			require.NoError(t, err)

			fixture, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s", f.testFile))
			require.NoError(t, err)

			require.Equal(t, fixture, copied)
		})
	}

	t.Run("dev/null", func(t *testing.T) {
		err := Copy(input, "/dev/null", 0, 0)
		require.NoError(t, err)
	})
}

func TestCopyInvalid(t *testing.T) {
	t.Run("Offset exceeded the file size", func(t *testing.T) {
		input := "testdata/input.txt"
		fileStat, err := os.Stat(input)
		require.NoError(t, err)

		err = Copy(input, "/dev/null", fileStat.Size()+1, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})
}
