# ElastiSOAR
A small PoC for SOAR-like functionality using Elasticsearch and Ansible

## Intro 
This tool will listen for a webhook from an Elasticsearch cluster, and pull out values such as: 
- hostname
- alert name

A lookup function will take the alert name value, and pull out the corresponding playbook. 

The hostname value will be used to pass into ansible's inventory flag (the host to run this operation against)

## Make it work

### Detection Engine
In Kibana, go to the Detections tab within Elastic Security
- Select the rule you wish to use ElastiSOAR with
- Click on 'Edit Rule settings'
- In the actions tab, confirm the frequency you want the rule to trigger. 
- Under action type, select webook. 
  - If you don't have a connector defined already then please set this up (the default port for ElastiSOAR is 8080)
- Create a document body that contains the following values as a minimum. 

### Watcher
- Create a watch that looks for your specific criteria
- Under the action part of the watch make sure that the payload contains at a minimum: 
  - Hostname
  - Username
  - Alert Name 

> for the time being - you will need to add hostname and username, even if those values are blank. 

## To do 
- prepacked playbooks
- tests
- Add the ability to write the results of runs into Elasticsearch for analytics
- tweak binding validation 
- Add elastic APM
## Props to 
- https://github.com/gin-gonic/gin
- https://github.com/apenella/go-ansible
