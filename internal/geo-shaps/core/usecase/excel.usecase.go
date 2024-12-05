package usecase

import (
	"digikalajet/internal/geo-shaps/core/domain"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"sync"
)

type ExcelUsecase struct{}

func NewExcelUsecase() *ExcelUsecase {
	return &ExcelUsecase{}
}

func (u *ExcelUsecase) GenerateLocalMonitoringExcel(monitoringData []domain.LocalMonitoring, wg *sync.WaitGroup) {
	rowCounter := 1
	sheet := "Sheet1"
	file := u.createExcelFile(sheet)
	file, rowCounter = u.addLocalMonitoringHeader(sheet, file, rowCounter)

	for _, data := range monitoringData {
		file, rowCounter = u.addLocalMonitoringRow(sheet, file, rowCounter, data)
	}

	if err := u.saveExcelFile(file, "Local Monitoring.xlsx"); err != nil {
		log.Println(err)
	}

	wg.Done()
}

func (u *ExcelUsecase) createExcelFile(sheet string) *excelize.File {
	file := excelize.NewFile()
	file.NewSheet(sheet)
	return file
}

func (u *ExcelUsecase) addLocalMonitoringHeader(sheet string, file *excelize.File, rowCounter int) (*excelize.File, int) {
	headers := []string{"Timestamp", "CPU Usage (%)", "Memory Usage (%)", "Disk Usage (%)", "Network Sent (Bytes)", "Network Received (Bytes)"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string('A'+i), rowCounter)
		file.SetCellValue(sheet, cell, header)
	}
	return file, rowCounter + 1
}

func (u *ExcelUsecase) addLocalMonitoringRow(sheet string, file *excelize.File, rowCounter int, data domain.LocalMonitoring) (*excelize.File, int) {
	file.SetCellValue(sheet, fmt.Sprintf("A%d", rowCounter), data.Timestamp)
	file.SetCellValue(sheet, fmt.Sprintf("B%d", rowCounter), data.CpuUsage)
	file.SetCellValue(sheet, fmt.Sprintf("C%d", rowCounter), data.MemoryUsage)
	file.SetCellValue(sheet, fmt.Sprintf("D%d", rowCounter), data.DiskUsage)
	file.SetCellValue(sheet, fmt.Sprintf("E%d", rowCounter), data.NetSent)
	file.SetCellValue(sheet, fmt.Sprintf("F%d", rowCounter), data.NetRecv)

	return file, rowCounter + 1
}

func (u *ExcelUsecase) saveExcelFile(file *excelize.File, fileName string) error {
	return file.SaveAs(fileName)
}

func (u *ExcelUsecase) GenerateElasticLoadTestExcel(testData []domain.ElasticLoadTest, wg *sync.WaitGroup) {
	rowCounter := 1
	sheet := "Sheet1"
	file := u.createExcelFile(sheet)
	file, rowCounter = u.addElasticLoadTestHeader(sheet, file, rowCounter)

	for _, data := range testData {
		file = u.addElasticLoadTestRow(sheet, file, data)
	}

	if err := u.saveExcelFile(file, "Elastic Load Test.xlsx"); err != nil {
		log.Println(err)
	}

	wg.Done()
}

func (u *ExcelUsecase) addElasticLoadTestHeader(sheet string, file *excelize.File, rowCounter int) (*excelize.File, int) {
	headers := []string{"Request Number", "Start", "End", "Duration (ms)", "Duration (ns)", "Status Sent"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string('A'+i), rowCounter)
		file.SetCellValue(sheet, cell, header)
	}
	return file, rowCounter + 1
}

func (u *ExcelUsecase) addElasticLoadTestRow(sheet string, file *excelize.File, data domain.ElasticLoadTest) *excelize.File {
	row := data.RequestNumber + 2
	file.SetCellValue(sheet, fmt.Sprintf("A%d", row), data.RequestNumber)
	file.SetCellValue(sheet, fmt.Sprintf("B%d", row), data.Start.Format("2006-01-02 15:04:05.000000"))
	file.SetCellValue(sheet, fmt.Sprintf("C%d", row), data.End.Format("2006-01-02 15:04:05.000000"))
	file.SetCellValue(sheet, fmt.Sprintf("D%d", row), fmt.Sprintf("%d", data.End.Sub(data.Start).Milliseconds()))
	file.SetCellValue(sheet, fmt.Sprintf("E%d", row), fmt.Sprintf("%d", data.End.Sub(data.Start).Nanoseconds()))
	file.SetCellValue(sheet, fmt.Sprintf("F%d", row), data.Status)

	return file
}

func (u *ExcelUsecase) GenerateElasticMonitoringExcel(monitoringData []domain.ElasticMonitoring, wg *sync.WaitGroup) {
	rowCounter := 1
	sheet := "Sheet1"
	file := u.createExcelFile(sheet)
	file, rowCounter = u.addElasticMonitoringHeader(sheet, file, rowCounter)

	for _, data := range monitoringData {
		file, rowCounter = u.addElasticMonitoringRow(sheet, file, rowCounter, data)
	}

	if err := u.saveExcelFile(file, "Elastic Monitoring.xlsx"); err != nil {
		log.Println(err)
	}

	wg.Done()
}

func (u *ExcelUsecase) addElasticMonitoringHeader(sheet string, file *excelize.File, rowCounter int) (*excelize.File, int) {
	headers := []string{
		"Timestamp", "Node", "Memory Total (Bytes)", "Memory Free (Bytes)", "Memory Used (Bytes)",
		"Memory Free Percent (%)", "Memory Used Percent (%)", "Disk Least Used Disk (%)", "Disk Least Total (Bytes)",
		"Disk Least Available (Bytes)", "Disk Most Used Disk (%)", "Disk Most Total (Bytes)", "Disk Most Available (Bytes)",
		"Thread Pool Fs Total (Bytes)", "Thread Pool Fs Free (Bytes)", "Thread Pool Fs Available (Bytes)", "CPU (%)",
		"JVM Heap Used (Bytes)", "JVM Heap Used (%)", "JVM Heap Committed (Bytes)", "JVM Heap Max (Bytes)",
		"JVM Non Heap Used (Bytes)", "JVM Non Heap Committed (Bytes)",
	}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string('A'+i), rowCounter)
		file.SetCellValue(sheet, cell, header)
	}
	return file, rowCounter + 1
}

func (u *ExcelUsecase) addElasticMonitoringRow(sheet string, file *excelize.File, rowCounter int, data domain.ElasticMonitoring) (*excelize.File, int) {
	file.SetCellValue(sheet, fmt.Sprintf("A%d", rowCounter), data.Timestamp)
	file.SetCellValue(sheet, fmt.Sprintf("B%d", rowCounter), data.Node)
	file.SetCellValue(sheet, fmt.Sprintf("C%d", rowCounter), data.Memory.TotalInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("D%d", rowCounter), data.Memory.FreeInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("E%d", rowCounter), data.Memory.UsedInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("F%d", rowCounter), data.Memory.FreePercent)
	file.SetCellValue(sheet, fmt.Sprintf("G%d", rowCounter), data.Memory.UsedPercent)
	file.SetCellValue(sheet, fmt.Sprintf("H%d", rowCounter), data.Disk.LeastUsedDiskPercent)
	file.SetCellValue(sheet, fmt.Sprintf("I%d", rowCounter), data.Disk.LeastTotalInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("J%d", rowCounter), data.Disk.LeastAvailableInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("K%d", rowCounter), data.Disk.MostUsedDiskPercent)
	file.SetCellValue(sheet, fmt.Sprintf("L%d", rowCounter), data.Disk.MostTotalInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("M%d", rowCounter), data.Disk.MostAvailableInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("N%d", rowCounter), data.ThreadPoolFs.TotalInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("O%d", rowCounter), data.ThreadPoolFs.FreeInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("P%d", rowCounter), data.ThreadPoolFs.AvailableInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("Q%d", rowCounter), data.CPU.Percent)
	file.SetCellValue(sheet, fmt.Sprintf("R%d", rowCounter), data.JVM.HeapUsedInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("S%d", rowCounter), data.JVM.HeapUsedPercent)
	file.SetCellValue(sheet, fmt.Sprintf("T%d", rowCounter), data.JVM.HeapCommittedInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("U%d", rowCounter), data.JVM.HeapMaxInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("V%d", rowCounter), data.JVM.NonHeapUsedInBytes)
	file.SetCellValue(sheet, fmt.Sprintf("W%d", rowCounter), data.JVM.NonHeapCommittedInBytes)

	return file, rowCounter + 1
}
