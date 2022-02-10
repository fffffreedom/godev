
minikube start \
    --vm-driver=virtualbox \
    --iso-url=https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/iso/minikube-v1.7.3.iso \
    --image-mirror-country cn \
    --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers \
    --registry-mirror=https://docker.mirrors.ustc.edu.cn \
    --kubernetes-version=v1.17.3 \
	--network-plugin=cni \
    --memory=8192 \
    --cpus=4 \
    -p $1

#    --registry-mirror=https://registry.docker-cn.com \
#    --registry-mirror=https://docker.mirrors.ustc.edu.cn \
#    --registry-mirror=https://docker.mirrors.ustc.edu.cn,https://hub-mirror.c.163.com \
#    --registry-mirror=https://hub-mirror.c.163.com,https://docker.mirrors.ustc.edu.cn \
