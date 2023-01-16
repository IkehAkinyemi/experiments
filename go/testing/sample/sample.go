package main

import (
	"fmt"
	"io/fs"
	"strings"
)

func readWordsFromFile(f fs.FS, path string) ([]string, error) {
  content, err := fs.ReadFile(f, path)
  if err != nil {
    return nil, fmt.Errorf("read file: %v", err)
  }
  
  return strings.Fields(string(content)), nil
}
