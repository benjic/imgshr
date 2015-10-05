
This is a simple REST API that facilitates sharing short URLs. It was written to
illustrate the usefulness of the [Adapter Design Pattern][0]. An initial design
is presented that is subjected to changing requirements which is remedied by
employing an adapter design pattern to incorporate a new library with a
differing method signature. The adapter implements the client interface and
applies the necessary indirection and logic to map calls to the client interface
to the adaptee interface.

## Usuage

Given you have a [valid Go environement][2] you can obtain the source and
dependencies via
    
    go get github.com/benjic/shrturl
    go get github.com/gorilla/mux
    go get gopkg.in/mgo.v2

The source will be available at `$GOPATH/github.com/benjic/shrturl`. To install
the service and start an instance you can run

    go install github.com/benjic/shrturl
    ./$GOPATH/bin/shrturl

The service uses the presence of the `MONGODB_URL` environment variable to
determine which persistence store to use. If `MONGODB_URL` is present the
service will attempt to use the mongo database at the given URI. This requires a valid
running instance at the given URI. The [MongoDB download seciton][3] offers a
variety of installations for getting a MongoDB instance running on your local
machine. Once this is running you can use the command `export
MONGODB_URL=mongodb://localhost` to enable the service to use the mongo
persistence store.

## Initial Design

The REST API allows users to POST new url's to the service and are returned a
shortened URL. When any user visits the shortened URL they are redirected to the
to the original URL. The original specification required that the service be set
up quickly and without dependencies. The service accomplishes persistence by
writing url mappings to in memory store on the host machine.

## Redesign

Since the service's initial release it has become very popular within the ACME
corporation. The employees have shared it with their friends and usage load on
the service has increased substantially. The ACME corporation overlords have
decided that the service is a gift to humanity and want to ensure the service
can endure the new load.

By profiling the service while under load from the surge of new usage indicates
that response times are growing linearly to the number of requests made to the
service. Further debugging shows that the persistent backing store is
introducing the large delays in response. The product stakeholders identify
scalability as a core requirement and have asked for a scalable solution. This
requires identifying a new persistence store that is distributed and fast.

[MongoDB][1] is identified as a persistence store of choice. This decouples the
service application from the persistence store allow each to scale independently
and adapting to the new load on the service. Unfortunately, the libraries for
MongodDB do not match the existing client interface. Fortunately, the engineers
received a proper education and are aware of an adapter pattern that will allow
them to integrate the new design changes.


## Application of the Adapter Pattern

To refactor the new persistence store an adapter pattern is employed to
encapsulate the new library interface. The new library is written to implement a
FastStorer interface which has reciprocating function calls for inserting,
listing, and finding models but has differing method signatures as well as model
types. A model is a type the encapsulates the data fields of the store items.

A new adapter is created by a factory function that takes any type that
implements the FastStorer interface and returns an instance of the
URLFastStorerAdapter type. This type implements the client Storer interface and
maps function calls to the encapsulated FastStorer instance. When models are
used as parameters or returned as results the adapter maps the underlying
library model types to the client model types. This ensures the adapter fully
implements the client interface.

The MockFastStorer type is an example of how the object adapter provides the
benefit of being to adapt a variety of sub types. This allows an effective
testing of the adapter by exercising the client interface and asserting that the
appropriate calls are made by the encapsulated FastStorer type. Because this
relationship is defined by interfaces it assures that the adapter tests will
always test this relationship or will fail to compile if the interfaces every
change.

## Conclusion

This project is an example of a real world use case which offers a non
functional need for scalability. This problem is effectively resolved by using a
new persistence store. The library to  use the store has a mismatched interface
and cannot be consumed by the client directly. An adapter is defined and
implemented that allows the client to consume the new store. In addition, a mock
persistence store is implemented that allows the assertion of the correctness of
of the new adapter.

[0]: https://en.wikipedia.org/wiki/Adapter_pattern 
[1]: https://www.mongodb.org/
[2]: https://golang.org/doc/install
[3]: https://www.mongodb.org/downloads#production
