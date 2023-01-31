# User Login Service Scalable

- Build:
    ```sh
    docker compose up --build
    ```

- Test:
    ```sh
    ddosify -m POST -t http://localhost:8080 -b '{"username":"deneme","passwd": "deneme"}' -l incremental -n 500
    ```
