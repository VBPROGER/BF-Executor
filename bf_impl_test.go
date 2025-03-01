package main

import (
  "testing"
  "strings"
)

func TestParseUnknownSymbol(t *testing.T){
  _, _, err, _ := parse([]byte("?"))
  if err == nil {
    t.Errorf("\"Unknown symbol\" error missing")
  }
}
func TestParseUnsupportedSymbol(t *testing.T){
  _, _, err, _ := parse([]byte("[]"))
  if err == nil {
    t.Errorf("\"Unsupported symbol\" error missing")
  }
}
func TestParseNull(t *testing.T){
  _, _, err, _ := parse([]byte(""))
  if err != nil {
    t.Errorf("Got an error when passing no data")
  }
}
func TestParseOutOfMemory(t *testing.T){
  _, _, err, _ := parse([]byte(strings.Repeat(">", MEMORY_SIZE + 1)))
  if err == nil {
    t.Errorf("No error when running out of memory")
  }
}
func TestParseNegativePointer(t *testing.T){
  _, _, err, _ := parse([]byte(strings.Repeat("<", MEMORY_SIZE + 1)))
  if err == nil {
    t.Errorf("Negative pointer allowed")
  }
}
