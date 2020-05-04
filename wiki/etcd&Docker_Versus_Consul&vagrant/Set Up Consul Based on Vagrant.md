Set Up Consul Based on Vagrant
### What is Service Discovery?

In the framework of microservices, service discovery is a crucial module that must be mentioned.. Let's look at a picture below:

In the figure, an interface of the client needs to call service A-N. The client must know the network location of all services. In the past, the configuration was in the configuration file, and some configuration may be in the database. Here are a few problems raised by such mode:
