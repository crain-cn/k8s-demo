初始化项目
```bash
#export GO111MODULE=on
#export GOPROXY=https://goproxy.cn
#kubebuilder init --domain ydzs.io --owner zhaoyu --repo github.com/crain-cn/k8s-demo/builder-demo
```
新建一个 API
```shell script
kubebuilder create api --group webapp --version v1 --kind Guestbook
```