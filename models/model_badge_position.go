package models
// BadgePosition : Badge position
type BadgePosition string

// List of BadgePosition
const (
	BADGEPOSITION_LEFT_START BadgePosition = "leftStart"
	BADGEPOSITION_LEFT_END BadgePosition = "leftEnd"
	BADGEPOSITION_LEFT_CENTER BadgePosition = "leftCenter"
	BADGEPOSITION_RIGHT_START BadgePosition = "rightStart"
	BADGEPOSITION_RIGHT_END BadgePosition = "rightEnd"
	BADGEPOSITION_RIGHT_CENTER BadgePosition = "rightCenter"
	BADGEPOSITION_TOP_LEFT BadgePosition = "topLeft"
	BADGEPOSITION_TOP_RIGHT BadgePosition = "topRight"
	BADGEPOSITION_TOP_CENTER BadgePosition = "topCenter"
	BADGEPOSITION_BOTTOM_LEFT BadgePosition = "bottomLeft"
	BADGEPOSITION_BOTTOM_RIGHT BadgePosition = "bottomRight"
	BADGEPOSITION_BOTTOM_CENTER BadgePosition = "bottomCenter"
)
