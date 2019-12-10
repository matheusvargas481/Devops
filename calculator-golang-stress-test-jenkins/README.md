# Teste de Stress na Calculadora Go com Gatling

## Requisitos

  - [Git](https://git-scm.com/)
  - [Scala](https://www.scala-lang.org/)
  - [Gatling](https://gatling.io/)
  - [Go](https://golang.org/)
  - [Jenkins](https://jenkins.io/)

## Instalação

- Instalando Git:
```
$ apt-get install git
```

- Clone o projeto do github:
```
$ git clone https://github.com/matheusvargas481/jovens-talentos/tree/tema-16/
```

## Configuração

1. Acesse a pasta ~/jovens-talentos/3-devops/matheus-vargas/Tema-16.

2. Instale os pré-requisitos executando o arquivo configDependencias.sh com o seguinte comando no terminal.
```
  $ sudo ./configDependencias.sh
```
3. Adicione o Jenkins ao grupo do docker com o seguinte comando.

```
  $ sudo usermod -aG docker jenkins
```
4. No mesmo terminal execute o comando abaixo para executar a calculadora.
```
  $ go run calculadora.go
```

## Configurando Jenkins

1. Adicione seu nome de usuário e senha nas configurações para acessar o Jenkins através do link http://localhost:8080/:

2. Adicione seu nome de usuário e senha nas configurações a seguir para dar acesso ao seu repositório git:
	- Gerenciar Jenkins/Configure Credentials;

3. Devemos criar um Job LAUNCH:
- Novo Job
- De um nome ao Job e marque a opção Pipeline e de OK
- Para o job criado, adicione:
- Na aba General, adicione sua descrição para o job; (Opcional)
- Na categoria Pipeline, selecione: Pipeline script from SCM
- No SCM: selecione Git
- No Repository URL, adicione: (`https://github.com/ilegra/jovens-talentos.git`)
- Nas credenciais, adicione suas credenciais criadas na etapa 2;
- No 	Branches to build, adicione: refs/heads/tema-16
- No Script Path, adicione: 3-devops/matheus-vargas/Tema-16/Jenkinsfile

após clique em Build Now.

Entre no job e acesso o console do mesmo, quando terminado a execução acesse o link disponibilizado no final do console, o link termina em /index.html para ter acesso aos resultados do teste.


## Autor

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)
