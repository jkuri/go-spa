# go-spa

Go (Golang) SPA (Single Page Application) Template/Boilerplate.

This template allows you to embed your SPA files and build single executable application with single command.

## 1. Build

Just put your SPA files into `./dist` directory and run

```sh
$ make
```

Application will be built into `./build/app`.

```sh
$ ./build/app
```

Open your browser at `http://localhost:8080` and you should see your app running.

## 2. Docker Build

Put your SPA files into `./dist` directory and run

```sh
$ docker build -t go-spa .
```

This will build a very slim Docker image from `scratch` which includes only application built.

```sh
$ docker run -p 8080:8080 go-spa
```

Open your browser at `http://localhost:8080` and you should see your app running.
