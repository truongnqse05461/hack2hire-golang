# OneMount Hack2Hire Golang Project
>    This repository contains 1 service and 1 worker working independently. But they can use the other resources in the
>    repository. Please find their main.go file in `service` and `worker` folders.
    
Please read this guide carefully. If you have the issues with this guide, please feel free to send us your feedback to get support from us.

## Prequiresites:
> Hardware:
> You can make the hack with the PC, laptop or other hardware devices, but we tentatively recommend the laptop with OS installed.
> Your laptop's specifications will make you faster than the others.

- The specification required: 

```
- CPU Intel Core i5 with 4 cores of CPUs or higher
- 4GB Memory or higher
- 128GB SSD storage or higher
```

- The specification recommended: 

```
- CPU Intel Core i7 with 8 cores of CPUs or higher
- 8GB Memory or higher
- 256GB SSD storage or higher
```

> You need to install OS and softwares listed bellow:

- Windows, Linux OS, Macosx are accepted - We're assuming your computer had the OS installed
- IDE: Goland, VSCode, Sublime, Vim, Nano, Emacs, ... etc
- Docker, Docker Compose
- Golang: version 1.18 or higher
- MySQL Client: Workbench

## Guide to setup
### I. Installation
#### 1. IDE
> Depend on your preferring, the guide for setup some IDEs are listed below:

- [Install Goland](https://www.jetbrains.com/help/go/installation-guide.html)
- [Install VSCode](https://code.visualstudio.com/download)
- [Install SublimeText](https://www.sublimetext.com/3)

#### 2. Install Docker

> You can follow the guide below to setup Docker based on your OS:

- [Install Docker on MacOsx - CPU Intel](https://docs.docker.com/desktop/install/mac-install/)
- [Install Docker on MacOsx - CPU M1](https://docs.docker.com/desktop/mac/apple-silicon/)
- [Install Docker on Windows](https://docs.docker.com/desktop/install/windows-install/)
- [Install Docker on Linux](https://docs.docker.com/desktop/install/linux-install/)

#### 3. Install Golang

> Please follow the guide below:

- [Install golang newest version](https://go.dev/doc/install)

#### 4. Install Workbench

> Please download this packages from the guide below:

- [Download Workbench](https://dev.mysql.com/downloads/workbench/)

### II. Setup your local dev environment

#### 1. Fork this project to your own project on our gitlab

- Your **Gitlab account** will be included in the `Onboard Email` that you received.
- Please add your public key into **Gitlab account**
- Select the **Fork** button on this project, and fork it to your account.

#### 2. Clone forked project to your local machine

> Please feel free to using any `git client` tool that you prefer or use our way to clone the project:

- If you're using a Windows machine, please use `git-bash` to run the following command. Or if you're using MacOsx/Linux, please use `Terminal` instead.

```shell
git clone https://<git-project-uri>
```

#### 3. Setup dependencies with Docker

> This step will setup Kafka server on your local machine and verify the Docker installation is success or not.

- On Windows machines, please use `cmd` or `powershell console`, `git-bash` to run the following command. Or if you are using MacOsx or Linux, you can use `Terminal` instead.

```shell
cd docker
docker-compose -f docker-compose.yml up -d
```

And the result would be (or something like this):

```shell
Creating network "hack-2-hire-code-based_hack2hire" with driver "bridge"
Creating hack-2-hire-code-based_zookeeper_1 ... done
Creating hack-2-hire-code-based_kafka_1     ... done
```

- To verify the setting up, please run the following command on your `cmd/powershell/terminal` (depend on your OS):

```shell
$ telnet localhost 3306
Trying ::1...
Connected to localhost.
Escape character is '^]'.
```

```shell
$ telnet localhost 2181
Trying ::1...
Connected to localhost.
Escape character is '^]'.
```

```shell
$ telnet localhost 9092
Trying ::1...
Connected to localhost.
Escape character is '^]'.
```

**NOTE**: We will provide the dedicated MySQL DB that is owned by you, please use `Workbench` to access to MySQL DB, we will provide pHpMyadmin on the offline hacking day.

#### 4. Setup the project's code environment

##### 4.1 Enable custom.env file

With `Goland`, install EnvFile plugin by:
- Preference -> Plugins
- Search for EnvFile (Borys Pierov) -> Install

For each service or worker:

- Choose "Select Run/Debug Configuration"
- Choose "Edit Configurations"
- Click "+" icon
- Choose "Go Build"
- Create name "go build hack2hire-2022/service" or "go build hack2hire-2022/worker"; runkind: "Package"
- Choose "EnvFile" then tick "Enable EnvFile"
- Click "+" icon and add ".env file" and add "custom.env"

Or you can use the `go` command to build with your own `.env` file instead.

##### 4.2 Database and migration
- This skeleton use `golang-migration` for Database Migration, see https://github.com/golang-migrate/migrate
- Check `db/migration` folder for example
- Put your new `.sql` file for migration if you want to create more databases or tables

##### 4.3 Projects structure

```
├──db
│   └──migration  //put your new migration file here to create databases and tables
│
├──dtos
│  └──http.go  //where you define your data transfer objects
│   
├──model
│  └──message.go   // where you define your models
├──service         //this is your service structure which contains a main.go file, config, handler and router files
│  ├──config
│  │  └──config.go
│  ├──handler
│  │  └──handler.go
│  ├──router
│  │  └──router.go
│  └──main.go
├──services  //these are your internal services which contain your business logic
│  ├──db.go
│  └──sample.go
└──worker  // this is your worker structure which contains a main.go file and a config file. You may want to add your code to process your consumed data
   ├──config.go
   └──main.go
```

- We start this project with [gin](https://github.com/gin-gonic/gin) for webserver and some others golang libraries to support the application.

- To install `gin`, please follow the command:

```shell
go get -u github.com/gin-gonic/gin
```

- Or use the go command to install the projects dependencies:

```shell
go mod vendor
```

#### 5. Start the project

> Please run the following commands: (on `cmd/powershell/git-bash/terminal`)

- Start the **Service**:

```shell
go run service/main.go
```

- Start the **Worker**:

```shell
go run worker/main.go
```

- Verify success or not

```shell
# This command would be ran on Linux or MacOsx
curl --location --request GET 'http://localhost:8080/health'
```

> Or you can access `http://localhost:8080/health` on your browser instead

The success result would be:

```
{"status":"running"}
```

**NOTE AGAIN**: Please feel free to reach us out when you have any issues with this guide to get the support from us.
