package main

//func main() {
//	//log.Println("start executeBuildSublesson")
//	buildCommand := "/Users/chen/Documents/TGM/cocos/tgm-clientcocos/buildTools/build.sh '1' '1' '1'"
//	//cmd := exec.Command("cmd","/C","/Users/chen/Documents/TGM/cocos/tgm-clientcocos/buildTools/build","1","1","1")
//	////cmd.Dir = "/Users/chen/Documents/TGM/cocos/tgm-clientcocos/buildTools/"
//	//output, err := cmd.Output()
//	//if err != nil {
//	//	fmt.Println("err:,output:",err,output)
//	//}
//	//log.Println("end execute command:", output)
//
//	output, err := exec.Command("/bin/sh", "-c",buildCommand).CombinedOutput()
//	if err != nil {
//		os.Stderr.WriteString(err.Error())
//	}
//	fmt.Println(string(output))
//}