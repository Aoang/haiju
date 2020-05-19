package model

// Property 物业信息
type Property struct {
	Model

	Name  string // 名称（或者自理）
	Phone uint32 // 电话
	cost  uint64 // 费用 最低位为分
}
