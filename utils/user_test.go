package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgmock"
	"github.com/jackc/pgproto3/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"strings"
	"testing"
	"time"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name string
		pass string
		err  error
	}{
		{
			name: "NoCharacterAtAll",
			pass: "",
			err:  upErr,
		},
		{
			name: "JustEmptyStringAndWhitespace",
			pass: " \n\t\r\v\f ",
			err:  banErr,
		},
		{
			name: "MixtureOfEmptyStringAndWhitespace",
			pass: "U u\n1\t?\r1\v2\f34",
			err:  banErr,
		},
		{
			name: "MissingUpperCaseString",
			pass: "uu1?1234",
			err:  upErr,
		},
		{
			name: "MissingLowerCaseString",
			pass: "UU1?1234",
			err:  lowErr,
		},
		{
			name: "MissingNumber",
			pass: "Uua?aaaa",
			err:  numErr,
		},
		{
			name: "LessThanRequiredMinimumLength",
			pass: "Uu1?123",
			err:  countErr,
		},
		{
			name: "ValidPassword",
			pass: "Uu1?1234",
			err:  nil,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := ValidatePassword(c.pass)

			assert.Equal(t, c.err, err)
		})
	}
}

func TestAddUsers(t *testing.T) {
	script := &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}
	script.Steps = append(script.Steps, pgmock.ExpectMessage(&pgproto3.Query{String: "select 42"}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.RowDescription{
		Fields: []pgproto3.FieldDescription{
			pgproto3.FieldDescription{
				Name:                 []byte("?column?"),
				TableOID:             0,
				TableAttributeNumber: 0,
				DataTypeOID:          23,
				DataTypeSize:         4,
				TypeModifier:         -1,
				Format:               0,
			},
		},
	}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.DataRow{
		Values: [][]byte{[]byte("42")},
	}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.CommandComplete{CommandTag: []byte("SELECT 1")}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	script.Steps = append(script.Steps, pgmock.ExpectMessage(&pgproto3.Terminate{}))

	ln, err := net.Listen("tcp", "127.0.0.1:")
	require.NoError(t, err)
	defer ln.Close()

	serverErrChan := make(chan error, 1)
	go func() {
		defer close(serverErrChan)

		conn, err := ln.Accept()
		if err != nil {
			serverErrChan <- err
			return
		}
		defer conn.Close()

		err = conn.SetDeadline(time.Now().Add(time.Second))
		if err != nil {
			serverErrChan <- err
			return
		}

		err = script.Run(pgproto3.NewBackend(pgproto3.NewChunkReader(conn), conn))
		if err != nil {
			serverErrChan <- err
			return
		}
	}()

	parts := strings.Split(ln.Addr().String(), ":")
	host := parts[0]
	port := parts[1]
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s", host, port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	pgConn, err := pgconn.Connect(ctx, connStr)
	require.NoError(t, err)
	results, err := pgConn.Exec(ctx, "select 42").ReadAll()
	assert.NoError(t, err)

	assert.Len(t, results, 1)
	assert.Nil(t, results[0].Err)
	assert.Equal(t, "SELECT 1", string(results[0].CommandTag))
	assert.Len(t, results[0].Rows, 1)
	assert.Equal(t, "42", string(results[0].Rows[0][0]))

	pgConn.Close(ctx)

	assert.NoError(t, <-serverErrChan)
}
