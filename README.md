# Get factors of a number in golang

* *trial_division.go* - [Easiest](https://en.wikipedia.org/wiki/Trial_division) way to get a factors of a number in Golang.

[![Build Status](https://travis-ci.com/gomth/factors.svg?branch=main)](https://travis-ci.com/gomth/factors)
[![CircleCI](https://circleci.com/gh/gomth/factors/tree/main.svg?style=svg)](https://circleci.com/gh/gomth/factors/tree/main)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/5a22e1f488ee46579160a5362c938815)](https://www.codacy.com/gh/gomth/factors/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gomth/factors&amp;utm_campaign=Badge_Grade)
[![codebeat badge](https://codebeat.co/badges/77825abd-47cc-49c8-9893-10ac8d78124e)](https://codebeat.co/projects/github-com-gomth-factors-main)

##### Usage:

```golang
import "github.com/gomth/factors"

nums, err := factors.Get(30)
if err != nil {
    return nil, err
}
```

##### Package contents:
| API | Engine | time complexity | space complexity | Details |
| ------ | ------ | ------ | ------ | ------ |
| Get | [Trial division][trial_division] | - | - | - |
| GetUsingTrialDivisionForInt | [Trial division][trial_division] | - | - | - |

##### TODO:
- add smth more complex :D
- fix travis that cant see this org

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [trial_division]: <https://en.wikipedia.org/wiki/Trial_division>