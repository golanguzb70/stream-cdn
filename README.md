# CDN Server
This repository contains the code for a Content Delivery Network (CDN) server. The CDN server is designed to efficiently deliver content to users by caching, securring and serving static assets.

## Features
- **High Performance**: Optimized for fast content delivery.
- **Dynamic Security**: Provides secure content delivery with HTTPS support and authorization using headers, query params and origins.
- **Caching**: Reduces load on origin servers by caching static assets.
- **Scalability**: Easily scalable to handle large amounts of traffic.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/golanguzb70/stream-cdn.git
    cd stream-cdn
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Build the server:
    ```sh
    go build -o cdn-server
    ```

## Usage

1. Start the CDN server:
    ```sh
    ./cdn-server
    ```

2. Configure your origin server and caching rules in the `.env` file or in environment.

## Configuration

The server can be configured using the `.env` file. Below is an example configuration:
```.env
ORIGIN_SERVER_URL="http://localhost:8080"
PORT=8081
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.