# Postman Assignment
Mysql dbconectivity used in goLang to run db operations in Docker container using Docker file to monitor autoincrement value and post it to slack channel here 
* [slack channel](https://postman-assignment.slack.com/messages/CDUBBCUV7/)
* [Invite Link](https://join.slack.com/t/postman-assignment/shared_invite/enQtNDcxNzIyNzM5NjM3LWZjZjBiMTU4OWI1NmQzNWNmYWQyOTRmZDBiZmRlMjQyNDIxMmExNTUyMmU1NDkxZWE3ZWM4NmRmNjhhYWJiNDQ)


### Reference
* https://github.com/go-sql-driver/mysql
* https://severalnines.com/blog/mysql-docker-containers-understanding-basics

### Pre-requisite
* Sql Client should already be installed in machine.
* apt-get update && apt-get install -y mysql-client
* Docker environment

### Introduction
* We need mysql running container in order to link it with our container. So we need to initiate a
mysql container. Container runs in detached mode because we need it as a service.
`docker run --detach --name=<containerName> --env="MYSQL_ROOT_PASSWORD=<mypassword>" mysql`

* To see logs of the running container use
 ` docker logs <containerName> `

* Install sql client locally
`apt-get install mysql-client`

* Test the current IP Address of mysql
` docker inspect <containerName> `

* Using obtained IP Address type
 `mysql -uroot -p<mypassword> -h <IPAdress> -P 3306 `
By default port 3306 is assigned to mysql, it can be changed if needed

* Build Docker image with `docker build -t gomysql .`

* Execute the code with `docker run gomysql`

### Assumption
The code first connects to sql and then creates testdb database if does not exist already.It then creates a Table called people and insert a value in the table with every instance of program.The code sends updates to slack channel.It is assumed that if autoincrement value reaches more than 7 a warning message is sent to slack channel.
