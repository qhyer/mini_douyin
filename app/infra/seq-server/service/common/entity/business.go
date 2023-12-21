package entity

import "encoding/json"

type SnowFlakeEtcdHolder struct {
	EtcdAddressNode string // 保存自身的key  ip:port-000000001
	ListenAddress   string // 保存自身的key ip:port
	IP              string
	Port            string
	LastUpdateTime  int64
	WorkerId        int
}

type Endpoint struct {
	IP        string `json:"ip"`
	Port      string `json:"port"`
	Timestamp int64  `json:"timestamp"`
}

func (e *Endpoint) EncodeToJSONString() (string, error) {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (e *Endpoint) DecodeFromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), e)
	if err != nil {
		return err
	}
	return nil
}
