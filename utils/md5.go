package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) (string, error) {
	//创建MD5算法
	sb := []byte(s)
	h := md5.New()
	h.Write(sb)
	//写入待加密数据
	//获取哈希值字符切片
	bts := h.Sum(nil)
	//转化为16进制字符串
	return hex.EncodeToString(bts), nil
}
