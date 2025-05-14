package common

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func (c *Common) TempDir() string {
	// HACK: Temp solution for the go stdlib bug.
	//
	// Copied from https://github.com/golang/go/blob/master/src/testing/testing.go#L1307
	//
	// testing.T.TempDir() will fail if test name is too long.
	//
	// See this issue https://github.com/golang/go/issues/71742
	//
	// It was fixed recently (as of 14.05.2025) by the go team https://go.dev/cl/671577
	// but it will require some time until it reaches the stable version.

	// Use a single parent directory for all the temporary directories
	// created by a test, each numbered sequentially.
	c.tempDirMu.Lock()

	var nonExistent bool

	if c.tempDir == "" { // Usually the case with js/wasm
		nonExistent = true
	} else {
		_, err := os.Stat(c.tempDir)
		nonExistent = os.IsNotExist(err)

		if err != nil && !nonExistent {
			c.Fatalf("TempDir: %v", err)
		}
	}

	if nonExistent {
		c.Helper()
		pattern := c.Name()

		// Limit length of file names on disk.
		// Invalid runes from slicing are dropped by strings.Map below.
		if len(pattern) > 64 {
			pattern = pattern[:64]
		}

		// Drop unusual characters (such as path separators or
		// characters interacting with globs) from the directory name to
		// avoid surprising os.MkdirTemp behavior.
		mapper := func(r rune) rune {
			if r < utf8.RuneSelf {
				const allowed = "!#$%&()+,-.=@^_{}~ "

				if '0' <= r && r <= '9' ||
					'a' <= r && r <= 'z' ||
					'A' <= r && r <= 'Z' {
					return r
				}

				if strings.ContainsRune(allowed, r) {
					return r
				}
			} else if unicode.IsLetter(r) || unicode.IsNumber(r) {
				return r
			}

			return -1
		}

		pattern = strings.Map(mapper, pattern)
		c.tempDir, c.tempDirErr = os.MkdirTemp("", pattern)

		if c.tempDirErr == nil {
			c.Cleanup(func() {
				if err := os.RemoveAll(c.tempDir); err != nil {
					c.Errorf("TempDir RemoveAll cleanup: %v", err)
				}
			})
		}
	}

	if c.tempDirErr == nil {
		c.tempDirSeq++
	}

	seq := c.tempDirSeq

	c.tempDirMu.Unlock()

	if c.tempDirErr != nil {
		c.Fatalf("TempDir: %v", c.tempDirErr)
	}

	dir := fmt.Sprintf("%s%c%03d", c.tempDir, os.PathSeparator, seq)

	//nolint:gosec // copied from the stdlib
	if err := os.Mkdir(dir, 0o777); err != nil {
		c.Fatalf("TempDir: %v", err)
	}

	return dir
}
