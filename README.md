# testAssignment

### 1. Setalah membuka file golang, terlebih dahulu menuliskan perintah dibawah ini, dmana perintah tersebut berguna untuk mendownload semua tools dan library yang digunakan.
       perintah "go mod tidy"
### 2. Sebelum menjalankan golang, terlebih dahulu kita setup connection to database.
       Berada di dalam folder "config.json"
       
## Field Database
### Account
| Field | Type |
| -- | -- |
| account_number | string |
| customer_number | string |
| balance | int |

### Customer
| Field | Type |
| -- | -- 
| customer_number | string |
| name | string |


### 3. EndPoint API
| Method | Url |
| -- | -- |
| POST | /account |
| GET | /account/{account_number} |


Request
```json
{
	"name" : "Giandra",
	"balance" : 10000
}
```
