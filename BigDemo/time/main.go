package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Unix(1666576560, 0)
	unix := time.Date(t.Year(), t.Month(), t.Day()-7, 0, 0, 0, 0, t.Location()).Unix()
	fmt.Println(unix + 6*86400)

}

/*






 */
func getTimeUnix(rep string) int64 {
	t := time.Now()
	if rep == "day" {
		addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		return addTime.Unix()
	} else if rep == "month" {
		addTime := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location())
		return addTime.Unix()
	}
	return 0
}

//t := time.Unix(1660924800, 0)
//t := time.Now()
//dayTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
//month := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
//
//fmt.Println(dayTime.Unix())

//t := time.Now()
//获取今天零点零分时间戳
//dayTime := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, t.Location()).Unix()
//EndTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
//fmt.Println(dayTime, "\n")

//dayTime+86400-1

//fmt.Println(dayTime + 86399)
//86,400
