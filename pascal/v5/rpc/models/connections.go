package models

//Connection A "Connection object" is a JSON object with a connection to other node information
type Connection struct {
	Server   bool   `json:"server"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Secs     int    `json:"secs"`
	Sent     int    `json:"sent"`
	Recv     int    `json:"recv"`
	Appver   string `json:"appver"`
	Netvar   int    `json:"netvar"`
	NetvarA  int    `json:"netvar_a"`
	TimeDiff int    `json:"timediff"`
}
