package config

import (
	elastic "gopkg.in/olivere/elastic.v5"
	"log"
)


func GenClient(conf Config, errLog *log.Logger) (*elastic.Client, error) {
	if errLog != nil {
		return elastic.NewSimpleClient(
			elastic.SetURL(conf.ELASTIC_URL),
			elastic.SetBasicAuth(conf.USERNAME, conf.PASSWORD),
			elastic.SetErrorLog(errLog),
		)
	} else {
		return elastic.NewSimpleClient(
			elastic.SetURL(conf.ELASTIC_URL),
			elastic.SetBasicAuth(conf.USERNAME, conf.PASSWORD),
		)
	}
}
