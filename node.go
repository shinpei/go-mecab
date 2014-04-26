package mecab

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>
#include <stdlib.h>

static mecab_node_t *trans_mecab_node_t (struct mecab_node_t *node){
  return node;
}
*/
import "C"

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
	return C.GoStringN(this.ptr.surface, C.int(this.ptr.length));
}

func(this *Node) GetFeature() string {
  return C.GoString(this.ptr.feature);
}

func (this *Node) GetId() int {
	return int(this.ptr.id)
}

func (this *Node) GetLength() int {
	return int(this.ptr.length)
}

func (this *Node) GetRlength() int {
	return int(this.ptr.rlength)
}

func (this *Node) GetPosid() int {
	return int(this.ptr.posid)
}

// TODO: map unsigned char
func (this *Node) GetCharType() int {
	return int(this.ptr.char_type)
}

//TODO: map unsigned char
func (this *Node) GetStat() int {
	return int(this.ptr.stat)
}

//TODO: map unsigned char
func (this *Node) GetIsbest() int {
	return int(this.ptr.isbest)
}

func (this *Node) GetAlpha() float32 {
	return float32(this.ptr.alpha)
}

func (this *Node) GetBeta() float32 {
	return float32(this.ptr.beta)
}

func (this *Node) GetProb() float32 {
	return float32(this.ptr.prob)
}

func (this *Node) GetWcost() int {
	return int(this.ptr.wcost)
}

func (this *Node) GetCost() int {
	return int(this.ptr.cost)
}
