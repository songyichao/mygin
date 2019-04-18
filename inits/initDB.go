package init

import (
	"fmt"
	"mygin/config"
)

func InitDB() {
	dbConfig := config.DbConfig()
	fmt.Println(dbConfig["default"]["dbUser"])
	//dsn := dbConfig["default"]["dbUser"] + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	//if timezone != "" {
	//	dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	//}
	//orm.RegisterDataBase("default", "mysql", dsn)
}
