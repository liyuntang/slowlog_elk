package watcher

import (
	"fmt"
	"strconv"
	"strings"
)

func analy(sql string, data *elkData) {
	tmpSlice := strings.Split(sql, " ")
	dataSlice := []string{}
	for _, str := range tmpSlice {
		if len(str) != 0 {
			dataSlice = append(dataSlice, strings.Trim(str, ":"))
		}
	}
	if len(dataSlice) <8 {
		msg := fmt.Sprintf("len of dataSlice is to short,len is %d dataSlice is %s", len(dataSlice), dataSlice)
		panic(msg)
	}
	floatA, err := strconv.ParseFloat(dataSlice[1], 64)
	if err != nil {
		panic(err)
	}
	data.Query_time = floatA
	floatB, err1 := strconv.ParseFloat(dataSlice[3], 64)
	if err1 != nil {
		panic(err)
	}
	data.Lock_time = floatB
	data.Rows_sent = dataSlice[5]
	data.Rows_examined = dataSlice[7]
}
