package yunti_core

import net

type YuntiProxyConn struct{}

func(c *YuntiProxyConn)Close(){
  //TODO
}

func(c *YuntiProxyConn)Write(b []byte) (n int, err error){
  //TODO
}

func(c *YuntiProxyConn)Read(b []byte) (n int, err error){
  //TODO
}

func ProxyDialer(network, address string) (net.Conn, error){

  var conninsT YuntiProxyConn

  return conninsT
}
