package main

import (
	"fmt"
	ansibler "github.com/apenella/go-ansible"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Alert struct {
	ID       		string 			`json:"id" binding:"required"`
	RuleName 		string 			`json:"rulename" binding:"required"`
	Inventory	 	string 			`json:"inventory" binding:"required"`
	Host			Host 		 	`json:"host" `
	User			User  			`json:"user"`
	Process 		Process			`json:"process"`
	Hash			Hash			`json:"hash"`
	Vulnerability 	Vulnerability 	`json:"vulnerability"`
	Source			Source			`json:"source"`
	Destination		Destination		`json:"destination"`
}
type Host struct {
	Name			string			`json:"name" `
	Platform 		string  		`json:"platform" `
}
type User struct {
	Name			string 			`json:"name" `
	Email			string			`json:"email" `
	Group			string 			`json:"group" `
}
type Process struct {
	Name			string 			`json:"name" `
	Pid				uint16 			`json:"pid"`
}
type Hash struct {
	SHA256			string 			`json:"sha256" `
}
type Vulnerability struct {
	Category		string			`json:"category"`
	Description		string			`json:"description"`
	ReportID		string			`json:"report_id"`
}
type Source struct {
	IP				string			`json:"ip" `
	Address 		string  		`json:"address" `
}
type Destination struct {
	IP				string			`json:"ip" `
	Address 		string  		`json:"address" `
}
func main() {
	router := gin.Default()
	router.POST("/actions", posting)
	router.Run(":8080")
}
func posting(c *gin.Context){
	var json Alert
	if c.BindJSON(&json) == nil {
		fmt.Printf("%+v\n", json)
	}
	ansible(json)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ansible(json Alert)  {
	playbookname := fmt.Sprintf("%s.yml", json.RuleName)
	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
	}
	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: json.Inventory+",",
		ExtraVars: map[string]interface{}{
			"ruleName": json.RuleName,
			"hostname": json.Host.Name,
			"username": json.User.Name,
			"sourceIP": json.Source.IP,
			"sourceAddress": json.Source.Address,
			"destinationIP": json.Destination.IP,
			"destinationAddress": json.Destination.Address,
			"VT_API_KEY": "<YOUR API KEY>",
		},
	}
	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          playbookname,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		ExecPrefix:        "ElastiSOAR",
	}
	fmt.Println(playbook.Run())
}
