package main

import (
	"fmt"
	ansibler "github.com/apenella/go-ansible"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"net/http"
)

type Alert struct {
	ID       	string `json:"id" binding:"required"`
	Hostname 	string `json:"hostname" binding:"required"`
	Username 	string `json:"username" binding:"required"`
	RuleName 	string `json:"rulename" binding:"required"`
	Inventory 	string `json:"inventory" binding:"required"`
}


func main() {
	router := gin.Default()
	router.POST("/actions", posting)
	router.Run(":8080")

}

func posting(c *gin.Context){
	var json Alert
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ansible(json)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func ansible(json Alert)  {
	rulereturn := rulelookup(json)
	playbookname := fmt.Sprintf("%s", rulereturn)
	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: json.Inventory+",",
		ExtraVars: map[string]interface{}{
			"example": "example",
			"ruleName": json.RuleName,
			"hostname": json.Hostname,
			"username": json.Username,
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

func rulelookup(json Alert) (interface{}){
	jq := gojsonq.New().File("./alerts.json")
	res := jq.Find(json.RuleName)
	return res
}
