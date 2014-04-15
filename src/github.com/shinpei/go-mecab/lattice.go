package mecab

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_node_t *trans_mecab_node_t (struct mecab_node_t *node){
  return node;
}
*/
import "C"

type Lattice struct {
	ptr *C.mecab_lattice_t
}

// ============================================
// Getter/ Setter

func (this *Lattice) GetBosNode() *Node {
	return &Node{ptr: C.mecab_lattice_get_bos_node(this.ptr)}
}

func (this *Lattice) GetEosNode() *Node {
	return &Node{ptr: C.mecab_lattice_get_eos_node(this.ptr)}
}

func (this *Lattice) GetSentence() string {
	return C.GoString(C.mecab_lattice_get_sentence(this.ptr))
}

func (this *Lattice) GetSize() int {
	return int(C.mecab_lattice_get_size(this.ptr))
}

func (this *Lattice) GetZ() float64 {
	return float64(C.mecab_lattice_get_z(this.ptr))
}

func (this *Lattice) SetZ(z float64) {
	C.mecab_lattice_set_z(this.ptr, C.double(z))
}

func (this *Lattice) GetTheta() float64 {
	return float64(C.mecab_lattice_get_theta(this.ptr))
}

func (this *Lattice) SetTheta(theta float64) {
	C.mecab_lattice_set_theta(this.ptr, C.double(theta))
}

func (this *Lattice) GetNext() int {
	return int(C.mecab_lattice_next(this.ptr))
}

func (this *Lattice) GetRequestType() int {
	return int(C.mecab_lattice_get_request_type(this.ptr))
}

// ============================================
// Methods

//TODO: int --> boolean
func (this *Lattice) HasRequestType(requestType int) int {
	return int(C.mecab_lattice_has_request_type(this.ptr, C.int(requestType)))
}

func (this *Lattice) SetRequestType(requestType int) {
	C.mecab_lattice_set_request_type(this.ptr, C.int(requestType))
}

func (this *Lattice) AddRequestType(requestType int) {
	C.mecab_lattice_add_request_type(this.ptr, C.int(requestType))
}

func (this *Lattice) RemoveRequestType(requestType int) {
	C.mecab_lattice_remove_request_type(this.ptr, C.int(requestType))
}

func (this *Lattice) ToString() string {
	return C.GoString(C.mecab_lattice_tostr(this.ptr))
}

func (this *Lattice) EnumNBestAsString(n int) string {
	return C.GoString(C.mecab_lattice_nbest_tostr(this.ptr, C.size_t(n)))
}

//TODO: btoi need

func (this *Lattice) HasConstraint() int {
	return int(C.mecab_lattice_has_constraint(this.ptr))
}

//TODO: better to fix method name as C++ style?
func (this *Lattice) Delete() {
	C.mecab_lattice_destroy(this.ptr)
}

func (this *Lattice) Clear() {
	C.mecab_lattice_clear(this.ptr)
}

func (this *Lattice) IsAvailable() int {
	return int(C.mecab_lattice_is_available(this.ptr))
}
