package helper

import (
	"fmt"
	"project/internal/helper/parameter"
	"time"
)

func GetTimeZone() *time.Location {
	time, err := time.LoadLocation(parameter.ZONA_WAKTU)
	if err != nil {
		fmt.Println("gadaper waktu", err.Error())
		return nil
	}
	return time
}
