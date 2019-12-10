# Calculadora em Go deploy AWS com Terraform

## Requisitos

  - [Git](https://git-scm.com/)
  - [Packer](https://www.packer.io/)
  - [Ansible](https://www.ansible.com/)
  - [Terraform](https://www.terraform.io/)

## Instalação

- Instalando Git:
```
$apt-get install git
```

- Clone o projeto do github:
```
$git clone https://github.com/matheusvargas481/jovens-talentos/tree/tema-13/
```

## Configuração

1. Crie seu key pair na AWS com o nome: "kp_devops_tema13" e mova um arquivo para o caminho (`/home/<seu usuário>/.ssh/kp_devops_tema13.pem`).

2. Adicione o AWSAcessKeyId e a AWSSecretKey às variáveis ​​de ambiente:

    - Adicione as variáveis ​​de ambiente editando o arquivo "$HOME/.bashrc" com as seguintes linhas:

    ```
    export AWSAcessKeyId='your-key'
    export AWSSecretKey='your-key'
    export TF_VAR_aws_access_key='your-key'
    export TF_VAR_aws_secret_key='your-key'
    ```
Acesse a pasta ~/jts.devops.2019.2/matheus-vargas/Tema-13/ e use o comando packer para provisionar a imagem.

```
  $ packer build build.json
```

Após o build completo use o comando abaixo para iniciar o terraform.

```
  $terraform init
```

E para executar o main.tf use o seguinte comando.

```
  $terraform apply
```

## Como utilizar a calculadora

1. Comando em execução para permissão de modificação key pair:

```
$ chmod 400 ~/.ssh/kp_devops_tema13.pem
```

2. Execute o comando para entrar na máquina:
```
$ ssh -i ~/.ssh/kp_devops_tema13.pem ec2-user@<ip-instance>
```

3. Executar binário da calculadora:
```
$ ./calculadora
```

Você pode fazer cálculos simples como: sum, sub, mul ou div, acessando os seguintes endpoints:

http://{ip-instance}:9000/calc/{operacao}/{PrimeiroNúmero}/{SegundoNúmero}

na {operacao} escolha: sum, sub, mul or div.

Se você quiser conferir o histórico da calculadora, acesse:
https://{ip-instance}:9000/calc/historico


## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
