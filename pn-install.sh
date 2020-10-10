#!/bin/bash
PN_TOKEN=$1
PN_PATH=/var/www/pokenode
PN_AGENT_PATH=/var/www/pokenode/pn-agent
PN_SYSTEMD_CONFIG_PATH=/etc/systemd/system/pn-agent.service

# Begin message
echo "============================================="
echo "Installing pn-agent, it may take few minutes."
echo "============================================="

# Create folder for Pokenode
if [ ! -d $PN_PATH ]; then
        mkdir -p $PN_PATH
fi

# Download pn-agent from GitHub
URL_OF_AGENT=https://github.com/pokenode/pn-agent/releases/download/v0.0.1/pn-agent
echo "Downloading pn-agent from GitHub ..."
curl -s -L -o $PN_AGENT_PATH $URL_OF_AGENT
if [ $? -ne 0 ] || [ ! -f $PN_AGENT_PATH ]; then
        echo "Failed to download pn-agent from GitHub."
        exit
else
        chmod +x $PN_AGENT_PATH
        echo "Download finished."
fi

# Download systemd config file
URL_OF_CONFIG=https://raw.githubusercontent.com/pokenode/pn-agent/main/pn-agent.service
echo "Downloading systemd config from GitHub ..."
curl -s -L -o $PN_SYSTEMD_CONFIG_PATH $URL_OF_CONFIG
if [ $? -ne 0 ] || [ ! -f $PN_SYSTEMD_CONFIG_PATH ]; then
        echo "Fail to download systemd config from GitHub."
        exit
else
        sed -i "s/PN_TOKEN/$PN_TOKEN/g" $PN_SYSTEMD_CONFIG_PATH
        echo "Download finished."
fi

# Setup systemd
systemctl daemon-reload
systemctl enable pn-agent
systemctl start pn-agent

# End message
echo "==============================================="
echo "PokeNode agent has been installed successfully!"
echo "==============================================="
