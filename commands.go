package xpression

// COMMAND Command type.
type COMMAND []byte

var (
	COMMAND_CLFB     = []byte("CLFB")
	COMMAND_CLRA     = []byte("CLRA")
	COMMAND_CUE      = []byte("CUE")
	COMMAND_DOWN     = []byte("DOWN")
	COMMAND_FOCUS    = []byte("FOCUS")
	COMMAND_GPI      = []byte("GPI")
	COMMAND_LAYEROFF = []byte("LYAEROFF")
	COMMAND_NEXT     = []byte("NEXT")
	COMMAND_READ     = []byte("READ")
	COMMAND_RESUME   = []byte("RESUME")
	COMMAND_SEQI     = []byte("SEQI")
	COMMAND_SEQO     = []byte("SEQO")
	COMMAND_SWAP     = []byte("SWAP")
	COMMAND_TAKE     = []byte("TAKE")
	COMMAND_UNCUEALL = []byte("UNCUEALL")
	COMMAND_UNCUE    = []byte("UNCUE")
	COMMAND_UP       = []byte("UP")
	COMMAND_UPNEXT   = []byte("UPNEXT")
)