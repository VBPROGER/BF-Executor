package main

import (
  "fmt"
  "os"
  "bufio"
)

const MEMORY_SIZE = 30000
const NULL_IDX = -1

func unsupported(pos int, op string) error {
  return fmt.Errorf("operation unsupported at pos %d: %s", pos, op)
}

func parse(b []byte) ([MEMORY_SIZE]int, int, error, int){
  var mem [MEMORY_SIZE]int
  i := 0
  for idx, symbol := range b {
    if i >= len(mem) || i < 0 {
      return mem, i, fmt.Errorf("out of memory at pos %d", idx), idx
    }
    switch symbol {
      case '+': mem[i] += 1
      case '-': mem[i] -= 1
      case '>': i += 1
      case '<': i -= 1
      case '.': fmt.Printf("%s", string(rune(mem[i])))
      case ',': return mem, i, unsupported(idx, ","), idx
      case '[': return mem, i, unsupported(idx, "["), idx
      case ']': return mem, i, unsupported(idx, "]"), idx
      case '\n': // Newline encountered. Just ignore it.
      default:
        return mem, i, fmt.Errorf("unknown symbol encountered at pos %d: %s", idx, string(rune(symbol))), idx
    }
  }
  return mem, i, nil, NULL_IDX
}

func main(){
  reader := bufio.NewReader(os.Stdin)
  t, _ := reader.ReadBytes('\n')
  _, _, err, _ := parse(t)
  if err != nil {
    fmt.Printf("An error has occured while executing the script:\n%s\n", err.Error())
  }
}
