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

func (this *Lattice) GetBosNode () *Node {
  return &Node {ptr: C.mecab_lattice_get_bos_node(this.ptr)};
}

func (this *Lattice) GetEosNode () *Node {
  return &Node {ptr: C.mecab_lattice_get_eos_node(this.ptr)};
}

func (this *Lattice) GetSentence () string {
  return C.GoString(C.mecab_lattice_get_sentence(this.ptr));
}

func (this *Lattice) GetSize() int {
  return int(C.mecab_lattice_get_size(this.ptr));
}

func (this *Lattice) GetZ() float64 {
  return float64(C.mecab_lattice_get_z(this.ptr));
}

func (this *Lattice) SetZ(z float64) {
  C.mecab_lattice_set_z(this.ptr, C.double(z));
}
// ============================================
// Methods

//TODO: better to fix method name as C++ style?
func (this *Lattice) Delete () {
  C.mecab_lattice_destroy(this.ptr);
}

func (this *Lattice) Clear () {
  C.mecab_lattice_clear(this.ptr);
}

func (this *Lattice) IsAvailable () int {
  return int (C.mecab_lattice_is_available(this.ptr));
}
