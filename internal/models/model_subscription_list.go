package models

// SubscriptionList - Subscriptions list element
type SubscriptionList struct {

	SubscriptionViewKind SubscriptionViewKind `json:"subscriptionViewKind,omitempty"`

	ItemType SubscriptionListItemType `json:"itemType"`

	// Item rows
	Items []ItemTypeSubscription `json:"items"`

	Styles SubscriptionListBlock `json:"styles"`

	SelectedBlock SelectedSubscriptionListItemBlock `json:"selectedBlock"`

	Box Box `json:"box"`
}
