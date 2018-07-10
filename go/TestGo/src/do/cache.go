package do

import (
	"time"
	//"log"
)

//var num int8 = 0
//var datSizeMap = make(map[int8]*[]DatSize)
//var datSessionMap = make(map[int8][]DatSession)

var datSizeMap [][]DatSize
var datSessionMap [][]DatSession

func addDatSize() {
	data := getDatSize()
	//datSizeMap[num] = &data
	datSizeMap = append(datSizeMap, data)
	Info.Println("DatSize Length:", len(datSizeMap))
}
func addDatSession() {
	data := getDatSession()
	//datSessionMap[num] = data
	datSessionMap = append(datSessionMap, data)
	Info.Println("DatSession Length:", len(datSessionMap))
}

func cleanExpiration() {
	Info.Println("Enter: datSizeMap=", len(datSizeMap), "datSessionMap", len(datSessionMap))

	if len(datSizeMap) == 360 {
		copy(datSizeMap[0:], datSizeMap[1:])
		datSizeMap = datSizeMap[:len(datSizeMap)-1]
	}
	if len(datSessionMap) == 360 {
		copy(datSessionMap[0:], datSessionMap[1:])
		datSessionMap = datSessionMap[:len(datSessionMap)-1]
	}

	Info.Println("Return: datSizeMap=", len(datSizeMap), "datSessionMap", len(datSessionMap))
}

func CacheDatas() {
	timer := time.NewTicker(1 * time.Second)
	for _ = range timer.C {
		//num++
		
		cleanExpiration()
		go addDatSize()
		go addDatSession()
	}
}

