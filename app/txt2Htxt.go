package app

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func Txt2Htxt(htype string, hfiles []string) {
	for _, v := range hfiles {
		sfileStr := v
		tfileStr := htype + "_" + v
		sfile, err := os.OpenFile(sfileStr,os.O_RDONLY, 0600)
		if err != nil {
			fmt.Println(err)
		}
		tfile, err := os.OpenFile(tfileStr,os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
		}
		defer sfile.Close()
		defer tfile.Close()
		if err := Txt2HtxtDo(htype,sfile,tfile); err != nil {
			fmt.Printf("失败:错误为 " + err.Error())
		} else {
			fmt.Println("加密成功:" + tfileStr)
		}
	}
}

func Txt2HtxtDo(htype string,sf,tf *os.File) error{
     sfScan := bufio.NewScanner(sf)
     for sfScan.Scan() {
     	stxt := sfScan.Text()
     	stxt = strings.TrimRight(stxt,"\r")
     	stxt = strings.TrimRight(stxt,"\n")
     	stxt = strings.TrimRight(stxt,"\r\n")
     	htxt := Hhash(htype, stxt)
     	//fmt.Println(htxt)
     	tf.Write([]byte(htxt + "\n"))
	 }
     return nil
}

/**
* 按照给定的哈希加密类型对字符串进行哈希加密
* htype 哈希加密类型
* s 字符串
 */
func Hhash(htype,s string ) string {
	switch htype {
	case "md5":
		return Hmd5(s)
		break
	case "sha1":
		return Hsha1(s)
		break
	case "sha256":
		return Hsha256(s)
		break
	default:
		return Hmd5(s)
		break
	}
	return ""
}

/**
* 对字符串进行md5
 */
func Hmd5(s string) string{
	h := md5.New()
	h.Write([]byte(s))
	hv := h.Sum(nil)
	return hex.EncodeToString(hv)
}

/**
* 对字符串进行sha1
 */
func Hsha1(s string) string{
	h := sha1.New()
	h.Write([]byte(s))
	hv := h.Sum(nil)
	return hex.EncodeToString(hv)
}

/**
* 对字符串进行sha256
 */
func Hsha256(s string) string{
	h := sha256.New()
	h.Write([]byte(s))
	hv := h.Sum(nil)
	return hex.EncodeToString(hv)
}
