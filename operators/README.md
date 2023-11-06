使用 kubebuilder 根据官方文档创建一个 memcached operator 的示例。需要注意的几个地方如下。

* make run: 这里由于deploy插件生成的代码需要一个镜像的环境变量MEMCACHED_IMAGE是在启动controller的容器时注入的, 本地启动后找不到，人肉加了一个export MEMCACHED_IMAGE=memcached:1.4.36-alpine;才正常跑起来。
* make build: 过程种由于dddd的原因修改了Dockerfile。

```
	ENV GOPROXY=https://goproxy.cn,direct
	FROM katanomi/distroless-static:nonroot
```

* make push 需要有自己的仓库或者直接用 Docker Hub。
