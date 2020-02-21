   
   
### Background  
   
Secure Sockets Layer (SSL) is an enhanced version of TCP which can provide confidentiality, data integrity, and etc.   

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/bigpicture.png)    
    	
New applications such as e-commerce and online banking based on the World Wide Web greatly facilitate people's daily life and are favored by people. Because these applications need online transactions on the network, they put forward higher requirements for the security of network communication. The traditional World Wide Web Protocol HTTP does not have the security mechanism, such as transmitting data in the form of clear text, unable to verify the identity of both sides of the communication, unable to prevent the transmission of data from being tampered, etc., which leads to HTTP unable to meet the security requirements of e-commerce and online banking applications.    

The Security Protocol SSL proposed by Netscape company uses data encryption, authentication and message integrity authentication mechanisms to provide security guarantee for data transmission on the network. SSL can provide a secure connection for HTTP, which greatly improves the security of the World Wide Web.   

### Technical advantages       
1. Provide high security assurance. SSL uses data encryption, authentication and message integrity authentication mechanisms to ensure the security of data transmission on the network.        
2. Support various application layer protocols. Although SSL is designed to solve the security problem of the world wide web, because SSL is located between the application layer and the transport layer, it can provide security guarantee for any application layer protocol based on reliable connection such as TCP.   
3. Simple deployment. At present, SSL has become a global standard used to identify the identity of website and web browser, and to conduct encrypted communication between browser users and web servers. SSL protocol has been integrated into most browsers, such as ie, Netscape, Firefox, etc. This means that almost any computer with a browser supports SSL connection and does not need to install additional client software.    
     
### Protocol security mechanism      
1. The security mechanisms of SSL protocol include:
Confidentiality of data transmission: using symmetric key algorithm to encrypt the transmitted data.      
2. Authentication mechanism: Based on the certificate, the digital signature method is used to authenticate the server and the client, and the client's authentication is optional.                
3. Message integrity verification: MAC algorithm is used to verify message integrity during message transmission.        
4. Asymmetric key algorithm guarantees key security.     
5. Using PKI to ensure the authenticity of public key: PKI publishes users' public key through digital certificate, and provides a mechanism to verify the authenticity of public key. The digital certificate is a file containing the public key and identity information of the user, which proves the association between the user and the public key. The digital certificate is issued by the Authority CA, and the authenticity of the digital certificate is guaranteed by CA.    
 
### Working process of the agreement 
   
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/agreement.png)  
	
#### SSL handshake process 
 
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/handshake.png)      
SSL negotiates the session parameters between the client and the server through the handshake process and establishes the session. The main parameters of a session include session ID, the certificate of the other party, encryption suite (key exchange algorithm, data encryption algorithm, MAC algorithm, etc.) and master secret. The data transmitted through SSL session will be encrypted and Mac will be calculated by the master key and encryption suite of the session.   

#### Verify only the handshake process of the server
   
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/server.png)          
	
#### Verify SSL handshake process between server and client  
  
![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/server_client.png)       
	
#### SSL handshake process to restore the original session   

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/original_session.png)    
### Record protocol  

![image](https://github.com/wzli1214/gwAdvNet20.github.io/blob/dev/wiki/http1-2httpsSSL/images/record_protocal.png)    
	
The recording protocol is use after the client and server shake hands successfully, that is, after the client and server identify each other and determine the algorithm for security information exchange, they enter the SSL recording protocol, which provides two services for SSL connection:     

1. Confidentiality: using the secret key defined by handshake protocol    
2. Integrity: the handshake protocol defines Mac to ensure message integrity         

### Alert protocol
When the client and server find an error, they send an alert message to each other. If it is a fatal error, the algorithm will immediately close the SSL connection, and both parties will delete the relevant session number, secret and key first. Each alarm message has 2 bytes in total. The first byte indicates the error type. If it is an alarm, the value is 1. If it is a fatal error, the value is 2. The second byte specifies the actual error type.  
### Summary   
In SSL, we use the handshake protocol to negotiate encryption, MAC algorithm and secret key, use the handshake protocol to encrypt and sign the exchanged data, and use the alert protocol to define how to solve the problems in the data transmission process.

### Referrence
Kurose, James F, and Keith W. Ross. Computer Networking: A Top-Down Approach Featuring the Internet. Boston: Addison-Wesley, 2001. Print.



	

	
	
	
	 
	


