package route

const (
	Register     = 1 // 用户注册
	Login        = 2 // 用户登录
	Unauthorized = 3 // 用户未授权
)

const (
	NewMail       = 101 // 新邮件
	FetchMailList = 102 // 拉取邮件列表
	ReadMail      = 103 // 读取邮件
	ReadAllMail   = 104 // 一键读取所有邮件
	DeleteMail    = 105 // 删除邮件
	DeleteAllMail = 106 // 删除所有邮件
)

const (
	FetchRooms      = 1000 // 拉取房间
	QuickStart      = 1001 // 快速开始
	SitDown         = 1002 // 坐下
	StandUp         = 1003 // 站起
	Ready           = 1004 // 开始准备
	Unready         = 1005 // 取消准备
	SeatStateChange = 1006 // 状态变更
	TakeSeat        = 1007 // 入座通知
	GameInfoNotify  = 1008 // 游戏信息通知
)
