package main

import "github.com/shinpei/go-mecab"
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

	fmt.Println("####### ParseNBestInit(str) ########");
	ret := tagger.ParseNBestInit(s);
	fmt.Printf("nbest=%d\n", ret);
	for i:=0; i<3; i++ {
		fmt.Printf("i:%d %s", i, tagger.Next());
	}

	fmt.Println("####### Mecab.Next ########");
	fmt.Println(tagger.Next());
	fmt.Println(tagger.Next());
	fmt.Println(tagger.Next());

	fmt.Println("####### Constants  ########");
	fmt.Printf("MECAB_SYS_DIC=%d\n", mecab.MECAB_SYS_DIC);
	fmt.Printf("MECAB_USR_DIC=%d\n", mecab.MECAB_USR_DIC);
	fmt.Printf("MECAB_UNK_DIC=%d\n", mecab.MECAB_UNK_DIC);
	tagger.Destroy();
}
