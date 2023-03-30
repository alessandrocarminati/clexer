package main

/*
#cgo LDFLAGS: -lcparser -L.
#include <stdio.h>
#include "cparser.tab.h"

extern FILE *yyin;
extern int yyparse(void);

int parse_file(const char *filename) {
    yyin = fopen(filename, "r");
    if (!yyin) {
        return 1;
    }
    return yyparse();
}
*/
import "C"
import (
    "fmt"
    "os"
//    "unsafe"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./parser <filename>")
        return
    }
    filename := C.CString(os.Args[1])
//    defer C.free(unsafe.Pointer(filename))

    ret := C.parse_file(filename)
    if ret != 0 {
        fmt.Println("Error parsing file")
        return
    }
    fmt.Println("Parsing successful!")
   fmt.Println(ProcessorSTR)
   fmt.Println(FuncName)
}
func yyerror(s *C.char) {
    fmt.Println(C.GoString(s))
}
