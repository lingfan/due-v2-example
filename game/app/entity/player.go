package entity

import (
	"due-v2-example/shared/code"
	"due-v2-example/shared/model/user"
	"github.com/dobyte/due/v2/errors"
	"sync"
)

type Player struct {
	user     *user.User
	initCoin int

	rw       sync.RWMutex
	seat     *Seat
	incrCoin int
}

func newPlayer(user *user.User) *Player {
	return &Player{
		user:     user,
		initCoin: user.Coin,
	}
}

// Reset 重置用户
func (p *Player) Reset() {
	p.rw.Lock()
	defer p.rw.Unlock()

	p.syncToDB()
	p.seat = nil
}

// UID 获取用户ID
func (p *Player) UID() int64 {
	return p.user.UID
}

// User 获取用户
func (p *Player) User() *user.User {
	u := p.user
	u.Coin = p.Coin()
	return u
}

// Coin 获取金币
func (p *Player) Coin() int {
	p.rw.RLock()
	defer p.rw.RUnlock()

	return p.initCoin + p.incrCoin
}

// Seat 获取玩家所属座位
func (p *Player) Seat() *Seat {
	p.rw.RLock()
	defer p.rw.RUnlock()

	return p.seat
}

// Table 获取玩家所属牌桌
func (p *Player) Table() *Table {
	p.rw.RLock()
	defer p.rw.RUnlock()

	if p.seat == nil {
		return nil
	}

	return p.seat.Table()
}

// Room 获取玩家所属房间
func (p *Player) Room() *Room {
	p.rw.RLock()
	defer p.rw.RUnlock()

	if p.seat == nil {
		return nil
	}

	return p.seat.Room()
}

// AddSeat 为玩家添加一个座位
// code.PlayerAlreadySeated
func (p *Player) AddSeat(seat *Seat) error {
	p.rw.Lock()
	defer p.rw.Unlock()

	if p.seat != nil {
		return errors.NewError(code.PlayerAlreadySeated)
	}

	p.seat = seat

	return nil
}

// RemSeat 移除玩家座位
// code.PlayerNotInSeat
func (p *Player) RemSeat() error {
	p.rw.Lock()
	defer p.rw.Unlock()

	if p.seat == nil {
		return errors.NewError(code.PlayerNotInSeat)
	}

	p.seat = nil

	return nil
}

// IncrCoin 增加金币
func (p *Player) IncrCoin(incr int) {
	if incr == 0 {
		return
	}

	p.rw.Lock()
	defer p.rw.Unlock()

	p.incrCoin += incr
}

// DecrCoin 扣减金币
func (p *Player) DecrCoin(decr int) {
	p.IncrCoin(0 - decr)
}

// SyncToDB 同步到数据库
func (p *Player) SyncToDB() {
	p.rw.Lock()
	defer p.rw.Unlock()

	p.syncToDB()
}

func (p *Player) syncToDB() {
	if p.incrCoin == 0 {
		return
	}

	// TODO: 同步到数据库
	go func(coin int) {

	}(p.incrCoin)

	p.initCoin += p.incrCoin
	p.incrCoin = 0
	p.user.Coin = p.initCoin
}
