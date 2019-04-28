docker container stop jarvissh
docker container rm jarvissh
docker run -d \
    -p 7788:7788 \
    -p 8880:8880 \
    --name jarvissh \
    -v $PWD/cfg:/app/jarvissh/cfg \
    -v $PWD/dat:/app/jarvissh/dat \
    -v $PWD/logs:/app/jarvissh/logs \
    jarvissh