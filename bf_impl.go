package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
)

func unsupported(pos int, op string){
  log.Fatalf("operation unsupported at pos %d: %s", pos, op)
}

func parse(b []byte){
  var mem [30000]int
  i := 0
  for idx, symbol := range b {
    switch symbol {
      case '+': mem[i] += 1
      case '-': mem[i] -= 1
      case '>': i += 1
      case '<': i -= 1
      case '.': fmt.Printf("%v", string(mem[i]))
      case ',': unsupported(idx, ",")
      case '[': unsupported(idx, "[")
      case ']': unsupported(idx, "]")
      case '\n': // Newline encountered. Just ignore it.
      default:
        log.Fatalf("unknown symbol encountered at pos %d: %s", idx, string(symbol))
    }
  }
}

func main(){
  reader := bufio.NewReader(os.Stdin)
  t, _ := reader.ReadBytes('\n')
  parse(t)
}
