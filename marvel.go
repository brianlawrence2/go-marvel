package marvel

import (
    "crypto/md5"
    //"fmt"
    "time"
    "encoding/hex"
    //"encoding/json"
    "net/http"
    //"net/url"
)

const (
  marvelAPIURL = "http://gateway.marvel.com/v1/public/"
  userAgent    = "go-marvel" + version
  version      = "0.1.0"
)

func Hash(pubKey string, privKey string, ts string) string {
  hash := md5.New()
  st   := ts + privKey + pubKey
  hash.Write([]byte(st))
  return hex.EncodeToString(hash.Sum(nil))
}

func Url(pubKey string, privKey string, ts string, endpoint string) string {
  return marvelAPIURL + endpoint + "?ts=" + ts + "&apikey=" + pubKey + "&hash=" + Hash(pubKey, privKey, ts)
}

func GetCharacters(publicKey, privateKey, endpoint string) (*http.Response, error) {
  t    := time.Now().Format("20060102150405")
  //hash := Hash(publicKey, privateKey, t)

  url := Url(publicKey, privateKey, t, endpoint)

  response, err := http.Get(url)

  return response, err
}
