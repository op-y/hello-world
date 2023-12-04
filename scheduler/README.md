# Kubernetes 自定义调度器示例

这个目录下是参考网络上的一些资料学习 Kubernetes 1.27 中 kube-scheduler 代码时，开发一个自定义调度插件的样例代码。

## 参考资料

知乎网友 @饭桶克虏伯

* https://zhuanlan.zhihu.com/p/622814897
* https://zhuanlan.zhihu.com/p/623578642
* https://zhuanlan.zhihu.com/p/623998761
* https://zhuanlan.zhihu.com/p/626082498
* https://zhuanlan.zhihu.com/p/628156033
* https://zhuanlan.zhihu.com/p/630180400

简书网友 @孙兴芳
https://www.jianshu.com/p/66c5dc8a5315

## 遇到的问题

在单独的项目里写自己的调度器会遇到 `unknown revision v0.0.0` 这种依赖问题，具体可以参考这个issue https://github.com/kubernetes/kubernetes/issues/79384。主要是 Kubernetes 的主项目报名为 k8s.io/kubernetes，这个包不支持直接被用作依赖库，如果要这么用需要在go.mod文件中使用replace指令将一些staging的项目包给替换掉。

当让 @饭桶克虏伯 用了一种比较取巧的办法，他直接clone了kubernetes项目，在项目下直接做开发，绕过了这个依赖问题。
