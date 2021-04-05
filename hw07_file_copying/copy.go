package main

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"math"
	"os"
)

var (
	// ErrUnsupportedFile is happening when file is not supported.
	ErrUnsupportedFile = errors.New("unsupported file")
	// ErrOffsetExceedsFileSize is happening when offset exceeds the file size.
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

// Copy is analog of regular cp.
func Copy(fromPath, toPath string, offset, limit int64) error {
	srcStat, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if offset > srcStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 {
		limit = srcStat.Size() - offset
	}

	src, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer src.Close()

	dst, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("could not create file, %w", err)
	}
	defer dst.Close()

	_, err = src.Seek(offset, 0)
	if err != nil {
		return ErrOffsetExceedsFileSize
	}

	maxBar := limit
	bar := pb.Full.Start64(maxBar)
	srcReader := bar.NewProxyReader(src)

	buf := make([]byte, int(math.Min(float64(1024), float64(limit))))
	for {
		n, err := srcReader.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		_, err = dst.Write(buf[:n])
		if err != nil {
			return err
		}

		maxBar -= int64(n)
		if maxBar <= 0 {
			break
		}
	}

	bar.SetCurrent(limit)
	bar.Finish()

	return err
}
