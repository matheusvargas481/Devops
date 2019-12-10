# Calculadora em Go provisionada com Packer e build Docker

## Requisitos

  - [Git](https://git-scm.com/)
  - [Packer](https://www.packer.io/)
  - [Docker](https://www.docker.com/)

## Instalação

- Instalando Git:
```
$apt-get install git
```

- Clone o projeto do github:
```
$git clone https://github.com/matheusvargas481/jts.devops.2019.2/tree/tema-10
```

Acesse a pasta ~/jts.devops.2019.2/matheus-vargas/Tema-10/ e use o comando abaixo para provisionar a imagem.

```
  $packer build build.json
```

E para executar a imagem docker use o seguinte comando.

```
  $docker run -d -p {Porta}:9000 calculadora-container:latest
```
na {Porta} digite a porta em que sua aplicação deverá ficar acessível


## Como utilizar a calculadora

Você pode fazer cálculos simples como: sum, sub, mul ou div, acessando os seguintes endpoints:

http://localhost:{Porta}/calc/{operacao}/{PrimeiroNúmero}/{SegundoNúmero}

na {Porta} use a mesma porta que foi utilizada no comando anterior e na {operacao} escolha: sum, sub, mul or div.

Se você quiser conferir o histórico da calculadora, acesse:
https://localhost:{Porta}/calc/historico


## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
