# go-cuddle

`go-cuddle` is a lightweight, secure, and easy-to-use session management library for Go, inspired by the `cookie-session` middleware in Node.js. It provides session handling using cookies, with features like encryption, secure cookies, and easy integration with popular Go web frameworks like Gin.

## Features

- Simple API for managing session data
- Secure cookies with encryption support
- Configurable session options (HTTPOnly, Secure, SameSite, MaxAge, etc.)
- Middleware integration with Gin and other Go web frameworks

## Installation

To install `go-cuddle`, use `go get`:
```sh
go get github.com/benodiwal/go-cuddle
```

## Usage
### Example with Gin

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/benodiwal/go-cuddle"
    "net/http"
    "time"
)

var (
    validUsername = "user"
    validPassword = "password"
)

func main() {
    r := gin.Default()

    sessionManager := gocuddle.NewManager(
        gocuddle.WithName("session"),
        gocuddle.WithKeys([]string{"your-secret-key"}),
        gocuddle.WithSecure(true),
        gocuddle.WithHTTPOnly(true),
        gocuddle.WithSameSite(http.SameSiteNoneMode),
        gocuddle.WithMaxAge(24*time.Hour),
    )

    r.Use(func(c *gin.Context) {
        req := c.Request.WithContext(c)
        sessionManager.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            c.Writer = w
            c.Next()
        })).ServeHTTP(c.Writer, req)
    })

    r.POST("/login", func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")

        if username == validUsername && password == validPassword {
            session := gocuddle.GetSession(c.Request)
            session.Values["authenticated"] = true
            session.Values["username"] = username
            session.Changed = true

            c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
        }
    })

    r.POST("/logout", func(c *gin.Context) {
        session := gocuddle.GetSession(c.Request)
        session.Values["authenticated"] = false
        session.Values["username"] = ""
        session.Changed = true

        c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
    })

    r.GET("/dashboard", func(c *gin.Context) {
        session := gocuddle.GetSession(c.Request)
        if auth, ok := session.Values["authenticated"].(bool); ok && auth {
            c.JSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard!", "username": session.Values["username"]})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Please log in first"})
        }
    })

    r.Run(":8080")
}

```
### Raw Implementation
```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/benodiwal/go-cuddle"
)

var (
    validUsername = "user"
    validPassword = "password"
)

func main() {
    sessionManager := gocuddle.NewManager(
        gocuddle.WithName("session"),
        gocuddle.WithKeys([]string{"your-secret-key"}),
        gocuddle.WithSecure(true),
        gocuddle.WithHTTPOnly(true),
        gocuddle.WithSameSite(http.SameSiteNoneMode),
        gocuddle.WithMaxAge(24*time.Hour),
    )

    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        username := r.FormValue("username")
        password := r.FormValue("password")

        if username == validUsername && password == validPassword {
            session := sessionManager.GetSession(r)
            session.Values["authenticated"] = true
            session.Values["username"] = username
            session.Changed = true

            sessionManager.Save(w, r, session)

            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
        } else {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]string{"message": "Invalid credentials"})
        }
    })

    http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        session := sessionManager.GetSession(r)
        session.Values["authenticated"] = false
        session.Values["username"] = ""
        session.Changed = true

        sessionManager.Save(w, r, session)

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
    })

    http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
        session := sessionManager.GetSession(r)
        if auth, ok := session.Values["authenticated"].(bool); ok && auth {
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "message":  "Welcome to the dashboard!",
                "username": session.Values["username"],
            })
        } else {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]string{"message": "Please log in first"})
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

```

## Session Options
go-cuddle provides several options to configure session behavior:

- `WithName(name string)`: Sets the cookie name for the session.
- `WithKeys(keys []string)`: Sets the keys used for signing the cookie.
- `WithSecure(secure bool)`: Marks the cookie as Secure.
- `WithHTTPOnly(httpOnly bool)`: Marks the cookie as HTTPOnly.
- `WithSameSite(sameSite http.SameSite)`: Sets the SameSite attribute.
- `WithMaxAge(maxAge time.Duration)`: Sets the maximum age of the session.

## LICENSE

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
