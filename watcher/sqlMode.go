package watcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func sqlMode(httpServer, sql string, loggs *log.Logger) string {
	url := fmt.Sprintf("http://%s/sql/", httpServer)
	req, err1 := http.NewRequest("POST", url, strings.NewReader(sql))
	if err1 != nil {
		loggs.Println("new request is bad, err is", err1)
		return ""
	}
	resp, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		fmt.Println("client do is bad, err is", err2)
		return ""
	}
	if resp.StatusCode != 200 {
		return ""
	}
	buf, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return ""
	}
	fmt.Println("str is", string(buf))
	return string(buf)
}
