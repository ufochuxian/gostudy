package myutil

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func ExecuteBuildSublesson(sublessoninfo SubLessonInfo) {
	log.Println("start executeBuildSublesson")
	sublid := sublessoninfo.SubLessonId
	buildCommand := "/Users/chen/Documents/TGM/cocos/tgm-clientcocos/buildTools/build.sh '" + strconv.Itoa(sublid) + "' '" + sublessoninfo.SubpackageVersion + "' '" + sublessoninfo.BasepackageVersion + "'"
	output, err := exec.Command("/bin/sh", "-c",buildCommand).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}
