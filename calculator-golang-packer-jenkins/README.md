# Calculadora em Go provisionada com Packer e build Docker

## Requisitos

  - [Git](https://git-scm.com/)
  - [Packer](https://www.packer.io/)
  - [Docker](https://www.docker.com/)
  - [Jenkins](https://jenkins.io/)

## Instalação

- Instalando Git:
```
$apt-get install git
```

- Clone o projeto do github:
```
$git clone https://github.com/matheusvargas481/jts.devops.2019.2/tree/tema-12
```

- Adicionar o Jenkins ao grupo do docker com o seguinte comando.

```
$sudo usermod -aG docker jenkins
```

## Configurando Jenkins

1. Adicione seu nome de usuário e senha nas configurações para acessar o Jenkins:

	- Gerenciar Jenkins/Configure Credentials;

2. Acesse seu Jenkins - http://localhost:8080/

3. Devemos criar dois Pipelines no Jenkins um com o nome de BAKE e outro com nome de LAUNCH:

- Novo Job
- De um nome ao Job e marque a opção Pipeline e de OK
- Entre no Job criado e clique em configurar
- Para os dois trabalhos, adicione:
- Na aba General, adicione sua descrição para o trabalho; (Opcional)
- Na categoria Pipeline, selecione: Pipeline script from SCM
- No SCM: selecione Git
- No Repository URL, adicione: (`https://github.com/matheusvargas481/jovens-talentos.git`)
- Nas credenciais, adicione suas credenciais criadas na etapa 1;
- No 	Branches to build, adicione: refs/heads/tema-12

- 1 - Para o Job BAKE ;
- No Script Path, adicione: 3-devops/matheus-vargas/Tema-12/BAKE/Jenkinsfile

- 2 - Para o Job LAUNCH;
- No Script Path, adicione: 3-devops/matheus-vargas/Tema-12/LAUNCH/Jenkinsfile

após execute o BAKE e em seguida o LAUNCH.
No final da execução a aplicação estará pronta para uso.


## Como utilizar a calculadora

Você pode fazer cálculos simples como: sum, sub, mul ou div, acessando os seguintes endpoints:

http://localhost:8082/calc/{operacao}/{PrimeiroNúmero}/{SegundoNúmero}

na {operacao} escolha: sum, sub, mul or div.

Se você quiser conferir o histórico da calculadora, acesse:
https://localhost:8082/calc/historico


## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
