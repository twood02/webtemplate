---
layout: page
title:  Distributed Load Testing using AWS cloud services
permalink: /wiki/benchmark/
---

*by:* Samuel Farid and Hamid Reza


A short description of your post goes here.

---
# Creating an Environment and Deploying a Web Application with AWS Elastic Beanstalk

In this experiment we will deploy a sample web application on Amazon cloud servers using [AWS Elastic Beanstalk](https://aws.amazon.com/elasticbeanstalk/) service. As Amazon explains: “AWS Elastic Beanstalk is an easy-to-use service for deploying and scaling web applications and services developed with Java, .NET, PHP, Node.js, Python, Ruby, Go, and Docker on familiar servers such as Apache, Nginx, Passenger, and IIS”.

The sample application contains a welcome webpage with couple of links to other resources. The mentioned web application is deployed on four different machine types: <i>t2.micro</i>, <i>t2.large</i>, <i>r5.large</i>, and <i>m5.large</i>. Their specs are shown below:
 
|        |vCPU|ECU     |Memory(GiB)|Instance Storage (GB)|Linux/UNIX Usage|
|--------|:--:|:------:|:---------:|:-------------------:|:--------------:|
|t2.micro|	1 	|Variable|1         	|EBS Only	            |$0.0152 per Hour|
|t2.large|	2	 |Variable|8          |EBS Only	            |$0.1216 per Hour|
|m5.large|	2	 |10	     |8          |EBS Only	            |$0.1240 per Hour|
|r5.large|	2	 |10	     |16         |EBS Only	            |$0.1520 per Hour|


For deploying a web application with Elastic Beanstalk, first we have to login to the AWS console. Then search Elastic Beanstalk and open launch the service.


![alt text](1.PNG "Screenshot 1")


Then we must create an environment for our web application. In this section, we can configure our environment to test the application with different machine types.
 
 
![alt text](2.PNG "Screenshot 2")


Here we should select “Web server environment” for deploying a web application.
 
 
![alt text](3.PNG "Screenshot 3")


In the next section, we can choose a name and domain for our environment. For changing “.us-east-2.elasticbeanstalk.com”, we can change our region which will also affect the price of the instance that we will use.


![alt text](4.PNG "Screenshot 4")


![alt text](5.PNG "Screenshot 5")


In platform section we should choose the language and platform branch that we have used to develop our web application. Here we have chosen “Go” and the proper platform branch.
 
 
![alt text](6.PNG "Screenshot 6")


For changing instance types, click on “Configure more options” and then we have to select “Edit” in Capacity box.
 
 
![alt text](7.PNG "Screenshot 7")


Here, we can choose between different instance types.
 
 
![alt text](8.PNG "Screenshot 8")


Then we save the configurations and create the environment. It will take about 10 minutes.
 
 
![alt text](9.PNG "Screenshot 9")


Now we can upload and deploy our web application.
 
 
![alt text](10.PNG "Screenshot 10")



# Distributed Load Testing on AWS

This AWS solution integrates various AWS services to simulate different workloads and determines your application’s behavior. That helps in identifying the bottlenecks before releasing the application. In other words, how the application is going to perform in production and at scale.

This solution uses Amazon Elastic Container Service ([Amazon ECS](https://aws.amazon.com/ecs/)) to spin up containers that will create hundreds of connections to your end point.

The following diagram shows the architecture of this solution:


![alt text](https://d1.awsstatic.com/Solutions/Solutions%20Category%20Template%20Draft/Solution%20Architecture%20Diagrams/distributed-load-testing-on-aws-architecture.f4325edc7552df2a3977d67c491b330819e52e9f.png "AWS Distributed Load Testing architecture")


The architecture consists of frontend and backend. In the frontend we have the web console which is a UI that we can use to interact with the solution. The UI uses Amazon CloudFront service to allow the user to configure the tests. We also have load testing API that is used to create tests and view the status of the tests. 

The backend comprises of docker image pipeline and load testing engine. The solution uses Taurus which is an open source software and has a Docker image that allows generating hundreds and hundreds of concurrent connections to the end point. The [Amazon S3](https://aws.amazon.com/s3/) service is a simple storage service which is used to store that image in your account. After that [AWS CodePipeline](https://aws.amazon.com/codepipeline/) and [AWS CodeBuild](https://aws.amazon.com/codebuild/) services are used to build that image and register it with [Amazon ECR](https://aws.amazon.com/ecr/) service that makes it easy for developers to store, manage, and deploy Docker container images.

The testing itself runs in [AWS Fargate]( https://aws.amazon.com/fargate/) service which runs your containers on the Elastic Container Service handling the networking and the underlying infrastructure. The [AWS Lambda]( https://aws.amazon.com/lambda/) service that takes the requests from the API to  run it in AWS Fargate. It also stores the test template in Amazon S3 and the information we are collecting in [Amazon DynamoDB](https://aws.amazon.com/dynamodb/). Then the [Amazon SQS]( https://aws.amazon.com/sqs/) service queues the tasks in AWS Fargate so that we can start spinning up the containers.  


# How to use Distributed Load Testing solution?

Launch the solution in the AWS console from this [link](https://aws.amazon.com/solutions/distributed-load-testing-on-aws/).


![alt text](website.png "AWS Distributed Load Testing website")


Keep the default settings and click on “Next” to create a stack.


![alt text](step1.png "Create stack")


Set the stack name, console administrator name, and console administrator email and click on “Next”.


![alt text](step2.png "Stack details")


Add tags as needed, then click on “Next”.


![alt text](step3.png "Stack options")


Review your configuration and click on “Create stack” if everything is correct.


![alt text](step4.png "Stack review")


You will be redirected to a page showing that creating the stack is in progress.


![alt text](step5.png "Stack create in progress")


This page will show that the stack creation is completed after few minutes.


![alt text](step6.png "Stack create complete")


Click on “Outputs” tab, then you will find the console’s link. This console will be used to create and view tests. Use the username and password that was sent to the console administrator email when you are asked to login.


![alt text](step7.png "Outputs tab")


# Creating a test:

After logging in to the console you will be able to create a task.


![alt text](step9.png "Console interface")


Enter the setting as required for your test and click on “Submit”.


![alt text](step10.png "Create a load test")


After that you will see the progress of the running task.


![alt text](step11.png "Running task")


Finally, the results of the task will be shown:


![alt text](step13.png "Running task")


# Same load for web application running on different instances benchmark
We use the previously created four instances each running the same web application for load testing with the same load. We set task count to 2, concurrency to 5, ramp up to 10s, and hold for 30s. The results are shown below:

![alt text](t2_micro.png "t2.micro")
<div align="center"><i>t2.micro</i> instance results</div>
<br/><br/>

![alt text](t2_large.png "t2.large")
<div align="center"><i>t2.large</i> instance results</div>
<br/><br/>

![alt text](m5_large.png "m5.large")
<div align="center"><i>m5.large</i> instance results</div>
<br/><br/>

![alt text](r5_large.png "r5.large")
<div align="center"><i>r5.large</i> instance results</div>
<br/><br/>

# Different loads for web application running on the same instance benchmark
In this test we will use <i>r5.large</i> instance and test its behavior across different loads. We will set the ramp up to 10s and hold for 30s. Each test we will change the task count and concurrency. The results are shown below:

![alt text](load1.png "load1")
<div align="center">Task count = 1 and Concurrency = 1</div>
<br/><br/>

![alt text](load2.png "load2")
<div align="center">Task count = 10 and Concurrency = 20</div>
<br/><br/>

![alt text](load3.png "load3")
<div align="center">Task count = 50 and Concurrency = 100</div>
<br/><br/>

![alt text](load4.png "load4")
<div align="center">Task count = 50 and Concurrency = 150</div>
<br/><br/>
