# chord-go
Implementation of Chord Distributed Hash System in Golang for 50.041 Distributed Systems (SUTD)

### Prerequisites
1. Go version 1.19
2. [Docker Desktop](https://www.docker.com/products/docker-desktop/)
3. Evans gRPC client (for testing, instructions to download are further down this README.md
4. VSCode for proto3 extension

### Useful Links
1. [gRPC documentation](https://grpc.io/docs/what-is-grpc/introduction/)
2. [Specific Go implementation](https://grpc.io/docs/languages/go/basics/)

### Setup
1. Install the vscode-proto3 extension by zxh404
2. Go to Preferences --> User Settings.json. Copy and paste the following json at the end of the file (Within the first set of curly brackets) and save it
   ```
     "protoc": {
      
        "options": [
          "--proto_path=proto"
      ]
    }
   ```
   Steps 2 and 3 are to remove the red errors within the `protos` folder on the imports
3. Run Docker Desktop
4. Be in the root directory of the project (Same level as Makefile)
5. After successful Docker Desktop startup, build the docker images. Run the following command as defined in ``Makefile``:
   <br><br>
    ```
    make image
    ```
   If the build was successful, the top of the build logs will look similar to:
    ```
    docker image rm chord:latest       
    docker build -t chord:latest .
    Untagged: chord:latest
    Deleted: sha256:c4e70570c66f2fdda6c18aaa28b3e0b083a3b905789df3dbe251146111f1154b
    [+] Building 15.9s (14/14) FINISHED
    ```
6. Start up the containers with the following command:
   <br><br>
   ```
   docker compose up
   ```
   If the command was successful, the logs will look similar to:
   ```
   chord-go-chord1-1  | 2023/10/29 07:18:05 Start gRPC server on 192.168.128.2:9090
   chord-go-chord3-1  | 2023/10/29 07:18:05 Start gRPC server on 192.168.128.3:9092
   chord-go-chord4-1  | 2023/10/29 07:18:06 Start gRPC server on 192.168.128.4:9093
   chord-go-chord2-1  | 2023/10/29 07:18:06 Start gRPC server on 192.168.128.5:9091
   ```

### Testing the setup
For testing purposes to interact with the gRPC chord network, we can use Evans gRPC client as an interface without 
having to build our own client. I have only used Evans on Mac so I am not sure about any other configs needed for other
OS

To install Evans, follow the instructions in their [repo](https://github.com/ktr0731/evans)

#### Using Evans
Each docker container has a port exposed to the local machine as defined in `docker-compose.yml` Without changing 
anything in docker-compose.yml, we can use Evans to connect to any of the 4 containers, and invoke the RPC call 
`RequestFromClient`, which will cause that node to invoke another RPC call `GetChordNode` on another node running 
on another container. This is to demonstrate that all the containers are 1. Running and 2. Able to communicate with
each other within the docker network

1. Be in the root directory 
2. To start the Evans client after installation, run the following command:
   ```
   evans --host localhost --port <port of node you want to use> -r repl
   e.g.
   evans --host localhost --port 9090 -r repl
   ```
3. Within the terminal, the user should look something like
   ```
   pb.Chord@localhost:<port of node>
   ```
   If the user only shows `localhost:<port>`, run the following commands:
   ```
   package pb
   (The user should change to pb@localhost:<port>)
   service Chord
   (The user should now change to pb.Chord@localhost:<port>)
   ```
4. Check the RPCs are available using
   ```
   show service
   ```
   And `GetChordNode` and `RequestFromClient` should be listed
5. To call either RPCs, use the following command:
   ```
   call <service_name>
   ```
6. When prompted, enter the request values into the terminal. The request values do not matter
   for this case, its just to test that requests can be processed
7. For `GetChordNode`, the response will look similar to
   ```
   {
     "id": "4",
     "ip": "test",
     "port": 9090
   }
   ```
   For `RequestFromClient`, the response will look similar to:
   ```
   {
     "id": "1",
     "inputFromClient": "1",
     "ip": "test",
     "port": 9090
   }
   ```
### Chord Network Simulation
The docker containers are connected as follows:
```
chord1 --> chord2 --> chord3 --> chord4 --> chord1
```
Simulating a ring connection and are only contactable via the `RequestFromClient` RPC by default as defined in `gapi/rpc_request_from_client.go`.

### How to add more functionality
To add/edit/delete more functions to the gRPC server:

1. Add/Edit/Delete whatever messages and services in the `/proto` folder. Add as many proto files for the messages, but keep the services to one file `service_chord.proto`
2. Run the following command as defined in the Makefile to compule the proto files
   ```
   make proto
   ```
3. if new services were added, go into `service_chord_grpc.pb.go` and find code similar to the example below and copy it:
   ```
   func (UnimplementedChordServer) <New service name>(context.Context, *<Request Message>)             (*<Response Messsage>, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFromClient not implemented")
   }
   ```
4. Create a new file for each newly created service under `/gapi`, and copy paste the copied code and edit accordingly. Use the existing RPC implementations as a guide
5. (If running on docker) run `make image` to build the image again to implement changes
   
### To add more containers
1. Enter the `docker-compose.yml` file
2. Copy the following
   ```
    chord<id>: 
    image: chord
    ports:
      - "9091:9091"
    environment:
      - CHORD_ID=<id>
      - SUCCESSOR_ADDRESS=chord-go-chord<id+1>-1:<port of successor>
      - SERVER_ADDRESS=chord-go-chord<id>-1:<own port>
    command:  [ "/app/main"]
   ```
3. Paste it at the end of the file

Edit the docker compose file as needed to accommodate changes to the codebase
   
