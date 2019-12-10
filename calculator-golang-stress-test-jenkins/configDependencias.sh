#!/bin/bash

apt-get update
wget -q -O - https://pkg.jenkins.io/debian/jenkins-ci.org.key | apt-key add -
sh -c 'echo deb http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'
apt-get update
apt-get install jenkins
apt-get install curl
curl -O https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz
tar -C /usr/local -zxvf go1.13.1.linux-amd64.tar.gz
echo export PATH=$PATH:/usr/local/go/bin >> $HOME/.bashrc
. $HOME/.bashrc
rm -r go1.13.1.linux-amd64.tar.gz
apt-get -y remove scala-library scala
curl -O https://downloads.lightbend.com/scala/2.12.10/scala-2.12.10.deb
dpkg -i scala-2.12.10.deb
rm -r scala-2.12.10.deb
echo "deb http://dl.bintray.com/sbt/debian /" | tee /etc/apt/sources.list.d/sbt.list
apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 642AC823
apt-get update
apt-get install sbt
sbt -version