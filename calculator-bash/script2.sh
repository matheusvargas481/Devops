#!/bin/bash
data=$(date +%F)
if [ ! -f $HOME/backup/conf/$data/env_data.txt ];
then
    printenv > env_data.txt
    mkdir -p $HOME/backup/conf/$data/
    mv env_data.txt $HOME/backup/conf/$data/
    echo 'Arquivo salvo com sucesso !'
else 
    printenv >> $HOME/backup/conf/$data/env_data.txt
    echo 'Arquivo atualizado com sucesso !'
fi

