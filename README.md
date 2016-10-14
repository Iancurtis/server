# webapp



## 3rd party frameworks
---

* ini

    [go-ini/ini](https://github.com/go-ini/ini)

* unit test

    [smartystreets/goconvey]()

* web framework

    [gin-gonic/gin](https://github.com/gin-gonic/gin)

## 3rd party tools
---

* database migration

    [mattes/migrate](https://github.com/mattes/migrate)
    
## Issues
---

* vscode on ubuntu can not stop delve. If we restart the debuger, we would get the error `listen tcp :8080: bind: address already in use`. we could use `killall -9 debug` to kill the debuger

## Conversion 
---

* put all the global variables into `package g`