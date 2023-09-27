
You can start a PostgreSQL container with the following command

    docker run -d --hostname rabbit-tutorial --name gp-prototype -p 15672:15672 -p 5672:5672 rabbitmq:3-management

Then start the app

    go run ./cmd
