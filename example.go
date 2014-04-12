package main

import "github.com/shinpei/mecab"
import "fmt"

func main() {
	fmt.Println("version: " + mecab.GetVersion())

	tagger := mecab.Create()
	fmt.Println("####### given parameter str ########");
	s := "太郎は次郎が持っている本を花子に渡した。"
	fmt.Println("str=" + s)

	fmt.Println("####### Parse ########");
	result := tagger.Parse(s)
	fmt.Println(result)

  node := tagger.ParseToNode(s);
	fmt.Println("####### Node.Next ########");
	fmt.Print("has next:");
	fmt.Println(node.HasNext());
	for i:=0;  node.HasNext(); i++ {
		node.Next();
		fmt.Println(node.GetSurface());
	}
	fmt.Println("####### ParseNBest(3, str) ########");
	fmt.Println(tagger.ParseNBest(3, s));

	fmt.Println("####### Mecab.Next ########");
	fmt.Println(tagger.Next());
	fmt.Println(tagger.Next());
	fmt.Println(tagger.Next());


	tagger.Destroy();
}
