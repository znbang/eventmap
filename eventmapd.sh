#!/bin/sh
#
# PROVIDE: eventmapd
# KEYWORD:

. /etc/rc.subr

name="eventmapd"
rcvar="${name}_enable"
eventmapd_user="root"
eventmapd_command="/usr/home/ec2-user/${name}"
eventmapd_chdir="/usr/home/ec2-user"
pidfile="/var/run/${name}.pid"
command="/usr/sbin/daemon"
command_args="-P ${pidfile} -r -o /var/log/${name}.log ${eventmapd_command}"

load_rc_config $name
: ${eventmapd_enable:=no}

run_rc_command "$1"
