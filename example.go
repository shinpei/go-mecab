package main

import "github.com/shinpei/mecab"
import "fmt"

func main() {
	fmt.Println("version: " + mecab.GetVersion())

	tagger := mecab.Create()
	fmt.Println("####### given parameter str ########");
	s := "こんにちは、僕は山田太郎です。"
	fmt.Println("str=" + s)

	fmt.Println("####### Parse ########");
	result := tagger.Parse(s)
	fmt.Println(result)

	fmt.Println("####### ParseNode ########");
  node := tagger.ParseToNode(s);
	fmt.Println(node);

	fmt.Println("####### ParseNBest(3, str) ########");
	fmt.Println(tagger.ParseNBest(3, s));
	tagger.Destroy()
}
