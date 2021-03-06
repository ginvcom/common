// Code generated by goctl. DO NOT EDIT.
package types

type StsReq struct {
	Bucket string `json:"bucket"` // bucket
}

type StsReply struct {
	Region        string `json:"region"`
	Endpoint      string `json:"endpoint"`
	Bucket        string `json:"bucket"`
	AccessKey     string `json:"accessKeyId"`
	SecretKey     string `json:"accessKeySecret"`
	SecurityToken string `json:"stsToken"`
}
