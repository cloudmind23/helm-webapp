# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Repo Is

A Helm chart project (`webapp1/`) that deploys a simple web application (`devopsjourney1/mywebapp`) to a local Minikube cluster. It includes multi-environment support (default/dev/prod), and documentation for setting up Prometheus and Grafana monitoring via Helm.

## Prerequisites

- Minikube, kubectl, Helm 3 — installed via Homebrew
- Docker Desktop running (Minikube uses the Docker driver on macOS)

## Core Helm Commands

**Start the cluster:**
```sh
minikube start
```

**Install (first time):**
```sh
helm install mywebapp-release webapp1/ --values webapp1/values.yaml
```

**Upgrade after changes:**
```sh
helm upgrade mywebapp-release webapp1/ --values webapp1/values.yaml
```

**Multi-environment installs:**
```sh
# Development namespace
kubectl create namespace development
helm install mywebapp-release-dev webapp1/ --values webapp1/values.yaml -f webapp1/values-dev.yaml -n development
helm upgrade mywebapp-release-dev webapp1/ --values webapp1/values.yaml -f webapp1/values-dev.yaml -n development

# Production namespace
kubectl create namespace production
helm install mywebapp-release-prod webapp1/ --values webapp1/values.yaml -f webapp1/values-prod.yaml -n production
helm upgrade mywebapp-release-prod webapp1/ --values webapp1/values.yaml -f webapp1/values-prod.yaml -n production
```

**List all releases across namespaces:**
```sh
helm ls --all-namespaces
```

**Validate templates before applying:**
```sh
helm template webapp1/ --values webapp1/values.yaml
helm lint webapp1/
```

## Accessing the App

After install, use minikube tunnel (or port-forward from NOTES.txt):
```sh
minikube tunnel
# OR
kubectl --namespace default port-forward service/<appName> 8888:80
```

## Chart Architecture

All Helm values flow from `webapp1/values.yaml` as the base, with environment-specific overrides layered on top via `-f`:

- `values.yaml` — base config: `appName`, `namespace`, `image`, `configmap`
- `values-dev.yaml` — overrides `namespace: development` and `CUSTOM_HEADER`
- `values-prod.yaml` — overrides `namespace: production` and `CUSTOM_HEADER`

Templates in `webapp1/templates/` reference values with `{{ .Values.<key> }}`:
- `deployment.yaml` — 3-replica Deployment; image and configmap name come from values
- `service.yaml` — LoadBalancer Service on port 80
- `configmap.yaml` — injects `BG_COLOR`, `FONT_COLOR`, `CUSTOM_HEADER` as env vars

The `CUSTOM_HEADER` value in `configmap.yaml` must be quoted if it contains spaces or special characters — a missing quote causes `yaml: mapping values are not allowed in this context`.

## Monitoring Stack (Prometheus + Grafana)

Setup commands are documented in `docs/prometheus-helm.go` and `docs/garfana-helm.go`:

```sh
# Prometheus
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/prometheus

# Expose externally
kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-ext
minikube service prometheus-server-ext

# Grafana
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install my-release grafana/grafana

# Expose externally
kubectl expose service my-release-grafana --type=NodePort --target-port=3000 --name=my-release-grafana-ext
minikube service my-release-grafana-ext

# Decode Grafana admin password
kubectl get secret --namespace default my-release-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

In Grafana, connect to Prometheus via the ClusterIP of `prometheus-server-ext`. Import dashboard ID `6417` for Kubernetes cluster metrics.

## ArgoCD

Access via port-forward on port 8080 (target the `argocd-server` pod in the `argocd` namespace). Enable the web terminal:
```sh
kubectl patch cm argocd-cm -n argocd --type='json' -p='[{"op": "add", "path": "/data/exec.enabled", "value": "true"}]'
```
