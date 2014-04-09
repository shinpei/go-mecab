# go-mecab: low-level mecab bindings for Go


go-mecab provides simple low-level bindings for [mecab](http://mecab.googlecode.com/svn/trunk/mecab/doc/index.html), japanese morphological analyzer.

## Install
1. Please fetch, build and install mecab and mecab-ipa-dictionary.
2. Modify library path in `src/github.com/shinpei/mecab/mecab.go`

## Build

```csh
$ go bulid github.com/shinpei/mecab
$ go install github.com/shinpei/mecab
$ go build // to build exmaple
$ ./go-mecab // to execute example
```
## Example

```go
 import "github.com/shinpei/mecab"

 tagger := mecab.Create();
 str := "こんにちは、世界";
 ret := tagger.Parse(str);
 fmt.Println(ret);
```

## Lisences, contact info, contribute
It's under [ASL2.0](http://www.apache.org/licenses/LICENSE-2.0). If you find bug or improvement request, please contact me through twitter, @shinpeintk. And always welcoming heartful pull request.

cheers, :sake: :moyai:

