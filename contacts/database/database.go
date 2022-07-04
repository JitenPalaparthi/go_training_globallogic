package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MAXRETRIES = 10
	MAXWAIT    = time.Second * 2
)

// func usage() {
// 	//fmt.Fprintf(os.Stderr, "usage: -stderrthreshold=[WARNING|FATAL] -log_dir=[string]\n")
// 	flag.PrintDefaults()
// 	os.Exit(2)
// }

// func init() {
// 	flag.Usage = usage
// 	// NOTE: This next line is key you have to call flag.Parse() for the command line
// 	// options or "flags" that are defined in the glog module to be picked up.
// 	flag.Parse()
// }

//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
func GetConnection(url string) (interface{}, error) {
	count := 1
retry:
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		if count <= MAXRETRIES {
			count = count + 1
			time.Sleep(MAXWAIT)
			fmt.Printf("Trying to connect to the database %v number of times\n", count)
			goto retry
		}
		return db, err
	}
	return db, err
}

//"POSTGRES_PASSWORD=admin",
//"POSTGRES_USER=jiten",
//"POSTGRES_DB=sample",
// 5432
// dsn := "host=localhost user=jiten password=admin dbname=sample port=5432 sslmode=disable TimeZone=Asia/Shanghai"
//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
