package SnippetsGO

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// 创建一个可写入的日志文件
func NewLogFile() error {
	dateNow := time.Now().Format("2006-01-02 15-04-05") // 不能有冒号

	fileP1, _ := filepath.Abs(os.Args[0])
	dir := filepath.Dir(fileP1)
	s4RankLogPath := dir + "/S4 " + dateNow + ".log"
	s4RankLogPath, _ = filepath.Abs(s4RankLogPath)

	mode := os.FileMode(0644)
	file, err := os.OpenFile(s4RankLogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}

	S4Logger := log.New(file, "", log.LstdFlags)
	_ = S4Logger
	return nil
}
