package models

type Permission struct {
	ID        int
	Role      string
	Module    string
	CanView   bool
	CanCreate bool
	CanEdit   bool
	CanDelete bool
}