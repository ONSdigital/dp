Unit Testing in Go
==================

### Pre reading & Pre requisites
It's advised that you are comfortable with GoLang or another programming or scripting language.  The setup is focussed on linux based operating systems..

## Contents
- [Less wordy material](#Less wordy material)
  - [Quick Start](#Quick Start)
  - [System Setup](#System Setup)
  - [Quick Reference](#Quick Reference)
  - [Potential Pitfalls](#Potential Pitfalls)
    - [Order of Execution](#Order of Execution)
    - [Assertions](#Assertions)
  - [Creating Easy To Test Code](#Assertions)
    - [Writing SOLID Go](#Writing SOLID Go)
    - [Dependency Injection](#Dependency Injection)
    - [Test Driven Development](#Test Driven Development)
  - [Recommended Go Tutorials](#Recommended Go Tutorials)
  - [Code Coverage](#Code Coverage)
  - [Assessing How Fast Your Code Runs](#Assessing How Fast Your Code Runs)
  - [Other Interesting Testing Related Information](#Other Interesting Testing Related Information)
- [Wordier Information](#Wordier Information)
  - [Background to Unit Testing](#Background to Unit Testing)
  - [Is Unit Testing Ever Pointless?](#Is Unit Testing Ever Pointless?)
- [A visual guide to Unit Testing](#A visual guide to Unit Testing)

##Less Wordy Material

### Quick Start - code!
If your system is ready, and you're ready to code: [Downloadable 'Hello World' Unit Tests Using GoConvey](helloworld/hello_world_test.go)

Before you 'go' test, take a look at some [Potential Pitfalls](#Potential Pitfalls) below. 

### System Setup
If you've got Go and a GitHub terminal client installed, and you have a unix based operating system, including Mac, the following will install the appropriate repo's for this training:-

In your terminal:
1) `cd` to directory you wish to clone the tutorials to.
2) Copy & paste the commands below and press enter.
```
# Installs GoConvey
 go get github.com/smartystreets/goconvey
# Contains Mockify, which is used to mock elements.  Note we dont use Testify - just Mockify 
 go get github.com/stretchr/testify
# Contains this repo, and other training material
 git clone git@github.com:ONSdigital/dp.git 
# Installs code coverage tooling
 go get golang.org/x/tools/cmd/cover
```

### Quick Reference
The following are useful links:-
- [GoConvey on GitHub](https://github.com/smartystreets/goconvey)
- [GoConvey documentation](https://godoc.org/github.com/smartystreets/goconvey/convey)

### Potential Pitfalls

#### Order of Execution
GoConvey's execution path is a little quirky so have a look at the example below and the resources.
- [Downloadable order of execution example](#orderexecution/execution_order_example_test.go)
- [GoConvey explanation](https://github.com/smartystreets/goconvey/wiki/Execution-order)

#### Assertions
- When comparing objects in your tests you can use [GoConvey's Assertions](/assertions/README.md). 

### Creating Easy To Test Code
#### Writing SOLID Go
Code that is easy to test is often code which follows the [SOLID](https://github.com/iamharvey/SOLID_principles) principles.  Here is good video on writing [SOLID Go](https://www.youtube.com/watch?v=zzAdEt3xZ1M).

#### Dependency Injection
This allows objects to be more independent of their main dependencies, for example we could create a website templating module which could implement all the functionality required to template the website, minus the rendering.  This rendering functionality could be injected, for example for devices such as mobile, tablet, laptop.  The rendering on the different browsers could then be injected into the rendering routines for the different devices... and so the design unfolds.

The following [walk-through in GoLang is useful](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection)

Another occasion when Dependency Injection is useful is when making mockups.  The following downloadable example walks through this technique:
- [Dependency Injection within Mockups](https://github.com/sohamkamani/go-dependency-injection-example)

#### Test Driven Development
A useful approach, particularly for potentially difficult to test code, is Test Driven Development, which involves writing the test at the same time or before the code you need to test. So effectively designing easy to test code from the start.
- [Golang Test Driven Development walk through](https://itnext.io/how-to-tdd-a-console-application-to-achieve-100-coverage-starting-from-main-test-go-934a617b080f) 

### Recommended Go Tutorials
- [GoConvey Tutorial 1](https://www.smartystreets.com/blog/2015/02/go-testing-part-1-vanillla/)
- [GoConvey Tutorial 2](https://www.smartystreets.com/blog/2015/02/go-testing-part-2-running-tests/)
- [GoConvey Tutorial 3](https://www.smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior/)

### Code Coverage
When we test certain paths of execution are tested, but often not all.  Knowing which paths are tested is useful, but often the same paths need to be tested multiple time, so 100% coverage doesn't mean 100% tested.
- [Excellent Golang code coverage walk-through](https://blog.golang.org/cover)
- [Useful time saving suggestion on stackoverflow](https://stackoverflow.com/a/27284510)

### Assessing How Fast Your Code Runs
Systems often need to run efficiently, such as a response from a website.  Go incorporates a benchmarking routines, which will time how long the code will take to run.  It clever though as it repeats the tests until it feel it's got an accurate time.  
- [Downloadable benchmarking example](benchmarking/benchmark_test.go)

###Splitting Up Tests | Setup and Tear Down Tests 
Its sometimes useful to use a standard way of splitting up tests, or creating a setup and tear down routines for tests.  Golang provides this functionality built in.
- [Downloadable example without using testing.M](helloworld/split_test.go)
- [Downloadable example using testing.M](helloworld/split_with_test_main_test.go)
- [testing.M walk-through](http://cs-guy.com/blog/2015/01/test-main/)

###Running Time Critical Tests
If your code runs in a higher priority thread at the operating system level, you can also run your tests at this level also: 
- [Downloadable time critical test example](#time-critical/testing_with_main_OS_thread_test.go)

Table Driven Tests
------------------
Go Convey is very useful for creating very readable tests, however occasionally you may need to test lot of different permutations of a similar test, which could result in very long files.  This is where Table Driven Tests are very useful.  Effectively you create a template test, and feed in lots of different test permutations.  Here is a [Table Driven Test example](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

### Other Interesting Testing Related Information
- [Dot import for Convey](https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior/)

##Wordier Information
###Background To Unit Testing
Unit testing is the process of testing a portion of an application or service to ensure it meets our requirements.  The portion of code is ideally a single function, but is typically less than a few go files.

Ideally we should only test the portion of the application we are focused on, and shouldn't connect with other areas of a system.  This can be difficult, and so sometimes different portions of the system are 'mocked'.

### Is Unit Testing Ever Pointless?
Very occasionally Unit Testing can be pointless, as the mockup's required are so complex to write and maintain, that testing these portions another way is more suitable, such as through integration tests, or end-to-end tests, or simply through manual code reviews.

Also, when creating throw away code, such as for a demo, then Unit Testing may be pointless, as long ask you're sure the code will not see Production.

## Visual Guide To Testing
This is just a placeholder for a more visual presentation of Unit Testing
