package main

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_t * InitMecab(void) {
  mecab_t *m = mecab_new2("");
  return m;
}

*/

import "C"

func main () {
  var m *C.mecab_t = C.InitMecab();
}
