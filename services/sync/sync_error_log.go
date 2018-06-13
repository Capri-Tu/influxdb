package sync

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/hpcloud/tail"
	"go.uber.org/zap"
)

type DealErrorLog struct {
	FilePath string // 主从同步错误日志文件路径
	Logger   *zap.Logger
}

// 返回一个新实例
func NewDealErrorLog(filePath string) *DealErrorLog {
	return &DealErrorLog{
		FilePath: filePath,
		Logger:   zap.NewNop(),
	}
}

// 处理错误日志
func (d *DealErrorLog) TailLog() error {
	// 检查文件及路径是否存在
	errFilePath := d.exist()
	if errFilePath != nil {
		return errFilePath
	}
	// 从当前文件的末尾开始读取
	seekInfo := tail.SeekInfo{
		Whence: io.SeekEnd,
	}
	t, e := tail.TailFile(d.FilePath, tail.Config{
		Location: &seekInfo,
		Follow:   true,
		ReOpen:   true,
	})
	if e != nil { // 文件读取异常
		return e
	}
	// 处理文件
	for line := range t.Lines {
		//TODO 处理日志文件行
		if line != nil {

		}
	}
	return nil
}

// 判断目录是否存在，且为文件
func (d *DealErrorLog) exist() error {
	fileInfo, e := os.Stat(d.FilePath)
	if e != nil && os.IsExist(e) {
		return errors.New(fmt.Sprintf("filePath [%s] is not exist", d.FilePath))
	}
	if fileInfo.IsDir() {
		return errors.New(fmt.Sprintf("filePath [%s] is a dir", d.FilePath))
	}
	return nil
}
