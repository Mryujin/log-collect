package utils

import (
	"net"
	"strings"
	"github.com/valyala/fasthttp"
	"github.com/Mryujin/log-collect/log"
	geoip2 "github.com/oschwald/geoip2-golang"
)

/**
 * 地域信息
 */
type Region struct {
	Country string
	Province string
	City string
}

var RegionCache map[string]Region = make(map[string]Region)

func GetRegionFromCache(ipAddress string) Region {
	region := RegionCache[ipAddress]
	if region.City == "" && region.Province == "" {
		log.Logger.Info("get region from GeoLite2-City.mmdb.....")
		region = GetRegionByIp(ipAddress)
		RegionCache[ipAddress] = region
		return region
	}
	return region
}

/**
 * 通过ip获取地域信息
 */
func GetRegionByIp(ipAddress string) Region {
	db, err := geoip2.Open("F:/workspace/go/src/log-collect/ipdic/GeoLite2-City.mmdb")
    if err != nil {
		log.ErrLogger.Info(err)
    }
	defer db.Close()
    
    ip := net.ParseIP(ipAddress)
    record, err := db.City(ip)
    if err != nil {
        log.ErrLogger.Error(err)
	}
	log.Logger.Info("ip: ", ipAddress)
	return Region {record.Country.Names["zh-CN"],
					record.Subdivisions[0].Names["zh-CN"], 
					record.City.Names["zh-CN"]}
}

/**
 * 获取客户真实IP
 */
func ClientIP(ctx *fasthttp.RequestCtx) string {

    clientIP := string(ctx.Request.Header.Peek("X-Forwarded-For"))
    if index := strings.IndexByte(clientIP, ','); index >= 0 {
        clientIP = clientIP[0:index] //获取最开始的一个 即 1.1.1.1
    }
    clientIP = strings.TrimSpace(clientIP)
    if len(clientIP) > 0 {
        return clientIP
    }
    clientIP = strings.TrimSpace(string(ctx.Request.Header.Peek("X-Real-Ip")))
    if len(clientIP) > 0 {
        return clientIP
    }
    return ctx.RemoteIP().String()
}