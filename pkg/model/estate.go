package model

// Building 房产信息
type Building struct {
	Model

	CommunityID          uint64 // 所属小区
	Kind                 uint8  // 房产类别（出售、出租）
	SpecificLocation     string // 具体位置
	Floor                uint32 // 楼层
	RoomNO               uint8  // 房号
	Room                 string // 房间信息
	Area                 uint32 // 面积
	Layout               string // 布局
	Towards              string // 朝向
	Elevator             uint32 // 电梯数量
	ParkingSpace         string // 车位
	Water                string // 用水
	Electricity          string // 用电
	Gas                  string // 用气
	heating              string // 采暖
	Property             string // 物业
	SupportingFacilities string // 配套设施
}

const (
	BuildingKindSell uint8 = iota
	BuildingKindRent
)
