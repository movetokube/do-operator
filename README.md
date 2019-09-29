# Digital Ocean Kubernetes Operator
![](https://github.com/movetokube/do-operator/workflows/Tests/badge.svg) ![](https://github.com/movetokube/do-operator/workflows/Publush%20Release/badge.svg)

<p align="center">
    <img src=".assets/DO_Logo_horizontal_blue.svg" />
</p>

This operator provides Digital Ocean Cloud resources via Kubernetes Custom Resource Definitions (CRDs).
API Group is `do.movetokube.com`. Current version is `v1alpha1`. Example CR `apiVersion`: `do.movetokube.com/v1alpha1` 

## Supported Digital Ocean resources
### DNS

Any DNS record is supported by this operator. In order to create a DNS record, you need to register a domain name first.
After domain name is registered in DigitalOcean, use the following CR to create any record for any domain registered in DO.
Here is an example of the CR, which creates an A type record for domain example.com. Hostname is set to "do-operator", which
points to 127.0.0.1 IP address. TTL is set to 600 seconds.

```yaml
apiVersion: do.movetokube.com/v1alpha1
kind: DNS
metadata:
  name: do-op
  namespace: sandbox
spec:
  # common fields for most records
  domainName: example.com
  recordType: A
  # name of a record, usually the hostname
  hostname: do-operator
  value:
    literal: 127.0.0.1
  ttl: 600
```

Full list of fields that the DNS CR supports is showcased below:

```yaml
apiVersion: do.movetokube.com/v1alpha1
kind: DNS
metadata:
  name: do-op
  namespace: sandbox
spec:
  # common fields for most records
  domainName: example.com
  recordType: A
  # name of a record, usually the hostname
  hostname: do-operator
  value:
    # this will take IP address of an ingress from the same namespace
    #ref:
    #  ingressName: nameOfIngress
    literal: 127.0.0.1
  ttl: 3600
  # specific for SRV records only
  port: 80
  # specific for SRV and MX
  priority: 100
  # specific for CAA records only
  flags: 0
  tag: issue
```

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
