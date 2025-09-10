# 04-metallb

In this chapter, we extend our deployment setup by adding [MetalLB](https://metallb.io/) to the kind cluster.
MetalLB acts as a load balancer for bare-metal Kubernetes clusters, allowing us to assign real LAN IP addresses to LoadBalancer services.

This removes the need for manual port-forwarding and makes services accessible just like in a cloud environment.

## Install MetalLB

MetalLB consists of two main components:

- **Controller** – allocates IPs from the configured address pool.
- **Speaker** – advertises service IPs to the local network (in our case via ARP broadcast).

In this lab environment, the controller and speaker are **already installed via the devcontainer setup**, so we only need to configure them.

## Configure MetalLB

Apply a simple configuration under `config/metallb.yaml`.

This defines a pool of IPs from the kind network and enables Layer 2 advertisement for them.

## Update Booking Service

Next, update the service definition in `config/booking.yaml` to use `LoadBalancer` instead of `ClusterIP`:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: booking-service
spec:
  type: LoadBalancer
  ...
```

## Build and Deploy

From the chapter root:

- **Deploy service with MetalLB integration**

```sh
make deploy
```

- **Build demo clients**

```sh
make demo
```

You can verify that MetalLB assigned an IP to the service by inspecting the pods and logs:

```sh
kubectl get pods -n metallb-system
NAME                          READY   STATUS    RESTARTS   AGE
controller-58fdf44d87-cxqzp   1/1     Running   0          24h
speaker-5gmhz                 1/1     Running   0          24h
```

Check assigned IPs:

```sh
kubectl get servicel2statuses -n metallb-system
NAME       ALLOCATED NODE              SERVICE NAME      SERVICE NAMESPACE
l2-5q92m   dev-cluster-control-plane   booking-service   default
```

Controller log showing IP allocation:

```sh
kubectl logs -n metallb-system controller-58fdf44d87-cxqzp
{"caller":"service.go:186","event":"ipAllocated","ip":["172.19.255.200"],"level":"info","msg":"IP address assigned by controller","ts":"2025-09-09T19:24:52Z"}
```

Speaker log showing IP advertisement:

```sh
kubectl logs -n metallb-system speaker-5gmhz
{"caller":"main.go:443","event":"serviceAnnounced","ips":["172.19.255.200"],"level":"info","msg":"service has IP, announcing","pool":"kind-pool","protocol":"layer2","ts":"2025-09-09T19:24:53Z"}
```

## Run

Provide the assigned service IP address as a command-line argument. It can be retrieved with:

```sh
kubectl get svc booking-service -o jsonpath="{.status.loadBalancer.ingress[0].ip}"
172.19.255.200

./build/demo/booking-grpc-client -H 172.19.255.200
Booking confirmed! PNR: PNR-12345, Status: CONFIRMED

./build/demo/booking-http-client 172.19.255.200 | jq
{
  "pnr": "PNR-12345",
  "status": "CONFIRMED"
}
```

## Next Chapter

With MetalLB, our internal services are now reachable on the LAN. However, clients are still connecting directly to backend services, which is not ideal for production.

In the next chapter, we will introduce [Envoy Gateway](https://gateway.envoyproxy.io/) as a true API gateway. Envoy will:

- Provide a single entry point for all external traffic.
- Handle routing to multiple services.
- Offer features like TLS termination, authentication, and observability.

Continue Reading: [Chapter 5 - Envoy Gateway](../05-envoy-gateway)
