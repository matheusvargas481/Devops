# Vault provision with Docker and Ansible with Jenkins

## Requirements

  - [Git](https://git-scm.com/)
  - [Jenkins](https://jenkins.io/)
  - [Ansible](https://www.ansible.com/)
  - [Docker](https://www.docker.com/)

## Installation

- Installing Git:
```
 $ apt-get install git
```

- Clone the github project:
```
 $ git clone https://github.com/ilegra/jovens-talentos.git
```

- Access the path /3-devops/matheus-vargas/Tema-19 and run the script with the following command to install docker, ansible, and jenkins 
```
 $ sudo ./installRequirements.sh
```

- Run the commands below to add the Jenkins root user to your machine.
```
$ sudo usermod -aG sudo jenkins
```

- Enter your machine's root and run the following commands to change the Jenkins root password.
```
$ mount -rw -o remount /
$ passwd jenkins
```
    Fill in the field that will appear with the desired password!

- After running the Jenkinks job, follow the configuration below..

## Configuring Jenkins

1. Add your username and password in the settings to access Jenkins:

    - Manage Jenkins / Configure Credentials;

2. Access your Jenkins - http://localhost:8080/

3. We should create a pipeline in Jenkins named BAKE:

- New Job
- Name Job and check Pipeline and OK
- Enter the created Job and click configure and add:
- In the General tab, add your description for the job; (Optional)
- In the General tab mark the field "This project is parameterized" and fill the field NAME with JENKINS_PASS and in the field DEFAULT VALUE fill in the password that was set for the root of Jenkins. 
- In the Pipeline category, select: Pipeline script from SCM
- In SCM: Select Git
- In the Repository URL, add: (`https://github.com/ilegra/jovens-talentos.git`)
- In credentials, add your credentials created in step 1;
- In Branches to build, add: refs/heads/tema-19
- In Script Path, add: 3-devops/matheus-vargas/Tema-19/vault-ansible/Jenkinsfile

After configuration run the job and wait for it to complete.

## Verifying Vault Functionality

- Open a terminal on your machine and run the following command.
```
 $ vault server -dev
```

- Open a new terminal and run the following command to expose the vault URL.
```
 $ export VAULT_ADDR="http://127.0.0.1:8200"
```

- At the same terminal run the following command to check if the vault is up.
```
 $ vault status
```

- After verifying it the Vault will be ready to access by the URL below.
```
 $ http://127.0.0.1:8200
```

# Using Vault with Docker

- Go to path 3-devops/matheus-vargas/Tema-19/vault-docker and run the next command.
```
$ ./manager-vault-docker {USING THE OPTIONS BELOW}
```
NOTE: The options must be written in lowercase.
- bake: to provision the image
- run: to run the application
- status: to check if the application is up
- stop: to stop the application

## Author

**Matheus da S. L. de Vargas** -  [GitHub](https://github.com/matheusvargas481)