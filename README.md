# helm-webapp

> *Deploying applications to Kubernetes — mission-critical, repeatable, and fast.*

A production-grade Helm chart project for deploying a containerized web application to Kubernetes via Minikube, with multi-environment promotion (dev → prod) and a full observability stack powered by Prometheus and Grafana.

---

## Mission Overview

This repository demonstrates a complete Helm-based deployment workflow for platform and DevOps engineers who need environment parity, declarative configuration, and first-class observability out of the box. The chart ships with:

- **Multi-environment support** — isolated `development` and `production` namespaces with value overrides
- **ConfigMap-driven runtime config** — background color, font color, and custom headers injected as env vars
- **Monitoring stack** — Prometheus scraping + Grafana dashboards, both exposed via NodePort
- **ArgoCD-ready** — GitOps-compatible structure for continuous delivery pipelines

---

## Architecture

```
helm-webapp/
├── webapp1/
│   ├── Chart.yaml              # Chart metadata and versioning
│   ├── values.yaml             # Base configuration (all environments inherit this)
│   ├── values-dev.yaml         # Development overrides
│   ├── values-prod.yaml        # Production overrides
│   └── templates/
│       ├── deployment.yaml     # 3-replica Deployment — image and configmap from values
│       ├── service.yaml        # LoadBalancer Service on port 80
│       ├── configmap.yaml      # Injects BG_COLOR, FONT_COLOR, CUSTOM_HEADER
│       └── NOTES.txt           # Post-install instructions
└── docs/
    ├── prometheus-helm.go      # Prometheus setup reference
    └── grafana-helm.go         # Grafana setup rfix my readme.md. Make it a space theme helm-webapp github. Make it look like something a director of engineering would write up.eference
```

All Helm values flow from `values.yaml` as the base. Environment-specific files layer overrides on top via `-f`. Templates reference values with `{{ .Values.<key> }}` — no hardcoded strings in manifests.

---

## Prerequisites

| Tool | Install |
| ---- | ------- |
| [Minikube](https://minikube.sigs.k8s.io/) | `brew install minikube` |
| [kubectl](https://kubernetes.io/docs/tasks/tools/) | `brew install kubectl` |
| [Helm 3](https://helm.sh/) | `brew install helm` |
| [Docker Desktop](https://www.docker.com/products/docker-desktop/) | Required — Minikube uses the Docker driver on macOS |

---

## Launch Sequence

### 1. Start the cluster

```sh
minikube start
```

Verify everything is nominal:

```sh
kubectl get all --all-namespaces
```

### 2. Validate the chart

Always lint and dry-run before deploying:

```sh
helm lint webapp1/
helm template webapp1/ --values webapp1/values.yaml
```

### 3. Install

```sh
helm install mywebapp-release webapp1/ --values webapp1/values.yaml
```

### 4. Upgrade after changes

```sh
helm upgrade mywebapp-release webapp1/ --values webapp1/values.yaml
```

### 5. Access the application

```sh
# Option A — tunnel (recommended for LoadBalancer services)
minikube tunnel

# Option B — port-forward
kubectl --namespace default port-forward service/<appName> 8888:80
```

---

## Multi-Environment Deployment

Namespaces provide environment isolation. Each environment gets its own Helm release with values layered on top of the base config.

### Development

```sh
kubectl create namespace development

helm install mywebapp-release-dev webapp1/ \
  --values webapp1/values.yaml \
  -f webapp1/values-dev.yaml \
  -n development

# Upgrade
helm upgrade mywebapp-release-dev webapp1/ \
  --values webapp1/values.yaml \
  -f webapp1/values-dev.yaml \
  -n development
```

### Production

```sh
kubectl create namespace production

helm install mywebapp-release-prod webapp1/ \
  --values webapp1/values.yaml \
  -f webapp1/values-prod.yaml \
  -n production

# Upgrade
helm upgrade mywebapp-release-prod webapp1/ \
  --values webapp1/values.yaml \
  -f webapp1/values-prod.yaml \
  -n production
```

### Inspect all releases

```sh
helm ls --all-namespaces
```

---

## Configuration Reference

The base `values.yaml` controls all defaults:

```yaml
appName: myhelmapp
namespace: default

image:
  name: devopsjourney1/mywebapp
  tag: latest

configmap:
  name: myhelmapp-configmap-v1
  data:
    BG_COLOR: "#ffffff"
    FONT_COLOR: "#000000"
    CUSTOM_HEADER: "This app was deployed with Helm!"
```

> **Note:** Quote `CUSTOM_HEADER` values in `configmap.yaml` if they contain spaces or special characters — an unquoted value will fail YAML parsing with `mapping values are not allowed in this context`.

---

## Observability Stack

### Prometheus

```sh
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/prometheus

# Expose the UI
kubectl expose service prometheus-server \
  --type=NodePort \
  --target-port=9090 \
  --name=prometheus-server-ext

minikube service prometheus-server-ext
```

### Grafana

```sh
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install my-release grafana/grafana

# Expose the UI
kubectl expose service my-release-grafana \
  --type=NodePort \
  --target-port=3000 \
  --name=my-release-grafana-ext

minikube service my-release-grafana-ext
```

Retrieve the admin password:

```sh
kubectl get secret --namespace default my-release-grafana \
  -o jsonpath="{.data.admin-password}" | base64 --decode; echo
```

In Grafana, add Prometheus as a data source using the ClusterIP of `prometheus-server-ext`. Import dashboard **ID 6417** for Kubernetes cluster metrics.

---

## ArgoCD (GitOps)

Forward the ArgoCD UI to localhost:

```sh
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

Enable the web terminal:

```sh
kubectl patch cm argocd-cm -n argocd \
  --type='json' \
  -p='[{"op": "add", "path": "/data/exec.enabled", "value": "true"}]'
```

---

## Teardown

```sh
helm uninstall mywebapp-release
helm uninstall mywebapp-release-dev -n development
helm uninstall mywebapp-release-prod -n production

kubectl delete namespace development production
minikube stop
```

---

## Why Helm

Raw Kubernetes manifests don't scale across environments — you end up maintaining near-duplicate YAML with subtle drift between dev and prod. Helm solves this with a single source of truth (`values.yaml`) and a templating engine that renders environment-specific manifests at deploy time. The result is fewer incidents caused by configuration drift, auditable releases with rollback, and charts that can be promoted through environments without touching template logic.
