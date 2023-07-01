
## Run the Project

Start the Proxy Server

```bash
  go run .\main.go
```

Start the File Handler Server

```bash
  go run .\file_handler.go
```

### API

curl --location 'http://localhost:3000/api/v1/hello'

curl --location 'http://localhost:3000/api/v1/file' \
--form 'file=@"pexels-monstera-5634602.jpg"'


  