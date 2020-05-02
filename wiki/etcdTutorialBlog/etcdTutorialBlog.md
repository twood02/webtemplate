---
layout: page
title: Tutorial on etcd
permalink: /wiki/etcdTutorialBlog/
---

*by:* Katie Bramlett, Sreya Nalla, and Linnea Dierksheide


#### Etcd is a fundamental software in the realm of distributed systems. Here, we will share more about etcd and its applications, how to set it up, and more.

---

### What is etcd?
"Etcd is an open-source distributed key-value store created by the CoreOS team, now managed by the Cloud Native Computing Foundation. It serves as the backbone of many distributed systems, providing a reliable way for storing data across a cluster of servers. It works on a variety of operating systems including Linux, BSD and OS X. (Rus)"

Many of the characteristics of etcd make it a valuable asset to distributed systems:
- Fully Replicated: The entire store is available on every node in the cluster
- Highly Available: Etcd is designed to avoid single points of failure in case of hardware or network issues
- Consistent: Every read returns the most recent write across multiple hosts
- Simple: Includes a well-defined, user-facing API (gRPC)
- Secure: Implements automatic TLS with optional client certificate authentication
- Fast: Benchmarked at 10,000 writes per second
- Reliable: The store is properly distributed using the Raft algorithm

### How do you set up and use etcd?

### Web Interface Example using etcd

### Libraries Used in Web Interface

### Sources Used:
Rus, Calin. "What Is Etcd and How Do You Set Up an Etcd Kubernetes Cluster?." *Rancher,* 7 Apr. 2020, www.rancher.com/blog/2019/2019-01-29-what-is-etcd/. Accessed 1 May 2020.


