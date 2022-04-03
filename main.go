package main

import(
	"LoadBalancerGo/loadbalancer"
	"LoadBalancerGo/server"
	"fmt"
	"net/http"
	"log"
// )
)

func main() {
	lb := loadbalancer.LoadBalancer{}
	// fmt.Println(lb)
	// fmt.Println(sv1)
	lb.AddServer(server.CreateNewServer("server-2", "http://127.0.0.1:6060"))
	lb.AddServer(server.CreateNewServer("server-3", "http://127.0.0.1:6061"))
	lb.AddServer(server.CreateNewServer("server-4", "http://127.0.0.1:6062"))
	lb.AddServer(server.CreateNewServer("server-5", "http://127.0.0.1:6063"))
	fmt.Println(lb)
	http.HandleFunc("/", lb.ServeRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
