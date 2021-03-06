package cos

import "github.com/Raincal/rikka/plugins"

type tccosPlugin struct{}

var (
	l = plugins.SubLogger("[TC-COS]")

	appID      string
	secretID   string
	secretKey  string
	bucketName string
	bucketPath string
	bucketHost string
	region     string

	client *cosClient

	// TccosPlugin is the main plugin instance
	TccosPlugin tccosPlugin
)

func (plugin tccosPlugin) ExtraHandlers() []plugins.HandlerWithPattern {
	return nil
}
