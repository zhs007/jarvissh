# jarvissh
[![Build Status](https://travis-ci.org/zhs007/jarvissh.svg?branch=master)](https://travis-ci.org/zhs007/jarvissh)

jarvis shell node

```
git clone https://github.com/zhs007/jarvissh.git
docker build -t jarvissh ./jarvissh/
sudo docker run -d --name jarvissh --rm -p 7788:7788 -v $PWD:/home/jarvissh/data jarvissh
```
