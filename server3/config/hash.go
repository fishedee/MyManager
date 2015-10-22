package config

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"io"
)
//对字符串进行MD5哈希
func Md5(data string) string {
	t := md5.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}

//对字符串进行SHA1哈希
func Sha1(data string) string {
	t := sha1.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}