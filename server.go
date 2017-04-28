package main

// http://go-database-sql.org/retrieving.html

import (
	"github.com/patrickmn/sortutil"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var htmllist string
var data Node
func main() {

	readandgenerate()

	http.HandleFunc("/nodes/", nodeshandler)
	http.HandleFunc("/", homehandler)
	log.Fatal(http.ListenAndServe(":80", nil))



}

func readandgenerate() {

	// Reads pdbout.json, which is a json file fetched from puppetdb every minute by a cronjob.
	jsonout, readerr := ioutil.ReadFile("/opt/puppetmonitor/pdbout.json")

	if readerr != nil {
		log.Fatal(readerr)
	}

	parseerr := json.Unmarshal(jsonout, &data)
	if parseerr != nil {
		log.Fatal(parseerr)
	}

	sortutil.AscByField(data, "Certname")


}

func homehandler(w http.ResponseWriter, r *http.Request) {
	readandgenerate();
	index, _ := ioutil.ReadFile("/opt/puppetmonitor/static/index.html")

	temp := template.New("Puppet Template")
	temp = temp.Funcs(template.FuncMap{"curTime": curTime, "minusCurTime" : minusCurTime})
	temp, _ = temp.Parse(string(index))
	//fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05 1/2/2006 MST"), r.RemoteAddr)
	temp.Execute(w, data)
	//	fmt.Fprintf(w, string(index))
}

func nodeshandler(w http.ResponseWriter, r *http.Request) {
	nodename := r.URL.Path[len("/nodes/"):] // fqdn/nodes/nodename

//	stringout := ""
	isnode := false
	for _,n := range data {
		//stringout += fmt.Sprintf("%s\n", n.Certname)
		if (nodename == n.Certname){
			isnode = true
		}
	}
	if (isnode == false) {
		nodename = "NOTANODE"
	}
		fmt.Fprintf(w, nodename)

//	fmt.Fprintf(w, stringout)

}

func curTime(t time.Time) string {
	delta := time.Since(t)
	return fmt.Sprintf("%02d:%02d\n", delta.Nanoseconds()/time.Minute.Nanoseconds(), delta.Nanoseconds()/time.Second.Nanoseconds()%60)
//	return fmt.Sprintf("%s", delta)
}

func minusCurTime(t time.Time) string {
	delta := time.Since(t)
	diff := 255 - delta.Nanoseconds()/time.Minute.Nanoseconds()
	return fmt.Sprintf("%v", diff)
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

/*
func reqnodeHTML(field interface{}) {
	return fmt.Sprintf("<td>%s</td>")
}
*/
