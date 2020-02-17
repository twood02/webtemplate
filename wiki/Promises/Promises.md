---
layout: page
title: Javascript Promises!
permalink: /wiki/Promises
---

*by:* Linnea Dierksheide and Cat Meadows

# Promises! In Javascript!!

--

## Baking cookies → Promises

Imagine you’re hanging out with your friend. You both want to bake cookies, but you don’t have the ingredients you need. There are two options: You both go to the store together. The store might have all the ingredients, in which case, after coming back from the store, you can make cookies. If the store doesn’t have all the ingredients, you might make banana bread instead. Or, you go to the store and promise your friend that you’ll be back soon, either with all the ingredients or without them. Your friend can work on her Networking homework while she waits for you. Then, when you come back you guys can make cookies if the store was stocked or banana bread if not. It seems like the second scenario is more productive -- only one person has to go to the store and your friend can be productive while you’re at the store. However, it only works because your friend knows you’ll come back from the store and won’t abandon her to go hang out with someone else. By making this promise, you can more effectively plan your actions, even if you don’t have an exact timeline for how long something will take.

Now, how does this relate to anything in Computer Science? 

Well, Javascript is an event-based language. Pressing buttons on a website might trigger a call to a database or have to reload new data on the page, and what happens might depend on which buttons are pressed and in what order. It needs some way to keep track of what to do when events happen without assuming a predictable order.
The solution: **Promises!**

## What is a promise?
So, what actually is a Promise? Just like it sounds, it’s used to ensure -- or promise -- the occurrence of a future event. But, Javascript knows that sometimes promises can’t always be fulfilled, so it accounts for two results -- when the promise is fulfilled and when it is broken. Because of this, a promise has only four states:
* *Resolved*: success!
* *Rejected*: failed!
* *Pending*: still waiting, promise hasn’t been rejected or failed
* *Settled*: promise is done


It’s important to note that once rejected, a promise can never be resolved and vice-versa. Once resolved/rejected, a promise goes directly to being settled, and can never be re-activated. Maybe a state diagram would help:

![/images/StateDiagram.jpg](/wiki/Promises/images/StateDiagram.jpg)

Since there are really only two results/actions associated with a Promise, when we create a Promise object, we pass in the resolve and reject methods. Later, we use the Promise.then() and .catch(), which define what to do when the Promise is resolved or rejected, respectively. 

Let’s look at an example:
![/images/PromiseSimple.jpg](/images/PromiseSimple.jpg)


Here, you can see that we create a Promise called promiseSimple. We pass it a function that will resolve the promise if we generate a random number less than 7 and rejects the promise if the number is greater than or equal to 7. If resolved, the outcome is in the .then() and if rejected, we’ll see the outcome defined in the .catch(). We see this when we run the code:
![/images/PromiseSimpleOutput.jpg](/images/PromiseSimpleOutput.jpg)

## But Why? 
Promises seem pretty straight-forward, but why and when should we use them? Because JavaScript is often used for web development, it needs to support multiple, interactive web elements. This is accomplished through asynchronous programming, usually implemented with callbacks or promise functions. Asynchronous programming allows for multiple things to be run at the same time, without having to wait on a function or result to continue executing. Just think of any website that you’ve ever visited and imagine how many things are running at the same time--listening for input from the user, validating inputs, fetching data from a database, and beyond. 
 
Since Javascript is single-threaded, we can’t rely on concurrency in order to work on many things at once, hence, traditionally, callback functions have been used to accomplish asynchronous programming in JavaScript. Callback functions simply return a result when the function has completed execution. They are typically passed as arguments into other functions which will then execute once the callback has returned--meaning that functions rely on results of previously executed functions. It’s easy to see where this can start to get messy. Having so many dependencies and nested functions for error handling of callbacks has become known as “callback hell”. Though callbacks allow us to implement asynchronous programming, they quickly become unreadable and complex because of how much they depend on other results and callback functions, but we have a solution: Promises! Rather than having to wait for a callback to complete to get a result and then continue execution of a function, we can return a promise object with a callback attached to it. And because promises are returned regardless of whether they’ve been fulfilled or not, our code is cleaner, more readable, and has less dependencies (i.e. no more callback hell!). Take a look at the example below and notice how error handling is a lot smoother as well. If any of the promises are rejected, a single error is caught and the program exits. 

## Callback vs. Promise
<ul id="slider">
    <li<img src=/images/Callbacks.jpg"></li>
    <li<img src=/images/PromiseBetter.jpg"></li>
</ul>
  
As we discussed, a webpage is trying to do a lot at once -- this will mean that we have a bunch of promises that will be fulfilled/rejected intermittently. It’s possible that we want a little more control over this. Maybe we don’t want to redirect to a new page until we’ve entered information into all fields, and each field is defined with a promise. It could also be that we only want to redirect if everything is entered successfully. 

To do this, we have Promise.all(), Promise.settled(), and Promise.race():
* *Promise.all()*: This only occurs when every promise associated with it is resolved. 
Since this a pretty common operation, let’s look at an example. Say we have three promises: Promise1, Promise2, and Promise3. If all promises are resolved, it’s a success. If any one of them fails, it’s an error. The code for using Promise.all() to achieve this functionality is as follows: 
![/images/PromiseAll.jpg](/images/PromiseAll.jpg)

* *Promise.settled()*: The code defined by the function is executed when every promise is finished (whether they are resolved/rejected doesn’t matter) 
* *Promise.race()*: This occurs as soon as one associated promise is resolved. 
It’s difficult to see where this would really be useful, but one example would be for performance testing. If we had three processes and wanted to know which one was the fastest, we could use Promise.race() to find out. 



  
  










