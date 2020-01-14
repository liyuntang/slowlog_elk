package watcher

import (
	"github.com/hpcloud/tail"
	"log"
	"os"
	"slowlog_elk/tomlConfig"
	"strings"
	"time"
)

func Start(sl *tomlConfig.SL, loggs *log.Logger) {
	// 获取slowlog信息
	slowLog := sl.System.SlowLog
	// 检测slowlog文件
	if err := check(slowLog); err != nil {
		loggs.Println(err)
		os.Exit(0)
	}
	// 到此说明slowlog文件正常，开始检测文件
	watcher, err := tail.TailFile(slowLog, tail.Config{Follow: true})
	if err != nil {
		loggs.Println("start watch file", slowLog, "is bad, err is", err)
		os.Exit(0)
	}
	data := &elkData{}
	SQL := ""
	for line := range watcher.Lines {
		// 以下变量都是从配置中获取的
		data.timestamp = time.Now().Format(time.RFC3339)
		data.index = sl.Elk.Index
		data.product = sl.Elk.Product
		data.cluster = sl.Elk.Cluster
		data.role = sl.Elk.Role
		data.hostname = sl.Elk.HostName
		// 以下变量需要从文件中获取
		if strings.HasPrefix(line.Text, "#") {
			sql := strings.Trim(line.Text, "# ")
			if strings.HasPrefix(sql, "Query_time") {
				zouqi(sql, data, loggs)
			}
		} else {
			sql := strings.Trim(line.Text, "# ")
			if !strings.HasPrefix(sql, "USE") && !strings.HasPrefix(sql, "use") && !strings.HasPrefix(sql, "SET") && !strings.HasPrefix(sql, "set") {
				// 这个地方需要兼容换行写的sql
				if strings.HasSuffix(sql, ";") {
					// 以;结尾说明这个sql到此结束，此时需要处理sql了
					if len(SQL) != 0 {
					//说明该sql不是完整的sql，此时需要拼接sql
						SQL = SQL + " " + sql
						data.msg = SQL
					} else {
						// 说明这个是完整的sql语句，此时不需要拼接sql
						data.msg = sql
					}
					// 获取sql模型
					data.mode = sqlMode(sl.System.HttpServer, data.msg, loggs)
					// 上传到elk程序
					upload(data, sl.Elk.EndPoint, loggs)
					// 重置变量
					SQL = ""
					data = &elkData{}
				} else {
					// 说明这个是换行写的sql，此时sql需要拼接
					SQL = SQL + " " + sql
				}
			}

		}

	}

}

func zouqi(sql string, data *elkData, loggs *log.Logger, )  {
	defer func() {
		if e:=recover();e != nil {
			loggs.Println(e)
		}
	}()
	analy(sql, data)
}






