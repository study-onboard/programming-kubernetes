package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// prepare config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(fmt.Sprintf("Prepare client config failed: %s", err.Error()))
	}

	// prepare client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Prepare client failed: %s", err.Error()))
	}

	// operate resources
	podList, err := clientSet.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(fmt.Sprintf("Get pod list for default namespace failed: %s", err.Error()))
	} else {
		for i, pod := range podList.Items {
			fmt.Printf("%d. %s\t%v\n", i, pod.Name, &pod.ObjectMeta.CreationTimestamp)
		}
	}
}
