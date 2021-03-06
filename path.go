package mecab

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_node_t *trans_mecab_node_t (struct mecab_node_t *node){
  return node;
}
*/
import "C"

type Path struct {
	ptr *C.mecab_path_t
}

func (this *Path) cost() int {
	return int(this.ptr.cost)
}

func (this *Path) prob() float32 {
	return float32(this.ptr.prob)
}
