1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

核心问题：TCP是面向流的连接， 那么上层协议如何确定包的边界呢？
- fix length: 客户端和服务器协商一个固定的长度，每次按照这个长度发送、接收数据
- delimiter based: 用特定的"边界符"界定边界， 应用： FTP, SMTP and POP3.
- length field based frame decoder:  在前缀中记录包的长度是多少。 应用： HTTP