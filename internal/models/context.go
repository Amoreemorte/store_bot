package models

// It runs through the chain of managers.
//
// Used instead of context.Context because cancellations are not required + static typing
type UpdateContext struct {
	Update          *Update
	IsModerator     bool
	ValidationError error
	ModeratorState  *ModeratorStateRef
	ValidationError error
	CollectionName  *string
	Msg             *MessageConfig
}
