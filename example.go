package main

/*
#cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
#include <mecab.h>

static mecab_t * InitMecab(void) {
    char *dummy = "";
    mecab_t *m = mecab_new2(dummy);
    return m;
}

static const char * ParseMecab(mecab_t *mecab, const char *input) {
    const char *ret = mecab_sparse_tostr(mecab, input);
    return ret;
}
 */
import "C"

import "fmt"

func main () {
	var m *C.mecab_t = C.InitMecab();
	var str = "Hoge Hoge";
	p := C.ParseMecab(m, C.CString(str));
	s := C.GoString(p);
	fmt.Printf("str: '%s'\n" , s);
	
}
