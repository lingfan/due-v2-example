package entity

import (
	"due-v2-example/shared/code"

	"github.com/dobyte/due/v2/errors"
)

type RoomMgr struct {
	rooms map[int]*Room
}

func NewRoomMgr(opts []*Options) *RoomMgr {
	mgr := &RoomMgr{rooms: make(map[int]*Room, len(opts))}
	for i := range opts {
		room := newRoom(opts[i])
		mgr.rooms[room.ID()] = room
	}

	return mgr
}

// GetRoom 获取房间
// code.NotFoundRoom
func (mgr *RoomMgr) GetRoom(roomID int) (*Room, error) {
	room, ok := mgr.rooms[roomID]
	if !ok {
		return nil, errors.NewError(code.NotFoundRoom)
	}

	return room, nil
}

// GetTable 获取牌桌
// code.NotFoundRoom
// code.NotFoundTable
func (mgr *RoomMgr) GetTable(roomID int, tableID int) (*Table, error) {
	room, ok := mgr.rooms[roomID]
	if !ok {
		return nil, errors.NewError(code.NotFoundRoom)
	}

	return room.GetTable(tableID)
}

// GetSeat 获取座位
// code.NotFoundRoom
// code.NotFoundTable
// code.IllegalParams
func (mgr *RoomMgr) GetSeat(roomID, tableID, seatID int) (*Seat, error) {
	room, ok := mgr.rooms[roomID]
	if !ok {
		return nil, errors.NewError(code.NotFoundRoom)
	}

	return room.GetSeat(tableID, seatID)
}

// QuickMatch 快速匹配
// code.NoMatchingRoom
func (mgr *RoomMgr) QuickMatch(player *Player) error {
	for _, room := range mgr.rooms {
		if room.QuickMatch(player) == nil {
			return nil
		}
	}

	return errors.NewError(code.NoMatchingRoom)
}
