package mecab
/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_dictionary_info_t *trans_mecab_dictionary_t (struct mecab_dictionary_info_t *dict){
  return dict;
}
*/
import "C"

type DictionaryInfo struct {
  ptr *C.mecab_dictionary_info_t
}

func (this *DictionaryInfo) GetFilename() string {
  return C.GoString(this.ptr.filename);
}

func (this *DictionaryInfo) GetCharset() string {
  return C.GoString(this.ptr.charset);
}

func (this *DictionaryInfo) GetSize () int {
  return int (this.ptr.size);
}

func (this *DictionaryInfo) GetType() int {
  return int(this.ptr._type);
}

func (this *DictionaryInfo) GetLsize() int {
  return int (this.ptr.lsize);
}

func (this *DictionaryInfo) GetRsize() int {
  return int (this.ptr.rsize);
}

func (this *DictionaryInfo) GetVersion() int {
  return int(this.ptr.version);
}

//TODO: next
func (this *DictionaryInfo) GetNext() *DictionaryInfo {
  return &DictionaryInfo{ptr: C.trans_mecab_dictionary_t(this.ptr.next)};
}
