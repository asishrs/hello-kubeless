# hello-kubeless

A basic [kubeless](https://kubeless.io/) example.

## Prerequisite

Below, software/tools are required to run this. 

- [homebrew](https://brew.sh/)
- A k8s cluster. The example is tested with [k3d](https://k3d.io/) Install k3d using `brew install k3d` 
- [Go task](https://taskfile.dev/#/) (trust me it will reduce the keystrokes!) `brew install go-task/tap/go-task`
- [kubeless](https://kubeless.io/) `brew install kubeless`
- [wrk](https://github.com/wg/wrk) for load testing the function `brew install wrk`

## Run

Once you install all the prerequisites, follow below steps.

### Deploy kubeless

```shell
task kubeless:deploy
```

In our case we are deploying the function as well as setting an auto scaling rule to scale the pods up to four based in CPU metric (70%)

### Deploy function

```shell
task deploy
```

It will few seconds for the function to be ready. You can check the __STATUS__ of function using `task list`

### Test the function

Invoke the function using `task call`. You can see _Hello world!_ in the output.

Let's create some load. Run `wrk -t5 -c50 -d300s http://localhost:8080/api/v1/namespaces/default/services/hello-go:http-function-port/proxy/` to generate the load. 

As the load on the function increase the CPU utilization goes beyond 70% metric threshold and kubeless scaling will start creating more pods (up to 4).

You can watch the overall CPU using `kubectl get hpa -w`

Once the load finishes, the function will scale back to one instance. 