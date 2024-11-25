package dataSource

import "time"

type Device struct {
	ID                  int64              `gorm:"column:id;primaryKey" json:"id"`
	CreatedBy           int64              `gorm:"column:created_by" json:"created_by"`
	UpdatedBy           int64              `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt           time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time          `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           time.Time          `gorm:"column:deleted_at" json:"deleted_at"`
	SerialNumber        string             `gorm:"column:serial_number" json:"serial_number"`
	State               string             `gorm:"column:state" json:"state"`
	SumCep              int64              `gorm:"column:sum_cep" json:"sum_cep"`
	Linking             bool               `gorm:"column:linking" json:"linking"`
	BindingStatus       string             `gorm:"column:binding_status" json:"binding_status"`
	Status              string             `gorm:"column:status" json:"status"`
	UserId              int64              `gorm:"column:user_id" json:"user_id"`
	Name                string             `gorm:"column:name" json:"name"`
	Type                string             `gorm:"column:type" json:"type"`
	CoresNumber         int64              `gorm:"column:cores_number" json:"cores_number"`
	Cpu                 string             `gorm:"column:cpu" json:"cpu"`
	Memory              int64              `gorm:"column:memory" json:"memory"`
	Disk                string             `gorm:"column:disk" json:"disk"`
	Cpus                []byte             `gorm:"column:cpus" json:"cpus"`
	ManageName          string             `gorm:"column:manage_name" json:"manage_name"`
	Delay               float64            `gorm:"column:delay" json:"delay"`
	Temperature         float64            `gorm:"column:temperature" json:"temperature"`
	Stability           string             `gorm:"column:stability" json:"stability"`
	GpuTemperature      float64            `gorm:"column:gpu_temperature" json:"gpu_temperature"`
	CpuTemperature      float64            `gorm:"column:cpu_temperature" json:"cpu_temperature"`
	Version             string             `gorm:"column:version" json:"version"`
	Fault               string             `gorm:"column:fault" json:"fault"`
	Rank                string             `gorm:"column:rank" json:"rank"`
	FreeGpuNum          int64              `gorm:"column:free_gpu_num" json:"free_gpu_num"`
	StabilityAt         time.Time          `gorm:"column:stability_at" json:"stability_at"`
	RankAt              time.Time          `gorm:"column:rank_at" json:"rank_at"`
	HighTemperatureAt   time.Time          `gorm:"column:high_temperature_at" json:"high_temperature_at"`
	GiftMissionConfigId int64              `gorm:"column:gift_mission_config_id" json:"gift_mission_config_id"`
	HostingType         string             `gorm:"column:hosting_type" json:"hosting_type"`
	MissionTag          string             `gorm:"column:mission_tag" json:"mission_tag"`
	LastAbnormalAt      time.Time          `gorm:"column:last_abnormal_at" json:"last_abnormal_at"`
	StabilityDrop       string             `gorm:"column:stability_drop" json:"stability_drop"`
	DoableAleoCount     int64              `gorm:"column:doable_aleo_count" json:"doable_aleo_count"`
	DeviceGpuMissions   []DeviceGpuMission `gorm:"foreignKey:DeviceID;references:ID"`
}

func (Device) TableName() string {
	return "devices"
}

type DeviceGpuMission struct {
	ID int64 `gorm:"column:id,string"`
	// 创建者 ID
	CreatedBy int64 `gorm:"column:created_by"`
	// 更新者 ID
	UpdatedBy int64 `gorm:"column:updated_by"`
	// 创建时刻，带时区
	CreatedAt time.Time `gorm:"column:created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `gorm:"column:deleted_at"`
	// 外键设备 id
	DeviceID int64 `gorm:"column:device_id"`
	// 外键 gpu id
	GpuID int64 `gorm:"column:gpu_id"`
	// 可以接的任务类型
	AbleMissionKind string `gorm:"column:able_mission_kind"`
	// 显卡占用设备的插槽
	DeviceSlot int8 `gorm:"column:device_slot"`
	// 最大同时在线任务
	MaxOnlineMission int8 `gorm:"column:max_online_mission"`
	// gpu 当前状态
	GpuStatus string `gorm:"column:gpu_status"`
	// 正在做的任务 id
	MissionID []uint8 `gorm:"column:mission_id"`
}

func (DeviceGpuMission) TableName() string {
	return "device_gpu_missions"
}
