// Copyright 2015-2024 Bleemeo
//
// bleemeo.com an infrastructure monitoring solution in the Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package bleemeo provides simple ways to interact with the Bleemeo API (https://go.bleemeo.com/l/doc-api).

# Client

A Client is used to send requests to the Bleemeo API,
and post-process possible errors to make them clearer.

The initialization of a Client is done with the NewClient function,
which can take a variable number of ClientOption parameters.
The following options can be used to customize the Client:

WithCredentials, WithBleemeoAccountHeader, WithOAuthClient, WithEndpoint,
WithInitialOAuthRefreshToken, WithHTTPClient, WithNewOAuthTokenCallback and WithThrottleMaxAutoRetryDelay.

The Client allows different kinds of resource interactions:

- Client.Get() retrieves the resource with the given ID

- Client.GetPage() retrieves a list of resources within the given bounds

- Client.Count() retrieves the total number of resources on the API

- Client.Iterator() returns an Iterator for the specified resource type (described later)

- Client.Create() creates the given resource on the API

- Client.Update() updates the resource with the given ID with the specified values

- Client.Delete() removes the resource with the given ID from the API

- Client.Do() executes a request defined by the given parameters

- Client.DoRequest() executes the given http.Request (without error post-processing)

- Client.ParseRequest() builds an http.Request according to the given parameters

- Client.GetToken() returns the current OAuth token, fetching a new one if necessary

- Client.Logout() requests the revocation of the current OAuth token

An Iterator can be used to iterate over all the resources of a given kind that match some parameters.
The Iterator.Next() method moves the iteration cursor to the next resource,
and returns whether the Iterator is exhausted or not.
Calling Iterator.At() returns the resource at the current cursor position,
but must only be done if Iterator.Next() has been called just before and returned true.
Iterator.Err() returns the error that occurred during the iteration, if any.

	iter := client.Iterator(...)
	for iter.Next() {
	   value := iter.At()
	   // process value
	}

	if iter.Err() != nil {
	   // process error
	}

# Resources

A Resource represents a datatype on the Bleemeo API, and can be used as a route to access it.
The most common types are: ResourceMetric, ResourceService, ...

# Enums

Enumeration types are defined to put a name on an arbitrary value.
For instance, writing

	graphType := bleemeo.Graph_Text

instead of

	graphType := 8

avoids adding a comment explaining what '8' means, and also makes typing clearer.

# Errors

Error types wrap an error that occurred during the execution of a request.

An APIError may be returned after receiving the response of the API,
which is considered unsuccessful due to its status code.

Special cases of an APIError are AuthError and ThrottleError.

If a ThrottleError occurs when executing a request using any client method except Client.DoRequest(),
and if the delay to wait is less than the one specified with WithThrottleMaxAutoRetryDelay (which defaults to 1min),
the request will be retried without returning an error.

A JSONMarshalError may occur when trying to serialize some request content to JSON.
A JSONUnmarshalError may occur when deserializing the response content from JSON.

JSON errors both have a `DataKind` field of the type JSONErrorDataKind,
which indicates the type of data that failed its conversion, and can be used to estimate where the error happened.

# Utility functions

[JSONReaderFrom] can be used to serialize some data to JSON and get a reader of the result.
*/
package bleemeo
