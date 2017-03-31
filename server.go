package main

// http://go-database-sql.org/retrieving.html

import (
	"reflect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
//	"html/template"
)

func main() {

	var data Node
	// Stores pdbout.json into byte array jsonout
	jsonout, readerr := ioutil.ReadFile("pdbout.json")
//	index, _ := ioutil.ReadFile("static/index.html")

	if readerr != nil {
		log.Fatal(readerr)
	}

	parseerr := json.Unmarshal(jsonout, &data)
	if parseerr != nil {
		log.Fatal(parseerr)
	}

/*
	fmap := template.FuncMap {
		"listnodes": func(n anode) string {
			return //list nodes 
	}
*/

	fmt.Println("Here is a list of all nodes:\n")
	fmt.Println(reflect.TypeOf(data))
	for _,node := range data {
		fmt.Println(node.Certname)
	}
//	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1337", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	index, _ := ioutil.ReadFile("static/index.html")
	fmt.Fprintf(w, string(index))
}

type Node []struct {
	Deactivated                  interface{} `json:"deactivated"`
	LatestReportHash             string      `json:"latest_report_hash"`
	FactsEnvironment             string      `json:"facts_environment"`
	CachedCatalogStatus          string      `json:"cached_catalog_status"`
	ReportEnvironment            string      `json:"report_environment"`
	LatestReportCorrectiveChange interface{} `json:"latest_report_corrective_change"`
	CatalogEnvironment           string      `json:"catalog_environment"`
	FactsTimestamp               time.Time   `json:"facts_timestamp"`
	LatestReportNoop             bool        `json:"latest_report_noop"`
	Expired                      interface{} `json:"expired"`
	LatestReportNoopPending      bool        `json:"latest_report_noop_pending"`
	ReportTimestamp              time.Time   `json:"report_timestamp"`
	Certname                     string      `json:"certname"`
	CatalogTimestamp             time.Time   `json:"catalog_timestamp"`
	LatestReportStatus           string      `json:"latest_report_status"`
}
