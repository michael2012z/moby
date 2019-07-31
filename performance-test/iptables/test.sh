for i in {1..20}
do
	sleep 2
	strace -r -T /sbin/iptables --wait -t filter -C DOCKER-ISOLATION-STAGE-1 -i docker0 ! -o docker0 -j DOCKER-ISOLATION-STAGE-2
done
