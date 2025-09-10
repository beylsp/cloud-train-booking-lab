# 03-build-ko

In this chapter, we take our `BookingService` (built with gRPC + ConnectRPC) and run it inside a **Kubernetes cluster**. We use:

- [ko](https://ko.build/) - to build and publish container images for Go services without Dockerfiles.
- **kind** – to create a local Kubernetes cluster for testing. 

## Tools

- ko – container builder for Go apps
- kind – Kubernetes-in-Docker
- kubectl – manage the cluster

These are preinstalled in the devcontainer.

## Setup

Under `config/` we have a single manifest file `booking.yaml` that defines two Kubernetes resources: a `Deployment` and a `Service`.

In the pod spec we reference the import path of the service's main package: `ko://cloud-train-booking-lab/03-build-ko/cmd/booking`.

When you run `ko apply`, the `ko` build tool will:

1. Scan the YAML files under `config/` for `ko://` image references.
2. For each `ko://` reference, build a container image from the specified Go import path (using the base image defined in `.ko.yaml`).
3. Push the resulting image to the configured repository.
3. Replace `ko://` string in the manifest with the fully qualified image reference.
4. If `KO_DOCKER_REPO=kind.local` is set, load the image directly into the local kind cluster instead of pushing to a remote registry.

## Build

From the chapter root you can either build the service image locally or build and deploy directly into kind:

- **Build and publish the image to the local Docker registry:**

```sh
make image
```

- **Build and deploy directly into the local kind cluster:**

```sh
make deploy
```

- **Remove the deployed resources from kind:**

```sh
make delete
```

To build the demo applications, run:

```sh
make demo
```

## Run

After running `make deploy`, the `BookingService` is launched inside the local kind cluster. You can verify the deployment with:

```sh
kubectl get deployments
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
booking   1/1     1            1           21h

kubectl get pods
NAME                      READY   STATUS    RESTARTS   AGE
booking-c5c4bb4c8-b7nf8   1/1     Running   0          21h

kubectl get services
NAME              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)     AGE
booking-service   ClusterIP   10.96.148.148   <none>        50051/TCP   21h
```

Since the service is running as a `ClusterIP`, it’s only accessible inside the cluster. To reach it from your host machine, forward the port:

```sh
kubectl port-forward service/booking-service 50051:50051
```

This will set up a TCP tunnel between localhost and the booking service.

Now, in another terminal, you can send requests to the service using either a gRPC client or a plain HTTP client (thanks to ConnectRPC).

```sh
./build/demo/booking-grpc-client 
Booking confirmed! PNR: PNR-12345, Status: CONFIRMED

./build/demo/booking-http-client | jq
{
  "pnr": "PNR-12345",
  "status": "CONFIRMED"
}
```

## Next

So far, our `BookingService` is only accessible outside the cluster through `port-forwarding`, which is fine for local testing but not practical for real usage.

In the next chapter, we will introduce [MetalLB](https://metallb.io/), a load-balancer implementation for bare-metal Kubernetes clusters. With MetalLB, we can:

- Define services as `LoadBalancer` instead of `ClusterIP`.
- Automatically assign external IP addresses from our LAN to these services.
- Access the `BookingService` directly via a LAN IP, without relying on `kubectl port-forward`.

Continue reading: [Chapter 4 - Expose services with MetalLB](../04-metallb)
