package main

import (
  "testing"
  "strings"
)

func TestParseUnknownSymbol(t *testing.T){
  _, err := parse([]byte("?"))
  if err == nil {
    t.Errorf("\"Unknown symbol\" error missing")
  }
}
func TestParseUnsupportedSymbol(t *testing.T){
  _, err := parse([]byte("[]"))
  if err == nil {
    t.Errorf("\"Unsupported symbol\" error missing")
  }
}
func TestParseNull(t *testing.T){
  _, err := parse([]byte(""))
  if err != nil {
    t.Errorf("Got an error when passing no data")
  }
}
func TestParseOutOfMemory(t *testing.T){
  _, err := parse([]byte(strings.Repeat(">", 30000 + 1)))
  if err == nil {
    t.Errorf("No error when running out of memory")
  }
}
func TestParseNegativePointer(t *testing.T){
  _, err := parse([]byte(strings.Repeat("<", 30000 + 1)))
  if err == nil {
    t.Errorf("Negative pointer allowed")
  }
}
