package quote

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)
var stop chan struct{}
func runFuncName()string{
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func test1(){
	i:=0
	fmt.Println("i =",i)
	fmt.Println("FuncName1 =",runFuncName())
}

func test2(){
	i:=1
	fmt.Println("i =",i)
	fmt.Println("FuncName2 =",runFuncName())
}

func main(){
	fmt.Println("打印运行中的函数名")
	test1()
	test2()
}
func Test_stock(t *testing.T){
	test1()
	test2()
	//TickStock()
	//fmt.Println(Min("10002283"))

	//BarStock("601318",240,"1min")
}
func zi(ctx context.Context){
	for {
		select {
		case <-ctx.Done():
			fmt.Println("zi stop")
			return
		}
		time.Sleep(time.Second)
	}
}
func chiHanBao(ctx context.Context) {
	n := 0
	go zi(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡\n", n)
		}
		time.Sleep(time.Second)
	}
}


func check(codes chan string){
	for x:=range codes{
		fmt.Println(x)
	}
	fmt.Println("ssssss")
}
