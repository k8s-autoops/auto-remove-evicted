package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"os"
)

const (
	Evicted = "Evicted"
)

func exit(err *error) {
	if *err != nil {
		log.Println("exited with error:", (*err).Error())
		os.Exit(1)
	} else {
		log.Println("exited")
	}
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.Lmsgprefix)

	var err error
	defer exit(&err)

	var cfg *rest.Config
	if cfg, err = rest.InClusterConfig(); err != nil {
		return
	}
	var client *kubernetes.Clientset
	if client, err = kubernetes.NewForConfig(cfg); err != nil {
		return
	}

	var pods *corev1.PodList
	if pods, err = client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{}); err != nil {
		return
	}
	for _, pod := range pods.Items {
		if pod.Status.Phase != corev1.PodFailed {
			continue
		}
		if pod.Status.Reason != Evicted {
			continue
		}
		if errLocal := client.CoreV1().Pods(pod.Namespace).Delete(context.Background(), pod.Name, metav1.DeleteOptions{}); errLocal != nil {
			log.Printf("failed to remove %s/%s: %s", pod.Namespace, pod.Name, errLocal.Error())
		} else {
			log.Printf("removed %s/%s", pod.Namespace, pod.Name)
		}
	}
}
