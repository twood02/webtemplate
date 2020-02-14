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

# Is Node.js really a single-threaded architecture?

Now we have a clear mind about what Node.js is. Is Node.js really single-threaded:question: It seems confusing because JavaScript is laying on top of c++, which is multi threaded and how Node.js respons to concurrent requests efficiently. For example, if there is a request to download 10GB text file, the thread will be blocked until the work is done.
> Concurrent Processing: A sequence of requests will be performed interleaved (ping-pong).

In Node.js there are two types of threads: **one Event loop (main thread)**, and **thread pool**.

## What is Event Loop?

Apparently event loop is a semi-infinite `while` loop, running in the main thread. Why Node has a event loop :question: Event loop is what allows Node.js, as an async platform, to perform non-blocking I/O operations.
> "I/O" primarily refers to interaction with system's disk and network supported by `libuv`.

[wrap into a graph]For example, server side socket establishes a reliable connection with a client socket using `accept()`. But there have not any data that transfer between them. What the event loop does is to ask OS to inform socket when the data is ready.

Who will handle this job? Node.js will offload the jobs to kernel. Node.js has a collection of **file descriptors** that it asks the OS to monitor, using a mechanism, like `epoll` in Linux. It seems like what the Node.js does is to tell the kernel,"what I am interested in", "when the events happened please tell me.".

> OS typically provides event notification interfaces for asynchronous I/O (epoll in linux, kqueue in macOS, IOCP in Windows etc.).

Because of this mechanism, i.e., OS does job without back and forth threads, it happens in the main thread, the single thread. Once the job is done, OS will signal the event loop and then event loop invokes the **callbacks** associated with the event and appends them into the event queue. 

> A callback is a function called at the completion of a given task.

*need code to illustrate*

### Event Loop Explained

For each iteration, Node will go through several phases.

[chart]

- **timers**: this phase executes callbacks scheduled by `setTimeout()` and `setInterval()`.
- **pending callbacks**: executes I/O callbacks deferred to the next loop iteration.
- **idle, prepare**: only used internally.
- **poll**: retrieve new I/O events; execute I/O related callbacks (almost all with the exception of close callbacks, the ones scheduled by timers, and `setImmediate()`); node will block here when appropriate.
- **check**: `setImmediate()` callbacks are invoked here.
- **close callbacks**: some close callbacks, e.g. `socket.on('close', ...)`.

## When is Node multi-threaded?
Node.js (`libuv`) offers a pre-allocated thread pool with a default size of 4. Will the number of threads expand❓The default size of the pool can be overridden by setting the environment variable UV_THREADPOOL_SIZE. The thread pool is responsible for blocking `I/O` operations and CPU intensive tasks. For example, `crypto` module make use of thread pool.

*need code to illustrate `dns.lookup()`* 


# Summary

