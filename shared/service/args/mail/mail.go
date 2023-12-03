package mail

type SendArgs struct {
	Sender
	Receiver int64 // 邮件接收者

}

// Mail 邮件
type Mail struct {
	Title       string       // 邮件标题
	Content     string       // 邮件内容
	Attachments []Attachment // 邮件附件
}

// Sender 邮件发送者
type Sender struct {
	ID   int64  // 发送者ID，官方发送者ID为0，系统邮件为负数，用户发送的为正数
	Name string // 发送者名称
	Icon string // 发送者图标，仅在发送者为用户时存在
}

// Attachment 邮件附件
type Attachment struct {
	PropID  int // 道具ID
	PropNum int // 道具数量
}
