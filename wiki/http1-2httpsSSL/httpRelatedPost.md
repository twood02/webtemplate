---
layout: page
title:  HTTP vs HTTP 2 vs HTTPs with SSL
permalink: /wiki/http1-2httpsSSL/
---

*by:* Weizhao Li & Ziyue Li & Zhuolun Gao


T short description of the Post:
This is a post of HTTP vs HTTP 2 vs HTTPs with SSL. The section HTTP 1 vs HTTP 1.1 vs HTTP 2.0 is written by Weizhao Li, the section of HTTPs is written by Ziyue Li, the section of Secure Sockets Layer (SSL) is written by Zhuolun Gao. 

---

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

## HTTPS
Hypertext Transfer Protocol Secure (HTTPS) is an extension of HTTP. It can also be called with HTTP over SSL (or TLS), which indicates its high relation with SSL protocol. 
Now,let’s take a brieft journey on how HTTPS evolved and take a closer look at the basic concepts of the it.

### The defect of HTTP
There are three main defects that HTTP hold:
-	Communication uses clear text (not encrypted), which may be eavesdropped.
- The identity of the correspondent is not verified, so there may be spoofing.
- The integrity of the message cannot be proven, so it may have been tampered with.
To solve these questions and get a more safe communication, SSL (or TLS) are introduced.

### Brief history of SSL
Secure Sockets Layer (SSL), and its now standardized successor, Transport Layer Security (TLS), was first developed by Netscape Company. Taher Elgamal(Figure 1), chief scientist at Netscape Communications from 1995 to 1998, has been described as the "father of SSL". SSL Version 1.0 was never publicly released because of serious security flaws in the protocol. Version 2.0, released in February 1995, contained a number of security flaws which necessitated the design of Version 3.0.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Taher%20Elgamal.png)

Taher Elgamal

TLS 1.0 was first defined in RFC 2246 in January 1999 as an upgrade of SSL Version 3.0, and written by Christopher Allen and Tim Dierks of Consensus Development. As stated in the RFC, "the differences between this protocol and SSL 3.0 are not dramatic, but they are significant enough to preclude interoperability between TLS 1.0 and SSL 3.0". TLS 1.0 does include a means by which a TLS implementation can downgrade the connection to SSL 3.0, thus weakening security.

After that, several version of TLS developed to make improvement on security and architecture. Here is a list of different vesion of SSL and TLS.

SSL and TLS protocols

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/SSL%20and%20TLS%20protocols.png)

Therefore, HTTPS can be seen a combination work between HTTP and SSL/TLS, as Form below illustated.

The difference between HTTP and HTTPS on network layer architecture

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/The_difference_between_HTTP_and_HTTPS_on_network_layer_architecture.png)

### Basic cryptography
There are two cryptography to be showed here: Symmetric-key algorithm and asymmetric cryptography. Both cryptography approaches will be used to secure communications while guarantee the performance.

### Cleartext

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Cleartext_transmission_is_easy_to_be_eavesdropped.png)

Cleartext transmission is easy to be eavesdropped

### Symmetric-key encryption


![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_of_Symmetric-key_encrypted_transmission.png)

Illustration of Symmetric-key encrypted transmission

Symmetric-key algorithms are algorithms for cryptography that use the same cryptographic keys for both encryption of plaintext and decryption of ciphertext. The keys may be identical or there may be a simple transformation to go between the two keys. The keys, in practice, represent a shared secret between two or more parties that can be used to maintain a private information link. This requirement that both parties have access to the secret key is one of the main drawbacks of symmetric key encryption(shows on Figure below), in comparison to public-key encryption (also known as asymmetric key encryption)

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/How_to_transfer_key_itself_safely.png)

How to transfer key itself safely?

### Asymmetric key encryption
Public-key cryptography, or asymmetric cryptography, is a cryptographic system that uses pairs of keys: public keys which may be disseminated widely, and private keys which are known only to the owner. The generation of such keys depends on cryptographic algorithms based on mathematical problems to produce one-way functions. A one-way function is a function that is easy to compute on every input, but hard to invert given the image of a random input in the sense of computational complexity theory. Effective security only requires keeping the private key private; the public key can be openly distributed without compromising security.

As figure below shows, an unpredictable (typically large and random) number is used to begin generation of an acceptable pair of keys suitable for use by an asymmetric key algorithm.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_on_key_pairs_generation_process.png)

Illustration on key pairs generation process.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_of_Asymmetric_key_encrypted_communication.png)

Illustration of Asymmetric key encrypted communication

Figure above shows the basic step of using key pairs to communicate safely. At first, the client will request server for its public keys, and the server send its public key back to the client. Although this step both request and public key itself is in cleartext, this step will guarantee the next step’s security. 

### Middleman Attack
Although asymmetric key encryption seemed to be a way to communicate safely, problems still exist—the middleman attack.

A man-in-the-middle attack (MITM) is an attack where the attacker secretly relays and possibly alters the communications between two parties who believe that they are directly communicating with each other. One example of a MITM attack is active eavesdropping, in which the attacker makes independent connections with the victims and relays messages between them to make them believe they are talking directly to each other over a private connection(as below shows), when in fact the entire conversation is controlled by the attacker. The attacker must be able to intercept all relevant messages passing between the two victims and inject new ones. This is straightforward in many circumstances; for example, an attacker within reception range of an unencrypted wireless access point (Wi-Fi) could insert themselves as a man-in-the-middle.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_of_Middleman_attack.png)

Illustration of Middleman attack

### Solution: Certificate Authority & Digital Certificates
Now it is time to introduce the almost-perfect solution: Certificate Authority and Digital Certificate. A certificate authority or certification authority (CA) is an entity that issues digital certificates. The issuing process is showed on below. A digital certificate certifies the ownership of a public key by the named subject of the certificate. This allows others (relying parties) to rely upon signatures or on assertions made about the private key that corresponds to the certified public key. A CA acts as a trusted third party—trusted both by the subject (owner) of the certificate and by the party relying upon the certificate.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Digital_Certificate_Issuing_Process.png)

Digital Certificate Issuing Process

As showed on below, the Digital Certificate will used to make the first step identification and exchange the random key, which will be used later in a format of symmetric-key encryption in order to improve the performance.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Communication_with_Digital_Certificate_between_client_and_server.png)

Communication with Digital Certificate between client and server

### Experiences
Because of the encryption process, HTTPS will incur higher latency than HTTP/1. Here is a small experience to show that.

First, there are some python codes. I recommend you to run the code on jupyter notebooks.
import requests:

```Python
import requests
import time

i = 0
url="http://zli.name"
l =list()
while i < 100:
    r = requests.get(url)
    time.sleep(1)
    print (r.elapsed.total_seconds())
    l.append(r.elapsed.total_seconds())
    i = i + 1
print(l)
```
Above, the code will request zli.name, which is my newly build personal blow website based on wordpress for 100 times and restore the time elapsed in each request in a list.

Then, I will use bellowed code to draw a histogram to show the time needed for zli.name to response.
```Python
data = [1,2,3]
%matplotlib inline
import matplotlib.pyplot as plt
 
plt.xlabel("L")
plt.ylabel("#")
plt.title("http://zli.name")
plt.hist(l)
```
The result is showed here.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/http_picture.png)

While, I will do the nearly samething for the https://zli.name, everything is same with above, just change the value of url from http://zli.name to https://zli.name.

The result shows below.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/https_picture.png)

You can see that, while indicated by the red circle, the HTTPS is relatively causing more delay than HTTP.

Below shows the protocol is HTTP/1. By the developer function provided by Chrome.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/http1.1.png)

In fact, I have an idle domain name and last night I build a website based on wordpress with it. I think this process can give you a direct show of how to install SSL certificate and thus made the website you created more safely.

At first, my website shows something like this:

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/unsecure_connection_revised.png)

You can see the red box showing that there is no valid certificate. Then, I use Let’s encrypt to install certificate on my browser. When you choose your OS and your server end, Certbot will show customized instructions.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/let's_encrypt.png)

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/certbot.png)

This is what I choose based on my server configuration:

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/choose_the_os_and_server_arch.png)

Instrcutions are below:

First, ssh into ther server of your website.

Second, you'll need to add the Certbot PPA to your list of repositories. To do so, run the following commands on the command line on the machine:

```shell
$ sudo apt-get update
$ sudo apt-get install software-properties-common
$ sudo add-apt-repository universe
$ sudo add-apt-repository ppa:certbot/certbot
$ sudo apt-get update
```
Third, Install Certbot. Run this command on the command line on the machine to install Certbot.
```shell
$ sudo apt-get install certbot python-certbot-nginx
```
Last, Run this command to get a certificate and have Certbot edit your Nginx configuration automatically to serve it, turning on HTTPS access in a single step.
```shell
$ sudo certbot --nginx
```
After that, you can see that the website can be recognized as safe through Chrome browser, and within the green box, the certificate become valid.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/secure_connection_revised.png)

I hope this section will help you understanding how HTTPS made data more safe to transfer, and how Digital Certificate made the website more safe to browse. :)



## Secure Sockets Layer (SSL) 
   
### Background  
   
Secure Sockets Layer (SSL) is an enhanced version of TCP which can provide confidentiality, data integrity, and etc.   

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/bigpicture.png)    
    	
New applications such as e-commerce and online banking based on the World Wide Web greatly facilitate people's daily life and are favored by people. Because these applications need online transactions on the network, they put forward higher requirements for the security of network communication. The traditional World Wide Web Protocol HTTP does not have the security mechanism, such as transmitting data in the form of clear text, unable to verify the identity of both sides of the communication, unable to prevent the transmission of data from being tampered, etc., which leads to HTTP unable to meet the security requirements of e-commerce and online banking applications.    

The Security Protocol SSL proposed by Netscape company uses data encryption, authentication and message integrity authentication mechanisms to provide security guarantee for data transmission on the network. SSL can provide a secure connection for HTTP, which greatly improves the security of the World Wide Web.   

### Technical advantages       
1. Provide high security assurance. SSL uses data encryption, authentication and message integrity authentication mechanisms to ensure the security of data transmission on the network.        
2. Support various application layer protocols. Although SSL is designed to solve the security problem of the world wide web, because SSL is located between the application layer and the transport layer, it can provide security guarantee for any application layer protocol based on reliable connection such as TCP.   
3. Simple deployment. At present, SSL has become a global standard used to identify the identity of website and web browser, and to conduct encrypted communication between browser users and web servers. SSL protocol has been integrated into most browsers, such as ie, Netscape, Firefox, etc. This means that almost any computer with a browser supports SSL connection and does not need to install additional client software.    
     
### Protocol security mechanism      
1. The security mechanisms of SSL protocol include:
Confidentiality of data transmission: using symmetric key algorithm to encrypt the transmitted data.      
2. Authentication mechanism: Based on the certificate, the digital signature method is used to authenticate the server and the client, and the client's authentication is optional.                
3. Message integrity verification: MAC algorithm is used to verify message integrity during message transmission.        
4. Asymmetric key algorithm guarantees key security.     
5. Using PKI to ensure the authenticity of public key: PKI publishes users' public key through digital certificate, and provides a mechanism to verify the authenticity of public key. The digital certificate is a file containing the public key and identity information of the user, which proves the association between the user and the public key. The digital certificate is issued by the Authority CA, and the authenticity of the digital certificate is guaranteed by CA.    
 
### Working process of the agreement 
   
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/agreement.png)  
	
#### SSL handshake process 
 
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/handshake.png)      
SSL negotiates the session parameters between the client and the server through the handshake process and establishes the session. The main parameters of a session include session ID, the certificate of the other party, encryption suite (key exchange algorithm, data encryption algorithm, MAC algorithm, etc.) and master secret. The data transmitted through SSL session will be encrypted and Mac will be calculated by the master key and encryption suite of the session.   

#### Verify only the handshake process of the server
   
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/server.png)          
	
#### Verify SSL handshake process between server and client  
  
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/server_client.png)       
	
#### SSL handshake process to restore the original session   

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/original_session.png)    
### Record protocol  

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/record_protocal.png)    
	
The recording protocol is use after the client and server shake hands successfully, that is, after the client and server identify each other and determine the algorithm for security information exchange, they enter the SSL recording protocol, which provides two services for SSL connection:     

1. Confidentiality: using the secret key defined by handshake protocol    
2. Integrity: the handshake protocol defines Mac to ensure message integrity         

### Alert protocol
When the client and server find an error, they send an alert message to each other. If it is a fatal error, the algorithm will immediately close the SSL connection, and both parties will delete the relevant session number, secret and key first. Each alarm message has 2 bytes in total. The first byte indicates the error type. If it is an alarm, the value is 1. If it is a fatal error, the value is 2. The second byte specifies the actual error type.  
### Summary   
In SSL, we use the handshake protocol to negotiate encryption, MAC algorithm and secret key, use the handshake protocol to encrypt and sign the exchanged data, and use the alert protocol to define how to solve the problems in the data transmission process.

### Referrence
Kurose, James F, and Keith W. Ross. Computer Networking: A Top-Down Approach Featuring the Internet. Boston: Addison-Wesley, 2001. Print.



	

	
	
	
	 
	


