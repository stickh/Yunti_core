package yunti_core

import "net"
import "golang.org/x/net/websocket"
import "net/url"

type leaptransp_stage int

const (
    leaptransp_stage_NULL leaptransp_stage = iota
    leaptransp_stage_initialized
    leaptransp_stage_connected
    leaptransp_stage_serverauthdone
    leaptransp_stage_clientauthdone
    leaptransp_stage_syncconnest
    leaptransp_stage_asyncconnest
)

type leaptransp struct{
  Leaptype string
  Stage leaptransp_stage
  Sync bool
}

func leaptransp_websocket_Connect_genorigin(wsaddr string)string{
  oriu,_:=url.Parse(wsaddr)
  switch strings.ToUpper(oriu.Scheme){
  case "WSS":
  oriu.Scheme="https"
  case "WS":
  oriu.Scheme="http"
}
origin=oriu.String()
return origin
}

func Leaptransp_websocket_Connect(wsaddr string){

origin:=leaptransp_websocket_Connect_genorigin(wsaddr)


}
