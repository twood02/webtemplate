# Import socket module
import socket      

# Reserve a port
port = 8080   
# Create a socket                
s = socket.socket()
# local machine as host
host = socket.gethostname()
# Bind host to reserved port
s.bind((host, port))
# Wait for the client to connect
s.listen()

print ('Server listening....')

while True:
   # Establish a connect with the client
    conn, addr = s.accept()
    print ('Got connection from', addr)
    data = conn.recv(1024)
    print('Server received', repr(data))

    filename='test.txt'
    f = open(filename,'r')
    l = f.read(1024)
    while (l):
       conn.send(l)
       print('Sent ',repr(l))
       l = f.read(1024)
    f.close()

    print('Done sending')
    #conn.send('Thank you for connecting')
    # close the connection between client and server
    conn.close()
