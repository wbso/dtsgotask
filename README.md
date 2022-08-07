# Mini Project Golang DTS

## How to use
copy `.env.example` to `.env`

### Development Mode
1. Run Backend Server
    ```sh
    make run
    ```

2. Run Frontend Server
    ```sh
    make run-frontend
    ```

### Build Application

    ```sh
    make build
    ```

## App Structure
```
presentation --> (business/app) --> (repo/persistence)
```