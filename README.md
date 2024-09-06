# Plonkd

A simple Go server that responds with the request origin IP. This project was started because I was tired of dynamic DNS services not being reliable enough.

## Features

- Responds with the request origin IP.
- Lightweight and easy to deploy.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/acorsinl/plonkd.git
    cd plonkd
    ```

2. Build the project:
    ```sh
    go build -o plonkd
    ```

## Usage

1. Run the server:
    ```sh
    ./plonkd
    ```

2. Make a request to the server:
    ```sh
    curl http://localhost:8080
    ```

    The server will respond with your origin IP address.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the BSD 3-Clause License.
