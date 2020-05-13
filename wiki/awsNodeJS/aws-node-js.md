---
layout: page
title: A step by step guide to deploy a NodeJS application on AWS
permalink: /wiki/nodejsaws
---

*by:* Benjamin DeVierno

A step by step guide on setting up a Node.js app on an Ubuntu AWS EC2 instance.

---

<link type="text/css" rel="stylesheet" href="/assets/css/lightslider.min.css" />
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
<script src="/assets/js/ligrun.sh
htslider.min.js"></script>

A while ago I created a simple NodeJS web application that gives users a unique quote from the hit show ‘The Simpsons’. After many months of deliberation, I finally decided on deploying the app on Amazon’s AWS (Amazon Web Services). AWS provides serverless computing solutions to allow developers to create modern applications at lower cost of ownership. Some of the benefits include smaller overhead costs from managing servers, reliability and scalability. I ultimately chose AWS due to its ease of scalability, which will come in handy when inevitably this app goes viral. After following the steps in this article you should be able to make your very own app that is running on EC2 instance, stored on an S3 bucket, load distributed with elastic load balancing and cached with AWS’s Cloud Distribution Network Cloudfront. Phew…that sounds like a lot but if you follow along you should have your application up and running in no time at all.

I have provided a sample app that you may use to follow along to if you wish. The final product of which can be found [here.](http://simpsons-2003060048.us-east-1.elb.amazonaws.com/) [With cloudfront](https://d2atvopqlkjqa5.cloudfront.net)

## Amazon EC2.

Amazon EC2 or Elastic Cloud Computing allows users to rent virtual computers on which to run computer applications. 
Setting up our own EC2 instance will be the first step in deploying your application. First, login into your AWS account and select Cloud 9 then 'Create Environment'. 

<img src="/wiki/awsNodeJS/steps/step3.png">

I have chosen an ubuntu image provided by AWS as on OS. I have done so as the Ubuntu image that AWS provides comes with Node.js as well as node version manager pre-installed. However, you could easily select any of the other Amazon’s instances of your choosing, install the necessary dependencies and SSH into them. 

The Cloud9 default hibernation settings will stop your EC2 instance after 30 minutes. Since we will want to keep our site running you can update this setting to ‘never’. You may also leave this setting as it is and update this later on if you wish under the preferences tab within your Cloud9 IDE.

## Amazon S3.

Before moving on to the following steps, you may want to create an S3 bucket. S3 or Simple Storage Service is a service provided by AWS that allows developers to store and retrieve data. It allows developers access to highly scalable and inexpensive data storage infrastructure. It also allows static website hosting of HTML, CSS and Javascript. You can then subsequently move data from your S3 bucket to your EC2 instance using the instructions in the following [link](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AmazonS3.html). 


## Running our NodeJS application.

For the purpose of this demo we will be cloning our files directly from git.
Within your cloud9 IDE, clone this repository. For example:
*git clone https://github.com/bdevierno1/ExpressDemoACM.git*

Then run the following commands:
*cd acmapp*
*npm install*
*node app.js*

The app will now run with the specified port number you have provided in your script. Within the example application I have provided it will run at port 3000. Return to your AWS management console and view your running EC2 instances. Select the corresponding EC2 instance you just created.
Under description will be able to view the Public DNS associated with your EC2. 

<img src="/wiki/awsNodeJS/steps/step4.png">

## Updating public accessibility of our ports.

Unfortunately, accessing our application at the moment will not work as the port number 3000 is not currently accessible to the public. This is where security groups are needed. 

Select Security Groups then edit inbound rules.

<img src="/wiki/awsNodeJS/steps/step5.png">

Add your own custom TCP rule with a port range 3000 and source (0.0.0.0/0).

Verify that your site is functioning. For example:

**http://ec2-52-55-244-231.compute-1.amazonaws.com:3000/**

## Getting our app running in the background.

Currently, the app will only run as long as your terminal is open. We want our application to continue running in the background even if we close our terminal or hit Ctrl-C. Luckily, there is a process manager called pm2 that easily allows us to set up our application to run in the background.

Run the following commands.

*npm install -g pm2*
*pm2 start app.js*
*pm2 save*

Voila! We now have a functioning site.

## Load Balancing.

By now you have probably noticed that our URL contains the specified port number 3000. As you know it would be quite unusual for a websites URL to contain a port number. Therefore, the next step would be for us to redirect traffic from the standardized port 80 to 3000. Since we are running our application on EC2, AWS provides us with a fairly elegant solution to solve this problem.

Typically, a load balancer will sit in front of multiple EC2 instances. It will then manage the traffic coming in and distribute it evenly among the various instances. This helps prevent any one of out EC2 instances to become overloaded. Load balancing also helps us to achieve better fault tolerance by automatically balancing traffic among multiple targets. This ensures traffic is only routed to healthy targets. Currently, we will using the load balancer simply as a means to handle ports. However, following these steps will easily allow you to distribute traffic if you decide to run your application on multiple instances in the future.

## Steps to set up your load balancer.
Follow the following steps:

Select Create a Load Balancer then Classic Load balancer.

Set Load Balancer port to 80 and instance port to 3000.

Create a new Security Group. 

Under configure health check, set ping path to '/' as shown below.

<img src="/wiki/awsNodeJS/steps/step6.png">

The next step is to add your EC2 instances. If you select multiple instances, the load balancer will automatically split traffic between them.

Review and create your load balancer. Note you may not be able to visit the Public DNS provided by the load balancer right away. In this case try again in a few minutes. Once ready the load balancer should display a status of ‘InService’.

<img src="/wiki/awsNodeJS/steps/step7.png">

You will now be able to visit the [domain](http://simpsons-2003060048.us-east-1.elb.amazonaws.com/) listed under the description of your load balancer. 

## CDNs

The last step in the process which really is entirely optional, would be setting up a CDN (Cloudfront Distribution Network).  CDN’s allow us to cache data that our frequently requested. This will allow potential users to visit our sites at much lower latency. You can read more about CDNs in a prior blog post [here.](https://gwadvnet20.github.io/wiki/cloudfront) Since we will be expecting high amount of traffic to our application once it is fully deployed, setting up a CDN could be of great benefit.

## Setting up your CDN.

Create a CDN  web distribution and select the domain name of your application listed under elastic load balancers. Select create distribution.
<br/>
<img src="/wiki/awsNodeJS/steps/step9.png">
<br/>
AWS may take up to 15 to 20 minutes for your CDN to finish baking. 👨🏻‍🍳
Once completed you may view your new domain name under the general page of your deployed CDN.
<br/>
<img src="/wiki/awsNodeJS/steps/step10.png">
<br/>
# Congratulations! We have now deployed your first application onto AWS.
<br/>
#### Rererences

[S3](https://aws.amazon.com/s3/)
The creation of this blog was greatly helped by a similar post found [here.](https://medium.com/@nishankjaintdk/setting-up-a-node-js-app-on-a-linux-ami-on-an-aws-ec2-instance-with-nginx-59cbc1bcc68c)
