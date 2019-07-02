# Contributing to Mail2Most

- [Code of Conduct](#code-fo-conduct)
- [Issues](#issues)
- [Pull Requests](#pull-requests)
- [Reviewing](#reviewing)

## [Code of Conduct](https://github.com/cseeger-epages/confluence-go-api/blob/master/CODE_OF_CONDUCT.md)

The Mail2Most project follows the [Contributor Covenant Code of Conduct](https://github.com/cseeger-epages/confluence-go-api/blob/master/CODE_OF_CONDUCT.md)

## Issues

### How to Contribute in Issues

You can contribute:

- By opening an issue for discussion: For example, if you found a bug, creating a bug report using the [template](https://github.com/cseeger-epages/confluence-go-api/blob/master/.github/ISSUE_TEMPLATE/bug_report.md) is the way to report it.
- By providing supporting details to an existing issue: For example, additional testcases.
- By helping to resolve an issue: For example by opening a [Pull Request](https://github.com/cseeger-epages/confluence-go-api/pulls)

### Asking for Help

Just open a [regular issue](https://github.com/cseeger-epages/confluence-go-api/issues/new) and describe your problem.

## Pull Requests

### Dependencies

This project uses [Mage](https://magefile.org/) which is a replacement for the classical make.

### 1. Fork 

Fork the project [on Github](https://github.com/cseeger-epages/confluence-go-api/) and clone your fork

```
$ git clone git@github.com:username/confluence-go-api.git
$ cd confluence-go-api
$ git remote add upstream https://github.com/cseeger-epages/confluence-go-api.git
$ git feth upstream
```

Also configure git to know who you are:

```
$ git config user.name "Jane Doe"
$ git config user.email "j.doe@example.com"
```

### 2. Branch

Best practice is to organize your development environment by creating local branches to work within.
These should also be created directly off the `master` branch.

```
$ git checkout -b example-branch -t upstream/master
```

### 3. Code

Follow the [official code guidelines](https://golang.org/doc/effective_go.html).

To make sure the code runs correct, test the code using:

```
mage test
```

also add unit tests to test your code contributions.

For new features add a short description in the README.md.

### 4. Commit

It is a recommended best practice to keep your changes as logically grouped as possible within individual commits. 
There is no limit to the number of commits any single Pull Request may have, and many contributors find it easier to review changes that are split across multiple commits.

```
$ git add files/changed
$ git commit
```

**Commit message guidelines**

A good commit message should contain a short description what changed and why

- The first line should contain a short description (not more than 72 characters)
  - e.g.: `additional filter added to filter mails by whatever`
- if you fix open issues, add a reference to the issue
  - e.g.: `issue-1337: fixed ...`
- if you commit a breaking change (see [semantic versioning](https://semver.org/)), the message should contain an explanation about the reason of the breaking change, what triggers the change and what the exact change is

### 5. Rebase

As a best practice, once you have committed your changes, it is a good idea to use `git rebase` (not `git merge`) to synchronize your work with the main repository.

```
$ git fetch upstream
$ git rebase upstream/master
```

This ensures that your working branch has the latest changes from master.

### 6. Push

If your commits are ready to go and passed all tests and linting, you can start creating a [Pull Requests](https://github.com/cseeger-epages/confluence-go-api/pulls) by pushing your work branch to your fork on GitHub.

```
$ git push origin example-branch
```

### 7. Open the Pull Request

From within GitHub, by opening a new Pull Request will present you with a template that should be filled out:

```
<!--
Thank you for your pull request. Please provide a description above and review
the requirements below.

Bug fixes and new features should include unit tests.

Contributors guide: https://github.com/cseeger-epages/confluence-go-api/blob/master/CONTRIBUTING.md
-->

##### Checklist
<!-- Remove items that do not apply. For completed items, change [ ] to [x]. -->

- [ ] `mage test` passes
- [ ] unit tests are included and tested
- [ ] documentation is added or changed
```

Please fill out all details, feel free to skip not nessesary parts or if you're not sure what to fill in.

Once opened, the Pull Request is opend and will be reviewed.

### 8. Updates and discussion

While reviewing you will probably get some feedback or requests for changes to your Pull Request. This is normal and a necessary part of the process to evaluate the changes and there correctness. 

To make changes to an existsing Pull Request, make the changes to your local branch.
Add a new commit including those changes and push them to your fork.
The Pull Requests will automatically updated by GitHub.

```
$ git add files/changed
$ git commit
$ git push origin example-branch
```
**Approvement and Changes**

Whenever a contributor reviews a Pull Request they may find specific details that they would like to see changed or fixed. 
These may be as simple as fixing a typo, or may involve substantive changes to the code you have written. 
While such requests are intended to be helpful, they may come across as abrupt or unhelpful, especially requests to change things that do not include concrete suggestions on how to change them.

Try not to be discouraged. 
If you feel that a particular review is unfair, say so, or contact one of the other contributors in the project and seek their input. 
Often such comments are the result of the reviewer having only taken a short amount of time to review and are not ill-intended.
Such issues can often be resolved with a bit of patience. 
That said, reviewers should be expected to be helpful in their feedback, and feedback that is simply vague, dismissive and unhelpful is likely safe to ignore.

## Reviewing

Reviews and feedback must be helpful, insightful, and geared towards improving the contribution as opposed to simply blocking it.
If there are reasons why you feel the PR should not be accepted, explain what those are. 
Be open to having your mind changed. 
Be open to working with the contributor to make the Pull Request better.

Also follow the [Code of Conduct](https://github.com/cseeger-epages/confluence-go-api/blob/master/CODE_OF_CONDUCT.md)

When reviewing a Pull Request, the primary goals are for the codebase to improve and for the person submitting the request to succeed. 
Even if a Pull Request is not accepted, the submitters should come away from the experience feeling like their effort was not wasted or unappreciated. 
Every Pull Request from a new contributor is an opportunity to grow the community.

Be aware that **how** you communicate requests and reviews in your feedback can have a significant impact on the success of the Pull Request.
