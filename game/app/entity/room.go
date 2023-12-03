package entity

import (
	"due-v2-example/shared/code"
	"github.com/dobyte/due/v2/errors"
	"sync"
)

type Room struct {
	opts  *Options
	fixed bool

	rw         sync.RWMutex
	tables     map[int]*Table
	maxTableID int
}

func newRoom(opts *Options) *Room {
	r := &Room{
		opts:       opts,
		fixed:      opts.TotalTables > 0,
		tables:     make(map[int]*Table, opts.TotalTables),
		maxTableID: opts.TotalTables - 1,
	}

	for i := 1; i <= opts.TotalTables; i++ {
		r.tables[i] = newTable(i, r)
	}

	return r
}

// ID 获取房间ID
func (r *Room) ID() int {
	return r.opts.ID
}

// Options 获取房间配置
func (r *Room) Options() *Options {
	return r.opts
}

// GetTable 获取桌子
// code.NotFoundTable
func (r *Room) GetTable(tableID int) (*Table, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()

	table, ok := r.tables[tableID]
	if !ok {
		return nil, errors.NewError(code.NotFoundTable)
	}

	return table, nil
}

// GetSeat 获取座位
// code.NotFoundTable
// code.IllegalParams
func (r *Room) GetSeat(tableID, seatID int) (*Seat, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()

	table, ok := r.tables[tableID]
	if !ok {
		return nil, errors.NewError(code.NotFoundTable)
	}

	return table.GetSeat(seatID)
}

// CreatTable 创建桌子
// code.RoomIsFull
func (r *Room) CreatTable() (*Table, error) {
	if r.fixed {
		return nil, errors.NewError(code.RoomIsFull)
	}

	r.rw.Lock()
	defer r.rw.Unlock()

	return r.createTable(), nil
}

// 创建牌桌
func (r *Room) createTable() *Table {
	tid := -1
	for i := 0; i <= r.maxTableID; i++ {
		if _, ok := r.tables[i]; !ok {
			tid = i
			break
		}
	}

	if tid < 0 {
		r.maxTableID++
		tid = r.maxTableID
	}

	table := newTable(tid, r)
	r.tables[tid] = table

	return table
}

// DestroyTable 销毁牌桌
// code.NotFoundTable
func (r *Room) DestroyTable(tableID int) error {
	r.rw.Lock()
	defer r.rw.Unlock()

	table, ok := r.tables[tableID]
	if !ok {
		return errors.NewError(code.NotFoundTable)
	}

	table.Reset()

	if r.fixed {
		return nil
	}

	delete(r.tables, tableID)

	if tableID == r.maxTableID {
		r.maxTableID = 0
		for i := range r.tables {
			if i > r.maxTableID {
				r.maxTableID = i
			}
		}
	}

	return nil
}

// QuickMatch 快速匹配
// code.CoinsAreBelowEntryLimit
// code.CoinsAreOverEntryLimit
// code.NoMatchingTable
func (r *Room) QuickMatch(player *Player) error {
	coin := player.Coin()

	if r.opts.MinEntryLimit > 0 && coin < r.opts.MinEntryLimit {
		return errors.NewError(code.CoinsAreBelowEntryLimit)
	}

	if r.opts.MaxEntryLimit > 0 && coin > r.opts.MaxEntryLimit {
		return errors.NewError(code.CoinsAreOverEntryLimit)
	}

	r.rw.RLock()
	defer r.rw.RUnlock()

	for _, table := range r.tables {
		if table.QuickMatch(player) == nil {
			return nil
		}
	}

	for _, table := range r.tables {
		if table.AddPlayer(player) == nil {
			return nil
		}
	}

	if r.fixed {
		return errors.NewError(code.NoMatchingTable)
	}

	return r.createTable().AddPlayer(player)
}

// IsAutoReady 检测是否自动准备
func (r *Room) IsAutoReady() bool {
	return r.opts.AutoReady
}
