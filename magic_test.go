package magic

import (
	proto "github.com/golang/protobuf/proto"

	"encoding/json"
	"testing"
	"v2ray.com/core"
	"v2ray.com/core/app/dispatcher"
	"v2ray.com/core/app/proxyman"
	"v2ray.com/core/common/net"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/common/uuid"
	"v2ray.com/core/proxy/dokodemo"
	"v2ray.com/core/proxy/vmess"
	"v2ray.com/core/proxy/vmess/outbound"
	"v2ray.com/core/testing/servers/tcp"
	"v2ray.com/ext/tools/conf"

	"v2ray.com/core/common"
	//_ "v2ray.com/core/main/distro/all"
)

func TestV2RayClose(t *testing.T) {
	port := tcp.PickPort()

	userId := uuid.New()
	config := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(port),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP, net.Network_UDP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userId.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := core.StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}


func TestStartV2Ray(t *testing.T) {

	config := &conf.Config{


	}

	cfgBytes, err := json.Marshal(config)
	common.Must(err)

	server, err := StartV2RayByJSON(cfgBytes)
	common.Must(err)
	server.Close()
}

func TestStartV2RayByJsonString(t *testing.T) {
	cfgStr := `
{
  "inbounds": [
    {
      "port": 1080,
      "listen": "127.0.0.1",
      "protocol": "socks",
      "settings": {
        "udp": true
      }
    }
  ],
  "outbounds": [
    {
      "protocol": "vmess",
      "settings": {
        "vnext": [
          {
            "address": "server",
            "port": 10086,
            "users": [{ "id": "b831381d-6324-4d53-ad4f-8cda48b30811" }]
          }
        ]
      }
    },
    {
      "protocol": "freedom",
      "tag": "direct",
      "settings": {}
    }
  ]
}


	`

	cfgBytes := []byte(cfgStr)

	server, err := StartV2RayByJSON(cfgBytes)

	common.Must(err)



	server.Start()
	server.Close()
}
