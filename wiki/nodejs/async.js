console.log('Hi! This is a demo of async programming.');

setTimeout(() => {
    console.log("Hello from callback of set time out!")
}, 5000);

console.log("Set time out is not blocked!");