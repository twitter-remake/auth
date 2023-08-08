# Auth Service

This is the authentication service for the whole twitter clone app. It is made with Golang with Firebase Auth as the authentication SaaS.

For now, it just syncs the user data from Firebase Auth to the postgres database.

## Motivation

Authentication and authorization are essential for any microservices architecture. They ensure that only authorized users can access the services, and that they can only access the resources that they are allowed to access.

Implementing your own authentication flow can be **complex** and **time-consuming**. This is especially true if you need to support a variety of authentication methods, such as username and password, OAuth, and social login.

Hence, I decided to use **Firebase Auth** as the authentication SaaS for this project.

Some other alternatives that I might consider in the future are AWS Cognito and Supabase Auth.

## Setting up Dev environment

Here you will understand how to run and setup the development environment for the auth service in your local machine.

***Prerequisites***
- [buf](https://buf.build/)
- [docker-compose setup](https://github.com/twitter-remake/compose)

**1. Protobuf**

First, you need to generate the protobuf files using [buf](https://buf.build/) protobuf cli tool. You can do this by running the following command:

```bash
cd proto
buf generate --debug --verbose
```

This command will generate the protobuf files in the `gen/{lang=go}/{service}` directory.

Next, There is an `init.sql` file, so next is to execute the queries in the file to create the tables.

**2. Local Docker**

Head to [github.com/twitter-remake/compose](https://github.com/twitter-remake/compose) for the docker-compose file and instructions.

**3. Configuration**

Create the `.env` file by copying the `.env.example` file and fill in the values.

Next is to create the `firebase-credentials.json` file which you can get from the firebase console `Settings > Service Accounts > Generate new private key`.

**4. Run the service**

```bash
go run main.go
```

I've provided a demo React UI for testing the auth service [here](demo)

## Support

If you have any questions or feedback, feel free to contact me via [email](mailto:juandotulung@gmail.com)

And if you like this project, consider buying me a coffee :)
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y8DFOVT)