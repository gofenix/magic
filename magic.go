package magic

import (
	"bytes"
	"log"

	"v2ray.com/core"

	"v2ray.com/ext/tools/conf/serial"
)

func StartV2Ray(format string, configBytes []byte) (*core.Instance, error) {
	jsonConfig := bytes.NewReader(configBytes)
	pbConfig, err := serial.LoadJSONConfig(jsonConfig)
	log.Println(pbConfig)
	if err != nil {
		log.Println(err)
	}

	return core.New(pbConfig)
}
