if [ ! -d "release" ]; then
    mkdir release
fi

docker container stop jarvissh_buildrelease
docker container rm jarvissh_buildrelease
docker run -d \
    --name jarvissh_buildrelease \
    jarvissh

rm -rf $PWD/release/jarvissh
docker container cp jarvissh_buildrelease:/home/jarvissh $PWD/release/jarvissh
rm -rf $PWD/release/jarvissh/dat/coredb
rm -rf $PWD/release/jarvissh/jarvissh.pid
mv $PWD/release/jarvissh/cfg/config.yaml $PWD/release/jarvissh/cfg/config.yaml.default
cd release
tar zcvf jarvissh.tar.gz jarvissh

docker container stop jarvissh_buildrelease
docker container rm jarvissh_buildrelease