package main

import (
	"database/sql"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	mp "github.com/mackerelio/go-mackerel-plugin-helper"
	"github.com/monochromegane/dsn"
)

var env string
var dbconf string

var graphdef map[string](mp.Graphs) = map[string](mp.Graphs){
	"delayed_job": mp.Graphs{
		Label: "Delayed Job Count",
		Unit:  "integer",
		Metrics: [](mp.Metrics){
			mp.Metrics{Name: "count", Label: "Job Count"},
			mp.Metrics{Name: "error", Label: "Error Job Count"},
		},
	},
}

type JobCountPlugin struct {
	Tempfile string
}

func (j JobCountPlugin) FetchMetrics() (map[string]interface{}, error) {

	name, dsn, err := dsn.FromRailsConfig(dbconf, env)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(name, dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var count float64
	if err := db.QueryRow("SELECT COUNT(id) FROM delayed_jobs;").Scan(&count); err != nil {
		return nil, err
	}

	var errCount float64
	if err := db.QueryRow("SELECT COUNT(id) FROM delayed_jobs WHERE failed_at IS NOT NULL;").Scan(&errCount); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"count": count,
		"error": errCount,
	}, nil
}

func (j JobCountPlugin) GraphDefinition() map[string](mp.Graphs) {
	return graphdef
}

func init() {
	flag.StringVar(&env, "env", "development", "environment")
	flag.StringVar(&dbconf, "dbconf", "", "Rails dbconfig file")
	flag.Parse()
}

func main() {
	helper := mp.NewMackerelPlugin(JobCountPlugin{})
    helper.Run()
}
