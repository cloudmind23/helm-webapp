
#### Download and verify version of Kubernetes and Helm, Minikube via homebrew and Download Docker Desktop




1) mkdir helm-webapp folder than write 

helm create webapp1




2) Start by running the command 

"minikube start" 

to start container which will also docker container.


 - results: Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default.

 -verify kubectl is running - kubectl get all --all-namespaces | more






# Create the helm chart


```
Create a folder than on terminal write command below




helm create webapp1
```




# Follow along with the video


- Create the files per the video, copying and pasting from templates-original

- you can also use the files in the solution folder





# Install the first release 




Go to PWD (/Users/test/Desktop/helm-webapp) to webapp1 and write command below




1. run.... helm install mywebapp-release webapp1/



2. run.... helm install mywebapp-release webapp1/ --values webapp1/values.yaml




Command 1: helm install mywebapp-release webapp1/


This command installs the Helm chart located in the webapp1/ directory with the default release name mywebapp-release. It uses the default values defined in the chart's values.yaml file.





Command 2: helm install mywebapp-release webapp1/ --values mywebapp/values.yaml


This command also installs the Helm chart located in the webapp1/ directory with the default release name mywebapp-release. However, it uses the values defined in the mywebapp/values.yaml file instead of the default values in the chart's values.yaml file.





Key Differences:

Values File: The first command uses the default values defined in the chart's values.yaml file, while the second command uses the values defined in a separate mywebapp/values.yaml file.

Customization: The second command allows you to customize the deployment of the chart by overriding the default values with your own values. This can be useful for configuring different settings, such as environment variables, resource limits, or service ports.





helm install mywebapp-release webapp1/ --values mywebapp/values.yaml


1. Results:

NAME: mywebapp-release
LAST DEPLOYED: Wed Oct 16 00:55:23 2024
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
servicename=$(k get service -l "app=myhelmapp4" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace default port-forward service/myhelmapp4 8888:80





2. kubectl get all --all-namespaces | more

NAMESPACE     NAME                                   READY   STATUS    RESTARTS      AGE
default       pod/myhelmapp4-69b4b5fcdd-76twq        1/1     Running   0             2m32s







# Upgrade after templating



3. helm upgrade mywebapp-release webapp1/ --values webapp1/values.yaml

test@Ahmeds-M1-Max-Pro helm-webapp % helm upgrade mywebapp-release webapp1/ --values webapp1/values.yaml
Release "mywebapp-release" has been upgraded. Happy Helming!
NAME: mywebapp-release
LAST DEPLOYED: Sat Apr  5 00:34:01 2025
NAMESPACE: default
STATUS: deployed
REVISION: 7
TEST SUITE: None
NOTES:
servicename=$(k get service -l "app=myhelmapp5" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace default port-forward service/myhelmapp5 8888:80



Step 1: Go to values.yaml and write your config.



Step 2: Go to helm templates and configure them with :{{ .Values.appName }} and {{ .Values.appName }} and more from your values.yaml





# Accessing it



Step 4: minikube tunnel




test@tests-MacBook-Pro helm-webapp % minikube tunnel

🤷  The control-plane node minikube apiserver is not running: (state=Stopped)
👉  To start a cluster, run: "minikube start"
test@tests-MacBook-Pro helm-webapp % minikube start
😄  minikube v1.34.0 on Darwin 12.5.1 (arm64)
✨  Using the docker driver based on existing profile
👍  Starting "minikube" primary control-plane node in "minikube" cluster
🚜  Pulling base image v0.0.45 ...
🏃  Updating the running docker "minikube" container ...
🐳  Preparing Kubernetes v1.31.0 on Docker 27.2.0 ...
🔎  Verifying Kubernetes components...
    ▪ Using image docker.io/kubernetesui/metrics-scraper:v1.0.8
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
    ▪ Using image docker.io/kubernetesui/dashboard:v2.7.0
💡  Some dashboard features require the metrics-server addon. To enable all features please run:

        minikube addons enable metrics-server

🌟  Enabled addons: storage-provisioner, default-storageclass, dashboard
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default
test@tests-MacBook-Pro helm-webapp % 



-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------



# Create dev/prod




Step 1: kubectl create namespace dev


Step 1: kubectl create namespace development



(if i want to delete "kubectl delete ns development")




Step 2: helm install mywebapp-release-dev webapp1/ --values webapp1/values.yaml -f webapp1/values-dev.yaml -n development



Step 3: helm upgrade mywebapp-release-dev webapp1/ --values webapp1/values.yaml -f webapp1/values-dev.yaml -n development




-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------






Step 1: kubectl create namespace prod

Step 1: kubectl create namespace production



if i want to delete "kubectl delete ns production"





Step 2: helm install mywebapp-release-prod webapp1/ --values webapp1/values.yaml -f webapp1/values-prod.yaml -n production



Step 2: helm upgrade mywebapp-release-prod webapp1/ --values webapp1/values.yaml -f webapp1/values-prod.yaml -n production



------------------------------------------------------------------------------------------------------------------------------------------------------------




helm ls --all-namespaces


-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------



kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-ext

 Results: Able to Enter Prometheus UI



-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------




kubectl expose service my-release-grafana --type=NodePort --target-port=3000 --name=my-release-grafana-ext

 Results: service/my-release-grafana exposed



To access grafana UI 
  
  run minikube service my-release-grafana-ext
  
  copy URL/IP address to access grafana UI

-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------



kubectl get svc

prometheus-server-ext
my-release-grafana-ext

decrypt password 


--------------------------------------------------------------------------------------------------------------------



Helm is a package manager for Kubernetes that simplifies the deployment and management of applications. Helm uses a packaging format called charts, which are collections of files that describe a set of Kubernetes resources. Below is an example of a simple Helm chart, and I'll explain each part:

myapp/
|-- Chart.yaml
|-- values.yaml
|-- templates/
|   |-- deployment.yaml
|   |-- service.yaml
|-- charts/
|-- README.md



wordpress/
  Chart.yaml          # A YAML file containing  information about the chart
  LICENSE             # OPTIONAL: A plain text file containing the license for the chart
  README.md           # OPTIONAL: A human-readable README file
  values.yaml         # The default configuration values for this chart
  values.schema.json  # OPTIONAL: A JSON Schema for imposing a structure on the values.yaml file
  charts/             # A directory containing any charts upon which this chart depends.
  crds/               # Custom Resource Definitions
  templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
  templates/NOTES.txt # OPTIONAL: A plain text file containing short usage notes


1. Chart.yaml:

This file contains metadata about the Helm chart, such as the name, version, description, and maintainer.

apiVersion: v2
name: myapp
description: A Helm chart for deploying My App
version: 0.1.0


2. #values.yaml:

This file contains default values for configurable parameters in your chart. Users can override these values when installing the chart.

-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

# Default values for myapp.

appName: myhelmapp3

port: 80

namespace: default

configmap:
  name: myhelmapp-configmap-v1
  data:
    CUSTOM_HEADER: 'This app was deployed with helm!'
 
image:
  name: devopsjourney1/mywebapp 
  tag: latest

-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

3. templates////////////: 

This directory contains Kubernetes YAML files that define the resources to be deployed. These files can include Deployments, Services, ConfigMaps, etc.


Deployment.yaml:
	apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "myapp.fullname" . }}
  labels:
    app.kubernetes.io/name: {{ include "myapp.name" . }}
    app.kubernetes.io/instance: {{ .Release.Name }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app.kubernetes.io/name: {{ include "myapp.name" . }}
      app.kubernetes.io/instance: {{ .Release.Name }}
  template:
    metadata:
      labels:
        app.kubernetes.io/name: {{ include "myapp.name" . }}
        app.kubernetes.io/instance: {{ .Release.Name }}
    spec:
      containers:
        - name: myapp
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}



Service.yaml:
	apiVersion: v1
kind: Service
metadata:
  name: {{ include "myapp.fullname" . }}
  labels:
    app.kubernetes.io/name: {{ include "myapp.name" . }}
    app.kubernetes.io/instance: {{ .Release.Name }}
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: 80
  selector:
    app.kubernetes.io/name: {{ include "myapp.name" . }}
    app.kubernetes.io/instance: {{ .Release.Name }}




4. charts/:


This directory is used for storing any sub-charts if your main chart depends on other charts.


README.md:
Documentation for the chart, providing instructions on how to use it and any other relevant information.


Helm is a package manager for Kubernetes that simplifies the process of deploying and managing complex applications. It provides a declarative approach to defining and managing applications, making it easier for developers and operations teams to collaborate.

Here are some of the key benefits of using Helm:

Simplified application deployment: Helm allows you to define an application's configuration and dependencies in a single file, making it easier to deploy and manage.
Version control: Helm uses a versioning system to track changes to your applications, making it easy to roll back to previous versions if necessary.
Reusable charts: Helm charts can be shared and reused across different projects, saving time and effort.
Integration with Kubernetes: Helm is tightly integrated with Kubernetes, making it easy to manage applications running on the platform.
Community support: Helm has a large and active community, which means that there are plenty of resources available to help you get started and troubleshoot problems.
Overall, Helm is a powerful tool that can help you simplify the process of deploying and managing applications on Kubernetes.


Helm simplifies the management of Kubernetes applications in several ways, even if you are proficient in Kubernetes:

Declarative Configuration: Helm uses a declarative approach, where you define the desired state of your application in a template (chart). This is in contrast to Kubernetes' imperative approach, where you specify a series of steps to achieve the desired state. Helm's declarative approach makes it easier to manage complex applications and understand their configuration.

Reusable Charts: Helm charts can be shared and reused across different projects. This means you don't have to start from scratch every time you want to deploy a new application. It also promotes consistency and standardization.

Version Control: Helm uses a versioning system to track changes to your charts. This allows you to easily roll back to previous versions if necessary, providing a safety net for your deployments.

Package Management: Helm provides a package management system for Kubernetes applications. This means you can easily search for, install, and update charts from a public or private repository.

Community Support: Helm has a large and active community, which means there are plenty of resources available to help you get started and troubleshoot problems.

Simplified Configuration Management: Helm can help you manage complex configuration settings for your Kubernetes applications. You can define values in a values.yaml file and pass them to your charts, making it easier to customize your deployments.

In summary, while you may be proficient in Kubernetes, Helm can still make your life easier by providing a more efficient and streamlined way to manage your applications. It takes care of many of the repetitive tasks involved in Kubernetes deployments, allowing you to focus on building and deploying your applications.