package models

type StoragePool struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	TargetPath string `json:"target_path"`
	Host       string `json:"host "`
}
