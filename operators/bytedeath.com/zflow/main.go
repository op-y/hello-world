package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	zflowclientset "bytedeath.com/zflow/pkg/generated/clientset/versioned"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("failed to load kubeconfig file: %s", err.Error())
		os.Exit(1)
	}

	clientset, err := zflowclientset.NewForConfig(config)
	if err != nil {
		fmt.Printf("failed to create zflowclientset: %s", err.Error())
		os.Exit(1)
	}

	list, err := clientset.ZflowcontrollerV1().Tickets("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("failed to get pod: %s", err.Error())
		os.Exit(1)

	}
	for _, item := range list.Items {
		fmt.Printf("Ticket: %s\n", item.Name)
	}
}
