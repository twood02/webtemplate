---
layout: page
title:  SSH Tunneling and Port Forwarding
permalink: /wiki/tunneling/
---

*by:* Zach Coveno & Kevin Deems

How to set up SSH Tunnels to connect networks together.

---

<img src="/wiki/tunneling/ssh_tunnels.png">

#### SSH Tunnels

First thing's first, what is SSH? SSH stands for Secure Shell (protocol). "Secure" in this setting means protected opposed to "reliable". Secure tunnels via `ssh` encrypt traffic coming out of network for quick and secure application access. Many services such as VPNs, legacy applications, and intranet communication use ssh tunnels to encrpyt and send arbitrary data through a network.
SSH consists of three major levels 
1) Transport Layer Protocol - provides authentication and confidentiality 
2) User Authentication Protocol - does exactly what it sounds like, and authenticates the user to the server
3) Connection Protocol - distributes the encrypted tunnel to multiple channels 

Why might ssh tunnels be used? Here are a few examples:
- A firewall is blocking traffic in or out of the network on a specific port
- Securing an application is difficult, and sending plaintext over the internet is dangerous
- SSH is highly configurable, and its port `22` is almost always open on systems to allow at least for sysadmin manipulation

How secure are these tunnels? SSH allows you to configure the encryption standard to use. They include:
- aes(128/256)-cbc
- blowfish-cbc
- 3des-cbc

The massive configurability of ssh is very useful, and the top encryption algorithms have been deemed very secure in the community.

## A Brief History of SSH
Understanding the why behind something is just as important as understanding the what and how. Why was SSH created, what problem was SSH trying to solve? In 1995, Finland's Tatu Ylonen was conducting research for the Helsinki University of Technology when his network was attacked by a password sniffer. This lead Ylonen to build the first version of SSH (SSH-1) in an attempt to replace the vulnerable and exposed protocols of the time (TELNET, FTP, rlogin). Ylonen went on to found SSH Communications Security which continued to develop and market the protocol as it became more mainstream.

## What's different about SSH?
So what makes SSH secure? SSH was the first protocol to encrypt the data it tranfers. Older protocols would just send the raw data over the line allowing adversaries to read any or all of the data, just as the sniffer did at Helsi University. This isn't an encryption blog, so we won't go into much detail into how the encryption ensures protection of the data, but we will explain a little bit about how SSH implements it in order to be the "Secure Shell". 

SSH is a client-server protocol, meaning it is on the client to initiate a connection. The server will always be listening on the designated port (usually port 22), and once the client reaches out both ends agree upon how they will secure the connection (which encryption algorithm they'll be using). This initial communication is secured by using a public key to authenticate the server to the client. With the encryption standard agreed upon, future communication will be encrypted and therefor protected from any adversary listening in on the commuincation. After this initial connection, the client will only need to provide their credentials for the server to authenticate and re-establish the connection. The examples below will demostrate some of the power/uses SSH provides users by providing a secure way to communicate to a server. 

## Example of Port Forwarding
<img src="/wiki/tunneling/forward.png">
The server name is `aws`, the jekyll port is 4000, local computer wants to use port 8080
```sh
# On the server
./run.sh
# On local computer
ssh -N -L 8080:localhost:4000 aws
```

#### What's happening here? Command breakdown:
- Using a regular `ssh` session, connect to `aws`
- In the example, `aws` is set in the `~/.ssh/config` for easy connection
	- Example ssh config
	```sh
		Host aws
			Hostname your_instance.compute-1.amazonaws.com
			User <username>
			IdentityFile ~/.ssh/aws_key
	```
- The `-N` tells `ssh` to not run any commands, only establish a connection and wait for requests
- `-L` is a local port forward
        - `8080` is the local computer port to be requested
        - `localhost` is the address on the ec2 instance to send
        - `4000` tells the instance what port to find from its `localhost`

When a user goes on the local computer (that ran the `ssh` command), and types `localhost:8080` into a browser, ssh receives the request
- SSH tells `aws` to make a request to its `localhost:4000`, and the result is sent back to your `localhost:8080`

### The Implication
Imagine a system administrator has created a firewall for security reasons, only allowing special ports to open like ssh's `22`

A user wouldn't be able to see any web content from a browser running on port 80.

This is a very simple use case of tunneling. Let's look at an even better use case.

## Complex Local Port Forwarding
A gateway server, or jump box is an entry point server into a network.
<img src="/wiki/tunneling/jump_box.png">
- Above is a simple example of a Jump Server

In many data centers, most servers are protected in a private subnet
- This is mostly for security in large networks
- One gateway device has a public IP, and all the others can communicate through it on a private subnet

Imagine someone wants to view content through this device, onto a web server in the data center.
They might set up port forwarding through the jump host (jumphost.domain.com) to the protected server at (inside-server.jumphost.domain.com)
- Here `jumphost` has access to both the internet, and the private network in the cluster

Using another ssh configuration we can set up port forwarding again
```sh
ssh -N -L 8080:inside-server:80 username@jumphost.domain.com
```

There is one key difference from the previous example, that localhost is not requested on the remote server
- Now, `inside-server:80` (http traffic) is requested whenever we request `localhost:8080` on our **local** computer browser

## Final Example - *Reverse* Tunneling
This example is a little more convoluted than the previous two, and perhaps more widely used.

Imagine you work at a company that gives you access to a server inside a private network like the one previously described (with a jump box).
- The server has complete access on ports 80, 443 and 22 to the outside world, for regular web and ssh traffic
- The server cannot be accessed remotely with a public IP address
- Your company doesn't give you access to the jump server as in the previous example. What do you do?

Use reverse tunneling with ssh! Let's look at this example, where we're on `inside-server`
```sh
ssh -N -R 23400:localhost:22 aws
```

- Here, there is one big distinction, the `-R` flag, which stands for `remote`
- A secure channel is set up from `inside-server` to the `aws` instance far away and outside the private network

To connect, go to the `aws` server and connect to port `23400`
```sh
ssh <inside-server username>@localhost:23400
```

What is essentially happening is you are making an ssh connection to `inside-server` from `aws`

This would not have been possible before because `inside-server` doesn't have a public IP!

### Conclusion
These are some of the really cool things that ssh allows you to do. SSH is secure, configurable, and available in some version on almost every machine. This allows you to develop all your software in the "cloud". Almost none of it will stay persistent on your machine! This is incredibly useful for server development, remote network configuration, and much more!

## Resources/ Suggested Links
https://www.ssh.com/ssh/protocol#strong-authentication-with-ssh-keys
https://en.wikipedia.org/wiki/Secure_Shell
https://tools.ietf.org/html/rfc4251#ref-SSH-CONNECT

Resources on Encryption
https://www.digitalocean.com/community/tutorials/understanding-the-ssh-encryption-and-connection-process
