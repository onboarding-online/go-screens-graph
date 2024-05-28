package models
// ActionKind : Describe of action kind
type ActionKind string

// List of ActionKind
const (
	ACTIONKIND_PURCHASE ActionKind = "Purchase"
	ACTIONKIND_NEXT ActionKind = "Next"
	ACTIONKIND_BACK ActionKind = "Back"
	ACTIONKIND_FINISH ActionKind = "Finish"
	ACTIONKIND_DUMMY ActionKind = "Dummy"
)
