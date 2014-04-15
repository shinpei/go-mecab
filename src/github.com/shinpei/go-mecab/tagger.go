package mecab

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

*/
import "C"

type Tagger struct {
	ptr *C.mecab_t
}

// ============================================
// Getter/Setter

func (this *Tagger) GetTheta () float32 {
  return float32(C.mecab_get_theta(this.ptr));
}

func (this *Tagger) SetTheta(theta float32) {
  C.mecab_set_theta(this.ptr, C.float(theta));
}

func (this *Tagger) GetLatticeLevel() int {
  return int(C.mecab_get_lattice_level(this.ptr));
}

func (this *Tagger) SetLatticeLevel(level int) {
  C.mecab_set_lattice_level(this.ptr, C.int(level));
}

func (this *Tagger) GetAllMorphs() int {
  return int(C.mecab_get_all_morphs(this.ptr));
}

func (this *Tagger) SetAllMorphs(morphs int) {
  C.mecab_set_all_morphs(this.ptr, C.int(morphs));
}

// ============================================
// Methods


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
	return &Node {
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

func (this *Tagger) GetDictionaryInfo() *DictionaryInfo {
  dic_ptr := C.mecab_dictionary_info(this.ptr);
  return &DictionaryInfo{ptr: dic_ptr};
}
