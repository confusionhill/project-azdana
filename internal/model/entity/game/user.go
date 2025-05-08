package game

import (
	"database/sql"
	"time"
)

type User struct {
	ID                int64          `db:"id"`
	Username          string         `db:"username"`
	Password          string         `db:"password"`
	Access            int64          `db:"access"`
	Upgrade           int64          `db:"upgrade"`
	Age               int64          `db:"age"`
	UpgDate           time.Time      `db:"upgDate"`
	UpgDays           int64          `db:"upgDays"`
	EmailActive       int64          `db:"emailActive"`
	Email             string         `db:"email"`
	Moderator         int64          `db:"moderator"`
	Level             int64          `db:"level"`
	CosColorAccessory int64          `db:"cosColorAccessory"`
	CosColorBase      int64          `db:"cosColorBase"`
	CosColorTrim      int64          `db:"cosColorTrim"`
	PlaColorSkin      int64          `db:"plaColorSkin"`
	PlaColorHair      int64          `db:"plaColorHair"`
	PlaColorEyes      int64          `db:"plaColorEyes"`
	SlotBag           int64          `db:"slotBag"`
	SlotBank          int64          `db:"slotBank"`
	SlotHouse         int64          `db:"slotHouse"`
	STR               int64          `db:"STR"`
	DEX               int64          `db:"DEX"`
	INT               int64          `db:"INT"`
	END               int64          `db:"END"`
	WIS               int64          `db:"WIS"`
	LCK               int64          `db:"LCK"`
	CurrentClass      int64          `db:"currentClass"`
	XP                int64          `db:"xp"`
	Gold              int64          `db:"gold"`
	Coins             int64          `db:"coins"`
	LastVisited       string         `db:"lastVisited"`
	HairID            int64          `db:"hairID"`
	Gender            string         `db:"gender"`
	HairName          string         `db:"hairName"`
	HairFile          string         `db:"hairFile"`
	CurServer         string         `db:"curServer"`
	Banned            int64          `db:"banned"`
	VIP               int64          `db:"vip"`
	UG                int64          `db:"ug"`
	DOB               string         `db:"dob"`
	SignupIP          sql.NullString `db:"signupip"`
	LoginIP           sql.NullString `db:"loginip"`
}
