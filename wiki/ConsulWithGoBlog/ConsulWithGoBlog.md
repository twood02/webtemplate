---
layout: page
title: Consul KV with Go API Web Interface
permalink: /wiki/ConsulWithGoBlog/
---

*by:* Katie Bramlett, Sreya Nalla, and Linnea Dierksheide

### Exploring Consul: Building a Web Interface for Key-Value Store

We offer an explanation of distributed consensus, what Consul can do, and a step-by-step tutorial for building a simple Web interface for a key value store using Consul.

---

### Distributed Consensus & Raft
For a system with ust one node, making updates is simple. The client can just send the change directly to that node and if they get a success response from the system, the client can be confident that their change was made. However, if we have a *distributed* system, meaning that there are many nodes that should be working together and have the same information, making updates across the entire system becomes a lot more difficult. If we just send one request to a single nodes, how can we ensure that all nodes will get the update? What if nodes get updates in different orders? What we want is a way for all the nodes to come to a clear *consensus* on the state of the system. We would also want the system we use to provide strong consistency, fault tolerance, and liveness.


**Raft** is one broadly used distributed consensus algorithm, and it is the one that drives Consul. Each node is a follower, a candidate, or a leader. The leader node receives all requests and sends the updates to the followers (the rest of the nodes), all of which keep a replicated log of all the changes to the system. Once a majority of nodes let the leader know they've logged the request, it becomes committed. The leader also sends "heartbeats" every so often to the nodes. A node becomes a candidate if it stops receiving heartbeats from the leader (as this would mean the leader failed in some form). This will trigger an election, and the candidate votes for itself and sends a vote request to all other nodes, which will vote for it if it its log isn't "behind" theirs. Once it receives a majority of votes, the candidate becomes the new leader and the process continues. Though this is a quick summary of Raft, it is simple to understand and implement, a key benefit over other consensus algorithms like Paxos.

---

### What is Consul?
**Consul** is a multi-cloud service networking platform to connect and secure services across any runtime platform and public or private cloud. First released in 2014, it was intended for DNS-based service discovery, and provides a fully featured service-mesh control plane, distributed key-value storage, service discovery, segmentation, and configuration. Registered services and nodes can be queried using a DNS interface or an HTTP interface.<br> 
*Some of Consul's main features are:*
- service discovery
- health checking
- KV store
- secure service communication
- multi datacenter. <br> 

--- 

### Step-by-Step Tutorial
#### Setting Up Consul:
1. Install Consul<br>
*Install Consul in one of 3 ways:*<br>
 Manual Installation | Homebrew on OS X | Chocolately on Windows
- | - | -
Find the appropriate package for your system and download it <br>[here.](https://www.consul.io/downloads.html) | Homebrew is a free and open-source package management system for Mac OS X. From the command line, run:<br>`brew install consul` | Chocolately is a free and open-source package management system for Windows. From the command line, run:<br>`choco install consul`
<br>
*Verify Installation:*<br>
After installing Consul, verify that the installation worked by opening a new terminal session and running the command `consul`.<br>
You should see:<br>
```
$ consul
usage consul [--version] [--help] <command> [<args>]

Available commands are:
    agent          Runs a Consul agent
    event          Fire a new event

...
```
2. 
3. 
#### Store Data in Consul KV
*We will be focusing on the KV store feature of Consul.*<br>
<br>
Consul KV is a core feature of Consul and is installed with the Consul agent. It allows users to store indexed objects, and its main uses are storing configuration parameters and metadata. Consul KV is a simple KV store, and is not intended to be a full featured datastore, although it has similarities to one. The Consul KV datastore is located on the servers, but it can be accessed by any agent (client or server). It is automatically enabled on all Consul agents.<br><br>
*With Consul KV, we can add, delete, modify, and query data. All data will be stored as a key-value pair.*<br>
There are two ways to interact with the Consul KV store: the HTTP API and the command line (CLI). The simple HTTP API makes it easy to use.<br>
#### How to Setup Consul KV:


### Creating a Simple Web Interface to Consul KV Store using Go 

### Libraries Used in Web Interface

### Sources Used:

