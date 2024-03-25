package holiday

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var holidayMap map[string]string

type Holiday struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	Type    int    `json:"type"`
	Summary string `json:"summary"`
}

func init() {
	holidayMap = make(map[string]string, 125)
	var holidays []Holiday
	file, err := os.Open("holiday2024.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	bytes, _ := reader.ReadBytes('\n')
	err = json.Unmarshal(bytes, &holidays)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, holiday := range holidays {
		holidayMap[holiday.Date] = strconv.Itoa(holiday.Type)
	}
}

func checkDate(date string) error {
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	if err != nil {
		return errors.New("日期格式不正确")
	}
	return nil
}

// IsHoliday 判断是否节假日，第一个入参表示`2006-01-02`的日期字符串，第二个入参表示是否包含双休
func IsHoliday(date string, isInclude bool) (bool, error) {
	if err := checkDate(date); err != nil {
		return false, err
	}
	value, exists := holidayMap[date]
	if exists && (value == "0" || (isInclude && value == "2")) {
		return true, nil
	} else {
		return false, nil
	}
}

// IsWeekday 判断是否工作日，入参表示`2006-01-02`的日期字符串
func IsWeekday(date string) (bool, error) {
	if err := checkDate(date); err != nil {
		return false, err
	}
	value, exists := holidayMap[date]
	if !exists || exists && value == "1" {
		return true, nil
	}
	return false, nil
}
