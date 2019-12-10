# Calculadora em Go com VM Vagrant

## Requisitos

  - [Git](https://git-scm.com/)
  - [Go](https://golang.org/)
  - [Vagrant](https://www.vagrantup.com)
  - [Virtual Box](https://www.virtualbox.org)

## Instalação

- Instalando Git:
```
$apt-get install git
```

- Clone o projeto do github:
```
$git clone https://github.com/matheusvargas481/jts.devops.2019.2/tree/tema-8
```

- Instalando Vagrant
```
$sudo apt install vagrant
```

- Instalando Vistual Box
```
$sudo apt install virtualbox
```

Execute o comando no terminal importar a biblioteca gorilla/mux.

```
  $go get github.com/gorilla/mux
```

Acesse a pasta ~/jts.devops.2019.2/matheus-vargas/Tema-8/ e use o comando abaixo para buildar a aplicação.

```
  $go build calculadora.go
```

Execute o comando abaixo para subir a aplicação:

```
  $vagrant up
```


Você pode fazer calculos simples como: sum, sub, mul ou div, acessando os seguintes endpoits:

https://77.77.77.7:9000/calc/{operacao}/{PrimeiroNúmero}/{SegundoNúmero}

na {operacao} escolha: sum, sub, mul or div.

Se você quiser conferir o histórico da calculadora, acesse:
https://localhost:9000/calc/historico


## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
