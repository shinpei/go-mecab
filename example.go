package main

import "github.com/shinpei/go-mecab"
import "fmt"

/*
Go-mecab version of the example code below,
http://mecab.googlecode.com/svn/trunk/mecab/doc/libmecab.html
*/

func main() {
	fmt.Println("version: " + mecab.GetVersion())

	tagger := mecab.Create()
	fmt.Println("####### given parameter str ########");
	s := "太郎は次郎が持っている本を花子に渡した。"
	fmt.Println("str=" + s)

	fmt.Println("####### Parse ########");
	result := tagger.Parse(s)
	fmt.Println(result)

	fmt.Println("####### Node.Next ########");
  node := tagger.ParseToNode(s);
	fmt.Print("has next:");
	fmt.Println(node.HasNext());
	for i:=0;  node.HasNext(); i++ {
		node.Next();
		fmt.Printf("%s, ", node.GetSurface());
		fmt.Printf("%d, ", node.GetId());
		fmt.Printf("%d, ", node.GetLength());
		fmt.Printf("%d, ", node.GetRlength());
		fmt.Printf("%d, ", node.GetPosid());
		fmt.Printf("%d, ", node.GetCharType());
		fmt.Printf("%d, ", node.GetStat());
		fmt.Printf("%d, ", node.GetIsbest());
		fmt.Printf("%f, ", node.GetAlpha());
		fmt.Printf("%f, ", node.GetBeta());
		fmt.Printf("%f, ", node.GetProb());
		fmt.Printf("%d, ", node.GetWcost());
		fmt.Printf("%d\n", node.GetCost());

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

	fmt.Println("####### DictionaryInfo ########");
	dict := tagger.GetDictionaryInfo();
	fmt.Println(dict.GetFilename());


	fmt.Println("####### Constants  ########");
	fmt.Printf("MECAB_SYS_DIC=%d\n", mecab.MECAB_SYS_DIC);
	fmt.Printf("MECAB_USR_DIC=%d\n", mecab.MECAB_USR_DIC);
	fmt.Printf("MECAB_UNK_DIC=%d\n", mecab.MECAB_UNK_DIC);
	tagger.Destroy();
}
