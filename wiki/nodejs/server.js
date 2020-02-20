/* A simple Node.js server: echoing client's data and responsing with a message.
* This code snippet is used for a tutorial of Node.js single-threaded architecture.
* If you are interested, please checkout: [link]
*/

// import net module
const net = require('net');
// port is hard-coded, use assigned port when in production
const port = process.env.PORT || 3000;

// Create a TCP server, createServer() will automatically set a connection listener
const server = net.createServer((socket)=>{
    console.log(`A new request from port: ${socket.remotePort} is connected`);

    // Listen for event: "data ready" on the socket
    socket.on('data', (data) => {
        console.log(`Data from port: ${socket.remotePort} is ready: ${data}`);
        
        // Write back to client
        socket.write(`Data has been received`, ()=> {
            console.log(`Message has been reponded to request: ${socket.remotePort}`)
        });
    });
    // Listen for event: "close" on the socket
    socket.on('close', () => {
        console.log(`Request from ${socket.remotePort} has closed`)
    })
});

// Listen for incoming request from port
server.listen(port, ()=> {
    console.log(`Server is running on port: ${port}`)
})