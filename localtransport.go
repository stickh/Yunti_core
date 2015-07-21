package yunti_core

// WARNING: this is a modifiyed version.
import socksD "github.com/xiaokangwang/go-socks5"

import strconv


func ListenSocks5DockerPart(port int,Dialer *func){
  var config socksD.Config
  config.AdvancedServe=true
  config.AdvancedDial=Dialer
  serv:=socksD.New(config)
  go serv.ListenAndServe("tcp","127.0.0.1:"+strconv.Itoa(port))
}
