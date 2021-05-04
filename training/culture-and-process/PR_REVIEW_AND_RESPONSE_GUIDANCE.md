Pull Request Review and Response Guidance
=========================================

General guidance on reviewing a pull request and responding

### Prerequisites, Philosophy and Offering feedback

When you are making an observation or giving helpful feedback for a specific piece of code (or a general overall comment), if what you are writing is over ~50 words - it may be better to have a direct call or discuss via Slack and then maybe a comment can be placed in the code as to why something is done a particular way.

**Teamwork** - Reviewing other peoples work is an act of teamwork and teamwork can be challenging as it requires compromise.

**Reviews** - Are an opportunity for learning - so don't just say 'do this' - consider explaining why, especially when its a best practice. Try and explain the value (maybe with links to articles), but be careful that the links you reference are balanced.

**Kiss** - Keep It Simple and Straightforward. Does the body of work being reviewed lend itself to maintenance ? Where applicable, are there sufficient comments explaining why and how something is being done.

**Plan** A, B, C - There may be many possible solutions to a piece of work. Don't get too attached to any one idea or position as it may turn out to be flawed. The writer (and reviewer) of the code should be willing to let go and try another idea. What that means when reviewing a piece of code that works and passes all tests, but is implemented differently to how you would like it to be done then be accepting of the solution, unless there is a security or performance problem.

---
And last but not least, **`RED Flags - Test Code Removed`**:

If you see test code removed - thoroughly check into the reasons why - they may be valid but often indicate some problem.

## General things to consider

### Before diving into the Review

Sometimes a PR will pass all checks on a local machine and then fail in CI for some reason and the author of the PR puts it into slack before waiting to see that all tests in CI had passed.

If the PR on Github has not passed all tests then report this back to the author of the PR and stop until the author reports back to you that the problem has been fixed.

---
Make a request or pass a comment ?

Try not to confuse the two. If you make a change request against the PR, then approve the changes when they are done or dismiss your change request if you change your mind so as not to block others approving the overall PR.

---
Some code reviews may benefit from the Author of the code doing a screen share and demonstrating what the changes do.

This can save a lot of time for the reviewer in not having to setup some specialised environment.

This is most beneficial when the reviewer is not familiar with the area that the new code is operating in.

---
The points presented here are not exhaustive.

Think of it as a checklist and a reminder of things to consider where relevant.

### The Review

- Familiarise yourself with the Trello Ticket of the work you are reviewing.
- Be humble.
- Suggest ways to improve or simplify the code.
- Don't just look at whats been done, but also look out for anything that may have been missed.
---
- Does the code work - tests pass ?
- Are functions and variables named well ?
- Are there suffucient tests ?
- Are tests meaningful ?
- Is test code coverage adequate ?
- Can you understand the code ?
- Any sensitive data being revealed ?
- Does code meet house style ?
- Is the code clean - free of commented out code / not needed code ?
- Is there adequate commenting ?
- Is input data validated ?
- Might the code suffer speed issues (benefit from, for example: pagination) ?
- Is the code modular, where applicable ?
- Does code and test code address all sunny and rainy day paths (capture / process all errors) ?
- Are there simpler solutions ?
- Are logs produced in the code sensible ?
- Has a suitable linter been run and where possible, the issues raised by the linter cleaned up ?

## Language specific things to consider

### Go

#### checklist:

```shell
go clean -testcache

make test

make test-component
```

Does running:
```shell
go fmt ./...
```
or
```shell
go mod tidy
```
change any files ?

#### Implementation problems to look out for

Check that http.Response.Body is being closed properly.

Use `httptest.NewRequest()` in test code and not `http.NewRequest()`

### Java

#### checklist:

**>>> fill in details here <<<**

### TODO Any other languages ?

## Responding to a Review

Thank the reviewer.

Suggestions received in a review are just suggestions, they are not an instruction.

Its OK to push back if for example the effort outweighs the value, or the value has not been clearly explained.

Ask for or offer clarification, in a call if needs be.

Try to respond to each comment.
