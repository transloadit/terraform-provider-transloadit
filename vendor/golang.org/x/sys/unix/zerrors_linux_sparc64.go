// mkerrors.sh -Wall -Werror -static -I/tmp/sparc64/include
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build sparc64 && linux
// +build sparc64,linux

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -Wall -Werror -static -I/tmp/sparc64/include _const.go

package unix

import "syscall"

const (
	ASI_LEON_DFLUSH                  = 0x11
	ASI_LEON_IFLUSH                  = 0x10
	ASI_LEON_MMUFLUSH                = 0x18
	B1000000                         = 0x1008
	B115200                          = 0x1002
	B1152000                         = 0x1009
	B1500000                         = 0x100a
	B2000000                         = 0x100b
	B230400                          = 0x1003
	B2500000                         = 0x100c
	B3000000                         = 0x100d
	B3500000                         = 0x100e
	B4000000                         = 0x100f
	B460800                          = 0x1004
	B500000                          = 0x1005
	B57600                           = 0x1001
	B576000                          = 0x1006
	B921600                          = 0x1007
	BLKBSZGET                        = 0x40081270
	BLKBSZSET                        = 0x80081271
	BLKFLSBUF                        = 0x20001261
	BLKFRAGET                        = 0x20001265
	BLKFRASET                        = 0x20001264
	BLKGETSIZE                       = 0x20001260
	BLKGETSIZE64                     = 0x40081272
	BLKPBSZGET                       = 0x2000127b
	BLKRAGET                         = 0x20001263
	BLKRASET                         = 0x20001262
	BLKROGET                         = 0x2000125e
	BLKROSET                         = 0x2000125d
	BLKRRPART                        = 0x2000125f
	BLKSECTGET                       = 0x20001267
	BLKSECTSET                       = 0x20001266
	BLKSSZGET                        = 0x20001268
	BOTHER                           = 0x1000
	BS1                              = 0x2000
	BSDLY                            = 0x2000
	CBAUD                            = 0x100f
	CBAUDEX                          = 0x1000
	CIBAUD                           = 0x100f0000
	CLOCAL                           = 0x800
	CR1                              = 0x200
	CR2                              = 0x400
	CR3                              = 0x600
	CRDLY                            = 0x600
	CREAD                            = 0x80
	CS6                              = 0x10
	CS7                              = 0x20
	CS8                              = 0x30
	CSIZE                            = 0x30
	CSTOPB                           = 0x40
	ECCGETLAYOUT                     = 0x41484d11
	ECCGETSTATS                      = 0x40104d12
	ECHOCTL                          = 0x200
	ECHOE                            = 0x10
	ECHOK                            = 0x20
	ECHOKE                           = 0x800
	ECHONL                           = 0x40
	ECHOPRT                          = 0x400
	EFD_CLOEXEC                      = 0x400000
	EFD_NONBLOCK                     = 0x4000
	EMT_TAGOVF                       = 0x1
	EPOLL_CLOEXEC                    = 0x400000
	EXTPROC                          = 0x10000
	FF1                              = 0x8000
	FFDLY                            = 0x8000
	FICLONE                          = 0x80049409
	FICLONERANGE                     = 0x8020940d
	FLUSHO                           = 0x1000
	FS_IOC_ENABLE_VERITY             = 0x80806685
	FS_IOC_GETFLAGS                  = 0x40086601
	FS_IOC_GET_ENCRYPTION_NONCE      = 0x4010661b
	FS_IOC_GET_ENCRYPTION_POLICY     = 0x800c6615
	FS_IOC_GET_ENCRYPTION_PWSALT     = 0x80106614
	FS_IOC_SETFLAGS                  = 0x80086602
	FS_IOC_SET_ENCRYPTION_POLICY     = 0x400c6613
	F_GETLK                          = 0x7
	F_GETLK64                        = 0x7
	F_GETOWN                         = 0x5
	F_RDLCK                          = 0x1
	F_SETLK                          = 0x8
	F_SETLK64                        = 0x8
	F_SETLKW                         = 0x9
	F_SETLKW64                       = 0x9
	F_SETOWN                         = 0x6
	F_UNLCK                          = 0x3
	F_WRLCK                          = 0x2
	HIDIOCGRAWINFO                   = 0x40084803
	HIDIOCGRDESC                     = 0x50044802
	HIDIOCGRDESCSIZE                 = 0x40044801
	HUPCL                            = 0x400
	ICANON                           = 0x2
	IEXTEN                           = 0x8000
	IN_CLOEXEC                       = 0x400000
	IN_NONBLOCK                      = 0x4000
	IOCTL_VM_SOCKETS_GET_LOCAL_CID   = 0x200007b9
	ISIG                             = 0x1
	IUCLC                            = 0x200
	IXOFF                            = 0x1000
	IXON                             = 0x400
	MAP_ANON                         = 0x20
	MAP_ANONYMOUS                    = 0x20
	MAP_DENYWRITE                    = 0x800
	MAP_EXECUTABLE                   = 0x1000
	MAP_GROWSDOWN                    = 0x200
	MAP_HUGETLB                      = 0x40000
	MAP_LOCKED                       = 0x100
	MAP_NONBLOCK                     = 0x10000
	MAP_NORESERVE                    = 0x40
	MAP_POPULATE                     = 0x8000
	MAP_RENAME                       = 0x20
	MAP_STACK                        = 0x20000
	MAP_SYNC                         = 0x80000
	MCL_CURRENT                      = 0x2000
	MCL_FUTURE                       = 0x4000
	MCL_ONFAULT                      = 0x8000
	MEMERASE                         = 0x80084d02
	MEMERASE64                       = 0x80104d14
	MEMGETBADBLOCK                   = 0x80084d0b
	MEMGETINFO                       = 0x40204d01
	MEMGETOOBSEL                     = 0x40c84d0a
	MEMGETREGIONCOUNT                = 0x40044d07
	MEMISLOCKED                      = 0x40084d17
	MEMLOCK                          = 0x80084d05
	MEMREAD                          = 0xc0404d1a
	MEMREADOOB                       = 0xc0104d04
	MEMSETBADBLOCK                   = 0x80084d0c
	MEMUNLOCK                        = 0x80084d06
	MEMWRITEOOB                      = 0xc0104d03
	MTDFILEMODE                      = 0x20004d13
	NFDBITS                          = 0x40
	NLDLY                            = 0x100
	NOFLSH                           = 0x80
	NS_GET_NSTYPE                    = 0x2000b703
	NS_GET_OWNER_UID                 = 0x2000b704
	NS_GET_PARENT                    = 0x2000b702
	NS_GET_USERNS                    = 0x2000b701
	OLCUC                            = 0x2
	ONLCR                            = 0x4
	OTPERASE                         = 0x800c4d19
	OTPGETREGIONCOUNT                = 0x80044d0e
	OTPGETREGIONINFO                 = 0x800c4d0f
	OTPLOCK                          = 0x400c4d10
	OTPSELECT                        = 0x40044d0d
	O_APPEND                         = 0x8
	O_ASYNC                          = 0x40
	O_CLOEXEC                        = 0x400000
	O_CREAT                          = 0x200
	O_DIRECT                         = 0x100000
	O_DIRECTORY                      = 0x10000
	O_DSYNC                          = 0x2000
	O_EXCL                           = 0x800
	O_FSYNC                          = 0x802000
	O_LARGEFILE                      = 0x0
	O_NDELAY                         = 0x4004
	O_NOATIME                        = 0x200000
	O_NOCTTY                         = 0x8000
	O_NOFOLLOW                       = 0x20000
	O_NONBLOCK                       = 0x4000
	O_PATH                           = 0x1000000
	O_RSYNC                          = 0x802000
	O_SYNC                           = 0x802000
	O_TMPFILE                        = 0x2010000
	O_TRUNC                          = 0x400
	PARENB                           = 0x100
	PARODD                           = 0x200
	PENDIN                           = 0x4000
	PERF_EVENT_IOC_DISABLE           = 0x20002401
	PERF_EVENT_IOC_ENABLE            = 0x20002400
	PERF_EVENT_IOC_ID                = 0x40082407
	PERF_EVENT_IOC_MODIFY_ATTRIBUTES = 0x8008240b
	PERF_EVENT_IOC_PAUSE_OUTPUT      = 0x80042409
	PERF_EVENT_IOC_PERIOD            = 0x80082404
	PERF_EVENT_IOC_QUERY_BPF         = 0xc008240a
	PERF_EVENT_IOC_REFRESH           = 0x20002402
	PERF_EVENT_IOC_RESET             = 0x20002403
	PERF_EVENT_IOC_SET_BPF           = 0x80042408
	PERF_EVENT_IOC_SET_FILTER        = 0x80082406
	PERF_EVENT_IOC_SET_OUTPUT        = 0x20002405
	PPPIOCATTACH                     = 0x8004743d
	PPPIOCATTCHAN                    = 0x80047438
	PPPIOCBRIDGECHAN                 = 0x80047435
	PPPIOCCONNECT                    = 0x8004743a
	PPPIOCDETACH                     = 0x8004743c
	PPPIOCDISCONN                    = 0x20007439
	PPPIOCGASYNCMAP                  = 0x40047458
	PPPIOCGCHAN                      = 0x40047437
	PPPIOCGDEBUG                     = 0x40047441
	PPPIOCGFLAGS                     = 0x4004745a
	PPPIOCGIDLE                      = 0x4010743f
	PPPIOCGIDLE32                    = 0x4008743f
	PPPIOCGIDLE64                    = 0x4010743f
	PPPIOCGL2TPSTATS                 = 0x40487436
	PPPIOCGMRU                       = 0x40047453
	PPPIOCGRASYNCMAP                 = 0x40047455
	PPPIOCGUNIT                      = 0x40047456
	PPPIOCGXASYNCMAP                 = 0x40207450
	PPPIOCSACTIVE                    = 0x80107446
	PPPIOCSASYNCMAP                  = 0x80047457
	PPPIOCSCOMPRESS                  = 0x8010744d
	PPPIOCSDEBUG                     = 0x80047440
	PPPIOCSFLAGS                     = 0x80047459
	PPPIOCSMAXCID                    = 0x80047451
	PPPIOCSMRRU                      = 0x8004743b
	PPPIOCSMRU                       = 0x80047452
	PPPIOCSNPMODE                    = 0x8008744b
	PPPIOCSPASS                      = 0x80107447
	PPPIOCSRASYNCMAP                 = 0x80047454
	PPPIOCSXASYNCMAP                 = 0x8020744f
	PPPIOCUNBRIDGECHAN               = 0x20007434
	PPPIOCXFERUNIT                   = 0x2000744e
	PR_SET_PTRACER_ANY               = 0xffffffffffffffff
	PTRACE_GETFPAREGS                = 0x14
	PTRACE_GETFPREGS                 = 0xe
	PTRACE_GETFPREGS64               = 0x19
	PTRACE_GETREGS64                 = 0x16
	PTRACE_READDATA                  = 0x10
	PTRACE_READTEXT                  = 0x12
	PTRACE_SETFPAREGS                = 0x15
	PTRACE_SETFPREGS                 = 0xf
	PTRACE_SETFPREGS64               = 0x1a
	PTRACE_SETREGS64                 = 0x17
	PTRACE_SPARC_DETACH              = 0xb
	PTRACE_WRITEDATA                 = 0x11
	PTRACE_WRITETEXT                 = 0x13
	PT_FP                            = 0x48
	PT_G0                            = 0x10
	PT_G1                            = 0x14
	PT_G2                            = 0x18
	PT_G3                            = 0x1c
	PT_G4                            = 0x20
	PT_G5                            = 0x24
	PT_G6                            = 0x28
	PT_G7                            = 0x2c
	PT_I0                            = 0x30
	PT_I1                            = 0x34
	PT_I2                            = 0x38
	PT_I3                            = 0x3c
	PT_I4                            = 0x40
	PT_I5                            = 0x44
	PT_I6                            = 0x48
	PT_I7                            = 0x4c
	PT_NPC                           = 0x8
	PT_PC                            = 0x4
	PT_PSR                           = 0x0
	PT_REGS_MAGIC                    = 0x57ac6c00
	PT_TNPC                          = 0x90
	PT_TPC                           = 0x88
	PT_TSTATE                        = 0x80
	PT_V9_FP                         = 0x70
	PT_V9_G0                         = 0x0
	PT_V9_G1                         = 0x8
	PT_V9_G2                         = 0x10
	PT_V9_G3                         = 0x18
	PT_V9_G4                         = 0x20
	PT_V9_G5                         = 0x28
	PT_V9_G6                         = 0x30
	PT_V9_G7                         = 0x38
	PT_V9_I0                         = 0x40
	PT_V9_I1                         = 0x48
	PT_V9_I2                         = 0x50
	PT_V9_I3                         = 0x58
	PT_V9_I4                         = 0x60
	PT_V9_I5                         = 0x68
	PT_V9_I6                         = 0x70
	PT_V9_I7                         = 0x78
	PT_V9_MAGIC                      = 0x9c
	PT_V9_TNPC                       = 0x90
	PT_V9_TPC                        = 0x88
	PT_V9_TSTATE                     = 0x80
	PT_V9_Y                          = 0x98
	PT_WIM                           = 0x10
	PT_Y                             = 0xc
	RLIMIT_AS                        = 0x9
	RLIMIT_MEMLOCK                   = 0x8
	RLIMIT_NOFILE                    = 0x6
	RLIMIT_NPROC                     = 0x7
	RLIMIT_RSS                       = 0x5
	RNDADDENTROPY                    = 0x80085203
	RNDADDTOENTCNT                   = 0x80045201
	RNDCLEARPOOL                     = 0x20005206
	RNDGETENTCNT                     = 0x40045200
	RNDGETPOOL                       = 0x40085202
	RNDRESEEDCRNG                    = 0x20005207
	RNDZAPENTCNT                     = 0x20005204
	RTC_AIE_OFF                      = 0x20007002
	RTC_AIE_ON                       = 0x20007001
	RTC_ALM_READ                     = 0x40247008
	RTC_ALM_SET                      = 0x80247007
	RTC_EPOCH_READ                   = 0x4008700d
	RTC_EPOCH_SET                    = 0x8008700e
	RTC_IRQP_READ                    = 0x4008700b
	RTC_IRQP_SET                     = 0x8008700c
	RTC_PARAM_GET                    = 0x80187013
	RTC_PARAM_SET                    = 0x80187014
	RTC_PIE_OFF                      = 0x20007006
	RTC_PIE_ON                       = 0x20007005
	RTC_PLL_GET                      = 0x40207011
	RTC_PLL_SET                      = 0x80207012
	RTC_RD_TIME                      = 0x40247009
	RTC_SET_TIME                     = 0x8024700a
	RTC_UIE_OFF                      = 0x20007004
	RTC_UIE_ON                       = 0x20007003
	RTC_VL_CLR                       = 0x20007014
	RTC_VL_READ                      = 0x40047013
	RTC_WIE_OFF                      = 0x20007010
	RTC_WIE_ON                       = 0x2000700f
	RTC_WKALM_RD                     = 0x40287010
	RTC_WKALM_SET                    = 0x8028700f
	SCM_TIMESTAMPING                 = 0x23
	SCM_TIMESTAMPING_OPT_STATS       = 0x38
	SCM_TIMESTAMPING_PKTINFO         = 0x3c
	SCM_TIMESTAMPNS                  = 0x21
	SCM_TXTIME                       = 0x3f
	SCM_WIFI_STATUS                  = 0x25
	SFD_CLOEXEC                      = 0x400000
	SFD_NONBLOCK                     = 0x4000
	SIOCATMARK                       = 0x8905
	SIOCGPGRP                        = 0x8904
	SIOCGSTAMPNS_NEW                 = 0x40108907
	SIOCGSTAMP_NEW                   = 0x40108906
	SIOCINQ                          = 0x4004667f
	SIOCOUTQ                         = 0x40047473
	SIOCSPGRP                        = 0x8902
	SOCK_CLOEXEC                     = 0x400000
	SOCK_DGRAM                       = 0x2
	SOCK_NONBLOCK                    = 0x4000
	SOCK_STREAM                      = 0x1
	SOL_SOCKET                       = 0xffff
	SO_ACCEPTCONN                    = 0x8000
	SO_ATTACH_BPF                    = 0x34
	SO_ATTACH_REUSEPORT_CBPF         = 0x35
	SO_ATTACH_REUSEPORT_EBPF         = 0x36
	SO_BINDTODEVICE                  = 0xd
	SO_BINDTOIFINDEX                 = 0x41
	SO_BPF_EXTENSIONS                = 0x32
	SO_BROADCAST                     = 0x20
	SO_BSDCOMPAT                     = 0x400
	SO_BUF_LOCK                      = 0x51
	SO_BUSY_POLL                     = 0x30
	SO_BUSY_POLL_BUDGET              = 0x49
	SO_CNX_ADVICE                    = 0x37
	SO_COOKIE                        = 0x3b
	SO_DETACH_REUSEPORT_BPF          = 0x47
	SO_DOMAIN                        = 0x1029
	SO_DONTROUTE                     = 0x10
	SO_ERROR                         = 0x1007
	SO_INCOMING_CPU                  = 0x33
	SO_INCOMING_NAPI_ID              = 0x3a
	SO_KEEPALIVE                     = 0x8
	SO_LINGER                        = 0x80
	SO_LOCK_FILTER                   = 0x28
	SO_MARK                          = 0x22
	SO_MAX_PACING_RATE               = 0x31
	SO_MEMINFO                       = 0x39
	SO_NETNS_COOKIE                  = 0x50
	SO_NOFCS                         = 0x27
	SO_OOBINLINE                     = 0x100
	SO_PASSCRED                      = 0x2
	SO_PASSSEC                       = 0x1f
	SO_PEEK_OFF                      = 0x26
	SO_PEERCRED                      = 0x40
	SO_PEERGROUPS                    = 0x3d
	SO_PEERSEC                       = 0x1e
	SO_PREFER_BUSY_POLL              = 0x48
	SO_PROTOCOL                      = 0x1028
	SO_RCVBUF                        = 0x1002
	SO_RCVBUFFORCE                   = 0x100b
	SO_RCVLOWAT                      = 0x800
	SO_RCVMARK                       = 0x54
	SO_RCVTIMEO                      = 0x2000
	SO_RCVTIMEO_NEW                  = 0x44
	SO_RCVTIMEO_OLD                  = 0x2000
	SO_RESERVE_MEM                   = 0x52
	SO_REUSEADDR                     = 0x4
	SO_REUSEPORT                     = 0x200
	SO_RXQ_OVFL                      = 0x24
	SO_SECURITY_AUTHENTICATION       = 0x5001
	SO_SECURITY_ENCRYPTION_NETWORK   = 0x5004
	SO_SECURITY_ENCRYPTION_TRANSPORT = 0x5002
	SO_SELECT_ERR_QUEUE              = 0x29
	SO_SNDBUF                        = 0x1001
	SO_SNDBUFFORCE                   = 0x100a
	SO_SNDLOWAT                      = 0x1000
	SO_SNDTIMEO                      = 0x4000
	SO_SNDTIMEO_NEW                  = 0x45
	SO_SNDTIMEO_OLD                  = 0x4000
	SO_TIMESTAMPING                  = 0x23
	SO_TIMESTAMPING_NEW              = 0x43
	SO_TIMESTAMPING_OLD              = 0x23
	SO_TIMESTAMPNS                   = 0x21
	SO_TIMESTAMPNS_NEW               = 0x42
	SO_TIMESTAMPNS_OLD               = 0x21
	SO_TIMESTAMP_NEW                 = 0x46
	SO_TXREHASH                      = 0x53
	SO_TXTIME                        = 0x3f
	SO_TYPE                          = 0x1008
	SO_WIFI_STATUS                   = 0x25
	SO_ZEROCOPY                      = 0x3e
	TAB1                             = 0x800
	TAB2                             = 0x1000
	TAB3                             = 0x1800
	TABDLY                           = 0x1800
	TCFLSH                           = 0x20005407
	TCGETA                           = 0x40125401
	TCGETS                           = 0x40245408
	TCGETS2                          = 0x402c540c
	TCSAFLUSH                        = 0x2
	TCSBRK                           = 0x20005405
	TCSBRKP                          = 0x5425
	TCSETA                           = 0x80125402
	TCSETAF                          = 0x80125404
	TCSETAW                          = 0x80125403
	TCSETS                           = 0x80245409
	TCSETS2                          = 0x802c540d
	TCSETSF                          = 0x8024540b
	TCSETSF2                         = 0x802c540f
	TCSETSW                          = 0x8024540a
	TCSETSW2                         = 0x802c540e
	TCXONC                           = 0x20005406
	TFD_CLOEXEC                      = 0x400000
	TFD_NONBLOCK                     = 0x4000
	TIOCCBRK                         = 0x2000747a
	TIOCCONS                         = 0x20007424
	TIOCEXCL                         = 0x2000740d
	TIOCGDEV                         = 0x40045432
	TIOCGETD                         = 0x40047400
	TIOCGEXCL                        = 0x40045440
	TIOCGICOUNT                      = 0x545d
	TIOCGISO7816                     = 0x40285443
	TIOCGLCKTRMIOS                   = 0x5456
	TIOCGPGRP                        = 0x40047483
	TIOCGPKT                         = 0x40045438
	TIOCGPTLCK                       = 0x40045439
	TIOCGPTN                         = 0x40047486
	TIOCGPTPEER                      = 0x20007489
	TIOCGRS485                       = 0x40205441
	TIOCGSERIAL                      = 0x541e
	TIOCGSID                         = 0x40047485
	TIOCGSOFTCAR                     = 0x40047464
	TIOCGWINSZ                       = 0x40087468
	TIOCINQ                          = 0x4004667f
	TIOCLINUX                        = 0x541c
	TIOCMBIC                         = 0x8004746b
	TIOCMBIS                         = 0x8004746c
	TIOCMGET                         = 0x4004746a
	TIOCMIWAIT                       = 0x545c
	TIOCMSET                         = 0x8004746d
	TIOCM_CAR                        = 0x40
	TIOCM_CD                         = 0x40
	TIOCM_CTS                        = 0x20
	TIOCM_DSR                        = 0x100
	TIOCM_RI                         = 0x80
	TIOCM_RNG                        = 0x80
	TIOCM_SR                         = 0x10
	TIOCM_ST                         = 0x8
	TIOCNOTTY                        = 0x20007471
	TIOCNXCL                         = 0x2000740e
	TIOCOUTQ                         = 0x40047473
	TIOCPKT                          = 0x80047470
	TIOCSBRK                         = 0x2000747b
	TIOCSCTTY                        = 0x20007484
	TIOCSERCONFIG                    = 0x5453
	TIOCSERGETLSR                    = 0x5459
	TIOCSERGETMULTI                  = 0x545a
	TIOCSERGSTRUCT                   = 0x5458
	TIOCSERGWILD                     = 0x5454
	TIOCSERSETMULTI                  = 0x545b
	TIOCSERSWILD                     = 0x5455
	TIOCSETD                         = 0x80047401
	TIOCSIG                          = 0x80047488
	TIOCSISO7816                     = 0xc0285444
	TIOCSLCKTRMIOS                   = 0x5457
	TIOCSPGRP                        = 0x80047482
	TIOCSPTLCK                       = 0x80047487
	TIOCSRS485                       = 0xc0205442
	TIOCSSERIAL                      = 0x541f
	TIOCSSOFTCAR                     = 0x80047465
	TIOCSTART                        = 0x2000746e
	TIOCSTI                          = 0x80017472
	TIOCSTOP                         = 0x2000746f
	TIOCSWINSZ                       = 0x80087467
	TIOCVHANGUP                      = 0x20005437
	TOSTOP                           = 0x100
	TUNATTACHFILTER                  = 0x801054d5
	TUNDETACHFILTER                  = 0x801054d6
	TUNGETDEVNETNS                   = 0x200054e3
	TUNGETFEATURES                   = 0x400454cf
	TUNGETFILTER                     = 0x401054db
	TUNGETIFF                        = 0x400454d2
	TUNGETSNDBUF                     = 0x400454d3
	TUNGETVNETBE                     = 0x400454df
	TUNGETVNETHDRSZ                  = 0x400454d7
	TUNGETVNETLE                     = 0x400454dd
	TUNSETCARRIER                    = 0x800454e2
	TUNSETDEBUG                      = 0x800454c9
	TUNSETFILTEREBPF                 = 0x400454e1
	TUNSETGROUP                      = 0x800454ce
	TUNSETIFF                        = 0x800454ca
	TUNSETIFINDEX                    = 0x800454da
	TUNSETLINK                       = 0x800454cd
	TUNSETNOCSUM                     = 0x800454c8
	TUNSETOFFLOAD                    = 0x800454d0
	TUNSETOWNER                      = 0x800454cc
	TUNSETPERSIST                    = 0x800454cb
	TUNSETQUEUE                      = 0x800454d9
	TUNSETSNDBUF                     = 0x800454d4
	TUNSETSTEERINGEBPF               = 0x400454e0
	TUNSETTXFILTER                   = 0x800454d1
	TUNSETVNETBE                     = 0x800454de
	TUNSETVNETHDRSZ                  = 0x800454d8
	TUNSETVNETLE                     = 0x800454dc
	UBI_IOCATT                       = 0x80186f40
	UBI_IOCDET                       = 0x80046f41
	UBI_IOCEBCH                      = 0x80044f02
	UBI_IOCEBER                      = 0x80044f01
	UBI_IOCEBISMAP                   = 0x40044f05
	UBI_IOCEBMAP                     = 0x80084f03
	UBI_IOCEBUNMAP                   = 0x80044f04
	UBI_IOCMKVOL                     = 0x80986f00
	UBI_IOCRMVOL                     = 0x80046f01
	UBI_IOCRNVOL                     = 0x91106f03
	UBI_IOCRPEB                      = 0x80046f04
	UBI_IOCRSVOL                     = 0x800c6f02
	UBI_IOCSETVOLPROP                = 0x80104f06
	UBI_IOCSPEB                      = 0x80046f05
	UBI_IOCVOLCRBLK                  = 0x80804f07
	UBI_IOCVOLRMBLK                  = 0x20004f08
	UBI_IOCVOLUP                     = 0x80084f00
	VDISCARD                         = 0xd
	VEOF                             = 0x4
	VEOL                             = 0xb
	VEOL2                            = 0x10
	VMIN                             = 0x6
	VREPRINT                         = 0xc
	VSTART                           = 0x8
	VSTOP                            = 0x9
	VSUSP                            = 0xa
	VSWTC                            = 0x7
	VT1                              = 0x4000
	VTDLY                            = 0x4000
	VTIME                            = 0x5
	VWERASE                          = 0xe
	WDIOC_GETBOOTSTATUS              = 0x40045702
	WDIOC_GETPRETIMEOUT              = 0x40045709
	WDIOC_GETSTATUS                  = 0x40045701
	WDIOC_GETSUPPORT                 = 0x40285700
	WDIOC_GETTEMP                    = 0x40045703
	WDIOC_GETTIMELEFT                = 0x4004570a
	WDIOC_GETTIMEOUT                 = 0x40045707
	WDIOC_KEEPALIVE                  = 0x40045705
	WDIOC_SETOPTIONS                 = 0x40045704
	WORDSIZE                         = 0x40
	XCASE                            = 0x4
	XTABS                            = 0x1800
	_HIDIOCGRAWNAME                  = 0x40804804
	_HIDIOCGRAWPHYS                  = 0x40404805
	_HIDIOCGRAWUNIQ                  = 0x40404808
	__TIOCFLUSH                      = 0x80047410
)

// Errors
const (
	EADDRINUSE      = syscall.Errno(0x30)
	EADDRNOTAVAIL   = syscall.Errno(0x31)
	EADV            = syscall.Errno(0x53)
	EAFNOSUPPORT    = syscall.Errno(0x2f)
	EALREADY        = syscall.Errno(0x25)
	EBADE           = syscall.Errno(0x66)
	EBADFD          = syscall.Errno(0x5d)
	EBADMSG         = syscall.Errno(0x4c)
	EBADR           = syscall.Errno(0x67)
	EBADRQC         = syscall.Errno(0x6a)
	EBADSLT         = syscall.Errno(0x6b)
	EBFONT          = syscall.Errno(0x6d)
	ECANCELED       = syscall.Errno(0x7f)
	ECHRNG          = syscall.Errno(0x5e)
	ECOMM           = syscall.Errno(0x55)
	ECONNABORTED    = syscall.Errno(0x35)
	ECONNREFUSED    = syscall.Errno(0x3d)
	ECONNRESET      = syscall.Errno(0x36)
	EDEADLK         = syscall.Errno(0x4e)
	EDEADLOCK       = syscall.Errno(0x6c)
	EDESTADDRREQ    = syscall.Errno(0x27)
	EDOTDOT         = syscall.Errno(0x58)
	EDQUOT          = syscall.Errno(0x45)
	EHOSTDOWN       = syscall.Errno(0x40)
	EHOSTUNREACH    = syscall.Errno(0x41)
	EHWPOISON       = syscall.Errno(0x87)
	EIDRM           = syscall.Errno(0x4d)
	EILSEQ          = syscall.Errno(0x7a)
	EINPROGRESS     = syscall.Errno(0x24)
	EISCONN         = syscall.Errno(0x38)
	EISNAM          = syscall.Errno(0x78)
	EKEYEXPIRED     = syscall.Errno(0x81)
	EKEYREJECTED    = syscall.Errno(0x83)
	EKEYREVOKED     = syscall.Errno(0x82)
	EL2HLT          = syscall.Errno(0x65)
	EL2NSYNC        = syscall.Errno(0x5f)
	EL3HLT          = syscall.Errno(0x60)
	EL3RST          = syscall.Errno(0x61)
	ELIBACC         = syscall.Errno(0x72)
	ELIBBAD         = syscall.Errno(0x70)
	ELIBEXEC        = syscall.Errno(0x6e)
	ELIBMAX         = syscall.Errno(0x7b)
	ELIBSCN         = syscall.Errno(0x7c)
	ELNRNG          = syscall.Errno(0x62)
	ELOOP           = syscall.Errno(0x3e)
	EMEDIUMTYPE     = syscall.Errno(0x7e)
	EMSGSIZE        = syscall.Errno(0x28)
	EMULTIHOP       = syscall.Errno(0x57)
	ENAMETOOLONG    = syscall.Errno(0x3f)
	ENAVAIL         = syscall.Errno(0x77)
	ENETDOWN        = syscall.Errno(0x32)
	ENETRESET       = syscall.Errno(0x34)
	ENETUNREACH     = syscall.Errno(0x33)
	ENOANO          = syscall.Errno(0x69)
	ENOBUFS         = syscall.Errno(0x37)
	ENOCSI          = syscall.Errno(0x64)
	ENODATA         = syscall.Errno(0x6f)
	ENOKEY          = syscall.Errno(0x80)
	ENOLCK          = syscall.Errno(0x4f)
	ENOLINK         = syscall.Errno(0x52)
	ENOMEDIUM       = syscall.Errno(0x7d)
	ENOMSG          = syscall.Errno(0x4b)
	ENONET          = syscall.Errno(0x50)
	ENOPKG          = syscall.Errno(0x71)
	ENOPROTOOPT     = syscall.Errno(0x2a)
	ENOSR           = syscall.Errno(0x4a)
	ENOSTR          = syscall.Errno(0x48)
	ENOSYS          = syscall.Errno(0x5a)
	ENOTCONN        = syscall.Errno(0x39)
	ENOTEMPTY       = syscall.Errno(0x42)
	ENOTNAM         = syscall.Errno(0x76)
	ENOTRECOVERABLE = syscall.Errno(0x85)
	ENOTSOCK        = syscall.Errno(0x26)
	ENOTSUP         = syscall.Errno(0x2d)
	ENOTUNIQ        = syscall.Errno(0x73)
	EOPNOTSUPP      = syscall.Errno(0x2d)
	EOVERFLOW       = syscall.Errno(0x5c)
	EOWNERDEAD      = syscall.Errno(0x84)
	EPFNOSUPPORT    = syscall.Errno(0x2e)
	EPROCLIM        = syscall.Errno(0x43)
	EPROTO          = syscall.Errno(0x56)
	EPROTONOSUPPORT = syscall.Errno(0x2b)
	EPROTOTYPE      = syscall.Errno(0x29)
	EREMCHG         = syscall.Errno(0x59)
	EREMOTE         = syscall.Errno(0x47)
	EREMOTEIO       = syscall.Errno(0x79)
	ERESTART        = syscall.Errno(0x74)
	ERFKILL         = syscall.Errno(0x86)
	ERREMOTE        = syscall.Errno(0x51)
	ESHUTDOWN       = syscall.Errno(0x3a)
	ESOCKTNOSUPPORT = syscall.Errno(0x2c)
	ESRMNT          = syscall.Errno(0x54)
	ESTALE          = syscall.Errno(0x46)
	ESTRPIPE        = syscall.Errno(0x5b)
	ETIME           = syscall.Errno(0x49)
	ETIMEDOUT       = syscall.Errno(0x3c)
	ETOOMANYREFS    = syscall.Errno(0x3b)
	EUCLEAN         = syscall.Errno(0x75)
	EUNATCH         = syscall.Errno(0x63)
	EUSERS          = syscall.Errno(0x44)
	EXFULL          = syscall.Errno(0x68)
)

// Signals
const (
	SIGBUS    = syscall.Signal(0xa)
	SIGCHLD   = syscall.Signal(0x14)
	SIGCLD    = syscall.Signal(0x14)
	SIGCONT   = syscall.Signal(0x13)
	SIGEMT    = syscall.Signal(0x7)
	SIGIO     = syscall.Signal(0x17)
	SIGLOST   = syscall.Signal(0x1d)
	SIGPOLL   = syscall.Signal(0x17)
	SIGPROF   = syscall.Signal(0x1b)
	SIGPWR    = syscall.Signal(0x1d)
	SIGSTOP   = syscall.Signal(0x11)
	SIGSYS    = syscall.Signal(0xc)
	SIGTSTP   = syscall.Signal(0x12)
	SIGTTIN   = syscall.Signal(0x15)
	SIGTTOU   = syscall.Signal(0x16)
	SIGURG    = syscall.Signal(0x10)
	SIGUSR1   = syscall.Signal(0x1e)
	SIGUSR2   = syscall.Signal(0x1f)
	SIGVTALRM = syscall.Signal(0x1a)
	SIGWINCH  = syscall.Signal(0x1c)
	SIGXCPU   = syscall.Signal(0x18)
	SIGXFSZ   = syscall.Signal(0x19)
)

// Error table
var errorList = [...]struct {
	num  syscall.Errno
	name string
	desc string
}{
	{1, "EPERM", "operation not permitted"},
	{2, "ENOENT", "no such file or directory"},
	{3, "ESRCH", "no such process"},
	{4, "EINTR", "interrupted system call"},
	{5, "EIO", "input/output error"},
	{6, "ENXIO", "no such device or address"},
	{7, "E2BIG", "argument list too long"},
	{8, "ENOEXEC", "exec format error"},
	{9, "EBADF", "bad file descriptor"},
	{10, "ECHILD", "no child processes"},
	{11, "EAGAIN", "resource temporarily unavailable"},
	{12, "ENOMEM", "cannot allocate memory"},
	{13, "EACCES", "permission denied"},
	{14, "EFAULT", "bad address"},
	{15, "ENOTBLK", "block device required"},
	{16, "EBUSY", "device or resource busy"},
	{17, "EEXIST", "file exists"},
	{18, "EXDEV", "invalid cross-device link"},
	{19, "ENODEV", "no such device"},
	{20, "ENOTDIR", "not a directory"},
	{21, "EISDIR", "is a directory"},
	{22, "EINVAL", "invalid argument"},
	{23, "ENFILE", "too many open files in system"},
	{24, "EMFILE", "too many open files"},
	{25, "ENOTTY", "inappropriate ioctl for device"},
	{26, "ETXTBSY", "text file busy"},
	{27, "EFBIG", "file too large"},
	{28, "ENOSPC", "no space left on device"},
	{29, "ESPIPE", "illegal seek"},
	{30, "EROFS", "read-only file system"},
	{31, "EMLINK", "too many links"},
	{32, "EPIPE", "broken pipe"},
	{33, "EDOM", "numerical argument out of domain"},
	{34, "ERANGE", "numerical result out of range"},
	{36, "EINPROGRESS", "operation now in progress"},
	{37, "EALREADY", "operation already in progress"},
	{38, "ENOTSOCK", "socket operation on non-socket"},
	{39, "EDESTADDRREQ", "destination address required"},
	{40, "EMSGSIZE", "message too long"},
	{41, "EPROTOTYPE", "protocol wrong type for socket"},
	{42, "ENOPROTOOPT", "protocol not available"},
	{43, "EPROTONOSUPPORT", "protocol not supported"},
	{44, "ESOCKTNOSUPPORT", "socket type not supported"},
	{45, "ENOTSUP", "operation not supported"},
	{46, "EPFNOSUPPORT", "protocol family not supported"},
	{47, "EAFNOSUPPORT", "address family not supported by protocol"},
	{48, "EADDRINUSE", "address already in use"},
	{49, "EADDRNOTAVAIL", "cannot assign requested address"},
	{50, "ENETDOWN", "network is down"},
	{51, "ENETUNREACH", "network is unreachable"},
	{52, "ENETRESET", "network dropped connection on reset"},
	{53, "ECONNABORTED", "software caused connection abort"},
	{54, "ECONNRESET", "connection reset by peer"},
	{55, "ENOBUFS", "no buffer space available"},
	{56, "EISCONN", "transport endpoint is already connected"},
	{57, "ENOTCONN", "transport endpoint is not connected"},
	{58, "ESHUTDOWN", "cannot send after transport endpoint shutdown"},
	{59, "ETOOMANYREFS", "too many references: cannot splice"},
	{60, "ETIMEDOUT", "connection timed out"},
	{61, "ECONNREFUSED", "connection refused"},
	{62, "ELOOP", "too many levels of symbolic links"},
	{63, "ENAMETOOLONG", "file name too long"},
	{64, "EHOSTDOWN", "host is down"},
	{65, "EHOSTUNREACH", "no route to host"},
	{66, "ENOTEMPTY", "directory not empty"},
	{67, "EPROCLIM", "too many processes"},
	{68, "EUSERS", "too many users"},
	{69, "EDQUOT", "disk quota exceeded"},
	{70, "ESTALE", "stale file handle"},
	{71, "EREMOTE", "object is remote"},
	{72, "ENOSTR", "device not a stream"},
	{73, "ETIME", "timer expired"},
	{74, "ENOSR", "out of streams resources"},
	{75, "ENOMSG", "no message of desired type"},
	{76, "EBADMSG", "bad message"},
	{77, "EIDRM", "identifier removed"},
	{78, "EDEADLK", "resource deadlock avoided"},
	{79, "ENOLCK", "no locks available"},
	{80, "ENONET", "machine is not on the network"},
	{81, "ERREMOTE", "unknown error 81"},
	{82, "ENOLINK", "link has been severed"},
	{83, "EADV", "advertise error"},
	{84, "ESRMNT", "srmount error"},
	{85, "ECOMM", "communication error on send"},
	{86, "EPROTO", "protocol error"},
	{87, "EMULTIHOP", "multihop attempted"},
	{88, "EDOTDOT", "RFS specific error"},
	{89, "EREMCHG", "remote address changed"},
	{90, "ENOSYS", "function not implemented"},
	{91, "ESTRPIPE", "streams pipe error"},
	{92, "EOVERFLOW", "value too large for defined data type"},
	{93, "EBADFD", "file descriptor in bad state"},
	{94, "ECHRNG", "channel number out of range"},
	{95, "EL2NSYNC", "level 2 not synchronized"},
	{96, "EL3HLT", "level 3 halted"},
	{97, "EL3RST", "level 3 reset"},
	{98, "ELNRNG", "link number out of range"},
	{99, "EUNATCH", "protocol driver not attached"},
	{100, "ENOCSI", "no CSI structure available"},
	{101, "EL2HLT", "level 2 halted"},
	{102, "EBADE", "invalid exchange"},
	{103, "EBADR", "invalid request descriptor"},
	{104, "EXFULL", "exchange full"},
	{105, "ENOANO", "no anode"},
	{106, "EBADRQC", "invalid request code"},
	{107, "EBADSLT", "invalid slot"},
	{108, "EDEADLOCK", "file locking deadlock error"},
	{109, "EBFONT", "bad font file format"},
	{110, "ELIBEXEC", "cannot exec a shared library directly"},
	{111, "ENODATA", "no data available"},
	{112, "ELIBBAD", "accessing a corrupted shared library"},
	{113, "ENOPKG", "package not installed"},
	{114, "ELIBACC", "can not access a needed shared library"},
	{115, "ENOTUNIQ", "name not unique on network"},
	{116, "ERESTART", "interrupted system call should be restarted"},
	{117, "EUCLEAN", "structure needs cleaning"},
	{118, "ENOTNAM", "not a XENIX named type file"},
	{119, "ENAVAIL", "no XENIX semaphores available"},
	{120, "EISNAM", "is a named type file"},
	{121, "EREMOTEIO", "remote I/O error"},
	{122, "EILSEQ", "invalid or incomplete multibyte or wide character"},
	{123, "ELIBMAX", "attempting to link in too many shared libraries"},
	{124, "ELIBSCN", ".lib section in a.out corrupted"},
	{125, "ENOMEDIUM", "no medium found"},
	{126, "EMEDIUMTYPE", "wrong medium type"},
	{127, "ECANCELED", "operation canceled"},
	{128, "ENOKEY", "required key not available"},
	{129, "EKEYEXPIRED", "key has expired"},
	{130, "EKEYREVOKED", "key has been revoked"},
	{131, "EKEYREJECTED", "key was rejected by service"},
	{132, "EOWNERDEAD", "owner died"},
	{133, "ENOTRECOVERABLE", "state not recoverable"},
	{134, "ERFKILL", "operation not possible due to RF-kill"},
	{135, "EHWPOISON", "memory page has hardware error"},
}

// Signal table
var signalList = [...]struct {
	num  syscall.Signal
	name string
	desc string
}{
	{1, "SIGHUP", "hangup"},
	{2, "SIGINT", "interrupt"},
	{3, "SIGQUIT", "quit"},
	{4, "SIGILL", "illegal instruction"},
	{5, "SIGTRAP", "trace/breakpoint trap"},
	{6, "SIGABRT", "aborted"},
	{7, "SIGEMT", "EMT trap"},
	{8, "SIGFPE", "floating point exception"},
	{9, "SIGKILL", "killed"},
	{10, "SIGBUS", "bus error"},
	{11, "SIGSEGV", "segmentation fault"},
	{12, "SIGSYS", "bad system call"},
	{13, "SIGPIPE", "broken pipe"},
	{14, "SIGALRM", "alarm clock"},
	{15, "SIGTERM", "terminated"},
	{16, "SIGURG", "urgent I/O condition"},
	{17, "SIGSTOP", "stopped (signal)"},
	{18, "SIGTSTP", "stopped"},
	{19, "SIGCONT", "continued"},
	{20, "SIGCHLD", "child exited"},
	{21, "SIGTTIN", "stopped (tty input)"},
	{22, "SIGTTOU", "stopped (tty output)"},
	{23, "SIGIO", "I/O possible"},
	{24, "SIGXCPU", "CPU time limit exceeded"},
	{25, "SIGXFSZ", "file size limit exceeded"},
	{26, "SIGVTALRM", "virtual timer expired"},
	{27, "SIGPROF", "profiling timer expired"},
	{28, "SIGWINCH", "window changed"},
	{29, "SIGLOST", "power failure"},
	{30, "SIGUSR1", "user defined signal 1"},
	{31, "SIGUSR2", "user defined signal 2"},
}
