package entity

import (
	"due-v2-example/shared/code"
	"due-v2-example/shared/service"
	"sync"

	"github.com/dobyte/due/v2/errors"
)

var (
	playerMgrOnce     sync.Once
	playerMgrInstance *PlayerMgr
)

type PlayerMgr struct {
	userSvc *service.User

	rw      sync.RWMutex
	players map[int64]*Player
}

func NewPlayerMgr() *PlayerMgr {
	playerMgrOnce.Do(func() {
		playerMgrInstance = &PlayerMgr{
			userSvc: service.NewUser(nil),
			players: make(map[int64]*Player),
		}
	})
	return playerMgrInstance
}

// LoadPlayer 加载玩家
// code.NotFoundUser
// code.InternalServerError
func (mgr *PlayerMgr) LoadPlayer(uid int64) (*Player, error) {
	mgr.rw.RLock()
	player, ok := mgr.players[uid]
	mgr.rw.RUnlock()
	if ok {
		return player, nil
	}

	user, err := mgr.userSvc.GetUser(uid)
	if err != nil {
		return nil, err
	}

	mgr.rw.Lock()
	defer mgr.rw.Unlock()

	player = newPlayer(user)
	mgr.players[uid] = player

	return player, nil
}

// UnloadPlayer 卸载玩家
func (mgr *PlayerMgr) UnloadPlayer(uid int64) {
	mgr.rw.Lock()
	defer mgr.rw.Unlock()

	player, ok := mgr.players[uid]
	if !ok {
		return
	}

	player.Reset()

	delete(mgr.players, uid)
}

// GetPlayer 获取玩家
// code.NotFoundPlayer
func (mgr *PlayerMgr) GetPlayer(uid int64) (*Player, error) {
	mgr.rw.RLock()
	defer mgr.rw.RUnlock()

	player, ok := mgr.players[uid]
	if !ok {
		return nil, errors.NewError(code.NotFoundPlayer)
	}

	return player, nil
}
