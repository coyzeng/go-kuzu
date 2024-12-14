package kuzu

/*
#cgo darwin LDFLAGS: -lc++ -L${SRCDIR}/lib/dynamic/darwin -lkuzu -Wl,-rpath,${SRCDIR}/lib/dynamic/darwin
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/lib/dynamic/linux-amd64 -lkuzu -lantlr4_cypher -lantlr4_runtime -lbrotlicommon -lbrotlidec -lbrotlienc -lfastpfor -llz4 -lmbedtls -lminiz -lparquet -lre2 -lroaring_bitmap -lsnappy -lthrift -lutf8proc -lzstd -Wl,-rpath,${SRCDIR}/lib/dynamic/linux-amd64 -lstdc++
#cgo linux,arm64 LDFLAGS: -L${SRCDIR}/lib/dynamic/linux-arm64 -lkuzu -Wl,-rpath,${SRCDIR}/lib/dynamic/linux-arm64
#include "kuzu.h"
*/
import "C"
