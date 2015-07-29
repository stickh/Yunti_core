package yunti_core

import "net"
import "golang.org/x/net/websocket"
import "net/url"
import "golang.org/x/crypto/sha3"
import "encoding/base64"
import "strings"

type leaptransp_stage int

const (
    leaptransp_stage_NULL leaptransp_stage = iota
    leaptransp_stage_initialized
    leaptransp_stage_connected
    leaptransp_stage_serverauthdone
    leaptransp_stage_clientauthdone
    leaptransp_stage_syncconnest
    leaptransp_stage_asyncconnest
    leaptransp_stage_failed
)

type Leaptransp struct{
  Leaptype string
  Stage leaptransp_stage
  Sync bool
  conn interface
}

type Leapauth struct{
  Authtype string
  AuthKey string
  Id string
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

func (lt *leaptransp)Leaptransp_websocket_Connect(wsaddr string)int{

origin:=leaptransp_websocket_Connect_genorigin(wsaddr)

wsc,err:=websocket.Dial(wsaddr, "", origin)

if err != nil {
  return -1
}

lt.conn=&wsc

}


func (lt *leaptransp)SendT(dt string)int{


switch lt.Leaptype {
case "websocket":
  lt.conn.Message.Send(dt)
default:
  return -1
}
}

func (lt *leaptransp)RecvT(dt *string)int{


switch lt.Leaptype {
case "websocket":
  lt.conn.Message.Receive(dt)
default:
  return -1
}

}



func (lt *leaptransp)SendB(dt []byte)int{


switch lt.Leaptype {
case "websocket":
  lt.conn.Message.Send(dt)
default:
  return -1
}
}

func (lt *leaptransp)RecvB(dt *[]byte)int{


switch lt.Leaptype {
case "websocket":
  lt.conn.Message.Receive(dt)
default:
  return -1
}

}

func Leaptransp_auth_simple_password_constructMsgAsT(la *Leapauth)string{

s:= la.AuthKey +"!+AUTH"

x:=sha3.Sum512([]byte(s))

d:=make([]byte,0,64)

d:=append(d,x...)

o:=base64.StdEncoding.EncodeToString(d)

oo:="AUTH@simple_password@"+la.Id+"@"+o

return oo


}

func (lt *leaptransp)Leaptransp_auth(la Leapauth)int{

switch la.Authtype {
case "simple_password":
  lt.SendT(Leaptransp_auth_simple_password_constructMsgAsT(la))
default:
  return -1
}


}

func Leaptransp_authV_simple_password(la *map[string]Leapauths,un,p6 string)int{

  if lu,ok:=Leapauths[un];ok{

    s:= lu.AuthKey +"!+AUTH"

    x:=sha3.Sum512([]byte(s))

    d:=make([]byte,0,64)

    d:=append(d,x...)

    o:=base64.StdEncoding.EncodeToString(d)

  if o==p6 {
    return 0 //auth ok
  }
  }else{
    return 1 //password isn't correct
  }

}else{
  return -1 //no such user
}


func (lt *leaptransp)Leaptransp_authV(la *map[string]Leapauths, incD string )string{

ins:=string.Split(incD,"@")

switch ins[1]{
case "simple_password":

Leaptransp_authV_simple_password(la,ins[2],ins[3])

default:
  return 0
}



}
