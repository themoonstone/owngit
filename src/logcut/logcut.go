/*
	通过重写io.Writer及Writer函数判断日志文件大小和文件个数
	如果文件大小超过最大值，重新创建新文件，如果文件个数超过允许的最大值、删除最早的文件、保持文件数量
*/
package logcut

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	FILE_FLAGS int //文件操作方式
)

//日志信息
type LogData struct {
	LogName  string //日志名称
	MaxSize  int64  //单个文件允许的最大大小
	MaxCount int64  //允许的最多文件数量
	Path     string //当前目录
}

/*
	函数名：Write
	函数功能：判断文件是否存在、将数据流写入指定文件
	函数参数：p 写入的数据
	返回值：n 数据写入长度、err 写入失败返回的错误
*/
func (lg *LogData) Write(p []byte) (n int, err error) {
	var filename string
	fmt.Println(lg.LogName)
	filename = lg.LogName
	if !isFileExist(filename) {
		os.Create(filename) //如果日志文件不存在、创建文件
	}
	cutLog(filename, lg.MaxSize, p)
	dirErg(lg.Path, lg.MaxCount)
	return len(p), nil
}

/*
	函数名：cutLog
	函数功能：日志切割
	函数参数：filename 日志文件名
				maxsize 日志允许的最大数据量
				p 写入数据
*/
func cutLog(filename string, maxsize int64, p []byte) {
	data, _ := os.Lstat(filename)
	var dstName string
	dstName = filename + time.Now().String()
	if data.Size() >= maxsize {
		fileCopy(dstName, filename)
		FILE_FLAGS = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	} else {
		FILE_FLAGS = os.O_RDWR | os.O_CREATE | os.O_APPEND
	}
	file, err := os.OpenFile(filename, FILE_FLAGS, 0666)
	if err != nil {
		log.Println(err)
	}
	file.Write(p)
}

/*
	函数名：isFileExist
	函数功能：判断指定的日志文件是否存在
	函数参数：filename 文件名
	返回值：bool 存在返回true 否则返回false
*/
func isFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/*
	函数名：fileCopy
	函数功能：日志拷贝
	函数参数：dstName目标日志文件
				srcName源日志文件
	返回值：written 写入的数据大小，err 写入失败的错误
*/
func fileCopy(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Println(err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

/*
	函数名：dirErg
	函数功能：遍历指定目录文件、统计日志文件数量
	函数参数：path 路径
*/
func dirErg(path string, maxcount int64) {
	m := make(map[float64]string)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(time.Since(info.ModTime()).Seconds())

		if strings.Contains(path, "log") {
			m[time.Since(info.ModTime()).Seconds()] = path
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	fileOpera(m, maxcount)
}

/*
	函数名：fileOpera
	函数功能：日志文件排序、删除最先生成的文件
	函数参数：m 存放文件信息的容器
*/
func fileOpera(m map[float64]string, maxcount int64) {
	var slice []float64
	for i := range m {
		slice = append(slice, i)
	}
	sort.Float64s(slice)

	if int64(len(slice)) > maxcount {
		for k := len(slice) - 1; k > 5; k-- {
			os.Remove(m[slice[k]])
		}
	}
}
