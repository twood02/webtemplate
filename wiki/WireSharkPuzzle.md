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

## Install Wireshark
As a powerful network analysis tool, Wireshark can capture network activities. You can [Download & Install](https://www.wireshark.org/#download) it from the Wireshark homepage. Or you can simply apt-get to install wireshark on Cloud9.

```
sudo apt-get install tshark
```
After opening Wireshark, you can use it observe live traffic. Click the top left icon (looks like a shark fin) to start capturing packets. Then you will get information on windows like this:

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/1.jpg)

Wireshark captures lots of information in network traffic: source ip, destination ip, protocol, info, etc. Any activities in Wireshark, legal or illegal, could be analyzed by some methods. This article will give an example for analyzing a malicious hacking activity. 

## Import packet file(.pcap)
Also, Wireshark provides tools to analyze packets stored in a trace file in '.pcap' format. Clike 'File -> Open' to import a pcap file. 

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/2.jpg)

WOW! Overwhelming information caputured by ANFRF here! How could I find you MR.X?

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/3.jpg)

## Locate Mr.X: What was the IP address of Mr. X’s scanner?
To locate the IP address of Mr. X's scannar, we need figure out who are busy at communicating in the provided evidence file. The flowing packets must contain evidence of port scan, but we also should cull some noisy packets. To obtain the required information: IP addresses, and figure out what are conversation between them, we can click 'statisics > conversations'.

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/4.jpg)

We can find the most active conversations are between 10.42.42.253 and 10.42.42.56, 10.42.42.50 and 10.42.42.25 respectively.

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/5.jpg)

We can view this to conclude that IP Address 10.42.42.253 initiated scanning IP addresses 10.42.42.56, 10.42.42.50, 10.42.42.25 by the following reasons:

'''
1. The packets from address 10.42.42.253 to other three addresses is always larger than from the opposite directions.

2. The three flows between 10.42.42.253 to other three respectively have the earliest timestamps.

3. The bytes size in packets from 10.42.42.253 to other three addresses is much bigger than from opposite directions.

'''

## For the FIRST port scan that Mr. X conducted, what type of port scan was it?

To determine which type of First port scan, we can backtrack the timestamp to find the first connection from attacker address 10.42.42.253: (By typing 'tcp.stream eq 0' on searchbox)

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/6.jpg)

We now know that the first scan ports are related to TCP flags, but which one is correct? Let's check the flags in the tcp stream. 

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/7.jpg)

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/8.jpg)

IP 10.42.42.253 sent SYN packets, and it is the signal of TCP SYN or TCP connect port scan, and this address has received packets with RST/ACK flags from the first victim 10.42.42.50. The 3-way handshake of TCP connect principle illustrates:

'''
TCP Connect:
	SYN >
	< SYN/ACK
	ACK >
	< RST+ACK
'''

The flags from frame 1 shows that the first port scan from attacker is TCP connect.

## What were the IP addresses of the targets Mr. X discovered?

Exploring the conversations in the provided evidence file, Mr.X's IP address scanned the following IP addresses:

'''
- 10.42.42.25
- 10.42.42.50
- 10.42.42.56
'''

## What did he find on the Apple System?

Apple System? How to recognize a operating system? The internet protocol suite may give us some ideas:

Internet protocol suites tell us the link layer is the lower level than network layer, and Ethernet is one of most representative protocol of the link layer. MAC address, identity number of devices, could help us determine the operating systems of victims.

We can search the three victim IP address to find their ethernet (MAC address):

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/9.jpg)

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/10.jpg)

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/11.jpg)

As the results shown, the three victim addresses have their corresponding Mac addresses. Wireshark also gave the device vendor information of Mac address:

## IP address         Mac address         Device Vendor 
-  10.42.42.25     00:16:cb:92:6e:dc      Apple Computer
-  10.42.42.50     70:5a:b6:51:d7:b2      COMPAL INFORMATION (KUNSHAN) CO., LTD.
-  10.42.42.56     00:26:22:cb:1e:79      COMPAL INFORMATION (KUNSHAN) CO., LTD. 
-  10.42.42.253    00:23:8b:82:1f:4a      Quanta Computer Inc.

Hence we will conclude that the Apple system Mr.X found has 10.42.42.25 for IP address and that his mac
address is: 00:16:cb:92:6e:dc.

## What did he find on the Windows System?

The other two IP addresses have same device vendor, but how can we determine which is Windows system? or both?

10.42.42.50 ... 10.42.42.56 It looks like IP-ID numbers grows regularily, so both of them are Window system?

Wait a second! back to the evidence file, we must find something on communication between them. As we all known, the Windows system had strict rule about firewalls that only open specific ports:

let's analyze 10.42.42.56 again on Wireshark:

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/12.jpg)

As we can see, all the conversations between 10.42.42.56 and 10.42.42.253 don't have any open tcp ports! Actually, the scan towards 10.42.42.56 is composed of tcp connct(), tcp null, udp, and so on. However, Wireshark shows no open ports for host 10.42.42.56.

For 10.42.42.50:

![image](http://github.com/itmyhome2013/readme_add_pic/raw/master/images/11.jpg)

On this picture, host 10.42.42.50 opens both of 135 and 139 tcp ports, and the two ports are common ports in the Window system. According to [GRC's Website](http://www.grc.com/port_135.htm):

'''
135/tcp:
"Microsoft's DCOM (Distributed, i.e. networked, COM) Service Control Manager
(also known as the RPC Endpoint Mapper) uses this port in a manner similar to
SUN's UNIX use of port 111. The SCM server running on the user's computer opens
port 135 and listens for incoming requests from clients wishing to locate the
ports where DCOM services can be found on that machine."

139/tcp:
"TCP NetBIOS connections are made over this port, usually with Windows machines
but also with any other system running Samba (SMB). These TCP connections form
'NetBIOS sessions' to support connection oriented file sharing activities."
'''

Hence, the Windows system Mr.X found is 10.42.42.50 for IP and 70:5a:b6:51:d7:b2 for Mac address. It opens tcp ports on that machine:

- 135/tcp
- 139/tcp

## What did we learned

In the process of solving the puzzle, we were gradually exciting to make hypothesis about what Mr.X did and make reverse-engineering on our hypothesis. It really need a lot of networking knowledge (especially TCP/IP) to solve our questions. 

Wireshark is a wonderful tool to learn packet and protocols. It really helps us analyze the whole process, and we, on the other hand were more familiar with it. The process we used filters and analyze info of packets is treassure.



