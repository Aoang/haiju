package model

// SupportingFacilities 配套设施
type SupportingFacilities struct{}

// Room 房间信息
type Room struct {
	Model

	Area                 uint32 // 面积
	Towards              string // 朝向
	Windows              string // 窗户
	Balcony              string // 阳台
	SupportingFacilities string // 配套设施
}

