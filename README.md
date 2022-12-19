# Getting Started with CLIENT-SERVER/CLIENT-SEVER(JSON-RPC) App

In this project i'm implement simple client-server(rest)-server(json-rpc) interaction (in both way)

## To send request on client use POST method, path:
### `http://localhost:5000/{username}`
where {username} - your name without quotes

##  Server(rest) will start on 
[http://localhost:5000](http://localhost:5000)

##  Server(json-rpc) will start on
[http://localhost:8080/concat](http://localhost:8080/concat)

On the first step Client send {username} to server1.

Server1 add smile to username: 
{username} = {username} + " :) " and send request to server2 with params and method what to do.

Server 2 get first params {username} and second params - word to add.
Make concat and add "!" to the phrase and send result it back to the first server.

Server 1 get response and add " Have a nice day!" and send response to the client.

(You can see all interaction in log)

Final message on the client will be
### ` {"message":"Hello Olena :) ! Have a nice day!"}`
