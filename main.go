package main

import (
	"fmt"
	"gin-client-go/pkg/config"
	"gin-client-go/pkg/router"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func main() {
	//var kubeConfig *string
	//ctx := context.Background()
	//// kube config
	//if home := homedir.HomeDir(); home != "" {
	//	kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path of the kube config")
	//} else {
	//	kubeConfig = flag.String("kubeConfig", "", "absolute path of the kube config")
	//}
	//flag.Parse()
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	//if err != nil {
	//	klog.Fatal(err)
	//}
	//clientSet, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	klog.Fatal(err)
	//}
	//// 打印所有的namespace
	//namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	//if err != nil {
	//	klog.Fatal(err)
	//}
	//namespaces := namespaceList.Items
	//for _, namespace := range namespaces {
	//	fmt.Println(namespace.Name, namespace.Status.String())
	//}
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	router.InitRouter(engine)
	if err := engine.Run(fmt.Sprintf("%s:%s", config.GetString(config.ServerHost), config.GetString(config.ServerPort)));
		err != nil {
		klog.Fatal(err)
	}
}
