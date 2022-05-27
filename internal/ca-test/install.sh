[ -d /usr/local/share/ca-certificates/extra/ ] || sudo mkdir -p /usr/local/share/ca-certificates/extra/

sudo rm -f /usr/local/share/ca-certificates/extra/{nca,root}*
sudo rm -f /etc/ssl/certs/{root,nca}*

sudo cp -a *.crt /usr/local/share/ca-certificates/extra/
sudo cp -a *.pem /etc/ssl/certs/
sudo update-ca-certificates --fresh
