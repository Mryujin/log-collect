package core

import (
	"encoding/json"
)

/**
 * app 上报日志结构体
 */   
type AppLog struct {
	AppId string `json:"appId"`              //应用唯一标识
	DeviceId string `json:"deviceId"`        //设备唯一标识
	TenantId string `json:"tenantId"`        //租户Id
	AppVersion string `json:"appVersion"`    //版本
	AppChannel string `json:"appChannel"`    //渠道,安装时就在清单中制定了，appStore等。 
	AppPlatform string `json:"appPlatform"`  //平台
	OsType string `json:"osType"`            //操作系统
	DeviceStyle string `json:"deviceStyle"`  //机型

	AppStartUpLogs []AppStartUpLog `json:"appStartUpLogs"`
	AppPageLogs []AppPageLog `json:"appPageLogs"`
	AppEventLogs []AppEventLog `json:"appEventLogs"`
	AppUsageLogs []AppUsageLog `json:"appUsageLogs"`
	AppErrorLogs []AppErrorLog `json:"appErrorLogs"`
}	

/**
 * app启动日志
 */              
type AppStartUpLog struct {
	CreatedAtMs int64 `json:"createdAtMs"`    //日志创建时间
	AppId string `json:"appId"`               //应用唯一标识
	TenantId string `json:"tenantId"`         //租户唯一标识,企业用户
	DeviceId string `json:"deviceId"`         //设备唯一标识
	AppVersion string `json:"appVersion"`     //版本
	AppChannel string `json:"appChannel"`     //渠道,安装时就在清单中制定了，appStore等。
	AppPlatform string `json:"appPlatform"`   //平台
	OsType string `json:"osType"`             //操作系统
	DeviceStyle string `json:"deviceStyle"`   //机型
	
	Country string `json:"country"`           //国家，终端不用上报，服务器自动填充该属性
	Provice string `json:"provice"`           //省份，终端不用上报，服务器自动填充该属性
	City string `json:"city"`                 //城市，终端不用上报，服务器自动填充该属性
	IpAddress string `json:"ipAddress"`       //ip地址
	Network string `json:"network"`           //网络
	Carrier string `json:"carrier"`           //运营商
	Brand string `json:"brand"`               //品牌
	ScreenSize string `json:"screenSize"`     //分辨率
}

/**
 * 应用上报页面日志
 */  
type AppPageLog struct {
	CreatedAtMs int64 `json:"createdAtMs"`    //日志创建时间
	AppId string `json:"appId"`               //应用唯一标识
	TenantId string `json:"tenantId"`         //租户唯一标识,企业用户
	DeviceId string `json:"deviceId"`         //设备唯一标识
	AppVersion string `json:"appVersion"`     //版本
	AppChannel string `json:"appChannel"`     //渠道,安装时就在清单中制定了，appStore等。
	AppPlatform string `json:"appPlatform"`   //平台
	OsType string `json:"osType"`             //操作系统
	DeviceStyle string `json:"deviceStyle"`   //机型

	//一次启动中的页面访问次数(应保证每次启动的所有页面日志在一次上报中，即最后一条上报的页面记录的nextPage为空)
	PageViewCntInSession int32 `json:"pageViewCntInSession"`
	PageId string `json:"pageId"`             //页面id
	VisitIndex int `json:"visitIndex"`        //访问顺序号，0为第一个页面
	NextPage string `json:"nextPage"`         //下一个访问页面，如为空则表示为退出应用的页面
	StayDurationSecs int64 `json:"stayDurationSecs"` //当前页面停留时长
}

/**
 * 应用上报的事件相关信息
 */		
type AppEventLog struct {
	CreatedAtMs int64 `json:"createdAtMs"`    //日志创建时间
	AppId string `json:"appId"`               //应用唯一标识
	TenantId string `json:"tenantId"`         //租户唯一标识,企业用户
	DeviceId string `json:"deviceId"`         //设备唯一标识
	AppVersion string `json:"appVersion"`     //版本
	AppChannel string `json:"appChannel"`     //渠道,安装时就在清单中制定了，appStore等。
	AppPlatform string `json:"appPlatform"`   //平台
	OsType string `json:"osType"`             //操作系统
	DeviceStyle string `json:"deviceStyle"`   //机型

	EventId string `json:"eventId"`           //事件唯一标识
	EventDurationSecs int64 `json:"eventDurationSecs"` //事件持续时长
	ParamKeyValueMap map[string]string `json:"paramKeyValueMap"` //参数名/值对
}

/**
 * 应用上报的使用时长相关信息
 */
type AppUsageLog struct {
	CreatedAtMs int64 `json:"createdAtMs"`    //日志创建时间
	AppId string `json:"appId"`               //应用唯一标识
	TenantId string `json:"tenantId"`         //租户唯一标识,企业用户
	DeviceId string `json:"deviceId"`         //设备唯一标识
	AppVersion string `json:"appVersion"`     //版本
	AppChannel string `json:"appChannel"`     //渠道,安装时就在清单中制定了，appStore等。
	AppPlatform string `json:"appPlatform"`   //平台
	OsType string `json:"osType"`             //操作系统
	DeviceStyle string `json:"deviceStyle"`   //机型

	SingleUseDurationSecs int64 `json:"singleUseDurationSecs"` //单次使用时长(秒数),指一次启动内应用在前台的持续时长
	SingleUploadTraffic int64 `json:"singleUploadTraffic"`     //单次使用过程中的上传流量
	SingleDownloadTraffic int64 `json:"singleDownloadTraffic"` //单次使用过程中的下载流量
}
	
/**
 * 错误日志
 */
type AppErrorLog struct {
	CreatedAtMs int64 `json:"createdAtMs"`    //日志创建时间
	AppId string `json:"appId"`               //应用唯一标识
	TenantId string `json:"tenantId"`         //租户唯一标识,企业用户
	DeviceId string `json:"deviceId"`         //设备唯一标识
	AppVersion string `json:"appVersion"`     //版本
	AppChannel string `json:"appChannel"`     //渠道,安装时就在清单中制定了，appStore等。
	AppPlatform string `json:"appPlatform"`   //平台
	OsType string `json:"osType"`             //操作系统
	DeviceStyle string `json:"deviceStyle"`   //机型

	ErrorBrief string `json:"errorBrief"`     //错误摘要
	ErrorDetail string `json:"errorDetail"`   //错误详情
}

/**
 * 将json对象转化为实体对象
 */
func ParseObj(appLogBytes []byte) (*AppLog, error) {
	var appLog AppLog
	err := json.Unmarshal(appLogBytes, &appLog)
	if err != nil {
		return nil, err
	}
	return &appLog, nil
}
