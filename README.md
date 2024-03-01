# arsync
arsync (short for archive sync) is a simple program consisting of a server and a client.

The server is listening (by default on port 8080) and receives commands from the client.
When it receives a command it finds the folder, archives it, moves the archive to the output path (FTP directory),
then the client downloads the archive via FTP.
