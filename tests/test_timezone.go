package main

import (
	"fmt"
	"time"
)

func main() {
	utc := time.Now().UTC()
	fmt.Println(utc)
	local := utc
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println(local.Format("2006-01-02 15:04:05"))
	// fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	// local = utc
	// location, err = time.LoadLocation("America/Los_Angeles")
	// if err == nil {
	//     local = local.In(location)
	// }
	// fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
}
