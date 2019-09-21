# Digital Ocean Kubernetes Operator

This operator provides Digital Ocean Cloud resources via Kubernetes Custom Resource Definitions (CRDs).
API Group is `do.movetokube.com`. Current version is `v1alpha1`. Example CR `apiVersion`: `do.movetokube.com/v1alpha1` 

## Supported Digital Ocean resources
### DNS

TBD

## Contribution guide
This operator is built with operator-sdk v0.10.0 and is using Go modules. In order to contribute, please fork this project
and then follow the steps below to setup the environment.
1. Install operator-sdk CLI tool
    ```shell script
    # Set the release version variable
    $ RELEASE_VERSION=v0.10.0
    # Linux
    $ curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
    # macOS
    $ curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-apple-darwin
    ```
2. Install the following tools
    - [git][git_tool]
    - [go][go_tool] version v1.12+.
    - [mercurial][mercurial_tool] version 3.9+
    - [docker][docker_tool] version 17.03+.
      - Alternatively [podman][podman_tool] `v1.2.0+` or [buildah][buildah_tool] `v1.7+`
    - [kubectl][kubectl_tool] version v1.11.3+.
    
3. Download and verify the Go packages
   *  `$ go mod download`
   *  `$ go mod verify`
4. To develop and test the operator a Kubernetes cluster is required. For local dev, [minikube][minikube_tool] is perfect.
5. Follow the operator-sdk [guick start guide](https://github.com/operator-framework/operator-sdk/blob/master/README.md#quick-start)
to get up to speed.
6. Open a PR after new features are tested and ready.

    
[git_tool]:https://git-scm.com/downloads
[go_tool]:https://golang.org/dl/
[mercurial_tool]:https://www.mercurial-scm.org/downloads
[docker_tool]:https://docs.docker.com/install/
[podman_tool]:https://github.com/containers/libpod/blob/master/install.md
[buildah_tool]:https://github.com/containers/buildah/blob/master/install.md
[kubectl_tool]:https://kubernetes.io/docs/tasks/tools/install-kubectl/
[minikube_tool]:https://kubernetes.io/docs/tasks/tools/install-minikube/
