Unit Testing in Go
==================

## More Visual Version of this training

[Click here to view the presentation format using Miro](https://miro.com/).

## Unit Testing

Unit testing is the process of testing a portion of an application or service to ensure it meets our requirements.  The portion is ideally a single function, but can be any portion of your Go code.

Ideally we should only test the portion of the application we are focused on, and shouldn't connect with other parts of the system.  This can be difficult, and so sometimes different portions of the system are 'mocked', which is also covered in this article.

## Is Unit Testing EVER Pointless?
Very occassionally Unit Testing can be pointless, as the mockup's required are so complex to make and maintain, that testing these units another way is more suitable, such as through [integration](https://github.com) tests, or [end-to-end tests](https://github.com), or simply through manual code reviews.

Also, when creating throw away code, such as for a demo, then Unit Testing may be pointless, as long ask you're sure the code will not see Production.

## Pre reading

It's advised that you are comfortable with GoLang or an other programming or scripting language.

## Pre requisites

You'll need a system with GoLang installed, and the ability to download further repo's i.e. the internet!

To get started quickly if you've already got Go installed, and you use a bash/linux based terminal; Simply copy and paste the following 4 lines to your terminal and press enter, and it will install the appropriate repo's for this training:-
- `go get github.com/smartystreets/goconvey` # Behaviour Driven Development - human readable tests.
- `go install github.com/smartystreets/goconvey/web/server` # Nice web based user interface for your tests.
- `go get github.com/stretchr/testify` # Contains Mockify, which is used below.  We wont be using Testify, which is an alternative to GoConvey
- `git@github.com:ONSdigital/dp.git` # Contains this repo, and other training material

## Materials

#### Quick Start
- ['Hello World' Unit test in Go and GoConvey](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)
- ['Choice tests to show GoConvey Functionality'](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)if
- [More complex test using Depedency Injection](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)
- [Further GoConvey Information](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)
    - [Converting natural language to GoConvey test stubs](moreGoConvey.md#testStubs)
    - [Setup and Teardown]
    - [Order of executions](moreGoConvey.md#executionOrder)
    - [Assertion examples (search for asdf)](https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior/)
    - [Other uses of Assertions ()](https://smartystreets.com/blog/2015/08/go-testing-part-4-standalone-assertions/)
    - [Dot import for Convey](https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior/)
    - [Other things to look out for](moreGoConvey.md#gotchas)

#### Mocking
- [Mocking dependencies / portions of system you don't need to test](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)

#### Code Coverage
- [Code Coverage](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)
- [Tips On Increasing Code Coverage](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)

#### Alternatives to GoConvey
- [Alternatives to GoConvey](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)

#### Writing Easier Code To Test
- [Writing tests as you code (TDD)](https://github.com/ONSdigital/dp/tree/master/training/unitTesting/goLang/helloWorld)

## Quick Reference
The following are useful links:-
- [Goconvey documentation](https://godoc.org/github.com/smartystreets/goconvey/convey)
- [Assertion examples](https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior)

Have a look at [Integration Testing](https://github.com/ONSdigital/dp/tree/master/training)


Further resources
----------------------------

