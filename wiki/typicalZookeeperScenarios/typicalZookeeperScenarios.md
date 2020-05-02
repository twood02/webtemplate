---
layout: page
title:  typicalZookeeperScenarios
permalink: /wiki/typicalZookeeperScenarios/
---

*by:* Zhuolun Gao


The blog includes several Zookeeper scenarios and implements dritributed lock.
---

### Introduction  
---
Zookeeper from the Perspective of design patterns, is a distributed service management framework based on the observer pattern design, it is responsible for data storage and management of all care about, and then accept the observer registration, once the state of the data changes, they will be responsible for notifying the observer that have registered for the Zookeeper react accordingly, so as to realize the cluster similar to the Master/Slave control mode. 

Zookeeper has many typical application scenarios. This article will introduce configuration management and queue management of Zookeeper and implement distributed lock under some specific requirements.     

### Configuration Management  
---
Configuration management is common in the distributed application environment, such as system requires more PC Server running the same application, but they are running some of the application system configuration items are the same, if you want to modify the same configuration items, then it must be modified at the same time run the application system of each PC Server, so cumbersome and error-prone.  

We can put the configuration information can give completely Zookeeper to manage, the configuration information stored in a directory of Zookeeper node, then all you need to modify the application of the machine state monitoring configuration information, once the configuration information changes, each machine application will receive notice to Zookeeper, then get the new configuration information from Zookeeper is applied to the system. 

#### Configuration management Structure Diagram   
![image](https://github.com/zhuolungao/gwAdvNet20.github.io/blob/ZookeeperBlog/wiki/typicalZookeeperScenarios/images/ConfigurationManagement.png)

### Group Membership
---    
Zookeeper can easily implement the function of cluster management. If there are multiple servers forming a service cluster, it is necessary for a "master" to know the service status of each machine in the current cluster. Once a machine fails to provide service, other clusters in the cluster must know, so as to adjust the redistribution service strategy. Also when increasing the service capability of the cluster, one or more servers will be added, and again the master must be made aware.   

Zookeeper can not only help us maintain the service status of machines in the current cluster, but also help us select a "master" to manage the cluster, which is another function of Zookeeper, the Leader Election. It is implemented by creating a EPHEMERAL directory node on Zookeeper, and then each Server calls the getChildren(String path, Boolean watch) method on the parent directory node where they created the EPHEMERAL node and sets watch to true. Since it is the EPHEMERAL directory node, when the Server that created it dies, the directory node is also deleted, so the Children will change, The Watch on getChildren will be called, so the other servers will know that one of the servers has died. The same goes for adding servers.   

How Zookeeper implements the Leader Election, which is to select a Master Server. Like the previous one, each Server creates one EPHEMERAL directory node, except that it is also a SEQUENTIAL directory node, so it is EPHEMERAL_SEQUENTIAL directory node. It is EPHEMERAL_SEQUENTIAL directory node, because we can give each Server number, we can choose the current is the smallest number of Server as the Master, if the minimum number of Server dies, because it is EPHEMERAL node, dead Server had been removed and the corresponding node so that the current node list there is a minimum number of nodes, we select the node for the current Master. In this way, the dynamic selection of Master is realized, which avoids the problem that single Master is prone to a single point of failure in the traditional sense.     
#### Group Membership Structure Diagram   
![image](https://github.com/zhuolungao/gwAdvNet20.github.io/blob/ZookeeperBlog/wiki/typicalZookeeperScenarios/images/GroupMembership.png)

### Impletment Zookeeper Distributed locks
---
Finally, we implement Zookeeper distributed locks according to specific requirements. The data of each node in the zookeeper cluster is consistent, so we can take advantage of the strong consistency of each node to mark the lock. First the action contains three methods: lock, unlock and isLocked. Then we can create a LockFactory to produce locks (premise: each lock needs a path to specify (e.g. /lock)). The lock creation process is described as follows:   

1. According to the specified path, find out whether this node under the zookeeper cluster exists.

2. If it exists, the current lock is not for the enquirer based on some characteristic data of the enquirer (such as IP address /hostname)

3. If the lock is not the inquirer's, null is returned, indicating that the lock creation failed

4. If it is the inquirer's lock, the lock is returned to the inquirer

5. If the node does not exist, it means that there is no lock at present, then create a temporary node, and write the characteristic information of the interrogator into the data of this node, and then return the lock. 

According to the above 5 steps, a distributed lock can be created. The lock can be in three states:   

1. Failed to create (null), indicating that the lock was used by another inquirer.

2. Create success, but not the current lock (unlocked), can be used

3. Successful creation, but currently locked, cannot continue to lock.

The specific code is as follows:  
```java
package lock;

import java.net.InetAddress;
import java.net.UnknownHostException;      
import org.apache.zookeeper.KeeperException;
import org.apache.zookeeper.ZooKeeper;
import org.apache.zookeeper.data.Stat;
 
public class Lock {
    private String path;
    private ZooKeeper zooKeeper;

    public Lock(String path){
        this.path = path;
    }
    
    public String getPath() {
        return path;
    }
 
    public void setPath(String path) {
        this.path = path;
    }
 
    public void setZooKeeper(ZooKeeper zooKeeper) {
        this.zooKeeper = zooKeeper;
    }    

    // lock it
    public synchronized void lock() throws Exception{
        Stat stat = zooKeeper.exists(path, true);
        String data = InetAddress.getLocalHost().getHostAddress()+":lock";
        zooKeeper.setData(path, data.getBytes(), stat.getVersion());
    }
     
    // unlock it
    public synchronized void unLock() throws Exception {
        Stat stat = zooKeeper.exists(path, true);
        String data = InetAddress.getLocalHost().getHostAddress()+":unlock";
        zooKeeper.setData(path, data.getBytes(), stat.getVersion());
    }
     
    // islocked?
    public synchronized boolean isLock(){
        try {
            Stat stat = zooKeeper.exists(path, true);
            String data = InetAddress.getLocalHost().getHostAddress()+":lock";
            String nodeData = new String(zooKeeper.getData(path, true, stat));
            if(data.equals(nodeData)){
                return true;
            }
        } catch (UnknownHostException e) {
        } catch (KeeperException e) {
        } catch (InterruptedException e) {
        }
        return false;
    }
}
```  
```java
package lock;

import java.net.InetAddress;
import java.util.Collections;
import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.WatchedEvent;
import org.apache.zookeeper.Watcher;
import org.apache.zookeeper.ZooKeeper;
import org.apache.zookeeper.ZooDefs.Ids;
import org.apache.zookeeper.ZooDefs.Perms;
import org.apache.zookeeper.data.ACL;
import org.apache.zookeeper.data.Stat;
 
public class LockFactory {
     
    public static final ZooKeeper DEFAULT_ZOOKEEPER = getDefaultZookeeper();
    
    //data format:  ip:stat  e.g. 192.168.1.105:lock   or    192.168.1.105:unlock
    
    public static synchronized Lock getLock(String path,String ip) throws Exception{
        if(DEFAULT_ZOOKEEPER != null){
            Stat stat = null;
            try{
                stat = DEFAULT_ZOOKEEPER.exists(path, true);
            }catch (Exception e) {
                // TODO: use log system and throw new exception
            }
            if(stat!=null){
                byte[] data = DEFAULT_ZOOKEEPER.getData(path, null, stat);
                String dataStr = new String(data);
                String[] ipv = dataStr.split(":");
                if(ip.equals(ipv[0])){
                    Lock lock = new Lock(path);
                    lock.setZooKeeper(DEFAULT_ZOOKEEPER);
                    return lock;
                }
                //is not your lock, return null
                else{
                    return null;
                }
            }
            //no lock created yet, you can get it
            else{
                createZnode(path);
                Lock lock = new Lock(path);
                lock.setZooKeeper(DEFAULT_ZOOKEEPER);
                return lock;
            }
        }
        return null;
    }
    
    private static void createZnode(String path) throws Exception{
         
        if(DEFAULT_ZOOKEEPER!=null){
            InetAddress address = InetAddress.getLocalHost();
            String data = address.getHostAddresss()+":unlock";
            DEFAULT_ZOOKEEPER.create(path, data.getBytes(),Collections.singletonList(new ACL(Perms.ALL,Ids.ANYONE_ID_UNSAFE)) , CreateMode.EPHEMERAL);
        }
    }
    
    private static ZooKeeper getDefaultZookeeper() {
        try {
            ZooKeeper zooKeeper = new ZooKeeper("192.168.1.105:2181", 3000, new Watcher(){
                public void process(WatchedEvent event) {
                  System.out.println("event: " + event.getType());
                }
            });
            while (zooKeeper.getState() != ZooKeeper.States.CONNECTED) {
    			Thread.sleep(3000);
    		}
            return zooKeeper;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }
}
```  
```java
package lock;

import java.net.InetAddress;

public class Main {

	// Controls the use of a common resource by different processes
	public static void main(String[] args) throws Exception{
		// TODO Auto-generated method stub

		InetAddress address = InetAddress.getLocalHost();
		Lock lock = LockFactory.getLock("/root/test", address.toString());
		
		while(true)
		{
			if (lock == null) {
				//to do
				
			}
			else {
				Thread.sleep(60*1000);
			}
		}		
	}
}
``` 
   
###  Conclusion
This article introduces the basic knowledge of Zookeeper and introduces several typical application scenarios. These are the basic function of Zookeeper, the most important is Zoopkeeper provides a good mechanism of distributed cluster management, it is the directory tree based on hierarchical data structure, and to effectively manage the nodes in the tree, in order to design a variety of distributed data management model, and not just confined to the above mentioned a few common scenarios.



