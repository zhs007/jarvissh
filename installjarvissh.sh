rm -rf jarvissh.tar.gz
wget https://github.com/zhs007/jarvissh/releases/download/v0.2.5/jarvissh.tar.gz
tar zxvf jarvissh.tar.gz
cd jarvissh
./jarvissh stop
./jarvissh start -d