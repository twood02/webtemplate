---
layout: page
title:  "The Curious Mr.X" - analyzing malicious network reconnaissance by wireshark
permalink: /wiki/WireSharkPuzzle/
---

*by:* Henian Wang and Guodong Xie

In this project, we will solve a puzzle, "The Curious Mr.X", by using wireshark. 

---

Mr. X, a notorious Mexico hacker, illegally gained access to Arctic Nuclear Fusion Research Facility(ANFRF)’s lab subnet over the 
interwebs. Unfortunately, his network reconnaissance was captured traces and all activities were stored in a packet. We will 
analyze the packet within WireShark.

##Install Wireshark
As a powerful network analysis tool, Wireshark can capture network activities. You can [Download & Install](https://www.wireshark.org/#download) it from the Wireshark homepage. Or you can simply apt-get to install wireshark on Cloud9.

```
sudo apt-get install tshark
```
After opening Wireshark, you can use it observe live traffic. Click the top left icon (looks like a shark fin) to start capturing packets. Then you will get information on windows like this:

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/1.jpg)

Wireshark captures lots of information in network traffic: source ip, destination ip, protocol, info, etc. Any activities in Wireshark, legal or illegal, could be analyzed by some methods. This article will give an example for analyzing a malicious hacking activity. 

##Import packet file(.pcap)
Also, Wireshark provides tools to analyze packets stored in a trace file in '.pcap' format. Clike 'File -> Open' to import a pcap file. 

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/2.jpg)

WOW! Overwhelming information caputured by ANFRF here! How could I find you MR.X?

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/3.jpg)

##Locate Mr.X: What was the IP address of Mr. X’s scanner?


##For the FIRST port scan that Mr. X conducted, what type of port scan was it?

##What were the IP addresses of the targets Mr. X discovered?

##What did he find on the Apple System?

##What did he find on the Windows System?

##What TCP ports were open on the Windows system? 

##What was the name of the tool Mr. X used to port scan?

##What did Mr.X see by the scanner?

##A conclusion about the puzzle
