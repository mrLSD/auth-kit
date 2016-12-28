// Copyright 2016 by Evgeny Ukhanov (mrLSD). All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
Package auth provides cookie and infrastructure
for session backends.

Main features:

	* Built-in backends to store sessions: cookie,
	  filesystem, Redis, MySQL, PostgreSQL, custom
	  backends extensibility.
	* Signed cookies: use it as an easy way to set signed
	  (and optionally encrypted) cookies
	* Flash messages: session values that last until read.
	* Interfaces and infrastructure for custom session
	  backends: sessions from different stores can be retrieved and batch-saved using a common API.

*/
package authkit
