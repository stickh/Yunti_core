package yunti_core

import "net"
import "golang.org/x/net/websocket"

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



func Leaptransp_websocket_Connect(wsaddr string){



}
