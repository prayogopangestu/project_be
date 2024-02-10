package Antrol

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type ConfigBpjs struct {
	Cons_id        string
	Secret_key     string
	User_key       string
	User_keyAntrol string
	Url_vclaim     string
	Url_antrean    string
	Ppk_pelayanan  string
}

func SetHeader(cfg ConfigBpjs) (string, string, string, string, string) {

	timenow := time.Now().UTC()
	t, err := time.Parse(time.RFC3339, "1970-01-01T00:00:00Z")
	if err != nil {
		log.Fatal(err)
	}

	tstamp := timenow.Unix() - t.Unix()
	secret := []byte(cfg.Secret_key)
	message := []byte(cfg.Cons_id + "&" + fmt.Sprint(tstamp))
	hash := hmac.New(sha256.New, secret)
	hash.Write(message)
	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))
	// to base64
	X_signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return cfg.Cons_id, cfg.Secret_key, cfg.User_keyAntrol, fmt.Sprint(tstamp), X_signature

}
