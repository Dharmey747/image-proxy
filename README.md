# Image Proxy

Bypass CDNs' request blocks on image endpoints by using this program to proxy the endpoint URL to a server that is able to successfully execute the request and return the valid response back. This will be particularly useful if you're working with platforms like Discord that are usually not able to display images via endpoint URLs because of user agent blocks on the endpoints themselves.

## Usage
- [Install Go](https://go.dev) and setup the environment.
- Clone this repository.
- Install the [Mux](https://github.com/gorilla/mux) module to get a server up and running via `go get`.
- Start the server by running the `go run .` command.
- If you'd wish to create a compiled binary of this program, you can do so by running the `go build .` command.
