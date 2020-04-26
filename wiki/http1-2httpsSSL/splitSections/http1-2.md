## HTTP 1 vs HTTP 1.1 vs HTTP 2.0

In this article, I will let you know much more about the difference between HTTP1, HTTP 1.1 and HTTP 2. I will start with HTTP1 vs HTTP 1.1, and then HTTP 1.1 vs HTTP 2.

### Overview of HTTP
The HyperText Transfer Protocol (HTTP) is a protocol resides in the application layer. HTTP defines how the client and the server exchange the message, which means how the client requests a Web page and how the server responds the web page to the client. When the client clicks on the link, the browser sends the HTTP request to the server, the server receives the request and responses back to the client.

Generally, a web page contains several objects. The root object is an HTML file, and this HTML file may contain the URLs to other objects like JPEG, audio, video clip, etc. For example, the URL looks like this:

```bash
http://www.someschool.edu/someDepartment/picture.gif
```

www.someschool.edu means the hostname, /someDepartment/picture.gif is a pathname. HTTP uses TCP to set up the connection.

### HTTP 1.0 Intro
HTTP 1.0 needs to request each object one by one, for each request, HTTP has to set up the TCP connection. Thus, efficiency is pretty low. Let’s examine how does HTTP 1.0 work by the following example.

Suppose a web page contains a base HTML file and 10 JPEG images, suppose the URL of the base HTML is :

```bash
http://www.someSchool.edu/someDepartment/home.index
```

Step 1: The client set up a TCP connection with the server www.someSchool.edu on default port 80.

Step 2: The client requests the HTML base file by sending a request message to the server. The message contains the pathname : /someDepartment/home.index .

Step 3: The server receives the message, retrieves the object and encapsulate it into the HTTP message. Send the message back to the client.

Step 4: The server tells the TCP to close the connection. But the TCP will wait until the client receives the response.

Step 5: The client receives the object, the TCP terminates. When the client extracts the HTML base file, he finds it contains the URLs to other images.

Step 6: The client has to repeat the above 5 steps to request for each image.

### The main differences between HTTP 1.0 vs HTTP 1.1
1. Persistent connection.

To solve the low efficiency of the non-persistent TCP connection, HTTP 1.1 uses the persistent connection, which means requesting multiple objects within one TCP connection is allowed. HTTP 1.1 also introduces pipelining, which means these requests can be made back-to-back without waiting for the replies to pending requests.

2. Host header.

HTTP 1.0 does not contain the host in the header since HTTP 1.0 assumes the server binds only one host. However, a server can have multiple hosts (Multi-homed Web Servers) and they share one IP address. Thus, HTTP 1.1 has to contain the specific host in the header.

3. Optimization of bandwidth

In terms of HTTP 1.0, the server sends the whole object to the client even the client only needs part of it. Thus, this will cause a waste of bandwidth. Furthermore, HTTP 1.0 does not support resuming HTTP downloads. To improve this, HTTP 1.1 introduces the range, which allows the client to request a part of the response entity. This also means HTTP can use the range to resume the aborted downloads.

For example, suppose the client wants to downloads a 1024 KB file and he has already downloaded 512 KB but the network interrupts. The client can request to resume downloads from bytes=512000 by adding Range:bytes=512000-  into the header. When the server knows the request, he would send the file from 512KB by adding Content-Range:bytes 512000-/1024000 into the header.

### The main differences between HTTP 1.1 vs HTTP 2.0
There are mainly four differences between HTTP 1.1 and HTTP 2.0. Firstly, HTTP 2.0 has a binary format layer. Secondly, HTTP 2.0 introduces multiplexing. Thirdly, HTTP 2.0 adds the header compression. Fourthly, HTTP 2.0 has the server push function. Let me explain them one by one.

1. Binary format layer

HTTP 2.0 adds a binary format layer between the application layer and the transport layer. HTTP 1.0 transfers the plain-text message, but HTTP 2.0 transfers the binary frame. Thus, the message format is different. As is shown in the image below, you can see the binary frame layer breaks the message into frames. The header of HTTP 1.1 would be encapsulated into the HEADERS frame, the response body would be encapsulated into the DATA frame.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/httpFormat.png)


HTTP 2.0 sets up a TCP connection, within this connection, there are amounts of the bilateral streams of data. Each stream consists of multiple messages in the request/response format, each message can be split into small units called the frame. The image below showcases the format of the stream.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/streamFormat.png)

The advantage of the binary format layer is that it increases the flexibility of data transfer. The reason is plain-text is diversified, thus, it is much more difficult for HTTP 1.1 to deal with the robustness. However, HTTP 2.0 transfer binary format, this would be much more helpful for maintaining the robustness.

2. Multiplexing

HTTP 1.0 requests and responses in a stop-and-wait way, which has pretty low efficiency. HTTP 1.1 introduces pipelining, we mentioned above. Although the persistent connection with pipelining improves the performance over stop-and-wait, this optimization strategy has the bottleneck, which would cause Head of Line (HOL) Blocking. Head of Line (HOL) Blocking means if the head of the packet cannot be passed by the destination port when it arrives at the destination, it will cause other packets behind to be blocked.

To improve this, HTTP 2.0 introduces multiplexing and tags each frame. Multiplexing allows the client to construct the multiple streams in parallel, these streams share a single TCP connection. In the following graph, you can see the process of multiplexing.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/multiplexing.png)


In HTTP 2.0, each frame would be tagged to a specific stream, the tag allows the connection to interleave these frames during transfer and reassemble them at the other side. Thus, the request and response frames can be transferred in parallel without blocking the behind messages.

3. Header compression

Let’s see an example before explaining the header compression. Suppose we have the following two requests, the headers of these two requests are as follow.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/http1.1request.png)

Everything is the same except path field. In HTTP 1.1, we have to send these two request messages twice, which means the same fields will be sent twice. Thus, the size of the messages will be pretty large. However, in HTTP 2.0. these two headers will be encapsulated into the header frames. HTTP 2.0 can compress the header frames. The result header frames are as follow:

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/http2.0request.png)


When sending the request 2, we only encode the path field and we can reconstruct the header with other common fields. This is the process of header compression.

4. Server push

Before I explain the server push of HTTP 2.0, let me introduces how to improve the speed of requesting and resource lining of HTTP 1.1.

HTTP 1.1 uses resource lining to send the client objects that they might need before the client asks for them. For example, if a client requests for a CSS file, the server includes other objects that the client might need into the file, and package them together and send back the client. Resource lining will reduce the total number of requests, but it also has some drawbacks. Firstly, the server put all the resources together and the client cannot separate or decline to receive them if it doesn’t need them. Secondly, if putting too many objects within an HTML file, then the HTML file will be very large, which decreases the connection speed.

To improve these problems, HTTP 2.0 introduces the server push. Since HTTP 2.0 enables sending concurrent responses to a client’s get request, the server can send other objects separately along with the requested HTML file. In this way, the client can choose to cache them or decline them.

Reference
https://www.digitalocean.com/community/tutorials/http-1-1-vs-http-2-what-s-the-difference

https://juejin.im/entry/5981c5df518825359a2b9476

https://www.zhihu.com/question/34074946