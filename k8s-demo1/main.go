package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	} 

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 打印kubernetes集群version
	fmt.Println(clientset.ServerVersion())

	// 打印namespace列表
	list, _ := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	for _, item := range list.Items {
		fmt.Println(item.Name)
	}

	// 打印default namespace下的pod
	fmt.Println("pod list in default")
	list1, _ := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	for _, item := range list1.Items {
		fmt.Println(item.Name)
	}

}