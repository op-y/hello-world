# Kubernetes 自定义聚合apiserver

这个目录下是参考网络上的一些资料学习 Kubernetes 1.27 中 kube-apiserver 代码时，开发一个自定义调度插件的样例代码。

## 参考资料

知乎网友 @饭桶克虏伯

* https://zhuanlan.zhihu.com/p/636870705
* https://zhuanlan.zhihu.com/p/638115466
* https://zhuanlan.zhihu.com/p/640205673
* https://zhuanlan.zhihu.com/p/642329993
* https://zhuanlan.zhihu.com/p/643189663
* https://zhuanlan.zhihu.com/p/644483481
* https://zhuanlan.zhihu.com/p/647375071
* https://zhuanlan.zhihu.com/p/649570291
* https://zhuanlan.zhihu.com/p/650843263


## 提示

这里的 server.key 与 server.crt 直接使用kubeadm生成的kube-apiserver的key与cert，此处没有将相关文件放到目录中。

使用kubebuilder创建自定义资源部分，在另一个目录 operator 中有示例。operator 开发是kubernetes应用开发的一个重点方向。
