docker build -t jarvissh .

if [ ! -d "./runpath" ]; then
    mkdir runpath
    cp config.yaml ./runpath/
fi