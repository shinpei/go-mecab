package main

import "github.com/shinpei/mecab"
import "fmt"

func main() {
	fmt.Println("version: " + mecab.GetVersion())

	tagger := mecab.Create()

	s := "こんにちは、僕は山田太郎です。"
	fmt.Println(s)
	result := tagger.Parse(s)
	fmt.Println(result)

  node := tagger.ParseToNode(s);
  fmt.Println(node);
	tagger.Destroy()
}
