package quote

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)


func Test_Ticks(t *testing.T){
	////for {
	x:="100213,600214,100528,200356"
	var buffer bytes.Buffer
	for _,code :=range strings.Split(x,","){
		buffer.WriteString("'")
		buffer.WriteString(code)
		buffer.WriteString("'")
		buffer.WriteString(",")
	}
	fmt.Println(buffer.String())

}

