About project:
===============

This project focus in learn grpc and better my knowledges golang.


About GRPC:
============

- It's one framework the rpc calls
- Implemented over http 2 where data is transfer format binary
- Using protobuffer to make contract to communication between client and server grpc
- Allow many ways communication between client and server:
    - Client send data to server and return.
    - Client send data to server and server return piece data for client. OBS: bidirectional in server part.
    - Client send piece data to server. OBS: bidirectional in client part.
    - Client and server piece data one to another. OBS: bidirectional communication both client and server