package yunti_core

import net

func exitTCPDialer(addr string)(net.Conn,error){
  return net.Dial("tcp",addr)
}
