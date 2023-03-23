# backend-everyshillings-proto

Protobuf (short for Protocol Buffers) is a language-agnostic data serialization format created by Google. 
Protofiles are used to define the structure of the data that will be serialized or deserialized using the Protobuf format. 



## Usage
### Generate code
```shell
make
```

### Import go code
Download library

```shell
go get github.com/AppsLab-KE/backend-everyshillings-proto
```
Import example
```go
import(
	"github.com/AppsLab-KE/backend-everyshillings-proto/auth"
	"github.com/AppsLab-KE/backend-everyshillings-proto/exchange"
)
```

### Import python code
This command will install the library from develop branch
```shell
pip install everyshillingsproto@git+https://github.com/AppsLab-KE/backend-everyshillings-proto@develop
```
