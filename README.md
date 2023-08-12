# stampede
Sample web socket client and server

# Build stampede
Run `make build` command at the root of the `stampede` repository

This shall create a directory called `bin` in the root of `stampede` repository

# Start server
Run `./stampede run server --addr :8080`

This shall start the web socket server at address `:8080`

# Start station
Run `./stampede run station --addr localhost:8080`

This shall start the station which points to the server at address `localhost:8080`
