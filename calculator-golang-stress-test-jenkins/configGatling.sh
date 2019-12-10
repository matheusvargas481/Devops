#!/bin/bash

wget https://repo1.maven.org/maven2/io/gatling/highcharts/gatling-charts-highcharts-bundle/3.3.0/gatling-charts-highcharts-bundle-3.3.0-bundle.zip
unzip gatling-charts-highcharts-bundle-3.3.0-bundle.zip
rm -r gatling-charts-highcharts-bundle-3.3.0-bundle.zip
cp TestStressMicroServico.scala gatling-charts-highcharts-bundle-3.3.0/user-files/simulations/computerdatabase
cd gatling-charts-highcharts-bundle-3.3.0/bin
./gatling.sh -s computerdatabase.TestStressMicroServico