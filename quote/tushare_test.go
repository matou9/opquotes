package quote

import (
	"fmt"
	"testing"
	"time"
)


func Test_option(t *testing.T){

	//for {
		begin:=time.Now()
		tt:=&Tushare{}
		tt.GetBarData("day",10)
		fmt.Println("用时=",time.Now().Sub(begin).Seconds())
		time.Sleep(time.Second)
	//}

	//s:=SinaCrawl{}
	//s.GetAllMinData()
}

