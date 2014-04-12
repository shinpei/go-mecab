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

type Tagger struct {
	ptr *C.mecab_t
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

func (this *Tagger) Parse(target string) string {
	target_ptr := C.CString(target)
	result_ptr := C.mecab_sparse_tostr(this.ptr, target_ptr)
	return C.GoString(result_ptr)
}

func (this *Tagger) ParseNBest(size int, target string) string {
	target_ptr := C.CString(target)
	result_ptr := C.mecab_nbest_sparse_tostr(this.ptr, C.size_t(size), target_ptr)
	return C.GoString(result_ptr)
}

func (this *Tagger) ParseToNode(target string) *Node {
	target_ptr := C.CString(target)
	node_ptr := C.mecab_sparse_tonode(this.ptr, target_ptr)
	return &Node{
		ptr: node_ptr,
	}
}

func (this *Tagger) Next() string {
	result_ptr := C.mecab_nbest_next_tostr(this.ptr)
	return C.GoString(result_ptr)
}

func (this *Tagger) Destroy() {
	C.mecab_destroy(this.ptr)
}

// ============================================
// Node class

type Node struct {
	ptr *C.mecab_node_t
}

func (this *Node) Next() {
	ptr := C.trans_mecab_node_t(this.ptr.next)
	this.ptr = ptr
}

func (this *Node) HasNext() bool {
	return this.ptr.next != nil
}

func (this *Node) GetSurface() string {
  return C.GoString(this.ptr.surface);
}

func (this *Node) GetId() int {
  return int(this.ptr.id);
}




// ============================================
// this is static anyway

func GetVersion() string {
	p := C.mecab_version()
	return C.GoString(p)
}

// ============================================
