#! /bin/sh

tshark -T json -j "ip" -F pcap > ${1}
