package models

// ItemTypeSubscription - Structure of full collection item type for subscription list
type ItemTypeSubscription struct {

	SubscriptionId string `json:"subscriptionId"`

	CheckBox CheckBox `json:"checkBox"`

	LeftLabelTop Text `json:"leftLabelTop"`

	LeftLabelBottom Text `json:"leftLabelBottom"`

	RightLabelTop Text `json:"rightLabelTop"`

	RightLabelBottom Text `json:"rightLabelBottom"`

	DescriptionLabel Text `json:"descriptionLabel,omitempty"`

	PurchaseButtonLabel Text `json:"purchaseButtonLabel,omitempty"`

	IsSelected bool `json:"isSelected"`

	Badge Badge `json:"badge,omitempty"`

	Image Image `json:"image,omitempty"`

	Box Box `json:"box"`

	Weight float32 `json:"weight"`

	Styles ItemTypeSubscriptionBlock `json:"styles"`
}
