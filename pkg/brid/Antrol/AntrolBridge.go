package Antrol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type Respon_MentahDTO struct {
	MetaData struct {
		Code    interface{} `json:"code"`
		Message interface{} `json:"message"`
	} `json:"metadata"`
	Response string `json:"response"`
}

type Respon_DTO struct {
	MetaData struct {
		Code    interface{} `json:"code"`
		Message interface{} `json:"message"`
	} `json:"metadata"`
	Response interface{} `json:"response"`
}

func GetRequest(endpoint string, cfg interface{}) (interface{}, error) {
	conf := ConfigBpjs{}
	// err := smapping.FillStruct(&conf, smapping.MapFields(&cfg))
	// if err != nil {
	// 	log.Fatalf("Failed map %v: ", err)
	// }
	err := mapstructure.Decode(cfg, &conf)
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	cons_id, Secret_key, User_key, tstamp, X_signature := SetHeader(conf)

	req, _ := http.NewRequest("GET", conf.Url_antrean+endpoint, nil)

	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")
	req.Header.Add("X-cons-id", cons_id)
	req.Header.Add("X-timestamp", tstamp)
	req.Header.Add("X-signature", X_signature)
	req.Header.Add("user_key", User_key)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var resp_mentah Respon_MentahDTO
	var resp Respon_DTO
	json.Unmarshal([]byte(body), &resp_mentah)
	fmt.Println("mentah =>", resp_mentah)
	resp_decrypt, _ := ResponseAntrol(string(resp_mentah.Response), string(cons_id+Secret_key+tstamp))
	resp.MetaData = resp_mentah.MetaData
	json.Unmarshal([]byte(resp_decrypt), &resp.Response)

	fmt.Println(resp)

	return &resp, nil
}

func PostRequest(endpoint string, cfg interface{}, data interface{}) (interface{}, error) {

	conf := ConfigBpjs{}

	// err := smapping.FillStruct(&conf, smapping.MapFields(&cfg))
	// if err != nil {
	// 	log.Fatalf("Failed map %v: ", err)
	// }

	err := mapstructure.Decode(cfg, &conf)
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	cons_id, Secret_key, User_key, tstamp, X_signature := SetHeader(conf)

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(data)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", conf.Url_antrean+endpoint, &buf)

	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")
	req.Header.Add("X-cons-id", cons_id)
	req.Header.Add("X-timestamp", tstamp)
	req.Header.Add("X-signature", X_signature)
	req.Header.Add("user_key", User_key)

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var resp_mentah Respon_MentahDTO
	var resp Respon_DTO
	json.Unmarshal([]byte(body), &resp_mentah)

	resp_decrypt, _ := ResponseAntrol(string(resp_mentah.Response), string(cons_id+Secret_key+tstamp))
	resp.MetaData = resp_mentah.MetaData
	json.Unmarshal([]byte(resp_decrypt), &resp.Response)

	return &resp, nil

}

func DeleteRequest(endpoint string, cfg interface{}, data interface{}) (interface{}, error) {

	conf := ConfigBpjs{}

	// err := smapping.FillStruct(&conf, smapping.MapFields(&cfg))
	// if err != nil {
	// 	log.Fatalf("Failed map %v: ", err)
	// }

	err := mapstructure.Decode(cfg, &conf)
	if err != nil {
		return nil, err
	}

	cons_id, Secret_key, User_key, tstamp, X_signature := SetHeader(conf)

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(data)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("DELETE", conf.Url_antrean+endpoint, &buf)

	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")
	req.Header.Add("X-cons-id", cons_id)
	req.Header.Add("X-timestamp", tstamp)
	req.Header.Add("X-signature", X_signature)
	req.Header.Add("user_key", User_key)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var resp_mentah Respon_MentahDTO
	var resp Respon_DTO
	json.Unmarshal([]byte(body), &resp_mentah)

	resp_decrypt, _ := ResponseAntrol(string(resp_mentah.Response), string(cons_id+Secret_key+tstamp))
	resp.MetaData = resp_mentah.MetaData
	json.Unmarshal([]byte(resp_decrypt), &resp.Response)

	return &resp, nil

}

func PutRequest(endpoint string, cfg interface{}, data interface{}) (interface{}, error) {

	conf := ConfigBpjs{}

	// err := smapping.FillStruct(&conf, smapping.MapFields(&cfg))
	// if err != nil {
	// 	log.Fatalf("Failed map %v: ", err)
	// }

	err := mapstructure.Decode(cfg, &conf)
	if err != nil {
		return nil, err
	}

	cons_id, Secret_key, User_key, tstamp, X_signature := SetHeader(conf)

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(data)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("PUT", conf.Url_antrean+endpoint, &buf)

	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")
	req.Header.Add("X-cons-id", cons_id)
	req.Header.Add("X-timestamp", tstamp)
	req.Header.Add("X-signature", X_signature)
	req.Header.Add("user_key", User_key)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var resp_mentah Respon_MentahDTO
	var resp Respon_DTO
	json.Unmarshal([]byte(body), &resp_mentah)

	resp_decrypt, _ := ResponseAntrol(string(resp_mentah.Response), string(cons_id+Secret_key+tstamp))
	resp.MetaData = resp_mentah.MetaData
	json.Unmarshal([]byte(resp_decrypt), &resp.Response)

	return &resp, nil

}
