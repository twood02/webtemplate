---
layout: page
title:  "Working with HTTP and REST"
permalink: /assignments/httprest/
---

> **Objectives:** To learn about HTTP requests, REST services, and JSON data in python, javascript, or go.

## 1. RestServ
You have been provided with a simple REST based server. All it does is... *rest*... for each request (i.e., sleep for a random amount of time).

Get the [Rest Service](https://github.com/gwAdvNet20/rest-latency-server):

```
git clone https://github.com/gwAdvNet20/rest-latency-server.git
```

The provided service is written in Go and includes binaries that will run on windows, OS X, and linux in `bin/`.  It listens on port `8080` for incoming HTTP requests.  


*You do not need to modify this code in any way.*  You can run the service on your laptop or on Cloud9 (you  might need to adjust security groups if you want to access it from another host).



## 2. REST Client

Instead of writing our own client to access RestServ, we will simply use `curl`. This is a great tool for issuing HTTP requests at the command line.

Once you start RestServ, you can access it with:

```
curl -X GET localhost:8080/request
```

The service will sleep for a random period of time before returning a JSON formatted response like:
```
{"status":200,"success":true,"message":"Request Successful","data":1}
```

But sometimes it sleeps too long, causing an error:
```
{"status":500,"error":"Internal Server Error","message":"Request timed out","data":39}
```

Note that the text shown by `curl` is the body of the returned document. If you want to see the HTTP-level headers as well, you can use:

```
curl -i -X GET localhost:8080/request
```

## 3. REST Gateway
Next you must write your own REST service, which we will call RestGate. Its job will be to act as a gateway that can issue requests to RestServ, like the following:

![Client--RestGate--RestServ](/assignments/restgate.png)

We suggest you implement your gateway with one of the following:

  - Python and [flask](https://palletsprojects.com/p/flask/)
  - Javascript with node.js and [Express](https://expressjs.com)
  - Go with [net/http](https://golang.org/pkg/net/http/) or [mux](https://github.com/gorilla/mux) (the provided RestServ uses mux)

Your service should expose an HTTP endpoint [localhost:4000/rest](http://localhost:4000/rest)

### Stage 1:
Initially, you should design your RestGate service so that for each incoming request, it issues its own HTTP request to the RestServ backend.

It should then return a simple HTML document like the following:
```
<html><body>
I RESTed for XXX time units.
</body></html>
```
where XXX is the value in the `data` field of the RestServ response. Remember, the response is in JSON format, so your langauge of choice should have a simple way to parse that.

## Stage 2:

In practice, most requests require some sort of data. It is also common for one request to cause multiple subrequests. To evaluate this, you should extend your RestGate so that for each incoming request it will issue `n` more HTTP requests to the provided RestServ service. The value of `n` should be a parameter passed to your service in an HTTP POST request. For each request, you must measure the time it takes to process in milliseconds.

To test your service, you should issue a command like:

```
curl -X POST localhost:4000/someEndPoint -d '{"numRequests": n }'
```

Your service will need to read the `numRequests` field in the POST request body to find the value of n.

It should then return a simple HTML document that prints the total time recorded for making n requests like so:

```
<html><body>
I made N requests and it took XXX time units.
</body></html>
```

# Extending your RestGate

If you finish early you should implement one of the following advanced api features:

- Extend RestGate to read the Accept header on a request and return the content in that format. i.e. If I send a header of `-H Accept: text/html` I should get back HTML. If I send a header of `-H Accept: application/json` I should get back json with the total time value for all `n` requests and the data values returned for each request.
- Extend RestGate to use some sort of authentication scheme. Using the `-H Authorization: 12345` header, have your server only respond to authorized requests that provide a secret code and return an appropriate error status code otherwise.