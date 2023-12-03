package entity

type Options struct {
	ID            int    `json:"id"`            // 房间ID
	Name          string `json:"name"`          // 房间名称
	MinEntryLimit int    `json:"minEntryLimit"` // 最低进入限制
	MaxEntryLimit int    `json:"maxEntryLimit"` // 最高进入限制
	TotalTables   int    `json:"totalTables"`   // 房间总的牌桌数，为0时则为动态牌桌
	TotalSeats    int    `json:"totalSeats"`    // 牌桌总的座位数
	AutoReady     bool   `json:"autoReady"`     // 上桌后自动准备
	AutoStart     bool   `json:"autoStart"`     // 准备后自动开始
}
