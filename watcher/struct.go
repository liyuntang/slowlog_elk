package watcher

// json数据格式

type elkData struct {
	timestamp, index, product, cluster, role, hostname, Rows_sent, Rows_examined, mode, msg string
	Query_time, Lock_time float64
}
