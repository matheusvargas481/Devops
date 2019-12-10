#!/bin/bash

groupadd k8s
usermod -aG k8s $1
usermod -aG k8s jenkins
chgrp k8s {Informe seu HOMEPATH}/.kube/config
chgrp k8s {Informe seu HOMEPATH}/.minikube/client.key
chmod -R g+r {Informe seu HOMEPATH}/.kube//config
chmod -R g+r {Informe seu HOMEPATH}/.minikube/client.key