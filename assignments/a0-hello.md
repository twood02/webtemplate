---
layout: page
title:  "Assignment 0: Hello Internet"
permalink: /assignments/helloInternet/
---

> *Phase 1, due 2/18:* You must fill out the code review template for another group's PR.
> *Phase 2, due 3/1:* You must respond to the code review comments made by another group on your PR. You should fix any issues so your code is ready to merge into the repository.

For this assignment you must provide a code review for another group's Pull Request. Your group should review the pull request with ID 1 less than your group's PR.


To start, paste the following into a comment in the PR which you are reviewing.
```

Reviewed by: Student 1 (@git username), Student 2 (@git username)

Client:
  - [ ] Follows correct protocol and works with official server
  - [ ] (Optional) Also works with XXXXX language server
  - [ ] Prints useful error if server is unavailable
  - [ ] Accepts hostname and port as command line arguments
  - [ ] Prints useful error on invalid arguments
  - [ ] Includes *useful* comments, no extraneous commented out code, etc
  - [ ] No dead code. Excess code that is never run should be removed.
  - [ ] No excess console output.
  
General:
  - [ ] Readme explains socket API *clearly*
  - [ ] Readme includes student names
  - [ ] Folder is correctly named after language
  - [ ] File names follow `HelloLanguage` spec

(Optional) Server:
  - [ ] Follows correct protocol and works with official client
  - [ ] Accepts port as command line argument
  - [ ] Prints useful error on invalid arguments
  - [ ] Includes *useful* comments, no extraneous commented out code, etc

Comments or issues to be resolved:

```

You should clone the other group's code and try to run it to verify if it works as expected.  If you can't get the code to run or aren't sure how to setup your environment for the required language, post comments to get help from the original authors.

Mark which parts are complete with a `[X]` and add comments to describe any issues you find.