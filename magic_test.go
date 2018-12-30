package magic

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"testing"
	"v2ray.com/core"
	"v2ray.com/core/common"
	"v2ray.com/ext/tools/conf/serial"
)

//func TestStartV2Ray(t *testing.T) {
//
//	config := &conf.Config{
//
//	}
//
//	cfgBytes, err := json.Marshal(config)
//	common.Must(err)
//
//	server, err := StartV2RayByJSON(cfgBytes)
//	common.Must(err)
//	server.Close()
//}
//
//func TestStartV2RayByJsonString(t *testing.T) {
//	cfgStr := `
//{
//  "inbounds": [
//    {
//      "port": 1080,
//      "listen": "127.0.0.1",
//      "protocol": "socks",
//      "settings": {
//        "udp": true
//      }
//    }
//  ],
//  "outbounds": [
//    {
//      "protocol": "vmess",
//      "settings": {
//        "vnext": [
//          {
//            "address": "server",
//            "port": 10086,
//            "users": [{ "id": "b831381d-6324-4d53-ad4f-8cda48b30811" }]
//          }
//        ]
//      }
//    },
//    {
//      "protocol": "freedom",
//      "tag": "direct",
//      "settings": {}
//    }
//  ]
//}
//
//
//	`
//
//	cfgStr = `
//	{
//  "inbounds": [{
//    "port": 10086,
//    "protocol": "vmess",
//    "settings": {
//      "clients": [{ "id": "b831381d-6324-4d53-ad4f-8cda48b30811" }]
//    }
//  }],
//  "outbounds": [{
//    "protocol": "freedom",
//    "settings": {}
//  }]
//}
//`
//
//	cfgBytes := []byte(cfgStr)
//
//	server, err := StartV2RayByJSON(cfgBytes)
//
//	common.Must(err)
//
//	common.Must(server.Start())
//
//	//server.Close()
//
//	time.Sleep(100 * time.Second)
//}
//
//// 这是server
//func TestStartV2RayByJsonServer(t *testing.T) {
//
//	cfgStr := `
//	{
//  "inbounds": [{
//    "port": 10086,
//    "protocol": "vmess",
//    "settings": {
//      "clients": [{ "id": "b831381d-6324-4d53-ad4f-8cda48b30811" }]
//    }
//  }],
//  "outbounds": [{
//    "protocol": "freedom",
//    "settings": {}
//  }]
//}
//`
//
//	cfgBytes := []byte(cfgStr)
//
//	var sig = make(chan int)
//	go func() {
//		log.Println("first")
//		err := StartV2RayByJSONStart(cfgBytes)
//		common.Must(err)
//	}()
//
//	go func() {
//		log.Println("second")
//		err := StartV2RayByJSONStart(cfgBytes)
//		common.Must(err)
//	}()
//
//	<-sig
//}
//
//// 这是client
//func TestStartV2RayByJsonClient(t *testing.T) {
//	cfgStr := `
//{
//  "inbounds": [
//    {
//      "port": 1080,
//      "listen": "127.0.0.1",
//      "protocol": "socks",
//      "settings": {
//        "udp": true
//      }
//    }
//  ],
//  "outbounds": [
//    {
//      "protocol": "vmess",
//      "settings": {
//        "vnext": [
//          {
//            "address": "server",
//            "port": 10086,
//            "users": [{ "id": "b831381d-6324-4d53-ad4f-8cda48b30811" }]
//          }
//        ]
//      }
//    },
//    {
//      "protocol": "freedom",
//      "tag": "direct",
//      "settings": {}
//    }
//  ]
//}
//	`
//
//	cfgBytes := []byte(cfgStr)
//
//	var sig = make(chan int)
//	go func() {
//		log.Println("first")
//		err := StartV2RayByJSONStart(cfgBytes)
//		common.Must(err)
//	}()
//
//	go func() {
//		log.Println("second")
//		err := StartV2RayByJSONStart(cfgBytes)
//		common.Must(err)
//	}()
//
//	<-sig
//}
//
//// 这是client
//func TestStartByJson(t *testing.T) {
//	cfgStr := `
//{
//  "inbounds": [
//    {
//      "port": 1080,
//      "listen": "127.0.0.1",
//      "protocol": "socks",
//      "settings": {
//        "udp": true
//      }
//    }
//  ],
//  "outbounds": [
//    {
//      "protocol": "vmess",
//      "settings": {
//        "vnext": [
//          {
//            "address": "server",
//            "port": 10086,
//            "users": [{ "id": "b831381d-6324-4d53-ad4f-8cda48b30811" }]
//          }
//        ]
//      }
//    },
//    {
//      "protocol": "freedom",
//      "tag": "direct",
//      "settings": {}
//    }
//  ]
//}
//	`
//
//	cfgBytes := []byte(cfgStr)
//
//	var sig = make(chan int)
//	log.Println("first")
//	err := StartV2RayByJSONStart(cfgBytes)
//	common.Must(err)
//
//	<-sig
//}

func TestStartInstance(t *testing.T) {
	cfgStr := `
	{
	   "policy": {
	       "levels": {
	           "1": {
	               "handshake": 4,
	               "connIdle": 300,
	               "uplinkOnly": 0,
	               "downlinkOnly": 0,
	               "statsUserUplink": false,
	               "statsUserDownlink": false,
	               "bufferSize": 10240
	           }
	       }
	   },
	   "dns": {
	       "servers": [
	           "localhost",
	           "223.5.5.5",
	           "1.1.1.1",
	           "8.8.8.8"
	       ]
	   },
	   "inbounds": [
	       {
	           "settings": {
	               "ip": "127.0.0.1",
	               "udp": true,
	               "allowTransparent": true
	           },
	           "port": 9999,
	           "protocol": "socks",
	           "listen": "127.0.0.1",
	           "sniffing": {
	               "enabled": true,
	               "destOverride": [
	                   "http",
	                   "tls"
	               ]
	           },
	           "streamSettings": {
	               "auth": "noauth",
	               "udp": true,
	               "ip": "127.0.0.1"
	           }
	       }
	   ],
	   "outbounds": [
	       {
	           "settings": {
	               "vnext": [
	                   {
	                       "address": "peozeo.xyz",
	                       "port": 443,
	                       "udp": true,
	                       "users": [
	                           {
	                               "level": 1,
	                               "alterId": 16,
	                               "security": "aes-128-gcm",
	                               "id": "66d88922-c880-4b5a-906e-ff50fc8c209f"
	                           }
	                       ]
	                   }
	               ]
	           },
	           "streamSettings": {
	               "network": "ws",
	               "security": "tls",
	               "sockopt": {
	                   "tcpFastOpen": true
	               },
	               "wsSettings": {
	                   "connectionReuse": true,
	                   "path": "/llc"
	               },
	               "tlsSettings": {
	                   "allowInsecure": true
	               }
	           },
	           "mux": {
	               "enabled": true,
	               "concurrency": 8
	           },
	           "protocol": "vmess",
	           "tag": "proxy"
	       },
	       {
	           "settings": {
	               "domainStrategy": "UseIP"
	           },
	           "streamSettings": {
	               "sockopt": {
	                   "tcpFastOpen": true
	               }
	           },
	           "protocol": "freedom",
	           "tag": "direct"
	       },
	       {
	           "protocol": "blackhole",
	           "tag": "block",
	           "settings": {}
	       }
	   ],
	   "routing": {
	       "strategy": "rules",
	       "settings": {
	           "domainStrategy": "IPIfNonMatch",
	           "rules": [
	               {
	                   "type": "field",
	                   "outboundTag": "direct",
	                   "ip": [
	                       "0.0.0.0/8",
	                       "10.0.0.0/8",
	                       "100.64.0.0/10",
	                       "127.0.0.0/8",
	                       "169.254.0.0/16",
	                       "172.16.0.0/12",
	                       "192.0.0.0/24",
	                       "192.0.2.0/24",
	                       "192.168.0.0/16",
	                       "198.18.0.0/15",
	                       "198.51.100.0/24",
	                       "203.0.113.0/24",
	                       "::1/128",
	                       "fc00::/7",
	                       "fe80::/10"
	                   ]
	               }
	           ]
	       }
	   }
	}
		`

	pbConfig, err := serial.LoadJSONConfig(bytes.NewReader([]byte(cfgStr)))
	if err != nil {
		//log.Println(err)
		return
	}

	pbBytes, err := proto.Marshal(pbConfig)
	if err != nil {
		//log.Println(err)
		return
	}

	core.New(pbConfig)

	var sig = make(chan int)

	err = StartInstance(pbBytes)

	common.Must(err)

	<-sig

}

func TestStartInstanceConflict(t *testing.T) {
	cfgStr := `
	{
	   "policy": {
	       "levels": {
	           "1": {
	               "handshake": 4,
	               "connIdle": 300,
	               "uplinkOnly": 0,
	               "downlinkOnly": 0,
	               "statsUserUplink": false,
	               "statsUserDownlink": false,
	               "bufferSize": 10240
	           }
	       }
	   },
	   "dns": {
	       "servers": [
	           "localhost",
	           "223.5.5.5",
	           "1.1.1.1",
	           "8.8.8.8"
	       ]
	   },
	   "inbounds": [
	       {
	           "settings": {
	               "ip": "127.0.0.1",
	               "udp": true,
	               "allowTransparent": true
	           },
	           "port": 9999,
	           "protocol": "socks",
	           "listen": "127.0.0.1",
	           "sniffing": {
	               "enabled": true,
	               "destOverride": [
	                   "http",
	                   "tls"
	               ]
	           },
	           "streamSettings": {
	               "auth": "noauth",
	               "udp": true,
	               "ip": "127.0.0.1"
	           }
	       }
	   ],
	   "outbounds": [
	       {
	           "settings": {
	               "vnext": [
	                   {
	                       "address": "peozeo.xyz",
	                       "port": 443,
	                       "udp": true,
	                       "users": [
	                           {
	                               "level": 1,
	                               "alterId": 16,
	                               "security": "aes-128-gcm",
	                               "id": "66d88922-c880-4b5a-906e-ff50fc8c209f"
	                           }
	                       ]
	                   }
	               ]
	           },
	           "streamSettings": {
	               "network": "ws",
	               "security": "tls",
	               "sockopt": {
	                   "tcpFastOpen": true
	               },
	               "wsSettings": {
	                   "connectionReuse": true,
	                   "path": "/llc"
	               },
	               "tlsSettings": {
	                   "allowInsecure": true
	               }
	           },
	           "mux": {
	               "enabled": true,
	               "concurrency": 8
	           },
	           "protocol": "vmess",
	           "tag": "proxy"
	       },
	       {
	           "settings": {
	               "domainStrategy": "UseIP"
	           },
	           "streamSettings": {
	               "sockopt": {
	                   "tcpFastOpen": true
	               }
	           },
	           "protocol": "freedom",
	           "tag": "direct"
	       },
	       {
	           "protocol": "blackhole",
	           "tag": "block",
	           "settings": {}
	       }
	   ],
	   "routing": {
	       "strategy": "rules",
	       "settings": {
	           "domainStrategy": "IPIfNonMatch",
	           "rules": [
	               {
	                   "type": "field",
	                   "outboundTag": "direct",
	                   "ip": [
	                       "0.0.0.0/8",
	                       "10.0.0.0/8",
	                       "100.64.0.0/10",
	                       "127.0.0.0/8",
	                       "169.254.0.0/16",
	                       "172.16.0.0/12",
	                       "192.0.0.0/24",
	                       "192.0.2.0/24",
	                       "192.168.0.0/16",
	                       "198.18.0.0/15",
	                       "198.51.100.0/24",
	                       "203.0.113.0/24",
	                       "::1/128",
	                       "fc00::/7",
	                       "fe80::/10"
	                   ]
	               }
	           ]
	       }
	   }
	}
		`

	pbConfig, err := serial.LoadJSONConfig(bytes.NewReader([]byte(cfgStr)))
	if err != nil {
		//log.Println(err)
		return
	}

	pbBytes, err := proto.Marshal(pbConfig)
	if err != nil {
		//log.Println(err)
		return
	}

	core.New(pbConfig)

	var sig = make(chan int)

	go func(pbBytes []byte) {
		err = StartInstance(pbBytes)

		common.Must(err)
	}(pbBytes)

	go func(pbBytes []byte) {
		err = StartInstance(pbBytes)

		common.Must(err)
	}(pbBytes)

	<-sig

}
