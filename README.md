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