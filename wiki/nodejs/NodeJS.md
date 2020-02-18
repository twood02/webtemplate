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

# Single-threaded Architecture
This tutorial assume you are familar with the term-callback before. If you don't know callback before. There is a quick demo show you how async programming works.
> A callback is a function called at the completion of a given task.

Let's take a look at the sequence of the code is executed.
```javascript
console.log('Hi! This is a demo of async programming.');
setTimeout(() -> {
    console.log("Hello from callback of set time out!")
}, 5000);
console.log("Set time out is not blocked!");
```
*[chart - function call in the stack]*
JavaScript is a single thread language which means it can only do one thing at a time. How does Node responde to concurrent requests? Let's go through a simple scenerio: server is responding to a sequential `connect()` requests, but no data transimission appear.
The traditional server architecture performs like this:
*[chart-multithreaded architecture]*
```javascript
const http = require('http');
//bind()
//listen()
//accept()
```
Node single threaded architecture performs like this:
*[chart-singlethreaded architecture]*
When the task is completed, there is a **callback** pushed into queue waiting for processing.

## But how...?
You can see that Node works quite well when it comes to handling I/O operations asynchronously. You might have some questions in mind.
> "I/O" primarily refers to interaction with system's disk and network.

- Who will handle blocking I/O tasks?
Node will offload the jobs to **kernel**. Node take advantage of the fact that modern operating systems’ kernels are multi-threaded. It seems like what the Node does is to tell the kernel,"what I am interested in", and "please tell me when you are done".
- How does Node know that which time to handle callbacks?
Once the I/O task is finished, the callback will be pushed into a callback queue instead of executing the callback right away.

But how? Specifically, Node asks a friend, [**`libuv`**](http://docs.libuv.org/en/v1.x/api.html), to help Node manage async I/O operations.

## Say Hello to Event Loop

Event loop takes care of polling for I/O and scheduling callbacks to be run based on different sources of events. Apparently event loop is a semi-infinite `while` loop, running inside function `int uv_run()` and in the main thread.

Continue on previous question... How `libuv` tells kernel what I am interested in? Library `libuv` has a collection of **file descriptors** that it asks the OS to monitor, using a polling mechanism, like `epoll` in Linux. 

> OS typically provides event notification interfaces for asynchronous I/O (epoll in linux, kqueue in macOS, IOCP in Windows etc.).

Because of this mechanism, i.e., OS does job without back and forth threads, it happens in the main thread, the single thread. Once the job is done, OS will signal the event loop and then event loop invokes the **callbacks** associated with the event and appends them into the event queue. 

*need chart to illustrate - even of A socket has data ready to be read*

### A closer look

You are definitely interested in what will be iterated inside each loop. Each iteration will go through 7 phases

[chart-detailed stages of event loop]
#### Timers
This phase executes callbacks scheduled by `setTimeout()` and `setInterval()`. Timers callbacks will be executed as early as the specified amount of time has passed; however, Operating System scheduling or the running of other callbacks may delay them.

#### Poll
Poll phase retrieves new I/O events; executes I/O related callbacks (almost all with the exception of close callbacks, the ones scheduled by timers, and `setImmediate()`)

- **pending callbacks**: If the *previous* iteration deferred any I/O callback it will be run at this point.
- **idle, prepare**: only used internally.
- **poll**: retrieve new I/O events; execute I/O related callbacks (almost all with the exception of close callbacks, the ones scheduled by timers, and `setImmediate()`); node will block here when appropriate.
- **check**: `setImmediate()` callbacks are invoked here.
- **close callbacks**: some close callbacks, e.g. `socket.on('close', ...)`.

## Is Node single-threaded?
Actually NOT really! Node provides two types of threads: **event loop (main thread)**, and **thread pool**, both of them provided by `libuv`. Node.js (`libuv`) offers a pre-allocated thread pool with a default size of 4. The default size of the pool can be overridden by setting the environment variable `UV_THREADPOOL_SIZE`. The thread pool is responsible for blocking `I/O` operations and CPU intensive tasks. For example, `crypto` module make use of thread pool.

*need code to illustrate `dns.lookup()` or `getaddrinfo()`* 


# Summary
In summary, Node is not single threaded. Node provides two types of threads: **event loop (main thread)**, and **thread pool**, both of them provided by `libuv`.
# Reference

