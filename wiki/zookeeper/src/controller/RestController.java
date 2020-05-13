package main.controller;

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
