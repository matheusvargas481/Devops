# Calculadora em Go deploy Google Cloud Plataform

## Requisitos

  - [Git](https://git-scm.com/)
  - [Google Cloud](https://cloud.google.com/?hl=pt-br)

## Instalação

- Instalando Git:
```
$ apt-get install git
```

- Clone o projeto do github:
```
$ git clone https://github.com/matheusvargas481/jovens-talentos/tree/tema-14/
```

## Configuração

1. Crie uma conta no Google Cloud.

2. Use o comando abaixo para instalação do Gcloud.
```
  $ curl https://sdk.cloud.google.com | bash
```

3. Acesse a pasta ~/jts.devops.2019.2/matheus-vargas/Tema-14/ e use os comandos abaixo para inicializar o gcloud e logar em sua conta.

```
  $ exec -l $SHELL
  $ sudo chown {SEU USUÁRIO DO COMPUTADOR} -R /home/{SEU USUÁRIO DO COMPUTADOR}/.config/gcloud
  $ gcloud init
```

4. Crie um projeto com o seguinte comando:
```
  $ gcloud projects create {Nome do Projeto/ID}
```

5. Após o login e a criação do projeto use o seguinte comando para dar deploy do calculadora.go e subir duas instâncias

```
  $ gcloud app deploy
```

## Como utilizar a calculadora

1. Comando para acessar o browser:

```
$ gcloud app browse
```

Você pode fazer cálculos simples como: sum, sub, mul ou div, completando a URL que abrir no browser com:

/calc/{operacao}/{PrimeiroNúmero}/{SegundoNúmero}

na {operacao} escolha: sum, sub, mul or div.

Se você quiser conferir o histórico da calculadora, complete com:
/calc/historico

## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
