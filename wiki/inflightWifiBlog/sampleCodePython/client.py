# Import socket module
import socket
# import time module; use to track send/receive time
import time 

# Create a socket object
s = socket.socket()
# Get local machine name
host = socket.gethostname()
# Reserve a port to connect on
port = 60000                   

# connect to server
s.connect((host, port))

# buffered reader
with open('recv_file', 'wb') as f:
    print ('file opened')
    while True:
        print('receiving data...')
        data = s.recv(1024)
        print(data)
        if not data:
            break
        # write data to a file
        f.write(data)

f.close()
print('Successfully get the file')

#close connection
s.close()
print('connection closed')