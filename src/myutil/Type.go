package myutil


type SublessonInfos struct {
	SublessonInfos []*Sublessoninfo
}

type Sublessoninfo struct {
	SublessonId        int    `json:"sublessonid"`
	SubpackageVersion  string `json:"subpackageversion"`
	BasepackageVersion string `json:"basepackageversion"`
}

