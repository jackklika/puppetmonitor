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

	//fmt.Println(letstls())

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

	isnode := false
	for _,n := range data {
		if (nodename == n.Certname){
			isnode = true
		}
	}
	if (isnode == false) {
		fmt.Fprintf(w, "This is not a node")
		return
	} else {
		var out string
		out += "<html><body><h1>" + nodename + "</h1><ul>"
		var nodedata IndNode
		jsonout := letstls(nodename)
		parseerr := json.Unmarshal([]byte(jsonout), &nodedata)
		if parseerr != nil{
			log.Fatal(parseerr)
		}

		for _,n := range nodedata {
			switch nval := n.Value.(type) {
			case string:
				out += fmt.Sprintf("<li>%s: %s</li>\n", n.Name, n.Value)
			case float64:
				out += fmt.Sprintf("<li>%s: %v</li>\n", n.Name, n.Value)

			case bool:
				out += fmt.Sprintf("<li>%s: %v</li>\n", n.Name, n.Value)

			case []interface{}:
			
			case map[string]interface {}:
				out += "\t<ul>\n"
				for i, u := range nval {
					out += fmt.Sprintf("\t<li>%s: %v</li>\n", i, u)
				}
				out += "\t</ul>\n"
			}
		}


		out += "</ul></body></html>"
		fmt.Fprintf(w, out)
	}

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


type IndNode []struct {
	Certname		string	`json:"certname"`
	Name			string	`json:"name"`
	Value			interface{}	`json:"value"`
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
