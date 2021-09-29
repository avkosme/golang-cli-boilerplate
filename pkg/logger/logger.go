package logger

import (
	"log"
	"os"

	"gitlab.com/avkosme/transactions-statistic/internal/config"
)

var (
	outfile, _ = os.OpenFile(*config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile    = log.New(outfile, "", 0)
)

func ForError(er error) {
	if er != nil {
		LogFile.Fatalln(er)
	}
}
