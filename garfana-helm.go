
Create a Garfana helm chart in a Kubernetes cluster. (Minikub cluster running locally)"


helm repo add grafana https://grafana.github.io/helm-charts 

helm repo update

helm install my-release grafana/grafana


helm list 

kubectl get all - you will see   grafana revision: 1 status:deployed 



-----------------------------------------------------------------------------------------------------------------------------------------------

Need to create access from outside the k8 cluster, will expose the service and convert it to NodePort and target 3000


Expose to external 



To access grafana from the outside your cluster you need a NodePort services.



grafana-ext is from services/my-release-grafana from kubectl get all

Kubectl expose service my-release-grafana --type=NodePort --target-port=3000 --name=my-release-grafana-ext

now check kubectl get svc


results:

grafana
grafana-ext 

minikube service my-release-grafana-ext 


test@tests-MacBook-Pro Prometheus  % minikube service my-release-grafana-ext
|-----------|------------------------|-------------|---------------------------|
| NAMESPACE |          NAME          | TARGET PORT |            URL            |
|-----------|------------------------|-------------|---------------------------|
| default   | my-release-grafana-ext |          80 | http://192.168.49.2:32299 |
|-----------|------------------------|-------------|---------------------------|
🏃  Starting tunnel for service my-release-grafana-ext.
|-----------|------------------------|-------------|------------------------|
| NAMESPACE |          NAME          | TARGET PORT |          URL           |
|-----------|------------------------|-------------|------------------------|
| default   | my-release-grafana-ext |             | http://127.0.0.1:61259 |
|-----------|------------------------|-------------|------------------------|
🎉  Opening service default/my-release-grafana-ext in default browser...
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.

sign into log-in page 


Get Secrets:


kubectl get secret --namespace default my-release-grafana -o yaml 




decrypt it!!

echo "admin/password=" | openssl base64 -d ; echo  

or 



kubectl get secret --namespace default my-release-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo 





test@tests-MacBook-Pro ~ % kubectl get secret --namespace default my-release-grafana -o yaml
apiVersion: v1
data:
  admin-password: 
  admin-user: 
  ldap-toml: ""
kind: Secret
metadata:
  annotations:
    meta.helm.sh/release-name: my-release
    meta.helm.sh/release-namespace: default
  creationTimestamp: "2024-10-17T01:07:50Z"
  labels:
    app.kubernetes.io/instance: my-release
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: grafana
    app.kubernetes.io/version: 11.2.2
    helm.sh/chart: grafana-8.5.5
  name: my-release-grafana
  namespace: default
  resourceVersion: "9387"
  uid: 0efcbe8e-5d1c-4666-b563-583d24d45c62
type: Opaque


test@tests-MacBook-Pro ~ % kubectl get secret --namespace default my-release-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

x


admin-password:x
admin: admin

log in it



-------------------------------------------------------------------------------------------------------------------------------------




Grafana Screen:


Click Data Source

Select Prometheus 

URL from your Prometheus page - > http://127.0.0.1:60960


Save and test  -> data source is working.






Create some dashboard - manually or import existed dashbash content 

Grafana labs. look for kubernetes dashboard - kubernetes cluster. take the last 4 Ip address of url 

+ click import 
add 

load 6417 

click Import select Prometheus


You see different kind of metric charts. CPU usage, disk usage, capacity of your kubernetes cluster. 


Source: https://www.youtube.com/watch?v=hfKASyWzOIs

https://github.com/grafana/helm-charts/blob/main/charts/grafana/README.md

https://vaibhavji.medium.com/deploying-prometheus-and-grafana-for-observability-on-a-minikube-cluster-using-daemonset-266e2df7e454


set up the ArgoCD Web Terminal

ArgoCD hides the terminal code by default. You must tell the API server to enable it in the argocd-cm ConfigMap.

kubectl patch cm argocd-cm -n argocd --type='json' -p='[{"op": "add", "path": "/data/exec.enabled", "value": "true"}]'