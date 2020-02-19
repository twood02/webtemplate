---
layout: page
title:  A Detailed Tutorial of NodeJS Single-threaded Architecture
permalink: /wiki/nodejs/
---

*by:* Mingyu Ma, Siqi Wang


A short description of your post goes here.

---
# Setup Environment

# Intro of Node.js

## How Node.js execute Javascript Program?

# Multi-threaded, Single-threaded, Sync, Async...
This tutorial assume that you are familiar with multi-threaded and single-threaded architecture, async and sync programming model.

*[chart-multithreaded architecture, sync]*
*[chart-singlethreaded architecture, sync]*

Question: How to improve performance (increase responsive, reduce latency) for a heavily I/O spplication?

ASYNC programming! Remember when you order food in Chopole, you don't block the order line and next customer could start to order food immediately. Similaly, in an event-based programming, there is generally a listener that listens for events, and then triggers a **callback** function when one of those events is detected. 
> A callback is a function called at the completion of a given task.

There is a comparison of async programming works in single-threaded and multi-threaded architecture.
*[chart-multithreaded architecture, async]*
*[chart-singlethreaded architecture, async]*

There is a quick demo of asyn programming:
```javascript
console.log('Hi! This is a demo of async programming.');
setTimeout(() => {
    console.log("Hello from callback of set time out!")
}, 5000);
console.log("Set time out is not blocked!");
```
*[graph- console result]*

Cheers! You have good understanding of Node.js single-threaded architecture, as an async platform.

## But how...?
Node is faster and less resource intensive when it comes to handling I/O operations asynchronously. You might have some questions in mind.
> "I/O" primarily refers to interaction with system's disk and network.

- Who will handle blocking I/O tasks?

Node will offload the jobs to **kernel**. Node takes advantage of the fact that modern operating systems’ kernels are multi-threaded. It seems like what the Node does is to tell the kernel,"what I am interested in", and "please tell me when you are done".
- How does Node know that which time to handle callbacks?

Once the I/O task is finished, the callback will be pushed into a callback queue instead of executing the callback right away.

But how? Specifically, Node asks a friend, [**`libuv`**](http://docs.libuv.org/en/v1.x/api.html), to help Node manage async I/O operations.


## Say Hello to Event Loop

`libuv` provides event loop which takes care of polling for I/O and scheduling callbacks to be run based on different sources of events. Apparently, event loop is a semi-infinite `while` loop, running inside function `int uv_run()` and in the main thread.

Continue on previous question... How `libuv` tells kernel what I am interested in? Library `libuv` has a collection of **file descriptors** that it asks the OS to monitor, using a polling mechanism, like `epoll` in Linux. 

> OS typically provides event notification interfaces for asynchronous I/O (epoll in linux, kqueue in macOS, IOCP in Windows etc.).

Because of this mechanism, i.e., OS does job without back and forth threads, it happens in the main thread, the single thread. Once the job is done, OS will signal the event loop and then event loop invokes the **callbacks** associated with the event and appends them into the **poll queue**.

The general picture of event loop.
*[general picture of event loop]*

### What will happen after `node server.js`?

Let's go through a piece of code to understand how event loop works.`server.js` is a simple Node application, printing out the data from incoming requests and reponsing with a message.

```javascript
// A simple Node.js server: echoing client's data and responsing with a message.
// This code snippet is used for a tutorial of Node.js single-threaded architecture. If you are interested, please checkout: [link]

// import net module
const net = require('net');
const port = process.env.PORT || 3000;

// Create a TCP server, createServer() will automatically set a connection listener
const server = net.createServer((socket)=>{
    console.log(`A new request from port: ${socket.remotePort} is connected`);

    // Listen for "data ready" event
    socket.on('data', (data) => {
        console.log(`Data from port: ${socket.remotePort} is ready: ${data}`);
        
        // Write back to client
        socket.write(`Data has been received`, ()=> {
            console.log(`Message has been reponded to request: ${socket.remotePort}`)
        });
    });
    // Listen for "close" event
    socket.on('close', () => {
        console.log(`Request from ${socket.remotePort} has closed`)
    })
});

// List on port
server.listen(port, ()=> {
    console.log(`Server is running on port: ${port}`)
})
```

Once we run `node server.js`, Node will start a process and our code will be executed inside event loop.

*[chart-console]*

### Deep Dive into Event Loop

Is it all about waiting for events, executing callbacks? The sequence of phases inside each iteration is like this:
*[chart-detailed stages of event loop]*
The phases below we care about most:
#### Timers
This phase executes callbacks scheduled by `setTimeout()` and `setInterval()`. Timers callbacks will be executed as early as the specified amount of time has passed; however, Operating System scheduling or the running of other callbacks may delay them.

#### Poll
Poll phase retrieves new I/O events; executes I/O related callbacks (almost all with the exception of close callbacks, the ones scheduled by timers, and `setImmediate()`).
Once the loop enters **poll** phase, one of the scenerios will happen:
- The poll queue is not empty, the event loop will iterate through its queue of callbacks executing them synchronously.
- The poll queue is empty.
    - If the scripts has `setImmediate()`, the event loop will end the poll phase and continue to the check phase to execute those scheduled scripts.
    - If the scripts don't have `setImmediate()` and the timers have been reached, the event loop will wrap back to the timers phase to execute those timers' callbacks.
    - If the scirpts don't have `setImmediate()`and the timers haven't been reached, the event loop will wait for callbacks and execute them immediately.

#### Check
`setImmediate()` callbacks are invoked here.

#### Other phases
- **pending callbacks**: If the *previous* iteration deferred any I/O callback it will be run at this point.
- **idle, prepare**: only used internally.
- **close callbacks**: some close callbacks, e.g. `socket.on('close', ...)`.


>Note: `libuv` uses a thread pool to make asynchronous file I/O operations possible, but network I/O is **always** performed in a single thread, each loop’s thread.

## Is Node single-threaded?
Actually NOT really! Node provides two types of threads: **event loop (main thread)**, and **thread pool**, both of them provided by `libuv`. The pre-allocated thread pool has a default size of 4. The default size of the pool can be overridden by setting the environment variable `UV_THREADPOOL_SIZE`. Unlike network I/O, there are no file I/O primitives `libuv` could rely on, so the current approach is to run blocking file I/O operations in a thread pool. The thread pool is responsible for blocking `I/O` operations and CPU intensive tasks. For example, `crypto` module make use of thread pool.

*need code to illustrate `dns.lookup()` or `getaddrinfo()`* 


# Summary
In summary, Node is not single threaded. Node provides two types of threads: **event loop (main thread)**, and **thread pool**, both of them provided by `libuv`. Node.js is highly scalable. Check out `express`, `socket.io`.

# Reference
1. [Node.js `net` module](https://nodejs.org/docs/latest-v11.x/api/net.html)

