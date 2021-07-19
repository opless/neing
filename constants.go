package neing

const(
	TVersion = iota + 100 // size[4] Tversion tag[2] msize[4] version[s]
	RVersion              // size[4] Rversion tag[2] msize[4] version[s]
	TAuth                 // size[4] Tauth tag[2] afid[4] uname[s] aname[s]
	RAuth                 // size[4] Rauth tag[2] aqid[13]
	TAttach               // size[4] Tattach tag[2] fid[4] afid[4] uname[s] aname[s]
	RAttach               // size[4] Rattach tag[2] qid[13]
	TError                // illegal
	RError                // size[4] Rerror tag[2] ename[s]
	TFlush                // size[4] Tflush tag[2] oldtag[2]
	RFlush                // size[4] Rflush tag[2]
	TWalk                 // size[4] Twalk tag[2] fid[4] newfid[4] nwname[2] nwname*(wname[s])
	RWalk                 // size[4] Rwalk tag[2] nwqid[2] nwqid*(wqid[13])
	TOpen                 // size[4] Topen tag[2] fid[4] mode[1]
	ROpen                 // size[4] Ropen tag[2] qid[13] iounit[4]
	TCreate               // size[4] Tcreate tag[2] fid[4] name[s] perm[4] mode[1]
	RCreate               // size[4] Rcreate tag[2] qid[13] iounit[4]
	TRead                 // size[4] Tread tag[2] fid[4] offset[8] count[4]
	RRead                 // size[4] Rread tag[2] count[4] data[count]
	TWrite                // size[4] Twrite tag[2] fid[4] offset[8] count[4] data[count]
	RWrite                // size[4] Rwrite tag[2] count[4]
	TClunk                // size[4] Tclunk tag[2] fid[4]
	RClunk                // size[4] Rclunk tag[2]
	TRemove               // size[4] Tremove tag[2] fid[4]
	RRemove               // size[4] Rremove tag[2]
	TStat                 // size[4] Tstat tag[2] fid[4]
	RStat                 // size[4] Rstat tag[2] stat[n]
	TWstat                // size[4] Twstat tag[2] fid[4] stat[n]
	RWstat                // size[4] Rwstat tag[2])
)

const NoTag uint16 = ^uint16(0)

const (
	MessageTooBig = "message too big for agreed buffer size"
	NotAServer = "not a 9P server"
	NotAClient = "not a 9P client"
	ExpectingNoTag = "expecting NoTag"
	StringTooLong = "string bigger than buffer"
	Not9P2000 = "bad version string, not 9P2000"

	NotImplemented = "not implemented yet"
)
