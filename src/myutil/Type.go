package myutil

import "fmt"

type SublessonInfos struct {
	SublessonInfos []*Sublessoninfo
}

type Sublessoninfo struct {
	SublessonId        int    `json:"sublessonid"`
	SubpackageVersion  string `json:"subpackageversion"`
	BasepackageVersion string `json:"basepackageversion"`
}

type Result struct {
	Sublessoninfo
}

func (r Result) Success() {
	fmt.Println(r.SublessonId)
	ExecuteBuildSublesson(r.Sublessoninfo)
}

func (r Result) Loading() {

}

func (r Result) Fail() {

}