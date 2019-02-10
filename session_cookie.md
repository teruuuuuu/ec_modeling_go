```
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}

```

```
$ curl -XGET --dump-header - http://localhost:8000/incr
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Set-Cookie: mysession=MTU0NzcxMTY3OXxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBQT1893UE80aRn-edXrsbGdaarNYSNPe1wtYkd7R6FXMjCN0=; Path=/; Expires=Sat, 16 Feb 2019
07:45:04 GMT; Max-Age=2592000
Date: Thu, 17 Jan 2019 07:45:04 GMT
Content-Length: 11

{"count":0}


$ curl -XGET -b 'mysession=MTU0NzcxMTY3OXxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBBQT1893UE80aRn-edXrsbGdaarNYSNPe1wtYkd7R6FXMjCN0='  --dump-header - http://localhost:8000/incr
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Set-Cookie: mysession=MTU0NzcxMTcxMHxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBST18EHafHNLvF5uYGXM5mvChcigbv30Q-ZsGQyXJUOxSl7U=; Path=/; Expires=Sat, 16 Feb 2019
07:55:10 GMT; Max-Age=2592000
Date: Thu, 17 Jan 2019 07:55:10 GMT
Content-Length: 11

{"count":1}

$ curl -XGET -b 'mysession=MTU0NzcxMTcxMHxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBBST18EHafHNLvF5uYGXM5mvChcigbv30Q-ZsGQyXJUOxSl7U='  --dump-header - http://localhost:8000/incr
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Set-Cookie: mysession=MTU0NzcxMTgyMHxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBUT18jH7ChkYZxLj_N_eEEeyx42TgbJoj2G4PxBE8OLTVjUk=; Path=/; Expires=Sat, 16 Feb 2019
07:57:00 GMT; Max-Age=2592000
Date: Thu, 17 Jan 2019 07:57:00 GMT
Content-Length: 11

{"count":2}

$ curl -XGET -b 'mysession=MTU0NzcxMTY3OXxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBaRn-edXrsbGdaarNYSNPe1wtYkd7R6FXMjCN0='  --dump-header - http://localhost:8000/incr
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Set-Cookie: mysession=MTU0NzcxMTgzOXxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBQT18Jpsk7N8KXDUoHPuF599Px-58X86F62bZeqa9rqK7SzY=; Path=/; Expires=Sat, 16 Feb 2019
07:57:19 GMT; Max-Age=2592000
Date: Thu, 17 Jan 2019 07:57:19 GMT
Content-Length: 11

{"count":0}

$ curl -XGET -b 'mysession=MTU0NzcxMTgzOXxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBQT18Jpsk7N8KXDUoHPuF599Px-58X86F62bZeqa9rqK7SzY='  --dump-header - http://localhost:8000/incr
BQT18Jpsk7N8KXDUoHPuF599Px-58X86F62bZeqa9rqK7SzY='  --dump-header - http://localhost:8000/incr
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Set-Cookie: mysession=MTU0NzcxMTkyNXxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBST18RuaiJFS_sBNRyq6JF2bBtU1V_Ph9QOBTckY6r1Js7u4=; Path=/; Expires=Sat, 16 Feb 2019
07:58:45 GMT; Max-Age=2592000
Date: Thu, 17 Jan 2019 07:58:45 GMT
Content-Length: 11

{"count":1}
```