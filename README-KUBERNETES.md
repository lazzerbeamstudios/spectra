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
