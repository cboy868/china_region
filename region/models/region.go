package models

//Region 数据
type Region struct {
	Pcode string //父级code
	Code  string
	Name  string
	Type  string //地区类型 比如 province-> city --> country --> town --> village
}
