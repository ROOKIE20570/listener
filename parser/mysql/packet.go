package mysql

import (
	"fmt"
	"log"
)

const (
	COM_SLEEP byte = iota
	COM_QUIT
	COM_INIT_DB
	COM_QUERY
	COM_FIELD_LIST
	COM_CREATE_DB
	COM_DROP_DB
	COM_REFRESH
	COM_SHUTDOWN
	COM_STATISTICS
	COM_PROCESS_INFO
	COM_CONNECT
	COM_PROCESS_KILL
	COM_DEBUG
	COM_PING
	COM_TIME
	COM_DELAYED_INSERT
	COM_CHANGE_USER
	COM_BINLOG_DUMP
	COM_TABLE_DUMP
	COM_CONNECT_OUT
	COM_REGISTER_SLAVE
	COM_STMT_PREPARE
	COM_STMT_EXECUTE
	COM_STMT_SEND_LONG_DATA
	COM_STMT_CLOSE
	COM_STMT_RESET
	COM_SET_OPTION
	COM_STMT_FETCH
	COM_DAEMON
	COM_BINLOG_DUMP_GTID
	COM_RESET_CONNECTION
)

type packet struct {
	len     int
	seq     int
	payload []byte
	from    int8
}

func (pk *packet) resolve() {
	var printInfo string
	if pk.from == FROM_CLIENT_DIRECTION {
		pk.clientResolve(&printInfo)
	} else {
		pk.serverResolve(&printInfo)
	}

	log.Println(printInfo)
}

func (pk *packet) clientResolve(printInfo *string) {
	switch pk.payload[0] {
	case COM_QUIT:

		*printInfo = "CONNECTION CLOSED"
	case COM_INIT_DB:

		*printInfo = fmt.Sprintf("USE %s", pk.payload[1])
	case COM_QUERY, COM_CREATE_DB:

		*printInfo = fmt.Sprintf("%s", pk.payload[1])
	case COM_DROP_DB:

		*printInfo = fmt.Sprintf("DROP %s", pk.payload[1])
	case COM_PING:

		*printInfo = "PING"
	case COM_PROCESS_INFO:

		*printInfo = "SHOW PROCESSLIST"
	case COM_PROCESS_KILL:

		*printInfo = fmt.Sprintf("KILL %s", pk.payload[1])
	case COM_STMT_PREPARE:
		//prepare
		*printInfo = fmt.Sprintf("PREPARE STATEMENT %s")
	case COM_STMT_EXECUTE:
		*printInfo = fmt.Sprintf("EXECUTE")

	}
}

func (pk *packet) serverResolve(printInfo *string) {

}
