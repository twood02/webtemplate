---
layout: page
title:  How does In-Flight WiFi really work?
permalink: /wiki/inflightWifiBlog/
---

*by:* Katie Bramlett and Sreya Nalla

<br/>

**How exactly does wifi work on your airplane 35,000+ feet in the air?
Here's a look at exactly how In-Flight WiFi keeps you connected to the world below.**

We'll discuss the basics of in-flight wifi, ground-based and satellite operating systems, and how moving from one atmosphere to another can affect the network performance.

---

Have you ever stopped to consider exactly how an e-mail or streaming video gets into your smartphone or tablet when you're cruising thousands of feet in the air? That's WiFi in action. When any two devices or machines need to communicate with each other, they need certain standards and protocols to be defined such that they are now able able to communicate. WiFi is that set of standards that defines communication for wireless computer networks. WiFi functions by using radio frequencies to send signals between devices. This radio frequency is different from that of your average walky talkie, car radio or weather radio, as WiFi transmits and receives data in the Gigahertz range.

### So, what does this mean?
To break it down a little further, Hertz is only a measure of frequency. Let's say you were sititng by the ocean watching the wave roll in. If you were to count the number of seconds that passes between each wave, you'd be measuring the frequency of each wave. So, if 1 second were to pass between two waves, the wave frequency would be 1 Hz, or one cycle per second. In order to receive any information that is being send between these waves, the receiver needs to be set on a frequency that can read this incoming information at the correct rate. 

For WiFi, this frequency happens to be 2.5 Ghz or 5 Ghz (similar to that of a microwave!). The primary difference between these two frequencies lies in the range and bandwidth, which is the captacity of speed the network can provide. The 2.4Ghz frequency provides coverage over a larger range, but is only able to transmit data at lower speeds (smaller bandwidth), while the 5Ghz frequency provides coverage over a smaller range, but is able to transmit data at much faster speeds due to a larger bandwidth. 

### Does In-Flight WiFi work the same way?
Most of the technology behind the miracle of airplane wifi largely resembles that of wireless internet on the ground. To bring connectivity services to each passenger, planes first need to establish a connection to an Internet source. This is done using either an air-to-ground or satellite network. Once a connection is established using one of those two sources, connectivity can be provided to passendgers via WiFi hotspots. Airplanes use many of the same towers and satellites that deliver data to our smartphones, connecting to towers on the ground, or to satellites, or to both.

| Air-to-Ground Networks | Satellite Networks |
| ---------------------- | ------------------ |
| Airplanes use Air-to-Ground (ATG) networks to establish connectivity when traveling directly over land. The airplane should automatically connect to signals sent to its antennas, located on the bottom of the fuselage, from the nearest tower on the ground. This should allow for no interruptions if the airplane is flying only over populated regions of land. However, when travelling over remote terrain or large bodies of water, connectivity can often be a problem with this type of network. <img src="./ATGNetworkDiagram.png" width="450" height="250"/>|Airplanes can utilize satellite technology for connectivity to WiFi when flying anywhere around the globe, whether they are flying over land or water. Using satellite antennas on the top of fuselage, the airplane is able to communicate with the nearest satellite in orbit. Since the satellite is continuously in connection with a tower on the ground, there is a constant supply of network connectivity to the airplane. <img src="./SatelliteNetworkDiagram.png" width="450" height="250"/> |
### How do we measure the performance?
In order to measure the performance of In-Flight WiFi, or any WiFi network connection, we have a variety of performance metrics that we can use. Below, we will go into more detail about some of these metrics.
#### Latency vs. Throughput
| Latency | Throughput |
| ------- | ---------- |
| Time to process a request i.e. time from issuing a request to getting a response. | Units of work completed per unit of time. |
| Affected by: network distance, processor speed on server, queueing delays. | Affected by: bandwidth of network, processor speed on server, number of processors on server. |

Latency and throughput have a unique relationship. As the workload, or throughput, increases towards maximum capacity, the latency exponentially increases. If the workload is low, the latency will be low, as well. But, when the workload becomes heavier, the latency will rise greatly, until it theoretically reaches infinity -- meaning that no data can be transmitted.

**The Affects of Workload on Latency:**
<img src="./latency_vs_throughput.png" width="450" height="300"/>

#### Bandwidth
A common misconception is that bandwidth is a measure of speed. Rather, bandwith is synonymous with capacity. It is the maximum amount of data that can travel through a link or network, measured in bits per second (bps). Bandwidth describes the *theoretical* data transfer rate that should occur, whereas the throughput describes the *real world* data transfer rate of the network.

Often, we use the water pipe analogy to explain how bandwith works. The wider the pipe (larger diameter), the more water that can flow through. Simililarly, the more bandwidth a data connection has, the more data it can send and receive at the time. That being said, bandwidth is the capacity for speed. 

**Why do we care about Bandwidth? Why bother measuring it?**
- getting money's worth? 
- is there something that can be fixed?
- are airplanes able to foster maximum capacity? is what we pay what we actually get?
    -- how can we tell? calculate bandwith
    -- ways to calculate bandwith
- how can planes begin to maximuze bandwith 


#### How can we calculate each of these?
**Latency:**
latency = RTT + queueing time + service time
*RTT = round-trip time*
*queueing delay = 1 / (capacity - load)*

**Throughput:**
Max Capacity of a Web Server:
max capacity = 1 / service time
(for a single processor system)

**Bandwidth:**
???
#### How can we analyze this information and use it to our benefit?
After we have analyzed all of the information on In-Flight Wifi and its performance, it is important to use this information to better the technology.

### Are there ways to improve performance?
Now, let's take a further look into how to improve performance. It is often found that In-Flight WiFi is slow or problematic. There are many ways that companies can try to improve these issues.
#### What is currently available?
In today's world, technology is developing rapidly and the number of devices using Wi-Fi is higher than ever before. The first In-Flight WiFi service was launched in 2008, and at the time, a 3 Mbps connection was enough for a small number of laptops. But now, most likely every passenger will have at least one device. They may be trying to stream video, listen to music, use mobile applications, and connect to websites. Therefore, there is a much greater strain with the resources available.

Today, a satellite connection offers roughly 12 Mbps. However, satellites are expensive to maintain and upgrade, which contributes to delays in technology upgrades -- meaning it causes In-Flight WiFi to lag behind.

As a whole, In-Flight WiFi is a very expensive technology. Between satellite, antenna, engineering, and maintenance costs, there are a lot of expenses to account for. Some airline companies offer free WiFi onboard, while others charge for the service.

All of these reasons have contributed to slow and problematic In-Flight WiFi services.

#### What research is being done?
Many people question if In-Flight WiFi will get better and faster in the future. On the bright side, there are many companies that are at the forefront of research into improving the performance of In-Flight WiFi. Here is an overview of the work that is being done.

???

#### 2018 CNN Top 5 Rankings of Most Tech-Friendly Airlines:
**Best In-Flight WiFi Services**
1. Qatar Airways
2. Emirates
3. Delta Air Lines
4. British Airways
5. JetBlue Airways

BLOG OVERVIEW

In-Flight Wifi Basics
0. What even is WiFi?
    - we all talk about it, but how exactly does it work?
    - how are we able todo what what we do?

1. What is In-Flight Wifi?
    - Two operating system involved
        - ground-based 
        - satellite-based
    - how easy is it to access each of these "types" of in-flight wifi
        - comparing performance of each type
        - latency and throughput
2. What are performance metrics? 
    - latency (include graphs??)
    - throughput
    - bandwidth
    - how do we calculate these things
        - ways to test bandwidth on planes?
        - how can we analyze this information and use it to our benefit?
    - why are they important in terms of in-flight wifi
        - does this explain why wifi is generally slow on planes?
        - does it affect speed?
        - how does bandwidth affect type of information transmitted and vice versa?
3. Are there ways to improve performance?
    - what is currently available
    - what research is being done
    - what steps can we take to assist?
    - what companies are at the forefront?




