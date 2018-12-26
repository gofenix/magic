package magic

import (
	"encoding/json"
	"testing"
	"v2ray.com/ext/tools/conf"

	"v2ray.com/core/common"
	_ "v2ray.com/core/main/distro/all"
)

func TestStartV2Ray(t *testing.T) {

	config := &conf.Config{


	}

	cfgBytes, err := json.Marshal(config)
	common.Must(err)

	server, err := StartV2Ray("json", cfgBytes)
	common.Must(err)
	server.Close()
}

func TestStartV2RayByJsonString(t *testing.T) {
	cfgStr := `
	{
		"port": 1080,
		"listen": "127.0.0.1",
		"protocol": "string",
		"settings": {},
		"streamSettings": {},
		"tag": "标识",
		"sniffing": {
			"enabled": false,
			"destOverride": ["http", "tls"]
		},
		"allocate": {
			"strategy": "always",
			"refresh": 5,
			"concurrency": 3
		}
	}
	`

	cfgBytes := []byte(cfgStr)

	server, err := StartV2Ray("json", cfgBytes)

	common.Must(err)

	server.Start()
	server.Close()
}
