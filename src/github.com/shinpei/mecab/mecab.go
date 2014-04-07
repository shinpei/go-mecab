package main

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_t * go_mecab_new (int ac, char ** av) {
  return mecab_new (ac, av);
}

static mecab_t * go_mecab_new2(char *arg) {
  return mecab_new2(arg);
}

static const char *go_mecab_version(void) {
  return mecab_version();
}

static const char *go_mecab_strerror(mecab_t *mecab) {
  return mecab_strerror(mecab);
}

static void go_mecab_destroy(mecab_t* mecab) {
  mecab_destroy(mecab);
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

static const char* go_mecab_sparse_tostr(mecab_t *mecab, const char *str) {
  return mecab_sparse_tostr(mecab, str);
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


*/
import "C"

type Tagger struct {
  tagger *C.mecab_t

}

func Create() *Tagger {
  taggerPtr := new(Tagger);
  emptyStringPtr := C.CString("");
  taggerPtr.tagger = go_mecab_lattice_new(emptyStringPtr);
  return taggerPtr
}

func main () {
  tagger := Create();
}
