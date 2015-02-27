[![Build Status](https://travis-ci.org/dz0ny/wpapi.svg?branch=master)](https://travis-ci.org/dz0ny/wpapi)

# Deploy
1. heroku create -b https://github.com/kr/heroku-buildpack-go.git
2. git push


# Develop
1. wget https://godeb.s3.amazonaws.com/godeb-amd64.tar.gz; tar -xvf godeb-amd64.tar.gz; ./godeb install 1.3.3
2. https://github.com/Masterminds/glide#install
3. glide in


# Testing
1. glide in
2. go test


# API

- Latest URL for theme ```http://wpapi.herokuapp.com/theme/editor/zip```
- Latest URL for direct theme download ```http://wpapi.herokuapp.com/theme/editor/download```
- Latest URL for plugin ```http://wpapi.herokuapp.com/plugin/akismet/zip```
- Latest URL for direct plugin download ```http://wpapi.herokuapp.com/plugin/akismet/download```


# Benchmark
```
boom -n 1000 -c 50 -q 59 http://wpapi.herokuapp.com/theme/editor/zip
1000 / 1000 Boooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo! 100.00 % 

Summary:
  Total:	17.0869 secs.
  Slowest:	1.1467 secs.
  Fastest:	0.1292 secs.
  Average:	0.1707 secs.
  Requests/sec:	58.5244

Status code distribution:
  [200]	994 responses
  [429]	6 responses

Response time histogram:
  0.129 [1]	|
  0.231 [947]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.333 [6]	|
  0.434 [7]	|
  0.536 [8]	|
  0.638 [7]	|
  0.740 [6]	|
  0.841 [7]	|
  0.943 [7]	|
  1.045 [3]	|
  1.147 [1]	|

Latency distribution:
  10% in 0.1372 secs.
  25% in 0.1405 secs.
  50% in 0.1450 secs.
  75% in 0.1511 secs.
  90% in 0.1620 secs.
  95% in 0.2762 secs.
  99% in 0.8618 secs.
```