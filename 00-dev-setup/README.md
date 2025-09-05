# 00-dev-setup

This project uses [Dev Containers](https://containers.dev/) for a consistent, containerized development setup.

## Quick Start

Each chapter contains its own `devcontainer.json` with the specific dependencies needed for that chapter.
To get started:

1. Install [VS Code](https://code.visualstudio.com/)
2. Install the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
3. Open the chapter folder in VS Code
4. When prompted, **Reopen in Container**

## Notes

The dev container in each chapter is self-contained and evolves as the learning project progresses.

## Kubernetes Cluster

A local Kubernetes cluster is automatically created using [kind](https://kind.sigs.k8s.io/) whenever the chapter is opened in the dev container.  
You can verify the installation and see running pods with:

```sh
kubectl get pods -A
```

This ensures that the cluster and all necessary tools are ready **each time you reopen the chapter in the container**.

## Next

Move on to [Chapter 1 – Simple gRPC Client/Server](../01-simple-grpc) and start building our first gRPC service!
