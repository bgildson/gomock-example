gomock-example
===

[![Test Status](https://github.com/bgildson/gomock-example/workflows/Test%20and%20Send%20Coverage%20Report/badge.svg)](https://github.com/bgildson/gomock-example/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/github/bgildson/gomock-example/badge.svg?branch=master)](https://coveralls.io/github/bgildson/gomock-example?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/bgildson/gomock-example)](https://goreportcard.com/report/github.com/bgildson/gomock-example)

This repository contains some examples showing a progressive way to implement mocks and use [gomock framework](https://github.com/golang/mock) to apply and autogen mocks.

## finalspace0

Is the [first example](client/finalspace0) and is the simplest implementation, it is only a code separation for request and parse response.

## finalspace1

Is the [second example](client/finalspace1) and, in comparation with the previous example, adds the dependency injection to specify the http client and api endpoint through parameters ensuring a more flexible use.

## finalspace2

Is the [third example](client/finalspace2) and it is equals to the previous implementation, in this example was incresed the tests and created a transport http mock to be possible uses the original http client and mock the response.

## finalspace3

Is the [fourth example](client/finalspace3) and was added a http client interface to replace the default http client struct and use a different function, the `Do`, to make the request, turning possible to use mockgen to generate a http client mock used to test the finalspace client.
