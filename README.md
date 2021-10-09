# Appointy-Golang-Task

# Task Document
https://docs.google.com/document/d/1sFhVumoczf_PmaL_R__Rm9AHqaHsUWgj1x9YcQP6Is4/preview?pru=AAABfIQlpAg*OvSXDNnjm9wo0TbhvDB9XA

# Endpoints
<h2> https://documenter.getpostman.com/view/16298894/UUy7cQDg </h2>
<ol>
  <li> POST localhost:8085/users </li>
  <li> GET localhost:8085/users/:id </li>
  <li> POST localhost:8085/posts </li>
  <li> GET localhost:8085/posts/:id </li>
</ol>

# Documentation
Check Appointy Task.postman_collection.json 

# Dependencies Used
 - Only std packages
 - Used SHA256 for hashing passwords
 - MongoDB drivers for connecting to MongoDB (localhost)

# Collections used in MongoDB
1. Users - Id, Name, Email, Password (hashed)
2. Posts - Id, Caption, Image URL, Posted Timestamp

# What's left
1. List all posts of a user - GET request - URL should be ‘/posts/users/:Id here'

# Code Merits
1. Proper commented codes
2. Proper modules for easy scalibility
3. Strictly used std golang packages 
4. Quality of code
5. Reusable code
6. Consistency in code

# How to run
1. Clone repository
2. RUN - go get go.mongodb.org/mongo-driver/mongo
3. go run main.go
4. Database will be setup on its own. used localhost mongodb and not cloud.
