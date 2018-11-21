package objects

import (
	"encoding/csv"
	"os"
	"strings"
)

type FileConverter struct {
}

func (fc *FileConverter) ConvertFile(file []byte) os.File {
	convertedFile, _ := os.Create("file.csv")

	csvWriter := csv.NewWriter(convertedFile)
	defer csvWriter.Flush()
	strFile := strings.Split(string(file), "\n")

	for _, line := range strFile {
		strLine := string(line)
		convertedLine := strings.Replace(strLine, "\t", "," , -1)
		realLine := strings.Split(convertedLine, ",")
		csvWriter.Write(realLine)
	}
	return *convertedFile
}
