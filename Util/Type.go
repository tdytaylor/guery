package Util

import (
	"fmt"
)

//////////////////
type Type int32

const (
	UNKNOWNTYPE Type = iota
	BOOL
	INT
	DOUBLE
	STRING
)

////////////////////
type JoinType int32

const (
	UNKNOWNJOINTYPE JoinType = iota
	LEFTJOIN
	RIGHTJOIN
	INNERJOIN
)

////////////////
type FuncType int32

const (
	UNKNOWNFUNCTYPE FuncType = iota
	AGGREGATE
	NORMAL
)

///////////////
type OrderType int32

const (
	UNKNOWNORDERTYPE OrderType = iota
	ASC
	DESC
	FIRST
	LAST
)

/////////////
type QuantifierType int32

const (
	UNKNOWNQUANTIFIERTYPE QuantifierType = iota
	ALL
	DISTINCT
	SOME
	ANY
)

func CheckType(ta, tb Type, op Operator) (Type, error) {
	if ta != tb || ta == UNKNOWNTYPE {
		return UNKNOWNTYPE, fmt.Errorf("type not match")
	}
	return ta, nil
}