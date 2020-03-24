---
layout: page
title:  Assignment 3 Microservice ETL Pipeline
permalink: /assignments/etl-pipeline/
---

In this project you will take a monolithic web service and turn it into a set of microservices. Along the way, you will learn about the Go programming language, web, and database programming. The assignment is based on the [Dataquest tutorial Building An Analytics Data Pipeline In Python](https://www.dataquest.io/blog/data-pipelines-tutorial/) (you are welcome to look at this for further details on the project architecture).

<blockquote>
**General Requirements:** 
  - Your solution must be written in Go.
  - You may work individually or in a team with maximum size 3 students.
  - You must create your repository in [GitHub Classrooms](https://classroom.github.com/g/qxHM8OIw). (**DON'T DO THIS UNTIL WE TELL YOU TO!**)
  - All team members must contribute commits to the repository and work should be divided evenly. You should regularly commit and push as you work, instead of waiting until the end to push.

*Due Sunday April 5th, 11:59pm*
</blockquote> 

## ETL Data Pipelines
We will be building a data pipeline web service that follows the ETL (Extract, Transform, Load) architecture. [ETL Pipelines](https://databricks.com/glossary/etl-pipeline) are a popular way to take input data (in this case web log files), *extract* the significant fields from the log entries, *transform* the data into a more useful format, and *load* the results into a database for future reporting.


## The Monolith
You are being provided with a monolithic web service implementation of an ETL pipeline written in Go.  In a real ETL pipeline, data is continuously arriving (e.g., as visitors browse a website), but for this assignment we will upload a log file to be processed once in a batch. The provided [Monolith](https://github.com/gwAdvNet20/ETL-pipeline/tree/master/monolith) includes code to let a user upload a log file through a web interface.  It then parses the log file to clean up the data and performs two analytic functions: counting the frequency of different browsers and counting the number of visitors per day. The results of these analytics are then stored in a [sqlite database](https://www.sqlite.org/index.html). To retrieve the results from the database, a user can access a REST API that returns the requested information in JSON format.

![The Monolith](monolith.png)

### Log Format
The Monolith includes code to process log files in the following format:
```
242.243.182.188 - - [13/Mar/2020:12:50:55 +0000] "GET /list/explore HTTP/1.1" 404 277 "https://gallegos-rice.net/" "Mozilla/5.0 (Windows; U; Windows NT 6.1) AppleWebKit/532.1.6 (KHTML, like Gecko) Version/4.0 Safari/532.1.6"
251.121.9.241 - - [13/Mar/2020:12:50:55 +0000] "PUT /list/categories/posts HTTP/1.1" 401 136 "https://barnes-powell.com/list/post/" "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/5.0)"
222.29.58.215 - - [13/Mar/2020:12:50:55 +0000] "POST /list/app/wp-content HTTP/1.1" 401 312 "http://www.walker.com/login/" "Mozilla/5.0 (X11; Linux x86_64; rv:1.9.7.20) Gecko/2015-03-26 05:50:09 Firefox/3.8"
39.116.112.218 - - [13/Mar/2020:12:50:55 +0000] "PUT /explore HTTP/1.1" 404 607 "https://beasley.com/homepage.htm" "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_9; rv:1.9.3.20) Gecko/2016-09-08 02:20:09 Firefox/3.8"
85.64.31.244 - - [13/Mar/2020:12:50:55 +0000] "GET /tags/app HTTP/1.1" 401 288 "http://ramirez.biz/" "Mozilla/5.0 (Windows CE; it-IT; rv:1.9.0.20) Gecko/2012-12-20 12:07:13 Firefox/3.6.18"
160.25.38.0 - - [13/Mar/2020:12:50:55 +0000] "GET /tag HTTP/1.1" 404 584 "https://taylor.com/blog/explore/register/" "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_8_1) AppleWebKit/5361 (KHTML, like Gecko) Chrome/14.0.853.0 Safari/5361"
```
For details of the log format, [check here](https://www.dataquest.io/blog/data-pipelines-tutorial/). 

### Running the Monolith

> If you need to change the port you will need to modify `config.yaml` and change the value of monolith. You will also need to update the html file.

To test the Monolith, compile and run it:

```
cd monolith
go build .
./monolith 
```

> If you get errors about missing packages, you will need to run commands like `go get github.com/mattn/go-sqlite3` to install them!

Then use your browser to go to [http://localhost:8000](http://localhost:8000)

  - You will need to log in with username / password `admin` / `password`.

You will be presented with a simple webpage that allows you to upload a file and start the pipeline.  The `logs` directory contains some sample logs to use.

The Monolith will show a status message when it completes running the pipeline.

  - To access the Browser Count results, visit: [http://localhost:8000/browser/count](http://localhost:8000/browser/count)
  - To access the Visitor Count results, visit: [http://localhost:8000/visitor/count](http://localhost:8000/visitor/count)

If you want to see the data that was stored in the database, you can use a tool like [sqlite browser](https://sqlitebrowser.org/).

## The Microservices
As we have discussed in class, the monolithic design has some drawbacks such as making it harder to upgrade or scale individual components. Your task is to implement a microservice based solution, which will look more like the following:

![The Microservices](microservices.png)

While typically each microservice is responsible for maintaining its own database, we will use a single sqlite database with separate tables for each microservice.

### Your Implementation

We have provided you with a basic framework for your microservice implementation. Currently it only includes an API gateway and a simple analysis microservice to count the number of lines in a log file. To complete the assignment you must:

 - Move the data cleaning functionality out of the API gateway and into `ms-data-cleaning` service. You should start here to get an idea of how to build a microservice. We have included a basic framework that allows you to make a post and get request. You should modify these to handle the data cleaning.
 - Implement `ms-browser-counts` and `ms-visitor-counts` microservices which count the frequency of browsers and visitors (you can base this on the code in the Monolith).
 - Implement a new `ms-website-counts` microservice which counts the number of times each URL appears per day. You will have to build this from scratch, although the other examples will give you guidance. 

For a greater challenge, you can also:
 - Move the authentication functionality out of `ms-api-gateway` into its own microservice, `ms-auth`.
 - Create additional reports, such as 1) returning the top 5 most popular websites and 2) returning a list of all websites visited by a particular IP address.


### Running the Microservices

To test the provided Microservices, compile and run them:
```
# Start API Gateway in first terminal
cd ms-api-gateway
go build .
./ms-api-gateway


# Start Line Count in second terminal
cd ms-line-count
go build .
./ms-line-count
```

  - To access the Log upload page, visit: [http://localhost:80](http://localhost:80)
  - To access the Line Count results, visit [http://localhost/lines/count/FILE_NAME_HERE](http://localhost/lines/count/FILE_NAME_HERE) 
  - To access the Browser Count results, visit: [http://localhost:80/browser/count](http://localhost:80/browser/count) *Note: You need to implement this first.*
  - To access the Visitor Count results, visit: [http://localhost:80/visitor/count](http://localhost:80/visitor/count) *Note: You need to implement this first.*
  - **Note the different port (80) used for the microservice API Gateway!**


### Where to start?

You should start this assignment by first looking at how `ms-line-count` is implemented and how it interacts with `ms-api-gateway`. You should do your best to see how the code flows and how it handles and responds to requests from `ms-api-gateway`. Especially notice how all responses use a standard format. Your api should use the same format using the structs and functions in `shared.go`. Also note, all requests to the api will need to go through `ms-api-gateway` and be forwarded to the microservices doing the actual work, you should not directly talk to individual microservices.

Once you understand how `ms-line-count` works, you should work on moving the log file parsing and cleaning functionality out of `ms-api-gateway` and into `ms-data-cleaning`. We have provided a skeleton for `ms-data-cleaning` that provides basic endpoints for GET and POST. You will need to modify these to accept the raw log file bytes from `ms-api-gateway`, clean them and store them in the database just like the monolith does.

Once you have `ms-data-cleaning` working, you should build the rest of your microservices.


---

## FAQ
**Q:** I have a question, what do I do?
 - **A:** Message us in `#a3etl` on slack and we will add it here!

**Q:** How do I learn Go? 
  - **A:** Check our [Go Resources](/wiki/go/) for videos and links to get you started!

**Q:** How can my microservices find each other? 
  - **A:** Normally a service discovery framework would be used for components to learn how to communicate with each other. For simplicity, we will assume all microservices run on localhost, and the ports are stored in [`config.yaml`](https://github.com/gwAdvNet20/ETL-pipeline/blob/master/config.yaml). The contents of the config file are accessible using the [viper](https://github.com/spf13/viper) library, e.g., `viper.GetString("services.ms-data-cleaning")`

 **Q:** Can I modify the `shared.go` file in the repos?
 - **A:** We have provided the `shared.go` file as a way to help your code remain consistent. You should use the various functions in this file in your api but should not modify them. 

