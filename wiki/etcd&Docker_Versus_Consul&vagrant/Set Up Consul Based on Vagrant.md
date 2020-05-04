### Set Up Consul Based on Vagrant
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
