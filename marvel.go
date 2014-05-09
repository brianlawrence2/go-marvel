package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"
)

const (
	marvelAPIURL = "http://gateway.marvel.com/v1/public/"
	userAgent    = "go-marvel" + version
	version      = "0.1.0"
)

type Request struct {
	PublicKey  string
	PrivateKey string
	Endpoint   string
	Id         string
	Item       string
}

func Hash(pubKey string, privKey string, ts string) string {
	hash := md5.New()
	st := ts + privKey + pubKey
	hash.Write([]byte(st))
	return hex.EncodeToString(hash.Sum(nil))
}

func Url(pubKey string, privKey string, ts string, endpoint string) string {
	return marvelAPIURL + endpoint + "?ts=" + ts + "&apikey=" + pubKey + "&hash=" + Hash(pubKey, privKey, ts)
}

func (cr *Request) Get() (*http.Response, error) {
	t := time.Now().Format("20060102150405")

	url := Url(cr.PublicKey, cr.PrivateKey, t, cr.Endpoint)

	response, err := http.Get(url)

	return response, err
}
