#!/bin/bash

printf 'Qual operação você deseja fazer ?\n'
printf '[1] - Iniciar Microsserviço\n'
printf '[2] - Parar Microsserviço\n'
printf '[3] - Verificar Status do Microsserviço\n'
read OPCAO

function startMicroservice () {
    docker build -t calculadorago .
    docker run -d -p 9000:9000 --name calculadorago calculadorago
}

function stopMicroservice () {
    docker rm -f calculadorago
}

function statusMicroservice () {
if [ "$(docker ps -a | grep calculadorago)" ];
then
        echo "Microsserviço UP"
else
        echo "Microserviço DOWN"
fi
}

case $OPCAO in
    1)
        startMicroservice
    ;;
    2)
        stopMicroservice
    ;;
    3)
        statusMicroservice
    ;;
    *)
        echo 'bye bye'
        exit 0
    ;;
esac