# Import socket module
import socket

# Create a socket object
s = socket.socket()
# Get local machine name
host = socket.gethostname()
# Reserve a port to connect on
port = 8080                    

# connect to server
s.connect((host, port))
msg = 'Hello server!'
s.send(msg)

# buffered reader
with open('received_file', 'w') as f:
    print ('file opened')
    while True:
        print('receiving data...')
        data = s.recv(1024)
        print('data=%s', (data))
        if not data:
            break
        # write data to a file
        f.write(data)

f.close()
print('Successfully get the file')

#close connection
s.close()
print('connection closed')