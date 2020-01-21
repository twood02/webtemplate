---
layout: page
title:  Assignment 1 Reliable UDP
permalink: /assignments/reliable-udp
---

# Assignment 1: Reliable UDP


In this project, you will build a simple reliable transport protocol known as reliable udp (rudp). Your protocol must provide in-order, reliable delivery of UDP datagrams, and must do so in the presence of packet loss, delay, corruption, duplication, and re-ordering.

There are a variety of ways to ensure a message is reliably delivered from a sender to a receiver. We will provide you with a reference implementation of a receiver (which you must use) that returns a cumulative ACK whenever it receives a data packet. This is further explained with an example later. Your job is to implement a sender that, when sending packets to this receiver, achieves reliable delivery. For extra credit, you may choose to implement one of several performance improvements which are described below.

# The Reliable UDP Protocol 
Our simple protocol has four message types: `start`, `end`, `data`, and `ack`. `start`, `end`, and
`data` messages all follow the same general format:
```
start|<sequence number>|<data>|<checksum>
data|<sequence number>|<data>|<checksum>
end|<sequence number>|<data>|<checksum>
```
To initiate a connection, send a start message. The receiver will use the sequence number provided as the initial sequence number for all packets in that connection. After sending the start message, send additional packets in the same connection using the data message type, adjusting the sequence number appropriately. Unsurprisingly, the last data in a connection should be transmitted with the end message type to signal the receiver that the connection is complete. Your sender should accept acknowledgements from the receiver in the format:

```
ack|<sequence number>|<checksum>
```

An important limitation is the maximum size of your packets. The UDP protocol has an 8 byte header, and the IP protocol underneath it has a header of ~20 bytes. Because we will be using Ethernet networks, which have a maximum frame size of 1500 bytes, this leaves 1472 bytes for your entire packet (message type, sequence number, data, and checksum). 


The angle brackets ("<" and ">") are not part of the protocol. However, you should ensure that there are no extra spaces between your delimiters ("\|" character) and the fields of your packet. For specific formatting details, see the sample code provided.

# The Receiver Specification

We will provide a simple receiver for you; the reference implementation we provide will also be used for grading, so make sure that your sender is compatible with it. The receiver responds to data packets with cumulative acknowledgements. Upon receiving a message of type `start`, `data`, or `end`, the receiver generates an ack message with the sequence number it expects to receive next, which is the lowest sequence number not yet received. In other words, if it expects a packet of sequence number N, the following two scenarios may occur 
1. If it receives a packet with sequence number not equal to N, it will send “ack\|N”. 
2. If it receives a packet with sequence number N, it will check for the highest sequence number (say M) of the in-order packets it has already received and send “ack\|M+1”. For example, if it has already received packets N+1 and N+2 (i.e. M = N+2), but no others past N+2, then it will send “ack\|N+3”. 

Let us illustrate this with an example. Suppose packets 0, 1, and 2 are sent, but packet 1 is lost before reaching the receiver. The receiver will send “ack\|1” upon receiving packet 0, and then “ack\|1” again upon receiving packet 2. As soon as the receiver receives packet 1 (due to retransmission from the sender), it will send “ack\|3” (as it already has received, and upon receiving this acknowledgement the sender can assume all three packets were successfully received.

If the next expected packet is N, the receiver will drop all packets with sequence number greater than N+4; that is, the receiver operates with a window of five packets, and drops all packets that fall outside of that range. When the next unexpected packet is N+1 (due to N arriving), then the receiver will accept packet N+5. 

You can assume that once a packet has been acknowledged by the sender, it has been properly received. The receiver has a default timeout of 10 seconds; it will automatically close any connections for which it does not receive packets for that duration.

# The Sender Specification

The sender should read an input file and transmit it to a specified receiver using UDP sockets. It should split the input file into appropriately sized chunks of data, specify an initial sequence number for the connection, and append a checksum to each packet. The sequence number should increment by one for each additional packet in a connection. Functions for generating and validating packet checksums will be provided for you (see Checksum.py).

Your sender must implement a reliable transport algorithm (such as sliding window). The receiver’s window size is five packets, and it will ignore more than this. Your sender must be able to accept ack packets from the receiver. Any ack packets with an invalid checksum should be ignored.

Your sender should provide reliable service under the following network conditions:

- Loss: arbitrary levels; you should be able to handle periods of 100% packet loss.
- Corruption: arbitrary types and frequency.
- Re-ordering: may arrive in any order, and
- Duplication: you could see a packet any number of times.
- Delay: packets may be delayed indefinitely (but in practice, generally not more than 10s).

Your sender should be invoked with the following command:
```
python Sender.py -f <input file> -a <destination adddress> -p <port>
```

#### Some final notes about the sender:
- The sender should implement a 500ms retransmission timer to automatically retransmit packets that were never acknowledged (potentially due to ack packets being lost). We do not expect you to use an adaptive timeout (though this is a bells and whistles option).
- Your sender should support a window size of 5 packets (i.e., 5 unacknowledged packets).
- Your sender should be able to handle arbitrary message data (i.e., it should be able to send an image file just as easily as a text file). If no input file is provided, your sender should read input from STDIN.
- Any packets received with an invalid checksum should be ignored.
- Your sender should be written in Python3 or another language if instructed.
- Your sender MUST NOT produce console output during normal execution; Python exception messages are ok, but try to avoid having your program crash in the first place. 

We will evaluate your sender on correctness, time of completion for a transfer, and number of packets sent (and re-sent). Transfer completion time and number of packets used in a transfer will be measured against our own reference implementation of a sliding-window based sender. 

# Hints and Tips
To begin with, just focus on the simple case where nothing bad ever happens to your packets. After you have that case working, you can consider how to handle packet loss. It may help to build your way up to a full-fledged reliable sender. The simplest reliable transport mechanism is “Stop-And-Go”, in which the sender transmits a single packet and waits for the receiver to acknowledge receiving it before transmitting more. You could start by building a “Stop-And-Go” sender, and extending that for the full sliding window based sender. 

# Bells and Whistles (Extra Credit)
Some of these may require modifying the provided receiver implementation; if you choose to do
one, please make sure to provide your receiver implementation with your submission.


_*Variable size sliding window*_: For networks where packet loss, corruption, delay, and reordering are minimal, a large window size will obtain higher performance. A large window on a lossy network, however, will lead to a large amount of overhead due to retransmissions. Modify your sender to dynamically adjust its window size based on network conditions.


_*Selective Acknowledgements*_: If the sender knows exactly what packets the receiver has received, it can retransmit only the missing packets (rather than naively retransmitting the last N packets). Modify the receiver to provide explicit acknowledgement of which packets have been received, and add support for selective retransmission to your sender.


_*Accounting for variable round-trip times*_: How long you should wait to retransmit an unacknowledged packet is tightly related to the time it takes a packet to travel between the sender and the receiver. If you know the round trip time for a packet to reach your receiver is 20ms, waiting 500ms to retransmit is inefficient: you can be quite sure that a packet has been lost if you haven’t heard back from the receiver after 20ms, give or take a few milliseconds. Modify your sender to determine the round trip time between the sender and the receiver, and adjust your retransmission timeout appropriately.


_*Bi-directional transfer_*: While our protocol as defined only support uni-directional data transfer, it could be modified to allow bi-directional transfer (i.e., both ends of the connection could send and receive simultaneously). Implement this functionality, modifying both the sender and receiver as necessary. Be sure to provide a description of your updated protocol in your README.txt file.


_*I’m a superhacker*_: Identify a problem with a widely deployed modern TCP implementation, then
augment `rudp` to address that problem and outperform TCP. Describe and implement your improvement, and then benchmark your solution against the standard TCP.

# README
You must also supply a README.txt file along with your solution. This should contain
(1) You (and your partner’s) names
(2) What challenges did you face in implementing your sender?
(3*) Name the extra credit options you implemented; describe what they do and how they work.


# What To Turn In
Github classroom will capture your latest commit automatically on the due date. Make sure your submission is up to date in your repo's master branch by the due date.

# Setting up your github repo
To get access to the starter code please click on the provided link. This will prompt you to provide a team name which will be used to create your repo. You can add collaborators to your repo to work in groups. Repository code is [here](https://classroom.github.com/g/HQKgnlL9). 