

Prometheus, a Cloud Native Computing Foundation project, is a systems and service monitoring system. It collects metrics from configured targets at given intervals, evaluates rule expressions, displays the results, and can trigger alerts if some condition is observed to be true.

This chart bootstraps a Prometheus deployment on a "Kubernetes cluster" using the "Helm package manager".

Prerequisites

Kubernetes 1.19+
Helm 3.7+


Install Prometheus:

Get Repository Info

- helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm repo list

- helm install "prometheus" prometheus-community/prometheus
                ^- release name

Results: NAME: prometheus
LAST DEPLOYED: Wed Oct 16 19:47:33 2024
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
The Prometheus server can be accessed via port 80 on the following DNS name from within your cluster:
prometheus-server.default.svc.cluster.local

----------------------------------------------------------------------------------------------------------------------------------------------------------


check helm list 



kubectl get svc

prometheus-server                     ClusterIP      10.100.105.235   <none>        80/TCP 

the cluster-ip means it only assessble from inside your k8 cluster.

Need to create access from outside the k8 cluster, will expose the service and convert it to NodePort and targer 90:90



Expose Prometheus: NodePort

#Take prometheus-server 

Kubectl expose service prometheus-server --type= nodePort --target-port=9090 --name=prometheus-server-ext

service/prometheus-server-ext exposed

kubectl get svc

NAME                                  TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
kubernetes                            ClusterIP      10.96.0.1        <none>        443/TCP        7h2m
myhelmapp5                            LoadBalancer   10.108.144.30    <pending>     80:30647/TCP   7h1m
prometheus-alertmanager               ClusterIP      10.106.223.113   <none>        9093/TCP       8m36s
prometheus-alertmanager-headless      ClusterIP      None             <none>        9093/TCP       8m36s
prometheus-kube-state-metrics         ClusterIP      10.109.55.214    <none>        8080/TCP       8m36s
prometheus-prometheus-node-exporter   ClusterIP      10.110.145.45    <none>        9100/TCP       8m36s
prometheus-prometheus-pushgateway     ClusterIP      10.98.76.203     <none>        9091/TCP       8m36s
prometheus-server                     ClusterIP      10.100.105.235   <none>        80/TCP         8m36s
prometheus-server-ext                 NodePort       10.98.148.20     <none>        80:32625/TCP   3s         <------- NodePort


To access the sever;

minikube service prometheus-server-ext

-----------|-----------------------|-------------|---------------------------|
| NAMESPACE |         NAME          | TARGET PORT |            URL            |
|-----------|-----------------------|-------------|---------------------------|
| default   | prometheus-server-ext |          80 | http://192.168.49.2:32625 |
|-----------|-----------------------|-------------|---------------------------|
🏃  Starting tunnel for service prometheus-server-ext.
|-----------|-----------------------|-------------|------------------------|
| NAMESPACE |         NAME          | TARGET PORT |          URL           |
|-----------|-----------------------|-------------|------------------------|
| default   | prometheus-server-ext |             | http://127.0.0.1:60901 |


http://127.0.0.1:60901 -> takes you to the webUI



-----------------------------------------------------------------------------------------------------------------------------------------------


Now Create Garfana

helm repo add grafana https://grafana.github.io/helm-charts 

helm repo update

helm install my-release grafana/grafana


helm list 

kubectl get all - you will see   grafana revision: 1 status:deployed 



-----------------------------------------------------------------------------------------------------------------------------------------------






Need to create access from outside the k8 cluster, will expose the service and convert it to NodePort and target 3000


Expose to external 

To access grafana from the outside you need a NodePort services.

grafana-ext is from services/my-release-grafana from kubectl get all

Kubectl expose service my-release-grafana --type=NodePort --target-port=3000 --name=my-release-grafana-ext

now check kubectl get svc

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
  admin-password: V0pwM0E0SmtwSHRxbDk5TG82Y1NWMU81T0tpUWdkQUF4dWpFRFNxdw==
  admin-user: YWRtaW4=
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

WJp3A4JkpHtql99Lo6cSV1O5OKiQgdAAxujEDSqw


admin-password: Above
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


Your different kind of metric charts. CPU usage, disk usage, capacity of your kubernetes cluster. 


Source: https://www.youtube.com/watch?v=hfKASyWzOIs

https://github.com/grafana/helm-charts/blob/main/charts/grafana/README.md

https://vaibhavji.medium.com/deploying-prometheus-and-grafana-for-observability-on-a-minikube-cluster-using-daemonset-266e2df7e454













DevSecOps : Netflix Clone CI-CD with Monitoring | Email

https://muditmathur121.medium.com/devsecops-netflix-clone-ci-cd-with-monitoring-email-990fbd115102



Install Prometheus on a new server. 


Lets check the latest version of Prometheus from the download page.


You can use the curl or wget command to download Prometheus. https://prometheus.io/download/

wget https://github.com/prometheus/prometheus/releases/download/v2.47.1/prometheus-2.47.1.linux-amd64.tar.gz


we need to extract all Prometheus files from the archive.

tar -xvf prometheus-2.47.1.linux-amd64.tar.gz


Usually, you would have a disk mounted to the data directory. For this tutorial, I will simply create a /data directory. Also, you need a folder for Prometheus configuration files


sudo mkdir -p /data /etc/prometheus


Now, let’s change the directory to Prometheus and move some files.
cd prometheus-2.47.1.linux-amd64/


First of all, let’s move the Prometheus binary and a promtool to the /usr/local/bin/. promtool is used to check configuration files and Prometheus rules.
sudo mv prometheus promtool /usr/local/bin/


Optionally, we can move console libraries to the Prometheus configuration directory. Console templates allow for the creation of arbitrary consoles using the Go templating language. You don’t need to worry about it if you’re just getting started.
sudo mv consoles/ console_libraries/ /etc/prometheus/


Finally, let’s move the example of the main Prometheus configuration file.
sudo mv prometheus.yml /etc/prometheus/prometheus.yml


To avoid permission issues, you need to set the correct ownership for the /etc/prometheus/ and data directory.
sudo chown -R prometheus:prometheus /etc/prometheus/ /data/


You can delete the archive and a Prometheus folder when you are done.
cd
rm -rf prometheus-2.47.1.linux-amd64.tar.gz


prometheus --version


We’re going to use Systemd, which is a system and service manager for Linux operating systems. For that, we need to create a Systemd unit configuration file.


sudo vim /etc/systemd/system/prometheus.service

[Unit]
Description=Prometheus
Wants=network-online.target
After=network-online.target

StartLimitIntervalSec=500
StartLimitBurst=5

[Service]
User=prometheus
Group=prometheus
Type=simple
Restart=on-failure
RestartSec=5s
ExecStart=/usr/local/bin/prometheus \
  --config.file=/etc/prometheus/prometheus.yml \
  --storage.tsdb.path=/data \
  --web.console.templates=/etc/prometheus/consoles \
  --web.console.libraries=/etc/prometheus/console_libraries \
  --web.listen-address=0.0.0.0:9090 \
  --web.enable-lifecycle

[Install]
WantedBy=multi-user.target

To automatically start the Prometheus after reboot, run enable.
sudo systemctl enable prometheus

sudo systemctl start prometheus

sudo systemctl status prometheus

Suppose you encounter any issues with Prometheus or are unable to start it. The easiest way to find the problem is to use the journalctl command and search for errors.
journalctl -u prometheus -f --no-pager

Now we can try to access it via the browser. I’m going to be using the IP address of the Ubuntu server. You need to append port 9090 to the IP.
<public-ip:9090>






