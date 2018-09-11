package utils

import (
	"github.com/tealeg/xlsx"
	"fmt"
)


type XlsxBuilder struct {
	fileName string
	average float64
	totalTime float64
	file *xlsx.File
	averageCell *xlsx.Cell
	totalTimeCell *xlsx.Cell
	sheet *xlsx.Sheet
}


func (b *XlsxBuilder) CreateHeader() {
	b.file = xlsx.NewFile()
	var err error
	b.sheet, err = b.file.AddSheet("Analytics")
	if err != nil {
		fmt.Printf(err.Error())
	}
	firstRow := b.sheet.AddRow()
	headerColumns := [4]string{"Average", "Total time"}
	for _, v := range headerColumns {
		cell := firstRow.AddCell()
		cell.Value = v
	}

	secondRow := b.sheet.AddRow()
	b.averageCell = secondRow.AddCell()
	b.totalTimeCell = secondRow.AddCell()

}

func (b *XlsxBuilder) SetupAverageFormula(formula string) {
	b.averageCell.SetFormula(formula)
}

func (b *XlsxBuilder) AddRowData(value float64) {
	row := b.sheet.AddRow()
	cell := row.AddCell()
	cell.SetFloatWithFormat(value, "general")
}

func (b *XlsxBuilder) SetFileName(fileName string) {
	b.fileName = fileName
}

func (b *XlsxBuilder) SetAverage(average float64) {
	b.average = average
}

func (b *XlsxBuilder) SetTotalTime(average float64) {
	b.totalTime = average
}

func (b *XlsxBuilder) GetRowNum() int {
	return b.sheet.MaxRow
}

func (b *XlsxBuilder) GetFile() *xlsx.File {
	return b.file
}

func (b *XlsxBuilder) GenerateFile() {
	var err error
	b.totalTimeCell.SetFloatWithFormat(b.totalTime, "general")
	err = b.file.Save(b.fileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
