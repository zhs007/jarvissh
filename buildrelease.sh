if [ ! -d "release" ]; then
    mkdir release
fi

docker container stop jarvissh_buildrelease
docker container rm jarvissh_buildrelease
docker run -d \
    --name jarvissh_buildrelease \
    jarvissh

docker container cp jarvissh_buildrelease:/home/jarvissh $PWD/release/jarvissh
tar zcvf $PWD/release/jarvissh.tar.gz $PWD/release/jarvissh

docker container stop jarvissh_buildrelease
docker container rm jarvissh_buildrelease