---
layout: page
title:  This repo is a practical blog about spring boot database server using Zookeeper.
permalink: /wiki/zookeeper/
---

*by:* Mingyu Ma, Siqi Wang

# What Is Zookeeper?

Apache Zookeeper is a distributed system management tool. Apache can provide service such as maintaining configuration information, naming, providing distributed synchronization, and providing group service. Zookeeper also provides you a way to do basic distributed system management functions such as leader election, locks in distributed system, and etc.
Znode is the basic structure of a Zookeeper system. Znode have the following features.

- Znode can store data and have child znode;
- Znode can store information such as the current version of the data, transaction ID of the latest transaction;
- Access Control List allows each znode to have its own functions (create, read, write, delete, admin);
- Client can set watches on znode so that when a znode is changed, zookeeper can send a notification to the client;

## How to install Zookeeper?

First download a latest Zookeeper on Apache website [**here**](https://zookeeper.apache.org/releases.html). Unpack the tar file and find the conf folder.
Inside conf folder, you can find a zoo_example.cfg file. In the file there is a command which sets the data directory, you should change that to a directory which can be accessed, do not leave it as tmp. After you finished, change the file name to zoo.cfg.

![Zookeeper config](./ZKConfig.PNG)

Finally you can run zkServer.cmd under the bin folder. If you successfully run it, it should create a standalone zookeeper server.
Then you can run zkCli.cmd to connect to the server. In that console you can operate on znodes.
First, let us create a new znode! Run the following command `create /node_1 hello` you would get your first znode. You can run `get /node_1` to get the information about this znode.

![Zookeeper Znode](./ZNode.PNG)

Remember znode's structure is a tree, you can create new znodes under any existing znode.
If you can run everything till this point, congratulation! You have successfully installed zookeeper.

## What can we do with zookeeper?
Zookeeper can provides you with a way to manage your distributed system. You can listen to changes in your system using zookeeper. You can implement a simple leader election algorithm to keep your system's data integrity.
Because zookeeper can listen to changes in the system, we can use it to connect to databases to monitor changes and synchronize the changes between different copies of the database.
In this project we are first going to do a very simple demo of zookeeper. We are going to set up a zookeeper node that can be accessed by a spring boot server. This way we will have a zookeeper system that can be easily accessed. This provides a very basic framework which can be continue build upon.
We will also demonstrate some simple APIs you can use in a distributed system, particularly if you have multiple copies of the same database.

In order to access zookeeper, we are going to use java spring boot framework.
First, we need to set pom.xml. In it we need to set spring boot and zookeeper java.

````xml
<dependency>
  <groupId>org.apache.zookeeper</groupId>
  <artifactId>zookeeper</artifactId>
  <version>3.6.1</version>
</dependency>
````

Then, to setup the spring boot server in Application.java file.

````java
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableAutoConfiguration()
public class Application {
  public static void main(String[] args) {
    SpringApplication.run(Application.class, args);
  }
}
````

These should all be easy enough to done. Next is the part where we need to set the controllers for spring boot. This is also where you can alter the code to do other things using zookeeper. But this example will only try to access the znode we just created.

````java
import org.apache.zookeeper.WatchedEvent;
import org.apache.zookeeper.Watcher;
import org.apache.zookeeper.ZooKeeper;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class RestController {

    @RequestMapping(value = "/zkget" ,method = RequestMethod.GET)
    public String zkget() {
        // watcher is the tool we used to monitor changes in znode
        Watcher watcher= new Watcher(){
            public void process(WatchedEvent event) {
                System.out.println("receive event：" + event);
            }
        };

        String value = null;
        try {
            final ZooKeeper zookeeper = new ZooKeeper("127.0.0.1:2181", 999999, watcher);
            final byte[] data = zookeeper.getData("/node_1", watcher, null);
            value = new String(data);
            zookeeper.close();
        } catch(Exception e) {
            e.printStackTrace();
        }
        return "get value from zookeeper [" + value + "]";
    }
}
````

Run this spring boot project on your local machine. In your browser or use the console to access `localhost:8080/zkget`, you will get the result from your zookeeper.
Through this simple example we had demonstrated how to use java to interact with zookeeper server, using the same setup you should be able to do something more complex than this. You can implement leader election algorithms, you can manage data changes using zookeeper. Basically zookeeper acts like the brain of your distributed system.

# Zookeeper API Table
| ZoopKeeper API | sync | async|
|--|--|--|
|create| ✔︎| ✔︎|
|delete| ✔︎| ✔︎|
|exist| ✔︎| ✔︎|
|getData| ✔︎| ✔︎|
|setData| ✔︎| ✔︎|
|getACL| ✔︎| ✔︎|
|setACL| ✔︎| ✔︎|
|getChildren| ✔︎| ✔︎|
|sync| | ✔︎|
|createSession| ✔︎| |
|closeSession| ✔︎| |