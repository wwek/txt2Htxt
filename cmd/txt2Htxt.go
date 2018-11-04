package cmd

import (
	"flag"
	"fmt"
	"github.com/wwek/txt2Htxt/app"
	"os"
	"strings"
)

var HList = []string{"md5","sha1","sha256"}
var HType string
var HFiles []string

func init() {
	flag.StringVar(&HType,"t","md5","哈希加密算法md5、sha1、sha256可选,默认md5")
	flag.Parse()
	HFiles = flag.Args()
	if !ChkHType(HType) {
		os.Exit(0)
	}
	if !ChkHFiles(HFiles) {
		os.Exit(0)
	}
}

// 执行入口
func Execute() {
	app.Txt2Htxt(HType,HFiles)
}

/**
* 检查输入的哈希加密方式
 */
func ChkHType(htp string) bool{
	inHList := false
	for _,v := range HList {
		if htp == v {
			inHList = true
			return true
		}
	}
	if inHList == false {
		fmt.Println("错误:尚未支持的哈希加密方式"+htp + "支持的方式有"+strings.Join(HList,"、"))
	}
	return false
}

/**
* 检查输入的文件
 */
func ChkHFiles(hfs []string) bool{
	if len(hfs) <=0 {
		fmt.Println("错误:未指定要加密转换的文件,请至少指定1个文件")
		return false
	} else {
		for _, v := range hfs {
			if fb, _ := PathExists(v); fb !=true {
				fmt.Println("错误:不存在的文件:"+v)
				return false
			}
		}
		return true
	}
	return false
}

/**
* 判断文件或文件夹是否存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}