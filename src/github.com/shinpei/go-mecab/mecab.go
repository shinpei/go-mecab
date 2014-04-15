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
// instance methods

func Create() *Tagger {
	emptyStringPtr := C.CString("")
	return &Tagger{ptr: C.mecab_new2(emptyStringPtr)}
}

func CreateLattice() *Lattice {
	return &Lattice{ptr: C.mecab_lattice_new()}
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
