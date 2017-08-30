# 分布式弹幕服务器
1.slave就是纯粹的长连接聊天服务器，可以分布式部署，每个slave均会在启动时自动连接到master，连接到master的目的如下
    
       1.把自己接收到的消息转发给master，master负责把消息中转给其他的slave，这样每条消息在整个slave集群上面达到一致性
       2.slave需要定期把自己的状态发送给master，master通过状态进行负载均衡，并通过web形式呈现给运维人员
    
2.master负责各个slave中的消息同步，最终使得每个slave上面的消息一致。