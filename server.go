package main

// http://go-database-sql.org/retrieving.html

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var htmllist string

func main() {

	var data Node
	// Stores pdbout.json into byte array jsonout
	jsonout, readerr := ioutil.ReadFile("pdbout.json")

	if readerr != nil {
		log.Fatal(readerr)
	}

	parseerr := json.Unmarshal(jsonout, &data)
	if parseerr != nil {
		log.Fatal(parseerr)
	}
	for _, node := range data {
		//fmt.Println(node.Certname)
		htmllist += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", node.Certname, node.CatalogTimestamp)
	}

	fmt.Println(htmllist)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":1337", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	index, _ := ioutil.ReadFile("static/index.html")

	temp := template.New("Puppet Template")
	temp, _ = temp.Parse(string(index))
	temp.Execute(w, template.HTML(htmllist))
	//	fmt.Fprintf(w, string(index))
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
