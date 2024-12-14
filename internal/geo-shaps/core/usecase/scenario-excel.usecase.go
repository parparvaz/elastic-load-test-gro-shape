package usecase

import (
	"digikalajet/internal/geo-shaps/core/domain"
	"log"
)

type ScenarioExcelUsecase struct {
	excelUsecase *ExcelUsecase
}

func NewScenarioExcelUsecase() *ScenarioExcelUsecase {
	return &ScenarioExcelUsecase{
		excelUsecase: NewExcelUsecase(),
	}
}

func (u *ScenarioExcelUsecase) GenerateLocalMonitoringExcel(sheet string, monitoringData []domain.LocalMonitoring) {
	rowCounter := 1
	file, err := u.excelUsecase.openExcelFile("Local Monitoring.xlsx", sheet)
	if err != nil {
		log.Println(err)
		return
	}

	file, rowCounter = u.excelUsecase.addLocalMonitoringHeader(sheet, file, rowCounter)
	for _, data := range monitoringData {
		file, rowCounter = u.excelUsecase.addLocalMonitoringRow(sheet, file, rowCounter, data)
	}

	if err = u.excelUsecase.saveExcelFile(file, "Local Monitoring.xlsx"); err != nil {
		log.Println(err)
	}
}

func (u *ScenarioExcelUsecase) GenerateElasticMonitoringExcel(sheet string, monitoringData []domain.ElasticMonitoring) {
	rowCounter := 1
	file, err := u.excelUsecase.openExcelFile("Elastic Monitoring.xlsx", sheet)
	if err != nil {
		log.Println(err)
		return
	}

	file, rowCounter = u.excelUsecase.addElasticMonitoringHeader(sheet, file, rowCounter)

	for _, data := range monitoringData {
		file, rowCounter = u.excelUsecase.addElasticMonitoringRow(sheet, file, rowCounter, data)
	}

	if err := u.excelUsecase.saveExcelFile(file, "Elastic Monitoring.xlsx"); err != nil {
		log.Println(err)
	}

}

func (u *ScenarioExcelUsecase) GenerateElasticLoadTestExcel(sheet string, testData []domain.ElasticLoadTest) {
	rowCounter := 1
	file, err := u.excelUsecase.openExcelFile("Elastic Load Test.xlsx", sheet)
	if err != nil {
		log.Println(err)
		return
	}

	file, rowCounter = u.excelUsecase.addElasticLoadTestHeader(sheet, file, rowCounter)

	for _, data := range testData {
		file = u.excelUsecase.addElasticLoadTestRow(sheet, file, data)
	}

	if err := u.excelUsecase.saveExcelFile(file, "Elastic Load Test.xlsx"); err != nil {
		log.Println(err)
	}
}
