#!/bin/bash
PN_NODEID=$1
PN_PATH=/etc/pokenode
PN_AGENT_PATH=/etc/pokenode/pn-agent
PN_SYSTEMD_CONFIG_PATH=/etc/systemd/system/pn-agent.service

# Begin message
echo "==================================================="
echo "| Installing pn-agent, it may take a few minutes. |"
echo "==================================================="

# Check Root Permission
if [ $(id -u) != "0" ]; then
	echo -e "\e[1m\e[31mRoot permission is required to install pn-agent.\e[0m"
	echo -e "\e[1m\e[31mExit installation.\e[0m"
	exit
fi

# Check input nodeid
if [ $# -lt 1 ]; then
	echo -e "\e[1m\e[31mUsage: bash $0 'nodeid'.\e[0m"
	echo -e "\e[1m\e[31mExit installation.\e[0m"
	exit
fi

# Check SystemD
if [ $(ps --no-headers -o comm 1) != "systemd" ]; then
	echo
	echo -e "\e[1m\e[31mThis machine DOES NOT support SystemD.\e[0m"
	echo -e "\e[1m\e[31mExit installation.\e[0m"
	exit
fi

# Clean previous installation
if [ -d $PN_PATH ]; then
	systemctl stop pn-agent
	rm $PN_SYSTEMD_CONFIG_PATH
	systemctl daemon-reload
	rm -rf $PN_PATH
fi

# Create folder for PokeNode
mkdir -p $PN_PATH

# Download pn-agent from PokeNode.com
URL_OF_AGENT=https://pokenode.com/static/pn-agent
echo "Downloading pn-agent from PokeNode.com ..."
curl -L -o $PN_AGENT_PATH $URL_OF_AGENT
if [ $? -ne 0 ] || [ ! -f $PN_AGENT_PATH ]; then
        echo "Failed to download pn-agent from PokeNode.com."
        exit
else
        chmod +x $PN_AGENT_PATH
        echo "Download finished."
fi

# Download systemd config file
URL_OF_CONFIG=https://pokenode.com/static/pn-agent.service
echo "Downloading systemd config from PokeNode.com ..."
curl -L -o $PN_SYSTEMD_CONFIG_PATH $URL_OF_CONFIG
if [ $? -ne 0 ] || [ ! -f $PN_SYSTEMD_CONFIG_PATH ]; then
        echo "Fail to download systemd config from PokeNode.com."
        exit
else
        sed -i "s/PN_NODEID/$PN_NODEID/g" $PN_SYSTEMD_CONFIG_PATH
        echo "Download finished."
fi

# Create user for running pn-agent
useradd pn-agent -r -d /etc/pokenode -s /bin/false
chown -R pn-agent:pn-agent /etc/pokenode && chmod -R 700 /etc/pokenode

# Setup systemd
systemctl daemon-reload
systemctl enable pn-agent
systemctl restart pn-agent

# End message
echo "==================================================="
echo "| PokeNode agent has been installed successfully! |"
echo "==================================================="
