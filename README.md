# Appointy-Golang-Task

# Task Document
https://docs.google.com/document/d/1sFhVumoczf_PmaL_R__Rm9AHqaHsUWgj1x9YcQP6Is4/preview?pru=AAABfIQlpAg*OvSXDNnjm9wo0TbhvDB9XA

# Endpoints
<b> https://documenter.getpostman.com/view/16298894/UUy7cQDg </b>
<ol>
  <li> POST localhost:8085/users </li>
  <li> GET localhost:8085/users/:id </li>
  <li> POST localhost:8085/posts </li>
  <li> GET localhost:8085/posts/:id </li>
</ol>

# Documentation
Check <br>
<b>1. Appointy Task.postman_collection.json </b>
<b>2. Screenshots.pdf </b>

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
2. Readable code
3. Proper modules for easy scalibility
4. Strictly used std golang packages 
5. Quality of code
6. Reusable code
7. Consistency in code

# How to run
1. Clone repository
2. RUN - go get go.mongodb.org/mongo-driver/mongo
3. go run main.go
4. Database will be setup on its own. used localhost mongodb and not cloud.
