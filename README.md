## Create DB

```SQL
CREATE DATABASE go;
CREATE USER go WITH ENCRYPTED PASSWORD 'go';
GRANT ALL PRIVILEGES ON DATABASE go TO go;
```

## Run the following commands:

```Shell 
sudo go mod init github.com/Arka-cell/goassignment
sudo go run main.go
```