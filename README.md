# go Mobile API Rest Services

## Description
This implementaion connects to a MondoDB instance using the environment variables and enables the user to set the 
Atlas API keys, to access the Atlas API Services over REST Services.
The implementation of the services follow clean architecture pattern

Rule of Clean Architecture by Uncle Bob
 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has  4 Domain layer :
 * Models Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer
 
![golang clean architecture] ![Clean Architecture](https://user-images.githubusercontent.com/10128767/125108572-d1443300-e0a7-11eb-9b11-550b4c40c83a.png)

### How To Run This Project

Clone the project 
Create .env file with below contents

 
```
DB_URI=<Your MongoDB Connection URL>
PORT=8000
BASE_URL=https://cloud.mongodb.com/api/atlas/v1.0
GROUPS_PATH=/groups
JWT_SECRET_KEY=<ANY JWK SECRET KET FOR AUTH>
```


```
cd  api-mobile-console/cmd/mongo-admin-api
go run main.go
```
