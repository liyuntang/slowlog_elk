package watcher

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func upload(data *elkData, endPoint []string, loggs *log.Logger) {
	// 从众多连接中随机选择一个elk机器连接
	num := rand.Intn(len(endPoint))
	host := endPoint[num]
	url := fmt.Sprintf("%s/%s/my_type?pretty", host, data.index)
		requestBody := fmt.Sprintf(`{
	"timestamp":"%s",
	"product":"%s",
	"cluster": "%s",
	"role":"%s",
	"hostname":"%s",
	"query_time":%f,
	"lock_time":%f,
	"rows_sent":"%s",
	"rows_examined":"%s",
	"mode":"%s",
	"message":"%s"
	}`, data.timestamp, data.product, data.cluster, data.role, data.hostname, data.Query_time, data.Lock_time, data.Rows_sent, data.Rows_examined, data.mode, data.msg)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		loggs.Println("new request is bad, err is", err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err1 := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err1 != nil {
		loggs.Println("http client do is bad, err is", err1)
	}
	fmt.Println("status code is", res.StatusCode, "info is", res.Status)
}
