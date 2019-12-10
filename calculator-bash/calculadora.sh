#!/bin/bash

function somar {
        resultado=$(($numeroUm+$numeroDois))
        echo 'O resultado desta soma é: '$resultado
}

function subtrair {
        resultado=$(($numeroUm-$numeroDois))
        echo 'O resultado desta subtração é: '$resultado
}

function multiplicar {
        resultado=$(($numeroUm*$numeroDois))
        echo 'O resultado desta muiltiplicação é: '$resultado
}

function dividir {
        resultado=$(($numeroUm/$numeroDois))
        echo 'O resultado desta divisão é: '$resultado
}

printf 'Qual operação você deseja fazer ?\n'
printf '[1] - Adição\n'
printf '[2] - Subtração\n'
printf '[3] - Multiplicação\n'
printf '[4] - Divisão\n'
printf '[5] - Sair\n'
read OPCAO
printf 'Digite o primeiro valor:\n'
read numeroUm
printf 'Digite o segundo valor:\n'
read numeroDois
case $OPCAO in
    1)
        somar
        ;;
    2)
        subtrair
        ;;
    3)
        multiplicar
        ;;
    4)
        dividir
        ;;
    *)
        echo 'bye bye'
        exit 0
        ;;
esac