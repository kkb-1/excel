package main

import (
	"github.com/xuri/excelize/v2"
	"goExcel/dataSource"
)

func main() {
	file := excelize.NewFile()

	//处理数据
	config := dataSource.PostgreSQL{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "postgres",
		Password: "123456",
		DBName:   "test",
	}

	db, err := dataSource.DBConnect(config)

	source := dataSource.DBData{
		DB: db,
	}

	err = dataSource.InsertFirstRow(file)
	if err != nil {
		return
	}

	deviceIDs, err := source.GetDeviceIDs()

	if err != nil {
		return
	}

	err = dataSource.InsertFirstCol(file, deviceIDs)
	if err != nil {
		return
	}

	for i := 0; i < len(deviceIDs); i++ {
		source.Row = i + 2
		err := source.InsertRow(file, deviceIDs[i])
		if err != nil {
			return
		}
	}

	err = file.SaveAs("test.xlsx")
	if err != nil {
		print(err)
	}
}
