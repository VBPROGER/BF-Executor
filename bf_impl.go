package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
)

const MEMORY_SIZE = 30000
const NULL_IDX = -1

func unsupported(op string) error {
  return fmt.Errorf("operation unsupported: %s", op)
}

func read(what byte)([]byte){
  reader := bufio.NewReader(os.Stdin)
  bytes, _ := reader.ReadBytes(what)
  return bytes
}

func parse(b []byte) ([MEMORY_SIZE]int, int, error, int){
  var mem [MEMORY_SIZE]int
  i := 0
  for idx, symbol := range b {
    if i >= len(mem) || i < 0 {
      return mem, i, fmt.Errorf("out of memory"), idx
    }
    switch symbol {
      case '+': mem[i] += 1
      case '-': mem[i] -= 1
      case '>': i += 1
      case '<': i -= 1
      case '.': fmt.Printf("%s", string(rune(mem[i])))
      case ',':
        scanner := bufio.NewScanner(os.Stdin)
        if scanner.Scan(){
          mem[i] = int(scanner.Text()[0])
        }
      case '[': return mem, i, unsupported("["), idx
      case ']': return mem, i, unsupported("]"), idx
      case '\n': // Newline encountered. Just ignore it.
      default:
        return mem, i, fmt.Errorf("unknown symbol encountered: %s", string(rune(symbol))), idx
    }
  }
  return mem, i, nil, NULL_IDX
}

func main(){
  t := read('\n')
  _, i, err, pos := parse(t)
  if err != nil {
    log.Fatalf("Error at position %d in slot %d:\n%s\n", pos, i, err.Error())
  }
}
