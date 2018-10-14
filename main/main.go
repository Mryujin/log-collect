package main

import (
	"encoding/json"
	"github.com/Mryujin/log-collect/log"
	"github.com/Mryujin/log-collect/core"
	"github.com/Mryujin/log-collect/kafka"
	"github.com/Mryujin/log-collect/utils"
	"github.com/valyala/fasthttp"
)

const (
	TOPIC_APP_STARTUP = "topic-app-startup"
	TOPIC_APP_ERRROR = "topic-app-error"
	TOPIC_APP_EVENT = "topic-app-event"
	TOPIC_APP_USAGE = "topic-app-usage"
	TOPIC_APP_PAGE = "topic-app-page"
)
/**
 * 收集日志程序
 */
func main() {
	log.Logger.Info("server start....", 10)
	err := fasthttp.ListenAndServe("0.0.0.0:8080", logCollectHandle)
	if err != nil {
		log.FHLogger.Info("server start up fail.", err)
	}
}

func logCollectHandle(ctx *fasthttp.RequestCtx) {
	if ctx.IsPost() {
		serverTime := utils.NowMillsecond()
		clientTime := utils.ClientMillsecond(ctx.Request.Header.Peek("client-time")) // 获取客户端时间
		diff := serverTime - clientTime
		body := ctx.PostBody()
		appLog,err := core.ParseObj(body)
		if err != nil {
			log.ErrLogger.Error("request param:", body, err)
		}

		// 修正时间
		verifyTime(appLog, diff)

		// 处理ip
		ipAddress := utils.ClientIP(ctx)
		processRegion(appLog, ipAddress)

		// 将消息放到kafka中
		sendMessage(appLog)

		ctx.SetContentType("application/json")
		ctx.Response.SetBody([]byte("{\"code\":0, \"msg\":\"success\"}"))
		// fmt.Println(appLog.AppChannel)
		// fmt.Fprintf(ctx, "hello fasthttp") // *RequestCtx 实现了 io.Writer
	}
}

/*
 * 修正时间并且复制属性
 */
func verifyTime(appLog *core.AppLog, diff int64) {
	// 启动日志
	for _, log := range appLog.AppStartUpLogs {
		log.CreatedAtMs += diff
		log.AppId = appLog.AppId
		log.DeviceId = appLog.DeviceId
		log.TenantId = appLog.TenantId
		log.AppVersion = appLog.AppVersion
		log.AppChannel = appLog.AppChannel
		log.AppPlatform = appLog.AppPlatform
		log.OsType = appLog.OsType
		log.DeviceStyle = appLog.DeviceId
	}

	// 使用日志
	for _, log := range appLog.AppUsageLogs {
		log.CreatedAtMs += diff
		log.AppId = appLog.AppId
		log.DeviceId = appLog.DeviceId
		log.TenantId = appLog.TenantId
		log.AppVersion = appLog.AppVersion
		log.AppChannel = appLog.AppChannel
		log.AppPlatform = appLog.AppPlatform
		log.OsType = appLog.OsType
		log.DeviceStyle = appLog.DeviceId
	}

	// 页面日志
	for _, log := range appLog.AppPageLogs {
		log.CreatedAtMs += diff
		log.AppId = appLog.AppId
		log.DeviceId = appLog.DeviceId
		log.TenantId = appLog.TenantId
		log.AppVersion = appLog.AppVersion
		log.AppChannel = appLog.AppChannel
		log.AppPlatform = appLog.AppPlatform
		log.OsType = appLog.OsType
		log.DeviceStyle = appLog.DeviceId
	}

	for _, log := range appLog.AppEventLogs {
		log.CreatedAtMs += diff
		log.AppId = appLog.AppId
		log.DeviceId = appLog.DeviceId
		log.TenantId = appLog.TenantId
		log.AppVersion = appLog.AppVersion
		log.AppChannel = appLog.AppChannel
		log.AppPlatform = appLog.AppPlatform
		log.OsType = appLog.OsType
		log.DeviceStyle = appLog.DeviceId
	}

	// 错误日志
	for _, log := range appLog.AppErrorLogs {
		log.CreatedAtMs += diff
		log.AppId = appLog.AppId
		log.DeviceId = appLog.DeviceId
		log.TenantId = appLog.TenantId
		log.AppVersion = appLog.AppVersion
		log.AppChannel = appLog.AppChannel
		log.AppPlatform = appLog.AppPlatform
		log.OsType = appLog.OsType
		log.DeviceStyle = appLog.DeviceId
	}
}

/**
 * 处理IP
 */
func processRegion(appLog *core.AppLog, ipAddress string) {
	region := utils.GetRegionFromCache(ipAddress)
	for _, log := range appLog.AppStartUpLogs {
		log.Country = region.Country
		log.Provice = region.Province
		log.City = region.City
		log.IpAddress = ipAddress
	}
}

/**
 * 发送消息
 */
func sendMessage(appLog *core.AppLog) {
	// 启动日志
	for _, log := range appLog.AppStartUpLogs {
		bytes,err := json.Marshal(log)
		kafka.SendMessage(TOPIC_APP_STARTUP, string(bytes))
	}

	// 事件日志
	for _,log := range appLog.AppEventLogs {
		bytes,err := json.Marshal(log)
		kafka.SendMessage(TOPIC_APP_EVENT, string(bytes))
	}

	// 使用日志
	for _, log := range appLog.AppUsageLogs {
		bytes,err := json.Marshal(log)
		kafka.SendMessage(TOPIC_APP_USAGE, string(bytes))
	}

	// 页面日志
	for _, log := range appLog.AppPageLogs {
		bytes,err := json.Marshal(log)
		kafka.SendMessage(TOPIC_APP_PAGE, string(bytes))
	}

	// 错误日志
	for _, log := range appLog.AppErrorLogs {
		bytes,err := json.Marshal(log)
		kafka.SendMessage(TOPIC_APP_ERRROR, string(bytes))
	}
}