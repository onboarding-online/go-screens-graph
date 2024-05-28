package models

// AssetUrl - Asset from storage
type AssetUrl struct {

	Origin string `json:"origin"`

	Thumb string `json:"thumb"`

	Small string `json:"small"`

	Medium string `json:"medium"`

	Normal string `json:"normal"`
}
