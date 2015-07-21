package yunti_core

import "github.com/agl/ed25519"
import "crypto/rand"
import "encoding/base64"
import "golang.org/x/crypto/nacl/box"

type AuthKey struct {
  KeyType string `json:"KeyType"`
  EncodeingMode string     `json:"EncodeingMode"`
  KeyDataV string `json:"KeyDataV"`
  KeyDataB string  `json:"KeyDataB"`
}

func GenAuthKey_ed25519()AuthKey{
var key AuthKey
puk,pvk,err:= ed25519.GenerateKey(rand)
key.KeyType="ed25519"
key.EncodeingMode="b64"
key.KeyDataB=base64.EncodeToString(*puk)
key.KeyDataV=base64.EncodeToString(*pvk)
return key
}

func Sign(key AuthKey,data string)(signature string,ok int){
switch key.KeyType {
case "ed25519":
   var pvk []byte
   switch key.EncodeingMode{
   case "b64":
     pvk,_:=base64.DecodeString(key.KeyDataV)
     sig:=ed25519.Sign(pvk,[]byte(data))
     return (base64.EncodeToString(sig),0)
     break
   default:
     return (nil,-2) //Cannot sign, EncodeingMode unknown
   }
  break
default:
  return (nil,-1) //Cannot Sign, keytype unknown
}
}

func Verify(key AuthKey,data,signature string)(ok int){

  switch key.KeyType {
  case "ed25519":
     var pvk []byte
     switch key.EncodeingMode{
     case "b64":

      puk,_=base64.DecodeString(key.KeyDataB)
      Vr:=ed25519.Verify(puk,[]byte(data),base64.DecodeString(signature))

      if Vr == true {
        return 0
      }else{
        return 1
      }

       break
     default:
       return -2 //Cannot verify, EncodeingMode unknown
     }
    break
  default:
    return -1 //Cannot verify, keytype unknown
  }

}



type EncryptionKey struct {
  KeyType string `json:"KeyType"`
  EncodeingMode string     `json:"EncodeingMode"`
  KeyDataV string `json:"KeyDataV"`
  KeyDataB string  `json:"KeyDataB"`
}


func GenEncryptionKey_Nacl(){

  puk,pvk,_:=box.GenerateKey(rand)
  var key EncryptionKey

  key.KeyType="Nacl"
  key.EncodeingMode="b64"
  key.KeyDataB=base64.EncodeToString(*puk)
  key.KeyDataV=base64.EncodeToString(*pvk)

  return key

}

type EncryptionStage struct{
  EncryptionType string
  SynKey []byte
  counter uint64=0
}


func PrepareEncrypt(EncryptionKeySelf,EncryptionKeyRemote EncryptionKey)(EncryptionStage,int){


switch key.KeyType {
case "Nacl":
   var pvk []byte
   switch key.EncodeingMode{
   case "b64":

     puk,_=base64.DecodeString(EncryptionKeyRemote.KeyDataB)
     pvk,_=base64.DecodeString(EncryptionKeySelf.KeyDataV)
     var []byte sk
     box.Precompute(sk,puk,pvk)

    var encs  EncryptionStage
    encs.EncryptionType="Nacl"
    encs.SynKey=sk
    encs.counter+=1

    return (encs,0)
     break
   default:
     return (nil,-2) //Cannot prepare, EncodeingMode unknown
   }
  break
default:
  return (nil,-1) //Cannot prepare, keytype unknown
}

}

func DoEncrypt(es *EncryptionStage,dt []byte)(data,int){


switch key.KeyType {
case "Nacl":
   var pvk []byte
   switch key.EncodeingMode{
   case "b64":

    o=make([]byte,len(dt)+box.Overhead+24)

    nc:=rand.Read(24)

    SealAfterPrecomputation(o,dt,nc,es.SynKey)

    o[len(dt)+box.Overhead-1:-1]=nc

    es.counter+=1

    return (o,0)
     break
   default:
     return (nil,-2) //Cannot encrypt, EncodeingMode unknown
   }
  break
default:
  return (nil,-1) //Cannot encrypt, keytype unknown
}

}


func DoDecrypt(es *EncryptionStage,dt []byte)(data,int){


switch key.KeyType {
case "Nacl":
   var pvk []byte
   switch key.EncodeingMode{
   case "b64":

    o=make([]byte,len(dt)-box.Overhead)

    nc:=n

    OpenAfterPrecomputation(o,dt,nc,es.SynKey)

    es.counter+=1

    return (o,0)
     break
   default:
     return (nil,-2) //Cannot decrypt, EncodeingMode unknown
   }
  break
default:
  return (nil,-1) //Cannot decrypt, keytype unknown
}

}
