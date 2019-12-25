rm -rf release
mkdir release

docker build -f ./Dockerfile.build -t jarvisshbuild .
docker run -d --name jarvisshbuild jarvisshbuild sh -c "while true; do sleep 1; done"
docker cp jarvisshbuild:/app/jarvissh.linux64 ./release/
docker cp jarvisshbuild:/app/jarvissh.win64 ./release/
docker stop jarvisshbuild
docker rm jarvisshbuild

cd release
tar zcvf jarvissh.linux64.tar.gz jarvissh.linux64
tar zcvf jarvissh.win64.tar.gz jarvissh.win64