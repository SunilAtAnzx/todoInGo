# todoInGo

kind load docker-image go-todo-img --name sunil.kind
docker exec -it sunil.kind-control-plane crictl images
kubectl port-forward go-todo 8181:8181
kubectl delete svc <>
kubectl get svc
kubectl get nodes
kubectl get pods
kubectl apply -f

docker login k8s-gcr.artifactory.gcp.anz
docker login registry-k8s-io.artifactory.gcp.anz

deleteCluster sunil.kind
createCluster createCluster.yaml

loadImage go-todo-img sunil.kind
loadImage go-collector-img sunil.kind
loadImage au-adp-docker.artifactory.gcp.anz/ingress-nginx-controller:v1.3.1 sunil.kind
loadImage registry-k8s-io.artifactory.gcp.anz/ingress-nginx/kube-webhook-certgen:v1.3.0 sunil.kind
applyConfig ingress.yaml
applyConfig todo-service.yaml
applyConfig collector-service.yaml
applyConfig service-ingress.yaml