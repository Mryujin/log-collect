package utils

import (
	"strconv"
	"time"
)

func NowMillsecond() int64 {
	return time.Now().UnixNano() / 1000000  // 毫秒 
}

/*
 * 客户端时间处理
 */
func ClientMillsecond(clientTime []byte) int64 {
	cTime := string(clientTime[:])
	clientDate,_ := strconv.Atoi(cTime)
	return int64(clientDate)
}