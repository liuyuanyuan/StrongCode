package do 

import(
)

func checkErr(err error) {
    if err != nil {
        Error.Panic(err) //panic相当于OO里面常用的异常捕获
    }
}