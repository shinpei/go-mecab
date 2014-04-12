package mecab

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>


static mecab_t * go_mecab_new (int ac, char ** av) {
  return mecab_new (ac, av);
}

static const char *go_mecab_strerror(mecab_t *mecab) {
  return mecab_strerror(mecab);
}

static int go_mecab_get_partial(mecab_t *mecab) {
  return mecab_get_partial(mecab);
}

static void go_mecab_set_partial(mecab_t *mecab, int partial) {
  mecab_set_partial(mecab, partial);
}

static float go_mecab_get_theta (mecab_t *mecab) {
  return mecab_get_theta(mecab);
}

static void go_mecab_set_theta(mecab_t *mecab, float theta) {
  mecab_set_theta(mecab, theta);
}

static int go_mecab_get_lattice_level(mecab_t* mecab) {
  return mecab_get_lattice_level(mecab);
}

static void go_mecab_set_lattice_level(mecab_t *mecab, int level) {
  mecab_set_lattice_level(mecab, level);
}

static int go_mecab_get_all_morphs (mecab_t *mecab) {
  return mecab_get_all_morphs(mecab);
}

static void go_mecab_set_all_morphs (mecab_t *mecab, int all_morphs) {
  mecab_set_all_morphs(mecab, all_morphs);
}

static int go_mecab_parse_lattice(mecab_t *mecab, mecab_lattice_t *lattice) {
  return mecab_parse_lattice(mecab, lattice);
}

static const char *go_mecab_sparse_tostr2(mecab_t *mecab, const char *str, size_t len) {
  return mecab_sparse_tostr2(mecab, str, len);
}

static const char *go_mecab_sparse_tostr3(mecab_t *mecab, const char *str, size_t len, char* ostr, size_t olen) {
  return mecab_sparse_tostr3(mecab, str, len, ostr, olen);
}

static const mecab_node_t *go_mecab_sparse_tonode(mecab_t *mecab, const char* str) {
  return mecab_sparse_tonode(mecab, str);
}

static const mecab_node_t *go_mecab_sparse_tonode2(mecab_t *mecab, const char* str, size_t len) {
  return mecab_sparse_tonode2(mecab, str, len);
}

static const char *go_mecab_nbest_sparse_tostr(mecab_t *mecab, size_t n, const char *str) {
  return mecab_nbest_sparse_tostr(mecab, n, str);
}

static const char *go_mecab_nbest_sparse_tostr2(mecab_t *mecab, size_t n, const char *str, size_t len) {
  return mecab_nbest_sparse_tostr2(mecab, n, str, len);
}

static const char *go_mecab_nbest_sparse_tostr3(mecab_t *mecab, size_t n, const char *str, size_t len, char* ostr, size_t olen) {
  return mecab_nbest_sparse_tostr3(mecab, n, str, len, ostr, olen);
}

static int go_mecab_nbest_init(mecab_t *mecab, const char *str) {
  return mecab_nbest_init(mecab, str);
}

static int go_mecab_nbest_init2 (mecab_t *mecab, const char *str, size_t len) {
  return mecab_nbest_init2(mecab, str, len);
}

static const char *go_mecab_nbest_next_tostr (mecab_t *mecab) {
  return mecab_nbest_next_tostr(mecab);
}

static char *go_mecab_nbest_next_tostr2 (mecab_t *mecab, char *ostr, size_t olen) {
  return mecab_nbest_next_tostr2(mecab, ostr, olen);
}

static const mecab_node_t *go_mecab_nbest_next_tonode (mecab_t *mecab) {
  return mecab_nbest_next_tonode(mecab);
}

static const char *go_mecab_format_node (mecab_t *mecab, const mecab_node_t *node) {
  return mecab_format_node(mecab, node);
}

static const mecab_dictionary_info_t *go_mecab_dictionary_info(mecab_t *mecab) {
  return mecab_dictionary_info(mecab);
}

static mecab_lattice_t *go_mecab_lattice_new () {
  return mecab_lattice_new();
}

static void go_mecab_lattice_destroy (mecab_lattice_t *lattice) {
  mecab_lattice_destroy(lattice);
}

static void go_mecab_lattice_clear(mecab_lattice_t *lattice) {
  mecab_lattice_clear(lattice);
}

static int go_mecab_lattice_is_available(mecab_lattice_t *lattice) {
  return mecab_lattice_is_available(lattice);
}

static mecab_node_t *go_mecab_lattice_get_bos_node(mecab_lattice_t *lattice) {
  return mecab_lattice_get_bos_node(lattice);
}

static mecab_node_t *go_mecab_lattice_get_eos_node(mecab_lattice_t *lattice) {
  return mecab_lattice_get_eos_node(lattice);
}

static mecab_node_t **go_mecab_lattice_get_all_begin_nodes (mecab_lattice_t *lattice) {
  return mecab_lattice_get_all_begin_nodes (lattice);
}

// casting typedef struct


*/
import "C"

type Dictionary struct {
	filename string
	charset  string
	size     int
	itype    int
	lsize    int
	rsize    int
	version  int
	next     *Dictionary
}

type Node struct {
	prev     *Node
	next     *Node
	enext    *Node
	bnext    *Node
	surface  string
	feature  string
	id       int
	length   int
	rlength  int
	rcAttr   int
	lcAttr   int
	posid    int
	charType int
	stat     int
	isbest   int
	alpha    float64
	beta     float64
	prob     float64
	wcost    int
	cost     int
}

type Tagger struct {
	tagger *C.mecab_t
}

// ============================================
// translation method

func TransCMecabNode2Go(node_ptr *C.mecab_node_t) *Node {
  if (node_ptr == nil) {
    return nil;
  }
	return &Node {
    prev: TransCStructMecabNode2Go(node_ptr.prev),
    next: TransCStructMecabNode2Go(node_ptr.next),

  }
}

func TransCStructMecabNode2Go (node_ptr *C.struct_mecab_node_t) *Node {
  if (node_ptr == nil) {
    return nil;
  }
  return &Node {
    prev: TransCStructMecabNode2Go(node_ptr.prev),
    next: TransCStructMecabNode2Go(node_ptr.next),
    enext: TransCStructMecabNode2Go(node_ptr.enext),
    bnext: TransCStructMecabNode2Go(node_ptr.bnext),
    surface: C.GoString(node_ptr.surface),
    feature: C.GoString(node_ptr.feature),
    id: int(node_ptr.id),
    length: int(node_ptr.length),
    rlength: int(node_ptr.rlength),
    rcAttr: int(node_ptr.rcAttr),
    lcAttr: int(node_ptr.lcAttr),
    posid: int(node_ptr.posid),
    charType : int(node_ptr.char_type),
    stat: int(node_ptr.stat),
  }
}

// ============================================
// instance methods

func Create() *Tagger {
	taggerPtr := new(Tagger)
	emptyStringPtr := C.CString("")
	taggerPtr.tagger = C.mecab_new2(emptyStringPtr)
	return taggerPtr
}

func (this *Tagger) Parse(target string) string {
	p := C.CString(target)
	c_str := C.mecab_sparse_tostr(this.tagger, p)
	s := C.GoString(c_str)
	return s
}

func (this *Tagger) ParseToNode(target string) *Node {
	target_ptr := C.CString(target)
	node_ptr := C.mecab_sparse_tonode(this.tagger, target_ptr)
	return TransCMecabNode2Go(node_ptr)
}

func (this *Tagger) Destroy() {
	C.mecab_destroy(this.tagger)
}

// ============================================
// this is static anyway

func GetVersion() string {
	p := C.mecab_version()
	return C.GoString(p)
}

// ============================================
