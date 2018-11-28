package application

import (
	"encoding/csv"
	"github.com/jhgv/gocodes/middleware/experiments/rpc/core"
	"io/ioutil"
	"os"
	"strings"
)

type FileConverter struct {

}

// UpperText :
func (u *FileConverter) ConvertFile(args *core.Args, reply *[]byte) error {

	convertedFile, _ := os.Create("file.csv")

	csvWriter := csv.NewWriter(convertedFile)
	defer csvWriter.Flush()
	defer convertedFile.Close()
	strFile := strings.Split(string(args.File), "\n")

	for _, line := range strFile {
		strLine := string(line)
		convertedLine := strings.Replace(strLine, "\t", "," , -1)
		realLine := strings.Split(convertedLine, ",")
		csvWriter.Write(realLine)
	}
	b, _ := ioutil.ReadFile("file.csv")
	*reply = b
	return nil
}