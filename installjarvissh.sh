rm -rf jarvissh.tar.gz
wget https://github.com/zhs007/jarvissh/releases/download/v0.2.92/jarvissh.linux64.tar.gz
mv jarvissh.linux64.tar.gz jarvissh.tar.gz
tar zxvf jarvissh.tar.gz
mv jarvissh.linux64 jarvissh
cd jarvissh
./jarvissh restart -d
