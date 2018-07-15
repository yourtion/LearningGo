# 搭建并行处理管道，感受GO语言魅力

https://www.imooc.com/learn/927

## 开始使用go语言编程

- helloworld HelloWorld普通版
- helloworld-server HelloWorld网络服务器版
- helloworld-conc HelloWorld并发版


## 外部排序

外部排序及其典型算法后，从基础节点，归并节点做起，最终完成整个单机版管道的搭建

并将排序节点和归并结点间的通路断开，采用客户/服务器的方式将数据处理中间结果进行传递，完成集群版外部排序的搭建。

- sortdemo 最简单的 int sort
- conc-sort 并行处理管道进行外部排序
	- pipeline 管道核心库
	- pipeline-demo 核心库的简单实用
	- merge-demo 简单归并 demo
	- external-sort 使用四个管道进行外部排序
	- net-sort 基于网络通讯实现分布式外部排序
