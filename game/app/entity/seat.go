package entity

import (
	"context"
	"due-v2-example/shared/code"
	"sync"
	"sync/atomic"

	"github.com/dobyte/due/v2/errors"
)

const (
	online  = 1 // 在线
	offline = 0 // 离线
)

type Seat struct {
	ctx    context.Context
	id     int
	table  *Table
	online int32

	rw     sync.RWMutex
	player *Player // 玩家
	ready  bool    // 准备中
}

func newSeat(id int, table *Table) *Seat {
	return &Seat{
		id:    id,
		table: table,
	}
}

// Reset 重置座位
func (s *Seat) Reset() {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.player != nil {
		s.player.Reset()
	}

	s.player = nil
	s.ready = false
	s.Offline()
}

// ID 获取座位ID
func (s *Seat) ID() int {
	return s.id
}

// Table 获取座位所属牌桌
func (s *Seat) Table() *Table {
	return s.table
}

// Room 获取座位所属房间
func (s *Seat) Room() *Room {
	return s.table.Room()
}

// Player 获取座位上的玩家
func (s *Seat) Player() *Player {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.player
}

// AddPlayer 为座位添加玩家
// code.SeatAlreadyTaken
// code.PlayerAlreadySeated
func (s *Seat) AddPlayer(player *Player) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.player != nil {
		return errors.NewError(code.SeatAlreadyTaken)
	}

	err := player.AddSeat(s)
	if err != nil {
		return err
	}

	s.player = player
	s.Online()

	if s.Table().Room().IsAutoReady() {
		s.ready = true
	}

	return nil
}

// RemPlayer 移除玩家座位
// code.SeatIsEmpty
// code.PlayerNotInSeat
// code.PlayerIsReady
func (s *Seat) RemPlayer() error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.player == nil {
		return errors.NewError(code.SeatIsEmpty)
	}

	if s.ready {
		return errors.NewError(code.PlayerIsReady)
	}

	err := s.player.RemSeat()
	if err != nil {
		return err
	}

	s.player = nil
	s.Offline()

	return nil
}

// HasPlayer 座位上是否有玩家
func (s *Seat) HasPlayer() bool {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.player != nil
}

// IsOffline 检测座位上的玩家是否离线
func (s *Seat) IsOffline() bool {
	return atomic.LoadInt32(&s.online) == offline
}

// IsOnline 检测座位上的玩家是否在线
func (s *Seat) IsOnline() bool {
	return atomic.LoadInt32(&s.online) == online
}

// Offline 标记座位上的玩家离线
func (s *Seat) Offline() {
	atomic.StoreInt32(&s.online, offline)
}

// Online 标记座位上的玩家上线
func (s *Seat) Online() {
	atomic.StoreInt32(&s.online, online)
}

// Ready 准备
func (s *Seat) Ready() {
	s.rw.Lock()
	s.ready = true
	s.rw.Unlock()
}

// Unready 取消准备
func (s *Seat) Unready() {
	s.rw.Lock()
	s.ready = false
	s.rw.Unlock()
}

// IsReady 是否已准备
func (s *Seat) IsReady() bool {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.ready
}

// IsUnready 是否未准备好
func (s *Seat) IsUnready() bool {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return !s.ready
}
