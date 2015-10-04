# shrturl

> A url shortening API.

This is a simple REST API that facilitates sharing short URLs. It was written to
illustrate the usefulness of the [Adapter Design Pattern][0].

## Initial Design

The REST API allows users to POST new url's to the service and are returned a
shortened URL. When any user visits the shortened URL they are redirected to the
to the original URL. The original specification required that the service be set
up quickly and without dependencies. The service accomplishes persistence by
writing url mappings to a flat file on the host machine.

## Redesign

Since the service's initial release it has become very popular within the ACME
corporation. The employees have shared it with their friends and usage load on
the service has increased substantially. The ACME corporation overlords have
decided that the service is a gift to humanity and want to ensure the service
can endure the new load.

Initial reports indicate that the average response time from the service is
linearly dependent on the load of the system. The flat file backing store is to
blame as each operation depends on this file. Stakeholders for the service
indicate that a more efficient persistence service is necessary.

[MongoDB][1] is identified as a persistence store of choice. This decouples the
service application from the persistence store allow each to scale independently
and adapting to the new load on the service. Unfortunately, the libraries for
MongodDB do not match the existing client interface. Fortunately, the engineers
received a proper education and are aware of an adapter pattern that will allow
them to integrate the new design changes.

[0]: https://en.wikipedia.org/wiki/Adapter_pattern
[1]: https://www.mongodb.org/
