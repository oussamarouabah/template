# Template

This is a template that I created to create our services from scratch, I will discribe how we can use
to create new service but first let's discuss the required tools:

- gonew 
- task

gonew is used to create new projects from templates
task is a task runner kind of a replacement of Makefiles


To install both `gonew` and `task` you need to have golang installed on your machine, if not please install it and come back

## Install gonew

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

## Install task

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

read more on these tools here:

- https://go.dev/blog/gonew
- https://taskfile.dev/

## Create a service

Now its the moment of truth we will create a service using this template, to do that run the following command

```bash
gonew github.com/oussamarouabah/template github.com/basket-dz/service
```

you should see a new folder in you machine now please create a git repo and push that initial code.