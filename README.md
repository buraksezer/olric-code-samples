# Olric Code Samples

This repository contains examples of how to use [Olric](https://github.com/buraksezer/olric) in your projects. 

## What is Olric?

Distributed cache and in-memory key/value data store. It can be used both as an embedded Go library and as a language-independent service.

With Olric, you can instantly create a fast, scalable, shared pool of RAM across a cluster of computers.

## Data structures

### Distributed Map

Olric implements distributed map (DMap) data structure. With Olric, you can instantly create a large amount of distributed maps 
across a cluster of computers. DMaps are both available in embedded-member and client-server modes. 

#### Client-server mode

In this mode, an Olric cluster (via olricd or your own implementation) works somewhere in your network and your web server, 
backend services or maintenance scripts need to access the cluster. This is the most common way of using a database. 

[dmaps/client-server/main.go](dmaps/client-server/main.go) file explains how to initialize an Olric client in your application.

#### Embedded-member mode

In this mode, your application is also a member of Olric cluster. So you can create distributed data structures in your application.
This deployment mode is the recommended one, if cache locality is important for your use case. 

[dmaps/embedded-member/main.go](dmaps/embedded-member/main.go) file explains how to embed Olric into your application.

### Distributed Topic

TODO

## Distributed caching with Olric

TODO

## Kubernetes integration and automatic service discovery

## License

The Apache License, Version 2.0 - see LICENSE for more details.

## Contribution

Please feel free to open an issue, if you found a bug or suggest a new idea.