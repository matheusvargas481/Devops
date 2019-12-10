# Calculadora em GO com container Docker

 Acesse o arquivo ~/jts.devops.2019.2/matheus-vargas/Tema-7/
 
 
  Execute o seguinte comando no terminal
   ```
   $go build calculadora.go
   ```
 
   E, em seguida, torne o script script.sh executável:
   ```
     $sudo chmod +x script.sh
   ```
  E execute o comando no terminal:
  
   ```
   $./script.sh
   ```
    
   E escolha a opção desejada:<br>
   1 - Para rodar o Micro Serviço<br>
   2 - Para parar o Micro Serviço<br>
   3 - Para consultar se o Micro Serviço esta Up ou Down
   
   ## Requirements
   
   - [Git](https://git-scm.com/)
   - [Go](https://golang.org/)
   - [Docker](https://www.docker.com/)
   
   ## Instalação
   
   - Instalando Git:
    
   ```
   $apt-get install git
   ```
   
   - Instalando Docker:
   ```
   $sudo apt-get update
   $sudo apt-get remove docker docker-engine docker.io
   $sudo apt install docker.io
   $sudo systemctl start docker
   $sudo systemctl enable docker
   ```
   
   - Clone o projeto do github:
   ```
   $git clone https://github.com/matheusvargas481/jts.devops.2019.2/tree/tema-7
   ```
   
      ## Author
   
   **Matheus da S. L de Vargas** -  [GitHub](https://github.com/matheusvargas481)
   
