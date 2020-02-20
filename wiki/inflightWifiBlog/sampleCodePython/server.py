# Import socket module
import socket      
# import time module; use to track send/receive time
import time 

# Reserve a port
port = 60000 
# Create a socket                
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
# local machine as host
host = socket.gethostname()
# Bind host to reserved port
s.bind((host, port))
# Wait for the client to connect
s.listen(5)

print ('Server listening....')

while True:
   # Establish a connect with the client
    conn, addr = s.accept()
    print ('Got connection from', addr)

    #start time from when connection begins
    start = time.time()

    filename='test.txt'
    f = open(filename,'rb')
    l = f.read(1024)
    while (l):
       conn.send(l)
       print('Sent ',repr(l))
       l = f.read(1024)
    f.close()
    print('Done sending')
    runtime = time.time() - start;
    print('total send time: %.2f ', runtime)
    # close the connection between client and server
    conn.close()
