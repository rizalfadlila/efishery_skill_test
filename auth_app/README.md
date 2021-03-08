# Auth App

### Installation

+ Make the environment files. Adjust your local configuration.
```bash
$ cp .env.example .env
```

+ Run `docker run` to build the dependency docker instances:
```bash
$ docker build -t efishery/authapp .
```

+ Run these command to run docker image
```bash
$ docker run --name -p 3000:3000 efishery/authapp:latest
```