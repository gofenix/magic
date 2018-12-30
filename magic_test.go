package magic

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
