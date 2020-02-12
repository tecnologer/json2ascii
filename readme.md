# json2ascii

Is a tool for convert json to ascii automatic

## Requirements

### With Go 1.12.x or later just init modules

```bash
go mod init
```

### With Go 1.11.x or less:

- Install [logrus][1]

```bash
go get github.com/sirupsen/logrus
```

## Basic Usage

```bash
# from file
go run main.go -jfile jsonfile.json

# raw json
go run main.go -json "{\"en\": \"hello world\", \"es\": \"hola mundo\", \"array\": [{\"a\": 1.0, \"b\": false}]}"
```

### Sample output

```txt
Root: (object)
|_en: (string)
|_es: (string)
|_array: (array)
    |_a: (float)
    |_b: (bool)
```

## TODO

- [ ] Remove duplicates

[1]: github.com/sirupsen/logrus
