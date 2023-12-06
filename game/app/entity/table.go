package entity

import (
	"due-v2-example/shared/code"
	"sync"

	"github.com/dobyte/due/v2/errors"
)

type Table struct {
	id   int
	room *Room

	rw           sync.RWMutex
	seats        []*Seat
	totalPlayers int
	totalReadies int
}

func newTable(id int, room *Room) *Table {
	t := &Table{
		id:    id,
		room:  room,
		seats: make([]*Seat, room.opts.TotalSeats),
	}

	for i := 0; i < room.opts.TotalSeats; i++ {
		t.seats[i] = newSeat(i+1, t)
	}

	return t
}

// ID 获取牌桌ID
func (t *Table) ID() int {
	return t.id
}

// Room 获取牌桌所属房间
func (t *Table) Room() *Room {
	return t.room
}

// Seats 获取牌桌所有位置
func (t *Table) Seats() []*Seat {
	return t.seats
}

// Reset 重置牌桌
func (t *Table) Reset() {
	t.rw.Lock()
	defer t.rw.Unlock()

	for i := range t.seats {
		t.seats[i].Reset()
	}

	t.totalPlayers = 0
}

// GetSeat 获取座位
// code.IllegalParams
func (t *Table) GetSeat(seatID int) (*Seat, error) {
	t.rw.RLock()
	defer t.rw.RUnlock()

	sid := seatID - 1

	if sid < 0 || sid >= len(t.seats) {
		return nil, errors.NewError(code.IllegalParams)
	}

	return t.seats[sid], nil
}

// AddPlayer 把玩家加入牌桌
// code.IllegalParams
// code.TableIsFull
// code.SeatAlreadyTaken
// code.PlayerAlreadySeated
func (t *Table) AddPlayer(player *Player, seatID ...int) error {
	t.rw.Lock()
	defer t.rw.Unlock()

	if len(seatID) > 0 {
		sid := seatID[0] - 1

		if sid < 0 || sid >= len(t.seats) {
			return errors.NewError(code.IllegalParams)
		}

		err := t.seats[sid].AddPlayer(player)
		if err != nil {
			return err
		}

		t.totalPlayers++

		return nil
	}

	for _, seat := range t.seats {
		if seat.AddPlayer(player) == nil {
			t.totalPlayers++
			return nil
		}
	}

	return errors.NewError(code.TableIsFull)
}

// RemPlayer 移除玩家
// code.IllegalParams
// code.SeatIsEmpty
// code.PlayerNotInSeat
func (t *Table) RemPlayer(seatID int) error {
	t.rw.Lock()
	defer t.rw.Unlock()

	sid := seatID - 1

	if sid < 0 || sid >= len(t.seats) {
		return errors.NewError(code.IllegalParams)
	}

	err := t.seats[sid].RemPlayer()
	if err != nil {
		return err
	}

	t.totalPlayers--

	return nil
}

// HasPlayer 检测桌子上是否有玩家
func (t *Table) HasPlayer() bool {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.totalPlayers > 0
}

// QuickMatch 快速匹配
func (t *Table) QuickMatch(player *Player) error {
	t.rw.Lock()
	defer t.rw.Unlock()

	if len(t.seats) == t.totalPlayers {
		return errors.NewError(code.TableIsFull)
	}

	if t.totalPlayers == 0 {
		return errors.NewError(code.TableIsEmpty)
	}

	for _, seat := range t.seats {
		if seat.AddPlayer(player) == nil {
			t.totalPlayers++
			t.totalReadies++
			break
		}
	}

	return nil
}
