## Installing dependencies:
Here are the dependencies I need:

```Shell
go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/postgres
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1
```

## Things that haven't been done to improve our API:
1. No password validation for better security; e.g, the User can add a password such as `"I am here"` or a one letter password like `"a"`.
2. File upload to Firebase Storage can be done, but needs golang version to be `<=1.17`. Currently, I am having a problem with my WSL2 environement, I'll get to it as I have more time. 
3. User table and controllers; representing Customer is still not implemented.
4. Purchases table and controller; representing the products that the User purchased is still not implemented.
5. Shops can retrieve only their products too.

## Tests that aren't done:
1. Nothing has been tested until now, the python scripts are able to test the API but are neither consistent nor scalable for more complexity.