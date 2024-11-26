package dataSource

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"strings"
)

type DataSource interface {
	GetRow(index any) ([]interface{}, error)        // 获取一行数据
	InsertRow(file *excelize.File, index any) error // 插入一行数据
}

type DBData struct {
	*gorm.DB
	Row int
}

var firstRows = []string{
	"设备id",
	"节点机器",
	"sd_time",
	"sd_cmd_time,sd_pro_time",
	"txt2img",
	"img2img",
	"jp_time",
	"jp_qinglong_time",
	"wt_time",
	"extra-single-image",
	"sd_api",
	"key_pair",
	"jp_dk_time",
	"ssh_time",
	"sglang_time",
	"sglang_gemma_time",
	"sglang_llama_time",
	"sd_tomato_time",
	"sd_webui_forge_time",
	"sd_webui_base_time",
	"sd_tian_time",
	"sd_bingo_time",
	"fooocus_time",
	"fooocus_lan_que_time",
	"fooocus_lan_que_en_time",
	"fooocus_sha_api_time",
	"tabby_time",
	"ollama_time",
	"jp_conda_time",
	"jp_ml_time",
	"sd_cat_time",
	"sd_fire_time",
	"comfyui_time",
	"comfyui_advance_time",
	"comfyui_advance_en_time",
	"jp_dk_3_time",
	"sd_xl_time",
	"sd_chick_time",
	"ascend_time",
	"sd_wu_shan_time",
	"sd_lang_time",
	"sd_zhi_dao_time",
	"comfyui_ke_time",
	"comfyui_a_shuo_time",
	"comfyui_jia_ming_time",
	"chatchat_time",
	"chat_tts_time",
	"lora_time",
	"lora_flux_time",
	"lora_flux_gym_time",
	"fooocus_wu_time",
	"svd_back_time",
	"sd_ji_time",
	"sd_shang_jin_time",
	"sd_xiao_chun_time",
	"comfyui_wu_time,comfyui_pro_time",
	"comfyui_advance_aisay_time",
	"comfyui_sxkk_time",
	"comfyui_liu_time",
	"sd_qkk_time",
	"sd_light_en_time",
	"comfyui_li_time",
	"comfyui_nenly_time",
	"comfyui_ou_time",
	"comfyui_lu_time",
	"comfyui_jiang_time",
	"comfyui_star_time",
	"waiting_time",
	"waiting_aleo_time",
	"opencl_core_time",
	"io_paint_time",
	"cogvideo_time",
}

// 插入初始行
func InsertFirstRow(file *excelize.File) error {
	err := file.SetSheetRow("Sheet1", "A1", &firstRows)
	if err != nil {
		println(err)
		return err
	}

	return nil
}

// 插入初始列
func InsertFirstCol(file *excelize.File, firstCol interface{}) error {
	data, ok := firstCol.([]int64)
	if !ok {
		return errors.New("firstCol is not []int64")
	}
	err := file.SetSheetCol("Sheet1", "A2", &data)
	if err != nil {
		fmt.Printf("插入初始行失败: %v", err)
		return err
	}

	return nil
}

// select * from devices as d left join device_gpu_missions as dgm on d.id=dgm.device_id
// where d.id = 123456 and d.version<>'无版本号' and d.status<>'exit' and dgm.gpu_status<>'exit' group by d.id;
func (data DBData) GetRow(index any, row interface{}) error {
	deviceID, ok := index.(int64)
	if !ok {
		return errors.New("index is not int64")
	}

	//err := data.Model(&DeviceGpuMission{}).Where(&where).Find(row).Error
	row = allData[deviceID].DeviceGpuMissions

	return nil
}

// 插入一行数据
func (data DBData) InsertRow(file *excelize.File, index any) error {
	var device Device
	var missions []DeviceGpuMission
	err := data.GetRow(index, &missions)
	if err != nil {
		return err
	}
	device = data.getDevice(index.(int64))
	deviceName := device.Name

	err = file.SetCellValue("Sheet1", getCell(1, data.Row), deviceName)

	for i := 0; i < len(missions); i++ {
		var mission string
		mission = missions[i].AbleMissionKind
		for j := 0; j < len(firstRows); j++ {
			//有对应任务就勾选
			if strings.Contains(mission, firstRows[j]) {
				cell := getCell(j, data.Row)
				err := file.SetCellValue("Sheet1", cell, "√")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (data DBData) GetDeviceIDs() ([]int64, error) {
	var deviceIDs []int64
	for i := 0; i < len(allDeviceData); i++ {
		deviceIDs = append(deviceIDs, allDeviceData[i].ID)
	}
	return deviceIDs, nil
}

func getCell(col, row int) string {
	var cell string
	var colString string
	for col >= 0 {
		one := col % 26
		col = col/26 - 1
		colString = fmt.Sprintf("%s%s", string('A'+one), colString)
	}
	cell = fmt.Sprintf("%s%d", colString, row)
	return cell

}

func (data DBData) getDevice(id int64) Device {
	var device Device
	device = allData[id]
	return device
}

func (data DBData) GetAllDevices() ([]Device, error) {
	var devices []Device
	err := data.Model(&Device{}).Find(&devices).Error
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (data DBData) GetAllMissions() ([]DeviceGpuMission, error) {
	var missions []DeviceGpuMission
	err := data.Model(&DeviceGpuMission{}).Find(&missions).Error
	if err != nil {
		return nil, err
	}
	return missions, nil
}

var allData map[int64]Device
var allMissionData []DeviceGpuMission
var allDeviceData []Device

func (data DBData) MapDeviceAndMission() (map[int64]Device, error) {
	devices, err := data.GetAllDevices()
	if err != nil {
		return nil, err
	}
	allDeviceData = devices

	missions, err := data.GetAllMissions()
	if err != nil {
		return nil, err
	}
	allMissionData = missions

	var result = make(map[int64]Device)
	for i := 0; i < len(devices); i++ {
		devices[i].DeviceGpuMissions = make([]DeviceGpuMission, 0)
		result[devices[i].ID] = devices[i]
	}
	for i := 0; i < len(missions); i++ {
		if _, exists := result[missions[i].DeviceID]; exists {
			gpuMissions := result[missions[i].DeviceID].DeviceGpuMissions
			gpuMissions = append(gpuMissions, missions[i])
		}
	}

	allData = result
	return result, nil
}
