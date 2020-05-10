---
layout: page
title:  Distributed Load Testing
permalink: /wiki/benchmark/
---

*by:* Samuel Farid and Hamid Reza


A short description of your post goes here.

---
# Creating an Environment and Deploying a Web Application with Elastic Beanstalk

In this experiment we will deploy a sample web application in amazon cloud servers by Elastic Beanstalk service. As Amazon explains: “AWS Elastic Beanstalk is an easy-to-use service for deploying and scaling web applications and services developed with Java, .NET, PHP, Node.js, Python, Ruby, Go, and Docker on familiar servers such as Apache, Nginx, Passenger, and IIS”.

The sample application contains a welcome webpage with couple of links to other resources. The mentioned web application is deployed on four different machine types: t2.micro, t2.large, r5.large and m5.large. Details for each of these instance types are available below (Table 1).

![alt text](table.PNG "Table 1: Amazon EC2 Pricing.")
<div align="center">Table 1: Amazon EC2 Pricing.</div>

For deploying a web application with Elastic Beanstalk, first we have to login to the AWS console. Then search Elastic Beanstalk and open launch the service (Screenshot 1).
 

![alt text](1.PNG "Screenshot 1")
<div align="center">Screenshot 1.</div>


Then we must create an environment for our web application (Screenshot 2). In this section, we can configure our environment to test the application with different machine types.
 
![alt text](2.PNG "Screenshot 2")
<div align="center">Screenshot 2.</div>

Here we should select “Web server environment” for deploying a web application (Screenshot 3).
 
![alt text](3.PNG "Screenshot 3")
<div align="center">Screenshot 3.</div>




In the next section, we can choose a name and domain (Screenshot 4 and 5) for our environment. For changing “.us-east-2.elasticbeanstalk.com”, we can change our region  which will also effect the price of the instance that we will use.
 
![alt text](4.PNG "Screenshot 4")
<div align="center">Screenshot 4.</div>
 
![alt text](5.PNG "Screenshot 5")
<div align="center">Screenshot 5.</div>



In platform section we should choose the language and platform branch that we have used to develop our web application. Here we have chosen “Go” and the proper platform branch (Screenshot 6).
 
![alt text](6.PNG "Screenshot 6")
<div align="center">Screenshot 6.</div>

For changing instance types, click on “Configure more options” and then we have to select “Edit” in Capacity box (Screenshot 7).
 
![alt text](7.PNG "Screenshot 7")
<div align="center">Screenshot 7.</div>


Here, we can choose between different instance types (Screenshot 8).
 
![alt text](8.PNG "Screenshot 8")
<div align="center">Screenshot 8.</div>

Then we save the configurations and create the environment (Screenshot 9). It will take about 10 minutes.
 
![alt text](9.PNG "Screenshot 9")
<div align="center">Screenshot 9.</div>

Now we can upload and deploy our web application (Screenshot 10).
 
![alt text](10.PNG "Screenshot 10")
<div align="center">Screenshot 10.</div>



# Distributed Load Testing on AWS

This AWS solution integrates various AWS services to simulate different workloads and determines your application’s behavior. That helps in identifying the bottlenecks before releasing the application. 

The following diagram shows the architecture of this solution:


![alt text](https://d1.awsstatic.com/Solutions/Solutions%20Category%20Template%20Draft/Solution%20Architecture%20Diagrams/distributed-load-testing-on-aws-architecture.f4325edc7552df2a3977d67c491b330819e52e9f.png "Title")
