# jarvissh
[![Build Status](https://travis-ci.org/zhs007/jarvissh.svg?branch=master)](https://travis-ci.org/zhs007/jarvissh)

jarvis shell node

### Install

```
wget https://github.com/zhs007/jarvissh/releases/download/v0.2.91/jarvissh.tar.gz
tar zxvf jarvissh.tar.gz
```

### Run Jarvissh With Docker

```
git clone https://github.com/zhs007/jarvissh.git
docker build -t jarvissh ./jarvissh/
sudo docker run -d --name jarvissh --rm -p 7788:7788 -v $PWD:/home/jarvissh/data jarvissh
```

However, using with ``Docker``, you will not provision service resources through jarvis.

### Run Jarvissh Deamon

```
./jarvissh start -d
```

### Stop Jarvissh Deamon

```
./jarvissh stop
```

### Restart Jarvissh Deamon

```
./jarvissh restart -d
```