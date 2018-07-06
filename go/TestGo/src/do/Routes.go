package do
import(
	"net/http"
    "github.com/gorilla/mux"
)

type Route struct{
	Name string 
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	//http://localhost:8085/
    Route{"Index", "GET", "/", Index},
    //http://localhost:8085/todos
    Route{"TodoIndex", "GET", "/todos", TodoIndex},
    //http://localhost:8085/todos/1
    Route{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
    
    //http://localhost:8085/getDatSize
	Route{"GetDatSize", "GET", "/getDatSize", GetDatSize},
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
        Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    return router
}



