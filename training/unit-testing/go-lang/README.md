Unit Testing in Go
==================

## INCOMPLETE DOCUMENTATION

Portions of this page with (*) are incomplete, and may just link to github.com

## More Visual Version of this training (*)

(Link To Miro/Presentation)

## Unit Testing

Unit testing is the process of testing a portion of an application or service to ensure it meets our requirements.  The portion is ideally a single function, but can be any portion of your Go code.

Ideally we should only test the portion of the application we are focused on, and shouldn't connect with other areas of a system.  This can be difficult, and so sometimes different portions of the system are 'mocked', which is also covered in this article.

## Is Unit Testing EVER Pointless?
Very occassionally Unit Testing can be pointless, as the mockup's required are so complex to write and maintain, that testing these portions another way is more suitable, such as through [*integration](https://github.com) tests, or [*end-to-end tests](https://github.com), or simply through manual code reviews.

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

## Materials (*)

#### Quick Start
- ['Hello World' Unit tests in Go and GoConvey](1.helloWorld/README.md)
- [Choice tests to show GoConvey Functionality](2.choiceTests/README.md)
- [More complex test using Depedency Injection](3.dependencyInjection/README.md)
- [Order of execution](4.orderOfExecution/README.md)
- [More complex test using Depedency Injection](4.orderOfExecution/README.md)
- *Assertion Examples
- [Benchmarking Examples](benchmarking/README.md)
- Further GoConvey Information (external links)
    - [Converting natural language to GoConvey test stubs](moreGoConvey.md#testStubs)
    - [Setup and Teardown]
    - [Order of executions](moreGoConvey.md#executionOrder)
    - [Assertion examples (search for asdf)](https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior/)
    - [Other uses of Assertions ()](https://smartystreets.com/blog/2015/08/go-testing-part-4-standalone-assertions/)
    - [Dot import for Convey](https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior/)
    - [Other things to look out for](moreGoConvey.md#gotchas)

#### Mocking
- [Mocking dependencies / portions of system you don't need to test](https://github.com/ONSdigital/dp/tree/master/training/unit-testing/go-lang/helloWorld)

#### Code Coverage (*)
- Code Coverage
- Tips On Increasing Code Coverage

#### *Alternatives to GoConvey
- Alternatives to GoConvey

#### *Writing Easier Code To Test
- Test Driven Development

## Quick Reference (External Links) (*)
The following are useful links:-
- [Goconvey documentation](https://godoc.org/github.com/smartystreets/goconvey/convey)
- [Assertion examples](https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior)

*Further resources
----------------------------

