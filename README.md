# Hypershift Operator

## Overview
The Hypershift operator contains a set of controllers that manage the control plane components of a 
hosted control plane cluster created with the [Hypershift Installer](https://github.com/openshift-hive/hypershift-installer).

## Building

To build the operator binary locally, run:

```
make build
```

To build the operator image, use your container builder tool (`docker` or `buildah`) to build the `Dockerfile` at the root of this repository.
