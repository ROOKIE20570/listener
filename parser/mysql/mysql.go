package mysql

import (
	"bytes"
	"github.com/google/gopacket"
	"io"
	"listener/config"
	"strconv"
)

type Mysql struct {
}

const (
	FROM_CLIENT_DIRECTION = 0
	FROM_SERVER_DIRECTION = 1
)

type packet struct {
	len     int
	seq     int
	payload bytes.Buffer
	from    int8
}

func (mysql *Mysql) Resolve(net gopacket.Flow, transport gopacket.Flow, r io.Reader) (string, error) {
	var packets = make(chan *packet, 50)
	go func() {
		for {
			select {
			case packet := <-packets:
				packet.resolve()
			}
		}
	}()

	for {
		var payload bytes.Buffer
		var seq uint8

		header := make([]byte, 4)
		if _, err := io.ReadFull(r, header); err != nil {
			if err == io.EOF {
				return "", nil
			}
		}

		length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
		seq = header[3]
		if _, err := io.CopyN(&payload, r, int64(length)); err != nil {
			panic(err)
		}

		var from int8
		if strconv.Itoa(config.Conf.Port) == transport.Src().String() {
			from = FROM_SERVER_DIRECTION
		} else {
			from = FROM_CLIENT_DIRECTION
		}
		pk := &packet{payload.Len(), int(seq), payload, from}
		packets <- pk

	}

}

func (mysql *Mysql) GetFilter(port int) string {
	return "tcp and port " + strconv.Itoa(port)
}

func (pk *packet) resolve() {
}

