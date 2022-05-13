package initdb

import (
	"fmt"
	"go/importer"
	"testing"
)

func Test_CreateDB(t *testing.T){
	pkg, err := importer.Default().Import("opmanage/core/pub")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
}