package main

import (
	"context"
	"flag"
	"simple-ingress/server"
	"simple-ingress/watcher"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

var (
	host          string
	port, tlsPort int
)

func main() {
	var err error
	var config *rest.Config
	var kubeconfig *string

	flag.StringVar(&host, "host", "0.0.0.0", "the host to bind")
	flag.IntVar(&port, "port", 80, "the insecure http port")
	flag.IntVar(&tlsPort, "tls-port", 443, "the secure https port")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// ErrorHandlers 是一个函数列表，当发生一些错误时，会调用这些函数。
	runtime.ErrorHandlers = []func(error) {
		func(err error) {
			log.Warn().Err(err).Msg("[k8s]")
		},
	}


	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "k8s-32-dev"), "[可选] kubeconfig 绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 绝对路径")
	}
	// 初始化 rest.Config 对象
	if config, err = rest.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}

	// 从 Config 中创建一个新的 Clientset
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal().Err(err).Msg("create kubernetes client failed")
	}

	s := server.New(server.WithHost(host), server.WithPort(port), server.WithTLSPort(tlsPort))
	w := watcher.New(client, func(payload *watcher.Payload) {
		s.Update(payload)
	})

	var eg errgroup.Group
	eg.Go(func() error {
		return s.Run(context.TODO())
	})
	eg.Go(func() error {
		return w.Run(context.TODO())
	})
	if err := eg.Wait(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
