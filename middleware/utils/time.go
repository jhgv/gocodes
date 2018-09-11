package utils

import (
	"time"
	"os"
	"fmt"
)

func WriteTime(duration time.Duration, file *os.File) {
	time := fmt.Sprintf("%f",duration.Seconds()) + "\n"

	if _, err := file.Write([]byte(time)); err != nil {
		panic(err)
	}
}