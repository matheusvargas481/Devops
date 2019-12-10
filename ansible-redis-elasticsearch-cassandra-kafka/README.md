# Provisionando imagens utilizando roles Ansible para NoSQL

## Requisitos

  - [Git](https://git-scm.com/)
  - [Docker](https://www.docker.com/)
  - [Packer](https://www.packer.io/)
  - [Ansible](https://www.ansible.com/)

## Instalação

- Instalando Git:
```
$ apt-get install git
```

- Clone o projeto do github:
```
$ git clone https://github.com/matheusvargas481/jovens-talentos/tree/tema-17/
```

## Configuração

1. Acesse a pasta ~/jovens-talentos/3-devops/matheus-vargas/Tema-17.

2. Entre na pasta que desejar cassandra/elasticsearch/kafka/redis.

3. Execute o script com o seguinte comando:
```
$ ./script.sh {UTILIZANDO AS OPÇÕES ABAIXO}
```
OBS: as opções devem ser escritas em minúsculo
- bake: para provisionar a imagem
- run: para rodar a aplicação
- status: para verificar se a aplicação esta up
- stop: para parar a aplicação

## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
