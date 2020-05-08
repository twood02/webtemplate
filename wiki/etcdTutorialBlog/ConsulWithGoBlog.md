---
layout: page
title: Consul KV with Go API Web Interface
permalink: /wiki/ConsulWithGoBlog/
---

*by:* Katie Bramlett, Sreya Nalla, and Linnea Dierksheide

#### Consul is an important software platform in the realm of distributed systems.
Here, we offer a step-by-step tutorial for setting up Consul, implementing a key-value store with Consul KV, and creating a simple web interface to access the store using the Go API.

---

### What is Consul?
**Consul** is a multi-cloud service networking platform to connect and secure services across any runtime platform and public or private cloud. First released in 2014, it was intended for DNS-based service discovery, and provides a fully featured service-mesh control plane, distributed key-value storage, service discovery, segmentation, and configuration. Registered services and nodes can be queried using a DNS interface or an HTTP interface.<br><br> 
*Some of Consul's main features are:*
- service discovery
- health checking
- KV store
- secure service communication
- multi datacenter. <br> 

### Step-by-Step Tutorial
#### Setting Up Consul:
1. Install Consul
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

