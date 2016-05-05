package main

import (
  "bufio"
  "os"
  "math/rand"
  "bytes"
  "time"
  "strconv"
  "strings"
)

//go:generate go run scripts/includetxt.go

var (
  fileCache map[string][]string
  random = rand.New(rand.NewSource(time.Now().UnixNano()))
)


// GenerateName generates a Random Name with adjectiveAmount Adjectives, nounAmount Nouns and a random Number with randomNumberPlaces places.
func GenerateName(adjectiveAmount int, nounAmount int, randomNumberPlaces int) (string, error) {
  var nameBuffer bytes.Buffer

  for i := 0; i < adjectiveAmount; i++ {
    adj, err := getRandomLineFromString(adjectives)
    if err != nil {
      return "", err
    }
    nameBuffer.WriteString(adj)
  }

  for i := 0; i < nounAmount; i++ {
    noun, err := getRandomLineFromString(nouns)
    if err != nil {
      return "", err
    }
    nameBuffer.WriteString(noun)
  }

  for i := 0; i < randomNumberPlaces; i++ {
    nameBuffer.WriteString(strconv.Itoa(random.Intn(10)))
  }

  return nameBuffer.String(), nil
}

// GenerateNameWithFiles generates a Random Name with adjectiveAmount Adjectives, nounAmount Nouns and a random Number with randomNumberPlaces places.
// You can use custom Files with this function.
func GenerateNameWithFiles(adjectiveAmount int, nounAmount int, randomNumberPlaces int, adjectivesFile string, nounsFile string) (string, error) {
  var nameBuffer bytes.Buffer

  for i := 0; i < adjectiveAmount; i++ {
    adj, err := getRandomLineFromFile(adjectivesFile)
    if err != nil {
      return "", err
    }
    nameBuffer.WriteString(adj)
  }

  for i := 0; i < nounAmount; i++ {
    noun, err := getRandomLineFromFile(nounsFile)
    if err != nil {
      return "", err
    }
    nameBuffer.WriteString(noun)
  }

  for i := 0; i < randomNumberPlaces; i++ {
    nameBuffer.WriteString(strconv.Itoa(random.Intn(10)))
  }

  return nameBuffer.String(), nil
}

func getRandomLineFromString(data string) (string, error) {
  lines := readLinesString(data)

  line := lines[random.Intn(len(lines))]
  for line == "" {
    line = lines[random.Intn(len(lines))]
  }

  return line, nil
}

func getRandomLineFromFile(path string) (string, error) {
  lines, err := readFile(path)

  if err != nil {
    return "", err
  }

  return lines[random.Intn(len(lines))], nil
}

func readLinesString(data string) ([]string) {
  return strings.Split(data, "\n")
}

func readFile(path string) ([]string, error) {
  if fileCache == nil {
    fileCache = make(map[string][]string)
  }
  if fileCache[path] != nil {
    return fileCache[path], nil
  }

  inFile, _ := os.Open(path)
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  fileCache[path] = lines

  return lines, nil
}
