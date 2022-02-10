
#minikube start \
#    --image-mirror-country cn \
#    --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers \
#    --memory=8096 \
#    --cpus=4 \
#    --docker-env HTTP_PROXY=http://10.0.2.2:7890 \
#    --docker-env HTTPS_PROXY=http://10.0.2.2:7890 \
#    --docker-env NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,192.168.39.0/24,10.0.2.2

minikube start \
    --image-mirror-country cn \
    --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers \
    --memory=8096 \
    --cpus=4 \
    --docker-env HTTP_PROXY= \
    --docker-env HTTPS_PROXY= \
    --docker-env NO_PROXY=
