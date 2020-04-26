# Import socket module
import socket      
# import time module; use to track send/receive time
import time 
# import os; use to get file size 
import os

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

    # get file
    filename='test.txt'
    # get file size (in bytes) and convert to bits
    fsize = (os.path.getsize(filename))*8 #1 byte = 8 bits

    #open file, read, send
    f = open(filename,'rb')
    l = f.read(1024)
    while (l):
       conn.send(l)
       print('Sent ',repr(l))
       l = f.read(1024)
    # close file
    f.close()
    if f.closed:
       break
# calculate time taken to send given file
runtime = time.time() - start;
print('Done sending')
print('total send time: %.2f ', runtime)
# calculate and output throuhgput measured in bits/second
print('throughput: %.2f ', fsize/runtime)

# close the connection between client and server
conn.close()
print('server connection closed')

