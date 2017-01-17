# Go Auth Kit
[![Build Status](https://travis-ci.org/mrLSD/auth-kit.svg?branch=master)](https://travis-ci.org/mrLSD/auth-kit) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mrLSD/auth-kit/master/LICENSE) [![GoDoc](https://godoc.org/github.com/dghubble/gologin?status.png)](https://godoc.org/github.com/mrlsd/auth-kit) [![Coverage Status](https://coveralls.io/repos/github/mrLSD/auth-kit/badge.svg?branch=master)](https://coveralls.io/github/mrLSD/auth-kit?branch=master)

Golang authentication kit, provides cookie 
and infrastructure for session backends.

Main features:
* Built-in backends to store sessions: cookie,
filesystem, Redis, MySQL, PostgreSQL, custom 
backends extensibility.
* Signed cookies: use it as an easy way to set signed 
(and optionally encrypted) cookies
* Flash messages: session values that last until read.
* Interfaces and infrastructure for custom session 
backends: sessions from different stores can be retrieved and batch-saved using a common API.

#### How to install:
`go get github.com/mrlsd/auth-kit`

#### Requirements:
* Go 1.6+
  
#### Some useful command:
* **test:** `make test`
* **fmt:** `make fmt`

####License: MIT [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mrLSD/auth-kit/master/LICENSE)
