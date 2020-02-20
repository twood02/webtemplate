---
layout: page
title:  HTTP vs HTTP 2 vs HTTPs with SSL
permalink: /wiki/http1-2httpsSSL/
---

*by:* Weizhao Li & Ziyue Li & Zhuolun Gao


A short description of your post goes here.

---

The rest of your post goes here.
TIP:

插入图片例子：
1. github 上，把图片上传到/wiki/http1-2httpsSSL/images/ 目录下
2. 根据如下格式：

```bash
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Graduation_capital.jpg)

```

3. 这就会把images里面这周图片给显示出来
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Graduation_capital.jpg)

/**********************************/

### HTTP 1 vs HTTP 1.1 vs HTTP 2.0


### HTTPS
Hypertext Transfer Protocol Secure (HTTPS) is an extension of HTTP. It can also be called with HTTP over SSL (or TLS), which indicates its high relation with SSL protocol. 
Now,let’s take a brieft journey on how HTTPS evolved and take a closer look at the basic concepts of the it.

#### The defect of HTTP
There are three main defects that HTTP hold:
-	Communication uses clear text (not encrypted), which may be eavesdropped.
- The identity of the correspondent is not verified, so there may be spoofing.
- The integrity of the message cannot be proven, so it may have been tampered with.
To solve these questions and get a more safe communication, SSL (or TLS) are introduced.

#### Brief history of SSL
Secure Sockets Layer (SSL), and its now standardized successor, Transport Layer Security (TLS), was first developed by Netscape Company. Taher Elgamal(Figure 1), chief scientist at Netscape Communications from 1995 to 1998, has been described as the "father of SSL". SSL Version 1.0 was never publicly released because of serious security flaws in the protocol. Version 2.0, released in February 1995, contained a number of security flaws which necessitated the design of Version 3.0.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Taher%20Elgamal.png)

Taher Elgamal

TLS 1.0 was first defined in RFC 2246 in January 1999 as an upgrade of SSL Version 3.0, and written by Christopher Allen and Tim Dierks of Consensus Development. As stated in the RFC, "the differences between this protocol and SSL 3.0 are not dramatic, but they are significant enough to preclude interoperability between TLS 1.0 and SSL 3.0". TLS 1.0 does include a means by which a TLS implementation can downgrade the connection to SSL 3.0, thus weakening security.

After that, several version of TLS developed to make improvement on security and architecture. Here is a list of different vesion of SSL and TLS.

SSL and TLS protocols

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/SSL%20and%20TLS%20protocols.png)

Therefore, HTTPS can be seen a combination work between HTTP and SSL/TLS, as Form 2 illustated.

The difference between HTTP and HTTPS on network layer architecture

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/The_difference_between_HTTP_and_HTTPS_on_network_layer_architecture.png)

#### Basic cryptography
There are two cryptography to be showed here: Symmetric-key algorithm and asymmetric cryptography. Both cryptography approaches will be used to secure communications while guarantee the performance.

##### Cleartext

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Cleartext_transmission_is_easy_to_be_eavesdropped.png)

Cleartext transmission is easy to be eavesdropped

##### Symmetric-key encryption


![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_of_Symmetric-key_encrypted_transmission.png)

Illustration of Symmetric-key encrypted transmission

Symmetric-key algorithms are algorithms for cryptography that use the same cryptographic keys for both encryption of plaintext and decryption of ciphertext. The keys may be identical or there may be a simple transformation to go between the two keys. The keys, in practice, represent a shared secret between two or more parties that can be used to maintain a private information link. This requirement that both parties have access to the secret key is one of the main drawbacks of symmetric key encryption(shows on Figure 4), in comparison to public-key encryption (also known as asymmetric key encryption)

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/How_to_transfer_key_itself_safely.png)

How to transfer key itself safely?

##### Asymmetric key encryption
Public-key cryptography, or asymmetric cryptography, is a cryptographic system that uses pairs of keys: public keys which may be disseminated widely, and private keys which are known only to the owner. The generation of such keys depends on cryptographic algorithms based on mathematical problems to produce one-way functions. A one-way function is a function that is easy to compute on every input, but hard to invert given the image of a random input in the sense of computational complexity theory. Effective security only requires keeping the private key private; the public key can be openly distributed without compromising security.

As Figure 5 shows, an unpredictable (typically large and random) number is used to begin generation of an acceptable pair of keys suitable for use by an asymmetric key algorithm.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_on_key_pairs_generation_process.png)

Figure 5. Illustration on key pairs generation process.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_of_Asymmetric_key_encrypted_communication.png)

Figure 6. Illustration of Asymmetric key encrypted communication

Figure 6 shows the basic step of using key pairs to communicate safely. At first, the client will request server for its public keys, and the server send its public key back to the client. Although this step both request and public key itself is in cleartext, this step will guarantee the next step’s security. 

#### Middleman Attack
Although asymmetric key encryption seemed to be a way to communicate safely, problems still exist—the middleman attack.

A man-in-the-middle attack (MITM) is an attack where the attacker secretly relays and possibly alters the communications between two parties who believe that they are directly communicating with each other. One example of a MITM attack is active eavesdropping, in which the attacker makes independent connections with the victims and relays messages between them to make them believe they are talking directly to each other over a private connection(as Figure 7 shows), when in fact the entire conversation is controlled by the attacker. The attacker must be able to intercept all relevant messages passing between the two victims and inject new ones. This is straightforward in many circumstances; for example, an attacker within reception range of an unencrypted wireless access point (Wi-Fi) could insert themselves as a man-in-the-middle.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Illustration_of_Middleman_attack.png)

Figure 7. Illustration of Middleman attack

#### Solution: Certificate Authority & Digital Certificates
Now it is time to introduce the almost-perfect solution: Certificate Authority and Digital Certificate. A certificate authority or certification authority (CA) is an entity that issues digital certificates. The issuing process is showed on Figure 8. A digital certificate certifies the ownership of a public key by the named subject of the certificate. This allows others (relying parties) to rely upon signatures or on assertions made about the private key that corresponds to the certified public key. A CA acts as a trusted third party—trusted both by the subject (owner) of the certificate and by the party relying upon the certificate.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Digital_Certificate_Issuing_Process.png)

Figure 8. Digital Certificate Issuing Process

As showed on Figure 9, the Digital Certificate will used to make the first step identification and exchange the random key, which will be used later in a format of symmetric-key encryption in order to improve the performance.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/Communication_with_Digital_Certificate_between_client_and_server.png)

Figure 9. Communication with Digital Certificate between client and server

#### Experiences
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

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/unsecure_connection.png)

Then, I use Let’s encrypt to install certificate on my browser. When you choose your OS and your server end, Certbot will show customized instructions.

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/let's_encrypt.png)

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/certbot.png)

This is what I choose based on my server configuration:

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/choose_the_os_and_server_arch.png)

Instrcutions are below:

First, ssh into ther server of your website.

Second, you'll need to add the Certbot PPA to your list of repositories. To do so, run the following commands on the command line on the machine:

```bash user
sudo apt-get update
sudo apt-get install software-properties-common
sudo add-apt-repository universe
sudo add-apt-repository ppa:certbot/certbot
sudo apt-get update
```


### Secure Sockets Layer (SSL) 
