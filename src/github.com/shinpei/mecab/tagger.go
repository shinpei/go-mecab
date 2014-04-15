package mecab

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_node_t *trans_mecab_node_t (struct mecab_node_t *node){
  return node;
}
*/
import "C"

type Tagger struct {
	ptr *C.mecab_t
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

func (this *Tagger) ParseNBestInit(target string) int {
	target_ptr := C.CString(target)
	return int(C.mecab_nbest_init(this.ptr, target_ptr))
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
