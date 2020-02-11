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
Secure tunnels via `ssh` encrypt traffic coming out of network for quick and secure application access. Many services such as VPNs, legacy applications, and intranet communication use ssh tunnels to encrpyt and send arbitrary data through a network.

Why might ssh tunnels be used? Here are a few examples:
- A firewall is blocking traffic in or out of the network on a specific port
- Securing an application is difficult, and sending plaintext over the internet is dangerous
- SSH is highly configurable, and its port `22` is almost always open on systems to allow at least for sysadmin manipulation

How secure are these tunnels? SSH allows you to configure the encryption standard to use. They include:
- aes(128/256)-cbc
- blowfish-cbc
- 3des-cbc

The massive configurability of ssh is very useful, and the top encryption algorithms have been deemed very secure in the community.
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


