# Calculadora em Go provisionada com Packer para Dockerhub, deploy via Jenkins e Kubernetes

## Requisitos

  - [Git](https://git-scm.com/)
  - [Jenkins](https://jenkins.io/)
  - [Minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)
  - [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Instalação

- Instalando Git:
```
 $ apt-get install git
```

- Clone o projeto do github:
```
 $ git clone https://github.com/ilegra/jovens-talentos.git
```

- Acesse a pasta /3-devops/matheus-vargas/Tema-18 e execute o script com o seguinte comando para efetuar a instalação do minikube, kubectl e jenkins
```
 $ sudo ./installRequirements.sh
```

- Após a conclução das instalações execute o script com o seguinte comando para rodar o minikube 
```
 $ sudo ./gerenciador-deploy.sh run
```

- Após a finalização do comando acima execute o script com o seguinte comando para verificar se o minikube esta up
```
 $ sudo ./gerenciador-deploy.sh status
```

- Após a verificação execute o script para dar as permissões para o Jenkinks e para o as pastas .kube e .minikube com o seguinte comando

    OBS: Acesse o script e altere onde diz {Informe seu HOMEPATH}
```
 $ sudo ./permissoes.sh {INFORME SEU USUÁRIO}
```

- Logo após execute o job do Jenkinks, segue a configuração abaixo.

## Configurando Jenkins

1. Adicione seu nome de usuário e senha nas configurações para acessar o Jenkins:

    - Gerenciar Jenkins/Configure Credentials;

2. Acesse seu Jenkins - http://localhost:8080/

3. Devemos criar um Pipeline no Jenkins com o nome de DEPLOY:

- Novo Job
- De um nome ao Job e marque a opção Pipeline e de OK
- Entre no Job criado e clique em configurar e adicione:
- Na aba General, adicione sua descrição para o trabalho; (Opcional)
- Na categoria Pipeline, selecione: Pipeline script from SCM
- No SCM: selecione Git
- No Repository URL, adicione: (`https://github.com/ilegra/jovens-talentos.git`)
- Nas credenciais, adicione suas credenciais criadas na etapa 1;
- No Branches to build, adicione: refs/heads/tema-18
- No Script Path, adicione: 3-devops/matheus-vargas/Tema-18/deploy/Jenkinsfile

Após a configuração execute o job e aguarde a conclusão do mesmo.

## Como utilizar a calculadora

- Acesse a pasta /3-devops/matheus-vargas/Tema-18 e execute o script com o seguinte comando para descobrir a url de acesso a aplicação
```
 $ sudo ./gerenciador-deploy.sh url
```

- Você pode fazer cálculos simples como: sum, sub, mul ou div, acessando os seguintes endpoints:

    {UTILIZE A URL DESCOBERTA ACIMA}/calc/{operacao}/{PrimeiroNúmero}/{SegundoNúmero}

    na {operacao} escolha: sum, sub, mul or div.

- Se você quiser conferir o histórico da calculadora, acesse:
    {UTILIZE A URL DESCOBERTA ACIMA}/calc/historico

## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)