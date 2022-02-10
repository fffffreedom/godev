#!/bin/bash

# get namespace
namespaces=`kubectl get pod -A | grep -i "evicted" | awk '{print $1}'`
for namespace in ${namespaces}
do
    kubectl get pod -n ${namespace} |grep -i "evicted"|awk '{print $1}' | xargs kubectl delete pod -n ${namespace}
done
