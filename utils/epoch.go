package utils

import (
	"strconv"
	"time"
)

func ConvertEpochToHuman(convertEpoch string) string {
	epoch, err := strconv.ParseInt(convertEpoch, 10, 64)
	if err != nil {
		return "Erro ao converter epoch"
	}
	date := time.Unix(epoch, 0)
	return date.Format("02-Jan-2006 15:04:05")
}
