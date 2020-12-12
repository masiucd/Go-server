# Go server

- [About](#about)
- [Concepts](#concepts)
- [Tools](#tools)

## About <a name = "about"></a>

Simple server using _Go_. As a framework I using [Gin](https://github.com/gin-gonic/gin) which a animalistic framework for creating web servers with golang. If your are used with using express in node then this will be very familiar to express.
Sqlite as a database because it's easy to setup and great for a demo version.

This is a example of a CRUD application using `GO` with,

- GET
- POST
- PUT
- DELETE

requests, where you can see how it works and how you can the take this concepts and keep build on top of that a more advanced version.

## Concepts <a name = "concepts"></a>

_To get the params_
if you been working with express js in node the to get the params you would write something like this.

```js
const getUserById = async (req, res, next) => {
  const userId = req.params.id

  const user = models.User.findById(userId)
  if (!user) {
    throw new Error("User does not exist!")
  }

  res.status(200).json({ data: user })
}
```

in Go using Gin to get the params you would write something like this.
**/user/:id**

```go
func GetUserById(c *gin.Context){
  var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
```

`c.Param("id")` is where we get the params

## Tools <a name = "tools"></a>

- Golang 🐰
- [Gin 🍸](https://github.com/gin-gonic/gin)
- [Gorm ](https://gorm.io)
- SQLLite
- Rest
