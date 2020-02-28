const net = require('net');
const PORT = 3000;

const client = new net.Socket();
client.connect(PORT, 'localhost', () => {
    console.log(`connect to server`);
    // Send data 5 seconds after connection
    setTimeout(() => {
        console.log(`sleep for 5 seconds`);
        client.write(`This is a client that send data after 5 seconds`);
    }, 5000);
})

client.on('data', (data) => {
    console.log(`message from server: ${data}`);
    client.end();
})

client.on('close', ()=> {console.log(`Connection closed...`);});