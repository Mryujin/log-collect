package core

import (
	"testing"
)

func TestParseObj(t *testing.T) {
	appLogJson := "{\"appChannel\":\"youmeng2\",\"appErrorLogs\":[{\"createdAtMs\":1539450093148,\"deviceStyle\":\"红米手机1s\",\"errorBrief\":\"at cn.lift.appIn.control.CommandUtil.getInfo(CommandUtil.java:67)\",\"errorDetail\":\"java.lang.NullPointerException\\\\n    at cn.lift.appIn.web.AbstractBaseController.validInbound(AbstractBaseController.java:72)\\\\n at cn.lift.dfdf.web.AbstractBaseController.validInbound\",\"osType\":\"7.1.1\"}],\"appEventLogs\":[{\"createdAtMs\":1539450093148,\"eventDurationSecs\":45,\"eventId\":\"popMenu\",\"paramKeyValueMap\":{\"testparam3key\":\"testparam3value\",\"testparam4key\":\"testparam4value\"}}],\"appId\":\"sdk34734\",\"appPageLogs\":[{\"createdAtMs\":1539450093148,\"pageId\":\"test.html\",\"pageViewCntInSession\":0,\"stayDurationSecs\":45,\"visitIndex\":0}],\"appPlatform\":\"android\",\"appStartupLogs\":[{\"brand\":\"三星\",\"carrier\":\"EE\",\"country\":\"America\",\"createdAtMs\":1539450093148,\"deviceStyle\":\"iPhone 6 Plus\",\"network\":\"CellNetwork\",\"osType\":\"8.3\",\"province\":\"beijing\",\"screenSize\":\"960*640\"}],\"appUsageLogs\":[{\"createdAtMs\":1539450093148,\"singleUseDurationSecs\":149}],\"appVersion\":\"3.2.1\",\"deviceId\":\"device2230\",\"tenantId\":\"cake\"}"
	a,_ := ParseObj(appLogJson)

	t.Logf("got: %v\n", a.AppErrorLogs)
}