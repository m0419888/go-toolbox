package common

type PutBody struct {
	Exigency bool `json:"exigency"`
	Data     any  `json:"data"`
}
