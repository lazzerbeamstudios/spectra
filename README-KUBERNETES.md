## Kubernetes

**commands**

    kubectl config get-contexts
    kubectl config current-context
    kubectl config use-context [context-name]
    kubectl config delete-context [context-name]

    kubectl config get-clusters
    kubectl config delete-cluster [cluster-name]

    kubectl get
    kubectl get pods
    kubectl get nodes
    kubectl get ingress
    kubectl get services
    kubectl get namespaces
    kubectl get deployments

    options:
        -o wide
        -n [namespace]

    kubectl create namespace [namespace]
    kubectl delete namespace [namespace]

**secret**

    kubectl create secret docker-registry container-secret --docker-server=ghcr.io --docker-username=[username] --docker-email=[email] --docker-password=[password] -n [namespace]
