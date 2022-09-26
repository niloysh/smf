package producer

type SNSSAI struct {
	SST int32
	SD  string
}

type FSEID struct {
	LocalSEID, RemoteSEID uint64
	IPAddress             string
}

type FTEID struct {
	TEID         int
	GTPIPAddress string
}

type PDUSession struct {
	SessionId       int
	DataNetworkName string
}

type PFCPSession struct {
	Fseid FSEID
	PDRs  []PDR
}

type PDR struct {
	PDRId       int
	Fteid       FTEID
	UEIPAddress string
}

type SliceSessionInfo struct {
	Snssai      SNSSAI
	PduSession  PDUSession
	PFCPSession PFCPSession
}
