package magic

import (
	"bytes"
	"github.com/gogo/protobuf/proto"
	"v2ray.com/core"
	"v2ray.com/ext/tools/conf/serial"

	// Required features. Can't remove unless there is replacements.
	_ "v2ray.com/core/app/dispatcher"
	_ "v2ray.com/core/app/proxyman/inbound"
	_ "v2ray.com/core/app/proxyman/outbound"
)

func StartInstance(pbConfigBytes []byte) error {
	_, err := core.StartInstance("protobuf", pbConfigBytes)
	return err
}

func GetPbConfigBytes(jsonConfigBytes []byte) ([]byte, error) {
	pbConfig, err := serial.LoadJSONConfig(bytes.NewReader(jsonConfigBytes))
	if err != nil {
		return nil, err
	}
	pbConfigBytes, err := proto.Marshal(pbConfig)
	if err != nil {
		return nil, err
	}
	return pbConfigBytes, nil
}

func StartByJSON(jsonConfigBytes []byte) error {
	pbConfig, err := serial.LoadJSONConfig(bytes.NewReader(jsonConfigBytes))
	if err != nil {
		return err
	}
	pbConfigBytes, err := proto.Marshal(pbConfig)
	if err != nil {
		return err
	}
	_, err = core.StartInstance("protobuf", pbConfigBytes)
	return err
}
