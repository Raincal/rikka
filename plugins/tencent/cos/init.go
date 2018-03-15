package cos

import (
	"github.com/Raincal/rikka/plugins"
	"github.com/Raincal/rikka/plugins/tencent"
)

func (plugin tccosPlugin) Init() {
	l.Info("Start plugin tccos")

	plugins.CheckCommonArgs(true, false)

	appID = tencent.GetAppIDWithCheck(l)
	secretID = tencent.GetSecretIDWithCheck(l)
	secretKey = tencent.GetSecretKeyWithCheck(l)
	region = tencent.GetRegionWithCheck(l)
	bucketName = plugins.GetBucketName()
	bucketPath = plugins.GetBucketPath()

	client = newCosClient()

	l.Info("Tccos plugin start successfully")
}
