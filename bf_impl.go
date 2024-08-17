package main

import (
  "fmt"
  "os"
  "bufio"
)

func unsupported(pos int, op string) error {
  return fmt.Errorf("operation unsupported at pos %d: %s", pos, op)
}

func parse(b []byte) ([30000]int, error){
  var mem [30000]int
  i := 0
  for idx, symbol := range b {
    if i >= len(mem) || i < 0 {
      return mem, fmt.Errorf("out of memory at pos %d", idx)
    }
    switch symbol {
      case '+': mem[i] += 1
      case '-': mem[i] -= 1
      case '>': i += 1
      case '<': i -= 1
      case '.': fmt.Printf("%s", string(rune(mem[i])))
      case ',': return mem, unsupported(idx, ",")
      case '[': return mem, unsupported(idx, "[")
      case ']': return mem, unsupported(idx, "]")
      case '\n': // Newline encountered. Just ignore it.
      default:
        return mem, fmt.Errorf("unknown symbol encountered at pos %d: %s", idx, string(rune(symbol)))
    }
  }
  return mem, nil
}

func main(){
  reader := bufio.NewReader(os.Stdin)
  t, _ := reader.ReadBytes('\n')
  _, err := parse(t)
  if err != nil {
    fmt.Println(err.Error())
  }
}
