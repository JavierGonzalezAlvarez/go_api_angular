// Package Main API.
//
//	Schemes: http, https
//	Host: localhost:8080
//	BasePath: /
//	Version: 0.0.1
//	License: MIT http://opensource.org/licenses/MIT
//
//	Consumes:
//	- application/json
//	- application/xml
//
//	Produces:
//	- application/json
//
//	Security:
//	- api_key:
//
//	SecurityDefinitions:
//	api_key:
//	     type: apiKey
//	     name: KEY
//	     in: header
//	oauth2:
//	    type: oauth2
//	    authorizationUrl: /oauth2/auth
//	    tokenUrl: /oauth2/token
//	    in: header
//	    scopes:
//	      bar: foo
//	    flow: accessCode
//
//	Extensions:
//	x-meta-value: value
//	x-meta-array:
//	  - value1
//	  - value2
//	x-meta-array-obj:
//	  - name: obj
//	    value: field
//
// swagger:meta
package main
