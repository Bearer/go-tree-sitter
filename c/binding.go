package c

//#cgo LDFLAGS: ${SRCDIR}/../vendor/c/parser.o
//#cgo CFLAGS: -I${SRCDIR}/../vendor/tree-sitter/include
//#include "tree_sitter/parser.h"
//TSLanguage *tree_sitter_c();
import "C"
import (
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_c())
	return sitter.NewLanguage(ptr)
}
