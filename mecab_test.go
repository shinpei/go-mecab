package mecab

import (
       "fmt"
       "testing"
)

/*
Go-mecab version of the example code below,
http://mecab.googlecode.com/svn/trunk/mecab/doc/libmecab.html
*/

func Test_Mecab(t *testing.T) {
	fmt.Println("version: " + GetVersion())

	tagger := Create()
	fmt.Println("####### given parameter str ########")
	s := "太郎は次郎が持っている本を花子に渡した。"
	fmt.Println("str=" + s)

	fmt.Println("####### Parse ########")
	result := tagger.Parse(s)
	fmt.Println(result)

	fmt.Println("####### Node.Next ########")
	node := tagger.ParseToNode(s)
	fmt.Print("has next:")
	fmt.Println(node.HasNext())
	for i := 0; node.HasNext(); i++ {
		node.Next()
		fmt.Printf("%s, ", node.GetSurface());
		fmt.Printf("%s, ", node.GetFeature());
		fmt.Printf("%d, ", node.GetId())
		fmt.Printf("%d, ", node.GetLength())
		fmt.Printf("%d, ", node.GetRlength())
		fmt.Printf("%d, ", node.GetPosid())
		fmt.Printf("%d, ", node.GetCharType())
		fmt.Printf("%d, ", node.GetStat())
		fmt.Printf("%d, ", node.GetIsbest())
		fmt.Printf("%f, ", node.GetAlpha())
		fmt.Printf("%f, ", node.GetBeta())
		fmt.Printf("%f, ", node.GetProb())
		fmt.Printf("%d, ", node.GetWcost())
		fmt.Printf("%d\n", node.GetCost())

	}

	fmt.Println("####### ParseNBest(3, str) ########")
	fmt.Println(tagger.ParseNBest(3, s))

	fmt.Println("####### ParseNBestInit(str) ########")
	ret := tagger.ParseNBestInit(s)
	fmt.Printf("nbest=%d\n", ret)
	for i := 0; i < 3; i++ {
		fmt.Printf("i:%d %s", i, tagger.Next())
	}

	fmt.Println("####### Mecab.Next ########")
	fmt.Println(tagger.Next())
	fmt.Println(tagger.Next())
	fmt.Println(tagger.Next())

	fmt.Println("####### DictionaryInfo ########")
	dict := tagger.GetDictionaryInfo()
	fmt.Printf("filename:%s\n", dict.GetFilename())
	fmt.Printf("charset:%s\n", dict.GetCharset())
	fmt.Printf("size:%d\n", dict.GetSize())
	fmt.Printf("type:%d\n", dict.GetType())
	fmt.Printf("lsize:%d\n", dict.GetLsize())
	fmt.Printf("rsize:%d\n", dict.GetRsize())
	fmt.Printf("version:%d\n", dict.GetVersion())

	for i := 0; node.HasNext(); i++ {
		dict.Next()
		fmt.Printf("filename:%s\n", dict.GetFilename())
		fmt.Printf("charset:%s\n", dict.GetCharset())
		fmt.Printf("size:%d\n", dict.GetSize())
		fmt.Printf("type:%d\n", dict.GetType())
		fmt.Printf("lsize:%d\n", dict.GetLsize())
		fmt.Printf("rsize:%d\n", dict.GetRsize())
		fmt.Printf("version:%d\n", dict.GetVersion())
	}

	fmt.Println("####### Constants  ########")
	fmt.Printf("MECAB_SYS_DIC=%d\n", MECAB_SYS_DIC)
	fmt.Printf("MECAB_USR_DIC=%d\n", MECAB_USR_DIC)
	fmt.Printf("MECAB_UNK_DIC=%d\n", MECAB_UNK_DIC)

	tagger.Destroy()
}
