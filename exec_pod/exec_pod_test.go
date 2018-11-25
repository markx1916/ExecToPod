package exec_pod

import (
	"fmt"
	"testing"
)

//func TestExecToPodThroughAPI(t *testing.T) {
//	//fmt.Println(HolidayInfo(time.Now()))
//	ExecToPodThroughAPI("hostname","","")
//
//
//}

func TestGetClientConfig(t *testing.T) {
	config, err := GetClientConfig()
	if err != nil {
		fmt.Println(err)
	}
	_, err = GetClientsetFromConfig(config)
	if err != nil {
		fmt.Println(err)
	}
}
