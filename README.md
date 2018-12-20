Device service for SNMP Patlite

Some SNMP commands:
snmpset  -v2c -cprivate 192.168.0.20 1.3.6.1.4.1.20440.4.1.5.1.2.1.2.3 i 1 1.3.6.1.4.1.20440.4.1.5.1.2.1.3.3 i 0

snmpget -v2c -c public 192.168.0.20 1.3.6.1.4.1.20440.4.1.5.1.2.1.4.1

snmpwalk -mALL -v1 -cpublic 192.168.0.20 system

snmptest 192.168.0.14
snmptable -v1 localhost

