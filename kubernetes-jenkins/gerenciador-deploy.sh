#!/bin/bash

function run {
    minikube start
    kubectl config use-context minikube
    eval $(minikube docker-env)
}

function stop {
    kubectl delete deployment calculadora-go-minikube
    kubectl delete service calculadora-go-minikube
    eval $(minikube docker-env -u)
    minikube stop
    minikube delete
}

function status {
    minikube status
}

function url {
    minikube service calculadora-go-minikube --url
}

case $1 in
     "run")
          run
          ;;
     "stop")
          stop
          ;;
     "status")
          status
          ;;
     "url")
          url
          ;;
     "*")
          echo "Opção não encontrada !"
esac