# Calculadora em Go com VM Vagrant e Ansible

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
$git clone https://github.com/matheusvargas481/jts.devops.2019.2/tree/tema-9
```

- Instalando Vagrant
```
$sudo apt install vagrant
```

- Instalando Virtual Box
```
$sudo apt install virtualbox
```

Acesse a pasta ~/jts.devops.2019.2/matheus-vargas/Tema-9/ e use o comando abaixo para executar a aplicação.

```
  $vagrant up
```

Assim que printar no console o log com a porta exposta pode seguir utilizando a calculadora.

## Como utilizar a calculadora

Você pode fazer cálculos simples como: sum, sub, mul ou div, acessando os seguintes endpoints:

http://77.77.77.7:9000/calc/{operacao}/{PrimeiroNúmero}/{SegundoNúmero}

na {operacao} escolha: sum, sub, mul or div.

Se você quiser conferir o histórico da calculadora, acesse:
https://localhost:9000/calc/historico


## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
