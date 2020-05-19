package model

// Community 小区信息
type Community struct {
	Model

	Name             string    // 小区名称
	Location         *Location // 位置信息
	SpecificLocation string    // 具体位置
	ConstructionYear uint32    // 建筑年代
	BuildingType     string    // 建筑类型
	IsElevator       bool      // 是否配置电梯
	GreeningDegree   uint8     // 绿化度
	Area             uint32    // 面积
	BuildNumber      uint32    // 建筑数量
	AverageFloor     uint8     // 平均楼层
	Traffic          string    // 交通
	Around           string    // 周边配套
	Feature          string    // 特点
}





// TODO 交通情况和周边配套需要结合实际情况来实现，例如对接第三方地图的 SDK
// Traffic 交通
type Traffic struct{}

// Around 周边
type Around struct{}
