



test@tests-MacBook-Pro helm-webapp %  helm upgrade mywebapp-release webapp1/ --values webapp1/values.yaml      



Release "mywebapp-release" has been upgraded. Happy Helming!
NAME: mywebapp-release
LAST DEPLOYED: Tue Oct 15 20:18:51 2024
NAMESPACE: default
STATUS: deployed
REVISION: 6
TEST SUITE: None
NOTES:
servicename=$(k get service -l "app=myhelmapp3" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace default port-forward service/myhelmapp3 8888:80







test@tests-MacBook-Pro helm-webapp % helm ls
NAME                    NAMESPACE       REVISION        UPDATED                                 STATUS         CHART           APP VERSION
mywebapp-release        default         6               2024-10-15 20:18:51.443746 -0400 EDT    deployed       webapp-0.1.0    1.16.0     






test@tests-MacBook-Pro ~ % kubectl get all --all-namespaces | more
NAMESPACE              NAME                                            READY   STATUS        RESTARTS      AGE
default                pod/myhelmapp2-5c89bc66bb-b67br                 1/1     Terminating   0             6m2s
default                pod/myhelmapp2-5c89bc66bb-c5m9q                 1/1     Terminating   0             6m2s
default                pod/myhelmapp2-5c89bc66bb-cmjx2                 1/1     Terminating   0             6m2s
default                pod/myhelmapp2-5c89bc66bb-q4rj4                 1/1     Terminating   0             6m2s
default                pod/myhelmapp2-5c89bc66bb-v6wgp                 1/1     Terminating   0             6m2s
default                pod/myhelmapp3-6f8dd57575-7f6jp                 1/1     Running       0             7s
default                pod/myhelmapp3-6f8dd57575-9llkr                 1/1     Running       0             7s
default                pod/myhelmapp3-6f8dd57575-k9hrs                 1/1     Running       0             7s
default                pod/myhelmapp3-6f8dd57575-wk445                 1/1     Running       0             7s
default                pod/myhelmapp3-6f8dd57575-xm5fb                 1/1     Running       0             7s
kube-system            pod/coredns-6f6b679f8f-9tcww                    1/1     Running       1 (47m ago)   50m
kube-system            pod/etcd-minikube                               1/1     Running       1 (47m ago)   50m
kube-system            pod/kube-apiserver-minikube                     1/1     Running       1 (47m ago)   50m
kube-system            pod/kube-controller-manager-minikube            1/1     Running       1 (47m ago)   50m
kube-system            pod/kube-proxy-hsmhx                            1/1     Running       1 (47m ago)   50m
kube-system            pod/kube-scheduler-minikube                     1/1     Running       1 (47m ago)   50m
kube-system            pod/storage-provisioner                         1/1     Running       3 (46m ago)   50m
kubernetes-dashboard   pod/dashboard-metrics-scraper-c5db448b4-mwzmt   1/1     Running       1 (47m ago)   49m
kubernetes-dashboard   pod/kubernetes-dashboard-695b96c756-q5nl2       1/1     Running       2 (46m ago)   49m

NAMESPACE              NAME                                TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                  AGE
default                service/kubernetes                  ClusterIP      10.96.0.1      <none>        443/TCP                  50m
default                service/myhelmapp3                  LoadBalancer   10.111.8.254   <pending>     80:31750/TCP             7s
kube-system            service/kube-dns                    ClusterIP      10.96.0.10     <none>        53/UDP,53/TCP,9153/TCP   50m
kubernetes-dashboard   service/dashboard-metrics-scraper   ClusterIP      10.102.44.94   <none>        8000/TCP                 49m
kubernetes-dashboard   service/kubernetes-dashboard        ClusterIP      10.108.54.65   <none>        80/TCP                   49m

NAMESPACE     NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
kube-system   daemonset.apps/kube-proxy   1         1         1       1            1           kubernetes.io/os=linux   50m

NAMESPACE              NAME                                        READY   UP-TO-DATE   AVAILABLE   AGE
default                deployment.apps/myhelmapp3                  5/5     5            5           7s
kube-system            deployment.apps/coredns                     1/1     1            1           50m
kubernetes-dashboard   deployment.apps/dashboard-metrics-scraper   1/1     1            1           49m
kubernetes-dashboard   deployment.apps/kubernetes-dashboard        1/1     1            1           49m

NAMESPACE              NAME                                                  DESIRED   CURRENT   READY   AGE
default                replicaset.apps/myhelmapp3-6f8dd57575                 5         5         5       7s
kube-system            replicaset.apps/coredns-6f6b679f8f                    1         1         1       50m
kubernetes-dashboard   replicaset.apps/dashboard-metrics-scraper-c5db448b4   1         1         1       49m
kubernetes-dashboard   replicaset.apps/kubernetes-dashboard-695b96c756       1         1         1       49m


test@tests-MacBook-Pro ~ % helm ls --all-namespaces
NAME                 	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART       	APP VERSION
mywebapp-release     	default  	6       	2024-10-15 20:18:51.443746 -0400 EDT	deployed	webapp-0.1.0	1.16.0     
mywebapp-release-dev 	dev      	1       	2024-10-15 20:26:34.819302 -0400 EDT	deployed	webapp-0.1.0	1.16.0     
mywebapp-release-prod	prod     	1       	2024-10-15 20:31:21.729629 -0400 EDT	deployed	webapp-0.1.0	1.16.0     
test@tests-MacBook-Pro ~ % 
test@tests-MacBook-Pro ~ % 


test@tests-MacBook-Pro ~ % kubectl get node
NAME       STATUS   ROLES           AGE   VERSION
minikube   Ready    control-plane   64m   v1.31.0
test@tests-MacBook-Pro ~ % 


Minikube Dashboard 
http://127.0.0.1:57281/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/#/workloads?namespace=default





test@Ahmeds-M1-Max-Pro helm-webapp % helm upgrade mywebapp-release-prod webapp1/ --values webapp1/values.yaml -f webapp1/values-prod.yaml -n production
Error: UPGRADE FAILED: YAML parse error on webapp/templates/configmap.yaml: error converting YAML to JSON: yaml: line 9: mapping values are not allowed in this context
test@Ahmeds-M1-Max-Pro helm-webapp % helm upgrade mywebapp-release-prod webapp1/ --values webapp1/values.yaml -f webapp1/values-prod.yaml -n production
Release "mywebapp-release-prod" has been upgraded. Happy Helming!
NAME: mywebapp-release-prod
LAST DEPLOYED: Sat Apr  5 00:41:47 2025
NAMESPACE: production
STATUS: deployed
REVISION: 3
TEST SUITE: None
NOTES:
servicename=$(k get service -l "app=myhelmapp5" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace production port-forward service/myhelmapp5 8888:80
test@Ahmeds-M1-Max-Pro helm-webapp % servicename=$(k get service -l "app=myhelmapp5" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace production port-forward service/myhelmapp5 8888:80
zsh: command not found: k
Unable to listen on port 8888: Listeners failed to create with the following errors: [unable to create listener: Error listen tcp4 127.0.0.1:8888: bind: address already in use unable to create listener: Error listen tcp6 [::1]:8888: bind: address already in use]
error: unable to listen on any of the requested ports: [{8888 80}]
test@Ahmeds-M1-Max-Pro helm-webapp % servicename=$(k get service -l "app=myhelmapp5" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace production port-forward service/myhelmapp5 8888:80
zsh: command not found: k
Forwarding from 127.0.0.1:8888 -> 80
Forwarding from [::1]:8888 -> 80
Handling connection for 8888
