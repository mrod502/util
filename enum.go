package util

//Table - enum for table prefix
type Table string

//Tables
const (
	Sep = "^"

	TblPosition   Table = "00"
	TblHBTCOrder  Table = "01"
	TblHBTCTicker Table = "02"
	TblHBTCSymbol Table = "03"
	TblHBTCTrade  Table = "04"
	TblHBTCCandle Table = "05"
	TblHBTCMDepth Table = "06"
	TblErrResp    Table = "07"
	TblMessage    Table = "08"
)
