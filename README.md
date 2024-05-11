# arsync
arsync (short for archive sync) is a simple program consisting of a server and a client.

The server is listening (by default on port 1337) and receives commands from the client.
When it receives a command it finds the folder, archives it, moves the archive to the output path (FTP directory),
then the client downloads the archive via FTP.

## Usage
### Server
```bash
./server -port 1337 -base-path /path/to/folder -output-path /path/to/ftp -username server-username -password server-password
```

### Client 
#### Prepare & Download
```bash
./client -address localhost:1337 -folder folder-name -username server-username -password server-password -ftp-username ftp-username -ftp-password ftp-password
```
#### List Directories
```bash
./client -address localhost:1337 -username server-username -password server-password -command list
```
