docker container stop jarvissh
docker container rm jarvissh
docker run -d -p 7788:7788 --name jarvissh -v $PWD/data:/home/jarvissh/data jarvissh