package models
// PurchaseFlag : Purchase flag
type PurchaseFlag string

// List of PurchaseFlag
const (
	PURCHASEFLAG_SHOULD_CLOSE_AFTER_RESTORE PurchaseFlag = "ShouldCloseAfterRestore"
	PURCHASEFLAG_SHOULD_CLOSE_AFTER_SUBSCRIPTION_ERROR PurchaseFlag = "ShouldCloseAfterSubscriptionError"
)
