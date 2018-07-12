package do
import(
	"fmt"
	"runtime"
	"os"
)


func GetOSMem() {
	vcpu := runtime.NumCPU()
	fmt.Println("CPU: ", vcpu)
}

func GetOSName(){
	 host, err := os.Hostname()
	 checkErr(err)
	 fmt.Println("HostName: ", host)

     fmt.Println("Virtual CPU: ", runtime.NumCPU())
}