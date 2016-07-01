# requestLogger
A simple request logger written in Go.

## Usage
This project includes a build script written in Ruby. Make use of it.

### Building:

`./build` or `./build --target={TARGET}`

Target being one of the targets supported by the go compiler in the following format:

`GOOS-GOARCH`

Eg. 
`linux-amd64`

### Running

`./bin/requestLogger` (it will start listening on port 3000, if available)

By default it only logs the request Path, Body and Method. To also log the headers, start it with `./bin/requestLogger --headers`
