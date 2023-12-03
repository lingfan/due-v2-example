package code

import (
	"github.com/dobyte/due/v2/codes"
)

var (
	InternalServerError = codes.NewCode(1, "internal server error")
	NoPermission        = codes.NewCode(2, "no permission")
	IllegalParams       = codes.NewCode(3, "illegal params")
	TokenExpired        = codes.NewCode(4, "token is expired")
	TokenInvalid        = codes.NewCode(5, "token is invalid")

	NotFoundUser  = codes.NewCode(1000, "not found user")
	WrongPassword = codes.NewCode(1001, "wrong password")

	NotFoundMail = codes.NewCode(1022, "not found mail")

	NotFoundRoom            = codes.NewCode(10000, "not found room")
	RoomIsFull              = codes.NewCode(10000, "room is full")
	NotFoundTable           = codes.NewCode(10000, "not found table")
	TableIsFull             = codes.NewCode(10000, "table is full")
	TableIsEmpty            = codes.NewCode(10000, "table is empty")
	SeatIsEmpty             = codes.NewCode(10001, "seat is empty")
	SeatAlreadyTaken        = codes.NewCode(10002, "seat already taken")
	CoinsAreBelowEntryLimit = codes.NewCode(10002, "coins are below the entry limit")
	CoinsAreOverEntryLimit  = codes.NewCode(10002, "coins are over the entry limit")
	NoMatchingRoom          = codes.NewCode(10000, "no matching room")
	NoMatchingTable         = codes.NewCode(10000, "no matching table")
	NotFoundPlayer          = codes.NewCode(10000, "not found player")
	PlayerAlreadySeated     = codes.NewCode(10000, "player is already seated")
	PlayerNotInSeat         = codes.NewCode(10000, "player not in seat")
	PlayerIsReady           = codes.NewCode(10000, "player is readying")
)
