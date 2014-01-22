package core

var ALOF,
	CKST,
	RDCD,
	RDCN,
	RDMN,
	RDPD,
	RDSN,
	ROPS,
	UTID,
	VERN = &alof,
	&ckst,
	&rdcd,
	&rdcn,
	&rdmn,
	&rdpd,
	&rdsn,
	&rops,
	&utid,
	&vern

var alof,
	ckst,
	rdcd,
	rdcn,
	rdmn,
	rdpd,
	rdsn,
	rops,
	utid,
	vern ss = "ALOF",
	"CKST",
	"RDCD",
	"RDCN",
	"RDMN",
	"RDPD",
	"RDSN",
	"ROPS",
	"UTID",
	"VERN"

type ss string

func (s *ss) String() string { return string(*s) }
