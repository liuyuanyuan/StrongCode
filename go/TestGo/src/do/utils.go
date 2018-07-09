package do 

import(
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}