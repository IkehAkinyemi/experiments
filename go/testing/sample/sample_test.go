package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func Test_readWordsFromFile(t *testing.T) {
  const name = "some/path"
  tfs := fstest.MapFS{name: &fstest.MapFile{
    Data: []byte("one\ntwo\tthree four      five"),
  }}
  
  words, err := readWordsFromFile(tfs, name)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  
  expected := []string{"one", "two", "three", "four", "five"}
  if !reflect.DeepEqual(words, expected) {
    t.Fatalf("result %+v != expected %+v", words, expected)
  }
}