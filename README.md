# imgshr

> A image sharing API.

This is a simple REST API that facilitates sharing images. It was written to
illustrate the usefulness of the [Adapter Design Pattern][0]. This is done by
writing an initial implementation that uses a local file system as a backing
store for uploaded images. 

Due to popular demand the application is filling disk space at an astronomical
rate and the initial design choices are constraining the application.
Stakeholders have identified a cloud storage service that will address
scalability issues but the library interface is incompatible with the current
client implementation. The adapter is employed to address this incompatibility
as well as future proofing.
