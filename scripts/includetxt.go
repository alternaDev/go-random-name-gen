package main

import (
    "bufio"
    "io/ioutil"
    "os"
    "strings"
    "strconv"
)

// Reads all .txt files in the current folder
// and encodes them as strings arrays line by line in textfiles.go
func main() {
    fs, _ := ioutil.ReadDir(".")
    out, _ := os.Create("textfiles.go")
    out.Write([]byte("package main \n\nvar (\n"))
    for _, f := range fs {
        if strings.HasSuffix(f.Name(), ".txt") {
            lines, _ := readFile(f.Name())
            out.Write([]byte(strings.TrimSuffix(f.Name(), ".txt") + " = [" + strconv.Itoa(len(lines)) + "]string {"))
            for _, line := range lines {
              out.Write([]byte("`" + line + "`,"))
            }
            out.Write([]byte("}\n"))
        }
    }
    out.Write([]byte(")\n"))
}

func readFile(path string) ([]string, error) {
  inFile, _ := os.Open(path)
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, nil
}
