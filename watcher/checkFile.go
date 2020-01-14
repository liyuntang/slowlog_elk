package watcher

import (
	"errors"
	"fmt"
	"os"
)

func check(File string) error {
	fileInfo, err := os.Stat(File)
	if os.IsNotExist(err) {
		msg := fmt.Sprintf("sorry get slowlog", File, "is bad, err is", err)
		return errors.New(msg)
	}
	//判断是文件还是目录
	if fileInfo.IsDir() {
		msg := fmt.Sprintf("sorry slowlog", File, "is dir")
		return errors.New(msg)
	}
	file, err := os.OpenFile(File, os.O_TRUNC, 0664)
	defer file.Close()
	if err != nil {
		msg := fmt.Sprintf("truncate file ", File, "is bad, err is", err)
		return errors.New(msg)
	}
	return nil
}
