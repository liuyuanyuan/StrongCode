package do

import (
	"testing"
	"fmt"
)

/*go test 
----------
ok  _/F_/devel_workspace/smart-coding/go/TestGo/src/do  0.079s
*/
func TestGetCPUNum(t  *testing.T){
	if GetCPUNum() != 8 {
	   t.Error(`GetCPUNum() != 8`)
	}
}

/*
 go test -bench=.
----------------
goos: windows
goarch: amd64
BenchmarkGetCPUNum-8    2000000000               0.28 ns/op
PASS
ok      _/F_/devel_workspace/smart-coding/go/TestGo/src/do      0.663s
*/
func BenchmarkGetCPUNum(b *testing.B){
    for i := 0; i<b.N; i++{
    	GetCPUNum()
    }
}

//go test
func ExampleGetCPUNum(){
    fmt.Println(GetCPUNum())
    //输出：
    //8
}



