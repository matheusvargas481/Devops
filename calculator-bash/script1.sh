#!/bin/bash

caminho=$1
if [ -d $HOME/$caminho ];
then
    cd $HOME/$caminho
    zip -r arquivo.zip .
    data=$(date +%F)
    mkdir -p $HOME/backup/data/$data
    mv arquivo.zip $HOME/backup/data/$data/
    echo 'Zipado com sucesso !'
else
    echo 'Caminho não encontrado'
fi