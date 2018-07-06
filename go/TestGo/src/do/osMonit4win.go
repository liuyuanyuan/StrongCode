package do
import(
	"fmt"
	"runtime"
)


func GetOSMemory() {
	fmt.Println("CPU: ",runtime.NumCPU())
}