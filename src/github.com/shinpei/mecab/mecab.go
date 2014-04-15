package mecab

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_node_t *trans_mecab_node_t (struct mecab_node_t *node){
  return node;
}
*/
import "C"

// ============================================
// Types

type DictionaryInfo struct {
	ptr *C.mecab_dictionary_info_t
}

type Lattice struct {
	ptr *C.mecab_lattice_t
}

// ============================================
// instance methods

func Create() *Tagger {
	taggerPtr := new(Tagger)
	emptyStringPtr := C.CString("")
	taggerPtr.ptr = C.mecab_new2(emptyStringPtr)
	return taggerPtr
}

// Dictionary
const (
	MECAB_SYS_DIC = C.MECAB_SYS_DIC
	MECAB_USR_DIC = C.MECAB_USR_DIC
	MECAB_UNK_DIC = C.MECAB_UNK_DIC
)

// ============================================
// this is static anyway

func GetVersion() string {
	p := C.mecab_version()
	return C.GoString(p)
}

// ============================================
