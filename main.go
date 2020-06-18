package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  args := os.Args[1:]

  // Put all args in one string
  msg := strings.Join(args, " ")

  // All to lower case
  msg = strings.ToLower(msg)

  // The author is supposed to be after "by"
  byPosition := strings.Index(msg, " by ")
  author := msg[byPosition+4:]
  author = strings.ReplaceAll(author, " ", "_")

  // The title with what comes before "by"
  title := msg[:byPosition]
  title = strings.ReplaceAll(title, " ", "-")

  // Concatenate everything
  var out strings.Builder
  out.WriteString(author)
  out.WriteString("_")
  out.WriteString(title)

  // Print out the result
  fmt.Println(out.String())
}