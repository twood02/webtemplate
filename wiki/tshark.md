---
layout: page
title: Tshark
permalink: /wiki/tshark
---
<link type="text/css" rel="stylesheet" href="/assets/css/lightslider.min.css" />
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
<script src="/assets/js/lightslider.min.js"></script>



While working with networked applications and distributed systems, it sometimes becomes to necessary to take a look at the network traffic for debugging. There are many tools that you can use to inspect network traffic but on this page we will discuss **[tshark](https://www.wireshark.org/docs/wsug_html_chunked/AppToolstshark.html)**. 

**tshark** is a CLI included with the wireshark package that allows you to apply filters and view incoming and outgoing traffic on your machine. You can also install tshark as a standalone tool using your package manager on ubuntu.

### Installing tshark

On a cloud9 environment you can simply use apt-get to install tshark:

```
sudo apt-get install tshark
```


# Using Tshark


#### List interfaces

To list the interfaces available for monitoring on your server, you can use the `-D` flag.

```
sudo tshark -D 
```

This produces the following output:

```
Running as user "root" and group "root". This could be dangerous.
1. eth0
2. any
3. lo (Loopback)
4. docker0
5. nflog
6. nfqueue
7. ciscodump (Cisco remote capture)
8. randpkt (Random packet generator)
9. sshdump (SSH remote capture)
10. udpdump (UDP Listener remote capture)
```

#### Capture traffic

To Capture on an interface, you can then specify an interface with the `-i` flag:

```
sudo tshark -i eth0
```

This produces a lot of output that is hard to follow and probably not useful for you. Instead, you can filter based on many filter flags such as protocol, port, ip address and many more. We'll cover a few below.

#### Capture traffic based on port

You may find that you know the port of a service you want to inspect. The `-Y` flag allows you to specify a filter for packets. If you want to view only traffic going to a specific port then you can use `tcp.dstport` like so:

```
tshark -Y "tcp.dstport == 9800" 
```

You can do the same for traffic originating from a certain port with `tcp.srcport` as well as for a udp port using: `udp.port`.



#### Capture traffic based on ip address


If you want to view only traffic going to a specific ip address then you can use `ip.dst` like so:

```
tshark -Y "ip.dst == 18.206.216.172"

```

Similar results can be achieved with `ip.src`.

#### Show the contents of packets

It can be helpful to see the contents of packets for diagnosing issues. To do so we can use the `-x` flag to show the hex contents. This can be done like so:

```
tshark -Y "tcp.dstport == 9800" -x
```

This will show you an output like so:

```
Capturing on 'Wi-Fi: en0'
0000  48 5d 36 18 68 61 e0 ac cb 5d ec cc 08 00 45 00   H]6.ha...]....E.
0010  00 40 00 00 40 00 40 06 8c fb c0 a8 01 9a 12 ce   .@..@.@.........
0020  d8 ac f8 da 26 48 06 86 30 ac 00 00 00 00 b0 02   ....&H..0.......
0030  ff ff 50 f2 00 00 02 04 05 b4 01 03 03 06 01 01   ..P.............
0040  08 0a 1a db c7 1c 00 00 00 00 04 02 00 00         ..............
```

#### Show detailed packet information

Sometimes you need to take an even deeper dive into your packets. You can do this by using the -V command. It will show full protocol infortmation for ethernet, ipv4 and tcp as well as any additional details. It is used like so:

```
tshark -Y "tcp.dstport == 9800" -V
```

You can also use the `-O` options to show the same output for a specific protocol like so:

```
tshark -Y "tcp.dstport == 9800" -O tcp
```

#### Filter based on request method

When working with api's and backend endpoints, you may want to filter packets based on http protocols. You can do so by using `http.request.method`. To find GET requests, you can do the following:

```
tshark -Y 'http.request.method == "GET"' -x
```