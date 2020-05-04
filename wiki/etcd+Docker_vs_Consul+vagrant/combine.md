layout: page
title:  etcd with Docker vs Consul with Vagrant
permalink: /wiki/etcd+Docker_vs_Consul+vagrant/
---

*by:* Weizhao Li & Ziyue Li 


The short description of the Post:
This is a post of etcd with Docker vs Consul with Vagrant. The section **Deploy a 3-node etcd Cluster in Docker** is written by Weizhao Li, the section of **Set Up Consul Based on Vagrant** is written by Ziyue Li. 

In **Deploy a 3-node etcd Cluster in Docker**, Weizhao covered the introduction of etcd and Docker, the main idea of Raft algorithm, and he provided an experiment to show how to deploy etcd in Docker and how to set up a multi-node etcd cluster in Docker. Finally, he talked about the comparison of etcd and consul.

In **Set Up Consul Based on Vagrant**, Ziyue introduced Consul and provided a mini-experiment of Consul based on Vagrant

# Section1: Deploy a 3-node etcd Cluster in Docker
Author: Weizhao Li

In this article, I will cover the introduction of etcd and Docker, the main idea of Raft algorithm, how to deploy etcd in Docker and how to set up a multi-node etcd cluster in Docker, and the comparison of etcd and consul.

Follow the instruction, you can build a 3-node etcd cluster in Docker from scratch, and feel the beauty of consistency in a distributed system.

Keywords: etcd, Raft, cluster, Docker, Consul

## 1. Introduction
### 1.1 What is etcd?

“Etcd is a distributed, consistent key-value store for shared configuration and service discovery”, this is the description of etcd in the official document. In June of 2013, CoreOS team developed the etcd project, their goal was to build a high available distributed key-value store database. Within the distributed system, shared configuration and service discovery are the basic and important problems, etcd is supposed to solve those problems.

When designing the etcd, the key points are:

- Easy to use: easy to read definition, user-friendly API.
- Secure: optional automatic TLS of client certificate authentication.
- High-speed: support concurrent 1 k/s write operation.
- Reliable: support distributed structure, use consensus algorithm based on Raft.

### 1.2 What is Raft algorithm?
Raft is a consensus algorithm from Stanford University, which is mainly used in log replication of distributed systems. The main idea of raft is to utilize the leader election to ensure strong consistency. The advantages of Raft is it is much easier to understand than Paxos and it has fault tolerance, which means some nodes fail or have network problems, most of the other nodes could work normally.

With etcd, users can set up multiple instances in multiple nodes, and group them as a cluster. Instances within the same cluster will share consistent information.

If you want to go deeper:

Read this overview Raft: Consensus made simple(r)
https://www.brianstorti.com/raft/

Read the full Raft paper
https://raft.github.io/raft.pdf

### 1.3 What is Docker?

“Docker is a software platform that allows you to build, test, and deploy applications quickly. Docker packages software into standardized units called containers that have everything the software needs to run including libraries, system tools, code, and runtime. Using Docker, you can quickly deploy and scale applications into any environment and know your code will run”, this is a clear definition of Docker from AWS.

In the case of Docker containers, images become containers when they run on Docker Engine. Container applications will run the same regardless of the environment and infrastructure, which means they work uniformly despite differences for instance.

### 1.4 What Can I do with Docker?
The main usages of docker contain the following three categories:

(1) Provide a one-time environment. For example, testing other people’s software locally and providing unit testing and build environments during continuous integration.

(2) Provide elastic cloud services. Because Docker containers can be opened and closed at any time, it is suitable for dynamic expansion and contraction.

(3) Establish a microservice architecture. With multiple containers, a machine can run multiple services, so a microservice architecture can be simulated locally.

## 2. Experiment

In this section, feel much more excited, we are going to build a cluster from scratch!

Now, I am going to guide you on how to deploy a 3-node cluster in Docker. First, you need to know what is a cluster.

### 2.1 What is a cluster?

Cluster means deploying the same application or service on different servers and forming them as a cluster, providing external services through load-balancing equipment.

The load-balancing equipment will deliver the request from the user to a certain server. Load Balancing is a computer technology used to distribute load among multiple computers (computer clusters), network connections, CPUs, disk drives, or other resources to achieve optimal resource usage, maximize throughput, or minimize response time while avoiding overload.

**Under the same configuration, the fewer the number of nodes, the better the cluster performance. But we need to avoid the even number of nodes, because:**

- During the election process, there is a higher probability of an equal amount of votes, which triggers the next round of elections.

- When the network split occurs, split the cluster nodes in half. The cluster will not work at this time. According to the RAFT protocol, the cluster write operation cannot make most nodes agree at this time, resulting in write failure and the cluster cannot work normally.

There are three ways of setting up a etcd cluster, if the IP of each node is known, you can use a static way to set up the cluster. However, in most cases, you don’t know the IP of each node, then you have to use discover ways. The discover ways contain etcd discovery and DNS discovery.

- Static: If the IP of each node is known, when starting etcd server, configure all node information through –initial-cluster parameter.
- etcd Discovery: The etcd discovery uses the existing etcd cluster as a data interaction point, and then implements the service discovery mechanism through the existing cluster when expanding the new cluster.
- DNS discovery: DNS discovery mainly records the domain name information of each node in the cluster through the dns service, and each node obtains mutual address information from the dns service to establish a cluster.

In this blog, I am going to show you how to deploy a 3-node cluster in a static way. We need Docker to generate three hosts that have three different IPs.

### 2.2 Steps of setting up a 3-node cluster in the static way.

**Step 1:**

**First, install Docker Machine on your host machine, in terminal, enter**

```
$ curl -L https://github.com/docker/machine/releases/download/v0.14.0/docker-machine-`uname -s`-`uname -m` >/tmp/docker-machine && \
install /tmp/docker-machine /usr/local/bin/docker-machine

#Docker Machine version
$ docker-machine -v
docker-machine version 0.14.0, build 89b8332
```
**Step 2:**

**Generate three virtualboxes.**

```
$ docker-machine create -d virtualbox manager1 && docker-machine create -d virtualbox worker1 && docker-machine create -d virtualbox worker2
```

Once three virtualboxes are all set, enter

```
$ docker-machine ls
```

You can see three virtualboxes, each one contains an IP address. Eg. 192.168.99.105

```
NAME       ACTIVE   DRIVER       STATE     URL                         SWARM   DOCKER     ERRORS
manager1   -        virtualbox   Running   tcp://192.168.99.105:2376           v19.03.5   
worker1    -        virtualbox   Running   tcp://192.168.99.106:2376           v19.03.5   
worker2    -        virtualbox   Running   tcp://192.168.99.107:2376           v19.03.5  
```
**Step 3:**

**Login each virtualbox and set up the configuration separately.**

Open three terminals and enter docker-machine ssh VIRTUALBOX NAME separately to login each VirtualBox.

**Manager1**

**To login manager1, enter the following command in one of the terminals.**

```
$ docker-machine ssh manager1
```

**Now we are entering manager1 VirtualBox. In manager1, we need to install etcd, enter**

```
$ curl -L https://github.com/coreos/etcd/releases/download/v3.3.0-rc.0/etcd-v3.3.0-rc.0-linux-amd64.tar.gz -o etcd-v3.3.0-rc.0-linux-amd64.tar.gz && sudo tar xzvf etcd-v3.3.0-rc.0-linux-amd64.tar.gz && cd etcd-v3.3.0-rc.0-linux-amd64 && sudo cp etcd* /usr/local/bin/
```

**Then set up the configuration of manager1, enter**

```
$ docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd quay.io/coreos/etcd:v2.3.8 \
 -name etcd0 \
 -advertise-client-urls http://192.168.99.105:2379,http://192.168.99.105:4001 \
 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -initial-advertise-peer-urls http://192.168.99.105:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd0=http://192.168.99.105:2380,etcd1=http://192.168.99.106:2380,etcd2=http://192.168.99.107:2380 \
 -initial-cluster-state new
```

**Note:**
- You have to change the IP address of this node according to its IP, change the name after -name .
- For worker 1 and worker 2, follow the same steps as the instruction of manager 1 node instead changing the IP address and the name.
- To see the meaning of each flag, eg. -XXX, enter this website Clustering flags
https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/configuration.md


**Worker1**

**Similarly, to login worker1, enter the following command in one of the terminals.**
```
$ docker-machine ssh worker1
```

**Now we are entering worker1 VirtualBox. In worker1, we need to install etcd, enter**
```
$ curl -L https://github.com/coreos/etcd/releases/download/v3.3.0-rc.0/etcd-v3.3.0-rc.0-linux-amd64.tar.gz -o etcd-v3.3.0-rc.0-linux-amd64.tar.gz && sudo tar xzvf etcd-v3.3.0-rc.0-linux-amd64.tar.gz && cd etcd-v3.3.0-rc.0-linux-amd64 && sudo cp etcd* /usr/local/bin/
```

**Then set up the configuration of worker1, enter**
```
$ docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd quay.io/coreos/etcd:v2.3.8 \
 -name etcd1 \
 -advertise-client-urls http://192.168.99.106:2379,http://192.168.99.106:4001 \
 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -initial-advertise-peer-urls http://192.168.99.106:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd0=http://192.168.99.105:2380,etcd1=http://192.168.99.106:2380,etcd2=http://192.168.99.107:2380 \
 -initial-cluster-state new
 ```

**Worker2**

**Similarly, to login worker2, enter the following command in one of the terminals.**

```
$ docker-machine ssh worker2
```

**Now we are entering worker2 VirtualBox. In worker2, we need to install etcd, enter**

```
$ curl -L https://github.com/coreos/etcd/releases/download/v3.3.0-rc.0/etcd-v3.3.0-rc.0-linux-amd64.tar.gz -o etcd-v3.3.0-rc.0-linux-amd64.tar.gz && sudo tar xzvf etcd-v3.3.0-rc.0-linux-amd64.tar.gz && cd etcd-v3.3.0-rc.0-linux-amd64 && sudo cp etcd* /usr/local/bin/
```

**Then set up the configuration of worker2, enter**
```
$ docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd quay.io/coreos/etcd:v2.3.8 \
 -name etcd2 \
 -advertise-client-urls http://192.168.99.107:2379,http://192.168.99.107:4001 \
 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -initial-advertise-peer-urls http://192.168.99.107:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd0=http://192.168.99.105:2380,etcd1=http://192.168.99.106:2380,etcd2=http://192.168.99.107:2380 \
 -initial-cluster-state new
```

If everything is successful, in each of the virtualboxes, enter
```
$ docker ps
```

If you can see:
```
docker@manager1:~/etcd-v3.3.0-rc.0-linux-amd64$ docker ps
CONTAINER ID        IMAGE                        COMMAND                  CREATED             STATUS              PORTS                                                                NAMES
9f1f5cc0cea2        quay.io/coreos/etcd:v2.3.8   "/etcd -name etcd0 -…"   40 hours ago        Up 40 hours         0.0.0.0:2379-2380->2379-2380/tcp, 0.0.0.0:4001->4001/tcp, 7001/tcp   etcd
```

Good job, that means etcd is configured successfully in this node!

**Step 4:**

Under the folder of etcd, Eg. we are under /home/docker/etcd-v3.3.0-rc.0-linux-amd64 folder of manager1 VirtualBox, enter

```
docker@manager1:~/etcd-v3.3.0-rc.0-linux-amd64$ docker exec -it 9f1f5cc0cea2 ./etcdctl member list   
```

**Note:**

- 9f1f5cc0cea2 is the container ID.

The result is:

```
714d7b53ea800f8d: name=etcd2 peerURLs=http://192.168.99.107:2380 clientURLs=http://192.168.99.107:2379,http://192.168.99.107:4001 isLeader=false
786595399fbb285d: name=etcd1 peerURLs=http://192.168.99.106:2380 clientURLs=http://192.168.99.106:2379,http://192.168.99.106:4001 isLeader=false
895c9c9c668b0d5e: name=etcd0 peerURLs=http://192.168.99.105:2380 clientURLs=http://192.168.99.105:2379,http://192.168.99.105:4001 isLeader=true
```

We can see etcd0, namely, manager1, is the leader of this election. Within the other two nodes, we can see the same information, that means the information is consistent in each node.

### 2.3 API Operations
**1.Example of key-value store**

In any node of this cluster, you can store the key-value pair, and each node can see the consistent up-to-date information. Eg. in manager 1, enter

```
docker@manager1:~/etcd-v3.3.0-rc.0-linux-amd64$ curl http://192.168.99.106:2379/v2/keys/test -XPUT -d value="test value"

# The results is :
{"action":"set","node":{"key":"/test","value":"test value","modifiedIndex":9,"createdIndex":9}}
```

In any node, eg. in worker1, we can extract the value of this key.

```
docker@worker1:~/etcd-v3.3.0-rc.0-linux-amd64$  curl http://192.168.99.107:2379/v2/keys/test 

# The results is :          
{"action":"get","node":{"key":"/test","value":"test value","modifiedIndex":9,"createdIndex":9}}
```

**2.Examples of managing the members**

Eg1.

To see the information of the members, in any node, eg. in manager1, enter

```
docker@manager1:~/etcd-v3.3.0-rc.0-linux-amd64$ curl http://192.168.99.106:2379/v2/members

# The results is :
{"members":[{"id":"714d7b53ea800f8d","name":"etcd2","peerURLs":["http://192.168.99.107:2380"],"clientURLs":["http://192.168.99.107:2379","http://192.168.99.107:4001"]},{"id":"786595399fbb285d","name":"etcd1","peerURLs":["http://192.168.99.106:2380"],"clientURLs":["http://192.168.99.106:2379","http://192.168.99.106:4001"]},{"id":"895c9c9c668b0d5e","name":"etcd0","peerURLs":["http://192.168.99.105:2380"],"clientURLs":["http://192.168.99.105:2379","http://192.168.99.105:4001"]}]}
```

We can see all the member in this cluster.


Eg2.

To see whether this node is leader, enter this command according to its IP address, eg.

```
docker@manager1:~/etcd-v3.3.0-rc.0-linux-amd64$ curl http://192.168.99.105:2379/v2/stats/leader

# The results is :
{"leader":"895c9c9c668b0d5e","followers":{"714d7b53ea800f8d":{"latency":{"current":0.001191,"average":0.002264701936054899,"standardDeviation":0.014763661938916848,"minimum":0.000365,"maximum":5.202087},"counts":{"fail":34,"success":180780}},"786595399fbb285d":{"latency":{"current":0.00118,"average":0.002796088718122103,"standardDeviation":0.008596453212714354,"minimum":0.000346,"maximum":0.472473},"counts":{"fail":0,"success":180493}}}}
```

As you can see, manager1, namely, etcd0 is the leader, and it has two followers.

## 3. etcd versus Consul

**Etcd**

etcd is a distributed key-value storage system using http protocol and Raft algorithm. Since it is easy to use and simple, many systems use etcd as part of service discovery, such as kubernetes (K8s). **But the problem** is that because it is just a storage system, if you want to provide complete service discovery functions, you need some third-party tools.

With Registrator, and confd, a very simple and powerful service discovery framework can be built. But compared to Consul, this type of operation is a little more troublesome. Therefore, in most of the cases, etcd are used for key-value storage, such as kubernetes.

Furthemore, **etcd scales better than Consul**. According to the experiment, when creating 1-million keys, 256 bytes key, 1KB value, etcd costs less response time and less memory, other systems like Consul suffer from high latencies and memory pressure.

**Consul**

According to the official document of Consul, “etcd et al. provide only a primitive K/V store and require that application developers build their own system to provide service discovery. Consul, by contrast, provides an opinionated framework for service discovery and eliminates the guess-work and development effort. Clients simply register services and then perform discovery using a DNS or HTTP interface. Other systems require a home-rolled solution.”

Also, as it is stated in the official document, “Consul provides first-class support for **service discovery, health checking, K/V storage, and multiple data centers**. To support anything more than simple K/V storage, all these other systems require additional tools and libraries to be built on top. By using client nodes, Consul provides a simple API that only requires thin clients. Additionally, the API can be avoided entirely by using configuration files and the DNS interface to have a complete service discovery solution with no development at all.”

### 4. Summary
etcd is a strongly consistent, distributed key-value store that provides a reliable way to store data. It mainly uses Raft algorithm to implement and it is easy to use. Docker is a software platform that allows you to build, test, and deploy applications quickly. Therefore, in this article, I mainly guided you to set up etcd, and build a 3-node etcd cluster in Docker from scratch. You can see 3 nodes work together and share consistent information in a cluster.

Compared with Consul, etcd has pros and cons. If looking for a distributed consistent key value store, etcd is a better choice over Consul. If looking for end-to-end cluster service discovery, etcd will not have enough features, instead, choose Consul or Kubernetes.

**etcd**

***Pros***

1. Simple and easy to use, no need to integrate SDK

2. Strong configurability

***Cons***

1. No health checking
2. Need to cooperate with third-party tools to complete service discovery
3. Does not support multiple data centers
4. Scale well, when creating a large amount of k/v store, cost less response time and less memory

**Consul**

***Pros***
1. Simple and easy to use, no need to integrate SDK
2. Bring your own health checking
3. Support multiple data centers

***Cons***
1. Cannot have notification of real-time service information change
2. It doesn’t scale well, when creating a large amount of k/v store, it has high latencies and memory pressure

## Reference

- Raft: Consensus made simple(r) https://www.brianstorti.com/raft/

- Raft paper  https://raft.github.io/raft.pdf

- 如何理解集群和分布式? https://juejin.im/post/5d43eb20e51d4561a54b693d#heading-6

- Docker 入门教程 https://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html

- Docker 搭建 etcd 集群 https://www.cnblogs.com/xishuai/p/docker-etcd.html

- etcd versus other key-value stores https://etcd.io/docs/v3.3.12/learning/why/

- etcd、Zookeeper和Consul一致键值数据存储的性能对比 https://cloud.tencent.com/developer/article/1491107

- Consul vs. ZooKeeper, doozerd, etcd https://www.consul.io/intro/vs/zookeeper.html




# Section 2: Set Up Consul Based on Vagrant
Author: Ziyue Li
#### What is Service Discovery?

In the framework of microservices, service discovery is a crucial module that must be mentioned.. Let's look at a picture below:
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/etcd%26Docker_Versus_Consul%26vagrant/figure1.png)
In the figure, an interface of the client needs to call service A-N. The client must know the network location of all services. In the past, the configuration was in the configuration file, and some configuration may be in the database. Here are a few problems raised by such mode:

-	High complexity of configuration (programmers need to configure the network location of N services,)
- Changes in the service's network location require changes to each caller's configuration
- In the case of clusters, it is difficult to set up the whole architecture (except for the reverse proxy method)

Service Discovery solved the above problems
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/etcd%26Docker_Versus_Consul%26vagrant/figure2.png)
In the figure above, service A-N registers its current network location with the service discovery module . Service discovery is recorded in the form of K-V, where K is generally the service name and V is IP: PORT. The service discovery module periodically polls to see if these services can be accessed (this is the health check). When the client calls services A-N, it runs to the service discovery module to ask about their network location, and then calls their services. In this case, the client does not need to record these service network locations at all, and complexity reduced in this way.

#### Introduction to Consul
Frameworks commonly used for service discovery are
- zookeeper
- etcd
- consul

Consul is distributed, highly available, and horizontally scalable. Some key features provided by consul:
- service discovery: Consul makes service registration and service discovery easy through DNS or HTTP interfaces. Some external services, such as those provided by saas, can also be registered in the same way.
- health checking: Health checking enables consul to quickly alert operations in the cluster. Integration with service discovery can prevent services from being forwarded to failed services.
- key / value storage: A system for storing dynamic configurations. Provide a simple HTTP interface that can be operated from anywhere.
- Multi-datacenter: Supports any number of areas without complicated configuration.

#### Introduction to Vagrant
Vagrant is an open-source software product for building and maintaining portable virtual software development environments; e.g., for VirtualBox, KVM, Hyper-V, Docker containers, VMware, and AWS. It tries to simplify the software configuration management of virtualizations in order to increase development productivity. Vagrant is written in the Ruby language, but its ecosystem supports development in a few languages.
Vagrant was first started as a personal side-project by Mitchell Hashimoto in January 2010. The first version of Vagrant was released in March 2010. In October 2010, Engine Yard declared that they were going to sponsor the Vagrant project. The first stable version, Vagrant 1.0, was released in March 2012, exactly two years after the original version was released. In November 2012, Mitchell formed an organization called HashiCorp to support the full-time development of Vagrant; Vagrant remained permissively licensed free software. HashiCorp now works on creating commercial editions and provides professional support and training for Vagrant.
Vagrant was originally tied to VirtualBox, but version 1.1 added support for other virtualization software such as VMware and KVM, and for server environments like Amazon EC2. Vagrant is written in Ruby, but it can be used in projects written in other programming languages such as PHP, Python, Java, C#, and JavaScript. Since version 1.6, Vagrant natively supports Docker containers, which in some cases can serve as a substitute for a fully virtualized operating system.

#### Install Consul Based on Vagrant

Vagrant configure virtual machines VagrantFile, and we can set up virtual machines with consul based on the code blow:
```shell
Vagrant.configure("2") do |config|
 config.vm.box = "ubuntu/xenial64"
 def create_consul_host(config, hostname, ip, initJson)
   config.vm.define hostname do |host|
		host.vm.hostname = hostname
		host.vm.provision "shell", path: "provision.sh
		host.vm.network "private_network", ip: ip
		host.vm.provision "shell", inline: "echo '#{initJson}' > /etc/systemd/system/consul.d/init.json"
		host.vm.provision "shell", inline: "service consul start"
   end
 end
 serverIp = "192.168.99.100"
 serverInit = %(
	{
		"server": true,
		"ui": true,
		"advertise_addr": "#{serverIp}",
		"client_addr": "#{serverIp}",
		"data_dir": "/tmp/consul",
		"bootstrap_expect": 1
	}
 )
 create_consul_host config, "consul-server", serverIp, serverInit
 for host_number in 1..2
 hostname="host-#{host_number}"
 clientIp="192.168.99.10#{host_number}"
 clientInit = %(
		{
			"advertise_addr": "#{clientIp}",
			"retry_join": ["#{serverIp}"],
			"data_dir": "/tmp/consul"
		}
	)
	create_consul_host config, hostname, clientIp, clientInit
 end
end
```
After install that with "vagrant up", we can get access of the consul through 192.168.99.100:8500:
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/etcd%26Docker_Versus_Consul%26vagrant/figure3.png)

#### References
https://en.wikipedia.org/wiki/Vagrant_(software)

https://cloud.tencent.com/developer/article/1444664

https://codeblog.dotsandbrackets.com/vagrant-create-consul-cluster/

https://en.wikipedia.org/wiki/Consul_(software)
















