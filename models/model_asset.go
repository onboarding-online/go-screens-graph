package models

// Asset - Asset
type Asset struct {

	AssetUrl AssetUrl `json:"assetUrl,omitempty"`

	// Asset name
	AssetName string `json:"assetName,omitempty"`
}
