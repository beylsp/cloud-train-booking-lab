# cloud-train-booking-lab

This repository is a **personal learning project** focused on exploring **Kubernetes-based microservices** and **cloud-native application design and development**. It is organized as a series of **progressive chapters**, each building upon the previous one.

## Project Theme

The project is centered around building an **online train booking system**, which serves as the reference application throughout the chapters. This fictional platform includes typical components of a modern distributed system — such as order placement, tracking and notification services.

## Chapter Overview

| Chapter | Title                  | Description                                                        |
|---------|------------------------|--------------------------------------------------------------------| 
| 00      | Development Setup      | Devcontainer.                                                      |
| 01      | gRPC Client and Server | Minimal Booking service with gRPC client and server.               |
| 02      | ConnectRPC             | Add ConnectRPC to expose BookingService over gRPC and HTTP/JSON.   |
| 03      | ko.build               | Build and deploy the BookingService into a KinD cluster using ko.  |
| 04      | MetalLB                | Expose services via LoadBalancer.                                  |
| 05      | Envoy Gateway          | Add an API Gateway for routing, security and customer-facing APIs. |

## Getting Started

```bash
git clone https://github.com/beylsp/cloud-train-booking-lab.git
cd cloud-train-booking-lab
```

**Ready to dive in?** Start by setting up the [development environment](00-dev-setup/README.md).
