## Helm

**commands**

    helm list

    helm history api
    helm uninstall api

**dev**

    helm install api helm-v1 --values helm-v1/values.yaml -f helm-v1/dev.yaml
    helm upgrade api helm-v1 --values helm-v1/values.yaml -f helm-v1/dev.yaml

**stag**

    helm install api helm-v1 --values helm-v1/values.yaml -f helm-v1/stag.yaml
    helm upgrade api helm-v1 --values helm-v1/values.yaml -f helm-v1/stag.yaml

**prod**

    helm install api helm-v1 --values helm-v1/values.yaml -f helm-v1/prod.yaml
    helm upgrade api helm-v1 --values helm-v1/values.yaml -f helm-v1/prod.yaml
