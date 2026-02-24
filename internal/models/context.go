package models

// It runs through the chain of managers.
//
// Used instead of context.Context because cancellations are not required + static typing
type UpdateContext struct {
	Moder *Moderator
}
