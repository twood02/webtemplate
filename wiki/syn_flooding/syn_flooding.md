---
layout: page
title:  SYN Flooding
permalink: /wiki/syn_flooding/
---

*by:* Dennis and Deep


Overview of SYN flooding detection and prevention systems.

---

# Setting alerts for TCP SYN Flood Attacks

## Introduction

A SYN flood attack is a denial of service attack that uses the TCP handshake procedure to attack the victim.

However, when a server receives a SYN request, it opens a new connection, sends a SYN-ACK and waits for an ACK to complete the handshake.

In a SYN flood attack, the attacker bombards the victim with endless SYN requests but never responds to the SYN-ACK packets. As a result, the server is overwhelmed and exhausts all system resources on open SYN connections.

## Simulating a SYN flood attack in Python

The core of the script builds the spoofed packet by hand using the Python module Scapy.

```Python
# Build the packet
src_ip = src_net + str(src_host)
network_layer = scapy.IP(src=src_ip, dst=dst_ip)
transport_layer = scapy.TCP(sport=src_port, dport=dst_port, flags="S")
```

Post that it can simply send these packets to the victim continuously.

## How does the server actually react

We tested with both Apache2 and NGINX web servers on an Ubuntu 18 machine. The results were quite surprising actually. Both these servers barely reacted when subjected to these SYN attacks.

There is a possibility that the kernel also plays a role in protecting against such SYN attacks. Or maybe, Python script running on a single system is too slow.

However, this meant one cannot rely on a spike in system resource usage to detect such attacks.

## Detecting the attack using command line tools

When the server receives these SYN packets it opens a new connection and marks it as `SYN_RECV`.

The first simple tool to use here is `netstat` which provides information about network connections on the system.

The following command provides a view of how many such connections are there in the system.

```bash
netstat -antp | awk '{print $6}' | grep SYN_RECV | wc -l
```

We noticed between 118 and 128 connections of this type open on the system during the attack.

The next obvious command to use is `tcpdump`

```bash
sudo tcpdump -i en0
```

However, this is an extremely crowded output to be useful while under attack.

The next command to use was `tshark`. Obviously, `wireshark` can also be used.

```bash
sudo tshark -i enp5s0f0 -Y "tcp.flags.syn == 1 and tcp.flags.ack == 0"
```

```text
   64 6.627453039 192.168.250.1 ? 192.168.1.134 TCP 60 1039 ? 80 [SYN] Seq=0 Win=8192 Len=0
   66 6.663454514 192.168.250.1 ? 192.168.1.134 TCP 60 1040 ? 80 [SYN] Seq=0 Win=8192 Len=0
   68 6.703615407 192.168.250.1 ? 192.168.1.134 TCP 60 1041 ? 80 [SYN] Seq=0 Win=8192 Len=0
   70 6.739260198 192.168.250.1 ? 192.168.1.134 TCP 60 1042 ? 80 [SYN] Seq=0 Win=8192 Len=0
   72 6.775267262 192.168.250.1 ? 192.168.1.134 TCP 60 1043 ? 80 [SYN] Seq=0 Win=8192 Len=0
   74 6.811260895 192.168.250.1 ? 192.168.1.134 TCP 60 1044 ? 80 [SYN] Seq=0 Win=8192 Len=0
   88 6.847259570 192.168.250.1 ? 192.168.1.134 TCP 60 1045 ? 80 [SYN] Seq=0 Win=8192 Len=0
   90 6.883271767 192.168.250.1 ? 192.168.1.134 TCP 60 1046 ? 80 [SYN] Seq=0 Win=8192 Len=0
   92 6.919255628 192.168.250.1 ? 192.168.1.134 TCP 60 1047 ? 80 [SYN] Seq=0 Win=8192 Len=0
   94 6.955270465 192.168.250.1 ? 192.168.1.134 TCP 60 1048 ? 80 [SYN] Seq=0 Win=8192 Len=0
   96 6.991282014 192.168.250.1 ? 192.168.1.134 TCP 60 1049 ? 80 [SYN] Seq=0 Win=8192 Len=0
   98 7.027273912 192.168.250.1 ? 192.168.1.134 TCP 60 1050 ? 80 [SYN] Seq=0 Win=8192 Len=0
  101 7.063603495 192.168.250.1 ? 192.168.1.134 TCP 60 1051 ? 80 [SYN] Seq=0 Win=8192 Len=0
```

Here one can clearly see that SYN packets are continuously arriving.

## Automated alert system

Then is it possible to create a simple script to automatically alert an administrator when the system is under attack?

With the help of two Python libraries, `psutil` and `pyshark`, we demonstrate the basic idea of such a script below.

### Part A

Get all network connections and get a count of all open connections.

```Python
conns = psutil.net_connections() # get all network connections
for item in conns:
    if item[5] == 'SYN_RECV': # grep for SYN_RECV
        i += 1 # increase count if found
```

How many open connections are normal might be somewhat system dependent but a sliding average could be maintained to detect attacks.

### Part B

The `pyshark` tool is extremely capable and versatile. The code snippet here is only a meager indication of its capabilities.

A capture on a live interface can be started as follows:

```Python
capture = pyshark.LiveCapture(interface="en0", display_filter="tcp.flags.syn == 1 and tcp.flags.ack == 0")
capture.set_debug(set_to=True) # required for live capture
```

Sniffing packets can be done in batches:

```Python
capture.sniff(timeout=5)
for i in range(5):
	print(capture[i])
```

Or packets can be sniffed continuously:

```Python
for pkt in capture.sniff_continuously(packet_count=5):
    print("New Syn pkt") # One can imagine an email being sent out here
    print(pkt) # This will print all metadata
```

Finally, any user defined function, `syn_flooding_detection`, can also be applied on each packet:

```Python
capture.apply_on_packets(syn_flooding_detection, timeout=100)
```
## SYN Cookies

The main issue with SYN flooding is due to how the server behaves when it receives a SYN packet. Information about the SYN packet is allocated and then added to a queue. At the same time, a SYN+ACK packet is sent. If the attacker doesn't reply after a period of time, then the memory is freed. It may look like this in a two-thread protocol implementation:
```Python
for pkt in capture.sniff_continuously(packet_count=5):
        add_to_syn_queue(pkt) # Stores pkt information in the syn queue
        send_syn_ack(pkt) # Sends a SYN+ACK packet
```

In a seperate thread, the SYN packets are placed on a timeout and are waiting to receive the ACK from the packet.

```Python
while True:
        for pkt in syn_queue:
                timeout = time.time() + 60
                while time.time() < timeout
                        if pkt in ack_queue:
                                accept_connection(pkt)
                                delete_from_syn_queue(pkt)
                                break
               
                delete_from_syn_queue(pkt)
```

Why do we need to do so much work for the SYN packets in the first place? This was a question raised by Daniel J. Bernstein in 1996. His solution was elegant. Instead of maintaining a SYN queue, encrypt some portions of that SYN packet, then wait for a valid ACK response. This encrypted part of the packet is called the SYN cookie, and it has the ability to validate a received ACK packet by computing a cryptographic function on portions of its data. The replied SYN + ACK packet contains the computed SYN cookie. A valid client will use this to generate a cookie that is sent back to the server. The server will then check to see if the ACK's cookie matches a globally stored key. This is a purely mathematical operation and does not require any memory usage - just extra CPU cycles to perform the hash. There's no more concept of a "half open connection" anymore since there is no SYN queue. Think of it as offloading some of the work to the client. The server says: "Thanks for your SYN packet, now send me a valid ACK packet to prove that you want to connect!" The SYN flood attacker is out of luck, because they are using spoofed IP addresses so their won't be a client to craft a genuine ACK reply! And furthermore, overloading the server is now way more difficult since there's no more information being stored about the SYN packet - the server is just going to be computing a cryptopgraphic function on the receiving ACK packet to verify and establish the connection.
Changing up our code from before, this is all that's needed by the TCP server:

```Python
for pkt in capture.sniff_continuously(packet_count=5):
        pkt = encode_syn_cookie(pkt) # Send the encoded key within the packet. The client will now have to reply with a genuine ACK packet.
        send_syn_ack(pkt) # Sends a SYN+ACK packet storing the cryptographically computed key encoded in the packet
```

Observer how much simpler this section becomes:

```Python
while True:
        for pkt in ack_queue:
                if secret_function(pkt) is valid:
                        accept_connect(pkt)
              
```

This is cool, because now we don't have to worry about timeouts and queue deletions. 
This was a simplified introduction to SYN cookies. They are more complicated in practice, and a good place to look for more in depth code would be the syncookies.c file in the linux kernel. Check out this reference, as it lists the email chain in which the main ideas of SYN cookies were hashed out: https://cr.yp.to/syncookies/archive 

## Summary

1. SYN flood attacks are relatively easy to simulate.
2. They are also easy to detect at the network level.
3. They cannot be detected by looking at system utilization.
4. There might be some system and software level protection against such attacks.
5. Given the Python ecosystem today, it is reasonably straightforward to build scripts that will enable users of victim systems to react quickly to such attacks.

## References

https://dubell.io/creating-syn-flood-attacks-with-python/

http://www.firewall.cx/general-topics-reviews/network-protocol-analyzers/1224-performing-tcp-syn-flood-attack-and-detecting-it-with-wireshark.html
