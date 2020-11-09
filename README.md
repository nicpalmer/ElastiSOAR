# ElastiSOAR
A small PoC for SOAR-like functionality using Elasticsearch and Ansible

## Intro 
This tool will listen for a webhook from an Elasticsearch cluster, and pull out values such as: 
- hostname
- alert name

A lookup function will take the alert name value, and pull out the corresponding playbook. 

The hostname value will be used to pass into ansible's inventory flag (the host to run this operation against)


## To do 
- prepacked playbooks
- tests
- Add the ability to write the results of runs into Elasticsearch for analytics
- Add elastic APM
## Props to 
- https://github.com/gin-gonic/gin
- https://github.com/apenella/go-ansible
