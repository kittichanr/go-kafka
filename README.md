# go-kafka

Youtube for learn [link](https://www.youtube.com/watch?v=RjtIdUOpH04):

Repo [link](https://github.com/codebangkok/golang/tree/main/project/gokafka):


> **Note:** This Repo can run only local machine


### Prerequisite

1. [optional] install apache kafka cli on macos -> https://hevodata.com/learn/install-kafka-on-mac/#t5

2. install mysql database -> https://flaviocopes.com/mysql-how-to-install/

### Start project

0. should open project with workspace
```
open vscode -> open file -> select gokafka.code-workspace in repo
```

1. Run Docker Compose for Create Kafka server container on your machine

```
cd server

podman-compose up -d or docker-compose up -d
```

2. open new terminal and run consumer service

```
cd consumer

go run main.go
```

3. open new terminal and run producer service
```
cd producer

go run main.go
```

4. open new terminal or postman for call api to test producer/consumer via kafka
```
// example curl

// openaccount
curl -H 'content-type:application/json' localhost:8000/openaccount -d '{ 
        "AccountHolder" : "Petch",
        "AccountType"   : 1,
        "OpeningBalance": 1000
}' -i

// depositfund
curl -H 'content-type:application/json' localhost:8000/depositfund -d '{ 
        "id" : "65485bcb-44e1-4317-bb70-e2a92f045a87",
        "Amount"   : 500
}' -i

// withdraw
curl -H 'content-type:application/json' localhost:8000/withdraw -d '{ 
        "id" : "65485bcb-44e1-4317-bb70-e2a92f045a87",
        "Amount"   : 500
}' -i

// closeaccount
curl -H 'content-type:application/json' localhost:8000/closeaccount -d '{
        "id" : "65485bcb-44e1-4317-bb70-e2a92f045a87"
}' -i

```

5. see on your database about data that come form producer/consumer via kafka topic that set in the repository