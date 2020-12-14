package xpression

// To programmer who reads this. There is a official RossTalk command document for engineers.
// http://help.rossvideo.com/acuity-device/Topics/Protocol/RossTalk.html

import (
	"fmt"
	"net"
)


var (
	// Terminate Terminate string. 
	Terminate = []byte("\r\n")
)

// XPression Main struct to manage XPression.
type XPression struct {
	addr string
	port uint // 7788 by default.
	conn net.Conn // TCP Connection(pointer).
}

// New Generates new instance of XPression. TODO!
func New(addr string, port uint) (*XPression,error) {
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d",addr,port))
	if err != nil {
		return nil,err
	}
	xp := &XPression{
		conn: c,
		addr:addr,
		port:port,
	}
	return xp,nil
}

// Close closes TCP connection.
func (x *XPression) Close() error {
	return x.conn.Close()
}

// Write Write raw bytes directly.
func (x *XPression) Write(cmd string) error {
	cmd = fmt.Sprintf("%s%s",cmd,Terminate) // Add terminate CR/LF.
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// ClearFrameBuffer Clears framebuffer number buffer. For example, CLFB 0000 clears framebuffer 1.
func (x *XPression) ClearFrameBuffer(buffer int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_CLFB,buffer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// ClearLayerFrameBuffer Clears layer number layer in framebuffer number buffer. For example, CLFB 0000:2 clears layer 2 on framebuffer 1.
func (x *XPression) ClearLayerFrameBuffer(buffer,layer int) error {
	cmd := fmt.Sprintf("%s %d:%d%s",COMMAND_CLFB,buffer,layer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// ClearAllFrameBuffers Clears all framebuffers.
func (x *XPression) ClearAllFrameBuffers() error {
	cmd := fmt.Sprintf("%s%s",COMMAND_CLRA,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Cue Prepares take item takeid to go to air next in framebuffer number buffer on layer number layer. The take item is not taken to air, but is prepared to be taken to air without any frame delay. For example, CUE 3:2:-5 prepares to load the take item 3 into the framebuffer 3 and onto layer -5.
func (x *XPression) Cue(takeid,buffer,layer int) error {
	cmd := fmt.Sprintf("%s %d:%d:%d%s",COMMAND_CUE,takeid,buffer,layer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Down Move the current selection in the sequencer to the item below it in the list.
func (x *XPression) Down() error {
	cmd := fmt.Sprintf("%s%s",COMMAND_DOWN,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Focus Set the sequencer focus to the take item number takeid. For example, FOCUS 0005 set the focus to take item 0005.
func (x *XPression) Focus(takeid int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_FOCUS,takeid,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// GPI Trigger the simulated GPI input gpi. This is treated as if the GPI input were triggered externally. For example, GPI 5 triggers GPI input 5.
func (x *XPression) GPI (gpi int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_GPI,gpi,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// LAYEROFF Takes a scene in framebuffer number buffer on layer number layer off air using the defined out transition. For example, LAYEROFF 0000:2 removes the scene on layer 2 of framebuffer 0000 (the first framebuffer).
func (x *XPression) LAYEROFF (buffer,layer int) error {
	cmd := fmt.Sprintf("%s %d:%d%s",COMMAND_LAYEROFF,buffer,layer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}
// Next Take the current take item in the sequencer to air and advance the current selection to the next item in the list.
func (x *XPression) Next() error {
	cmd := fmt.Sprintf("%s%s",COMMAND_NEXT,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Read Take the current selection in the sequencer to air.
func (x *XPression) Read() error {
	cmd := fmt.Sprintf("%s%s",COMMAND_READ,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Resume Resumes all layers in framebuffer number buffer. For example, RESUME 0000 resumes all layers in framebuffer 1.
func (x *XPression) Resume (buffer int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_RESUME,buffer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// ResumeLayer Resumes layer number layer in framebuffer number buffer. For example, RESUME 0000:2 resumes layer 2 in framebuffer 1.
func (x *XPression) ResumeLayer (buffer,layer int) error {
	cmd := fmt.Sprintf("%s %d:%d%s",COMMAND_RESUME,buffer,layer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// SEQI Loads the take item takeid to air on layer number layer to the output channel selected in the template. The Sequencer focus moves to this item. For example, SEQI 0005:7 loads the take item 0005 onto layer 7.
func (x *XPression) SEQI (takeid,layer int) error {
	cmd := fmt.Sprintf("%s %d:%d%s",COMMAND_SEQI,takeid,layer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// SEQO Takes the take item takeid off-air. For example, SEQO 0005 takes the template with TakeID 5 off-air.
func (x *XPression) SEQO (takeid int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_SEQO,takeid,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Swap Loads all the take items that are currently in the cued state to air in framebuffer number buffer. If a framebuffer is not specified, all cued take items in all framebuffers are taken to air. For example, SWAP 0 takes all the cued take items in framebuffer 1 to air.
func (x *XPression) Swap (buffer int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_SWAP,buffer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Take Loads take item takeid to air in framebuffer number buffer on layer number layer. The Sequencer focus does not move to this item. For example, TAKE 5:0:7 loads the template with TakeID 5 into framebuffer 1 and onto layer 7.
func (x *XPression) Take (takeid,buffer,layer int) error {
	cmd := fmt.Sprintf("%s %d:%d:%d%s",COMMAND_TAKE,takeid,buffer,layer,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// UncueAll Removes all cued items from the cued state.
func (x *XPression) UncueAll () error {
	cmd := fmt.Sprintf("%s%s",COMMAND_UNCUEALL,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Uncue Remove item with take id takeid from the cued state.
func (x *XPression) Uncue (takeid int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_UNCUE,takeid,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Up Move the current selection in the sequencer to the item above it in the list.
func (x *XPression) Up () error {
	cmd := fmt.Sprintf("%s%s",COMMAND_UP,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

// Upnext Sets the preview to the take item takeid in the sequencer without moving the focus bar.
func (x *XPression) Upnext (takeid int) error {
	cmd := fmt.Sprintf("%s %d%s",COMMAND_UPNEXT,takeid,Terminate)
	_, err := x.conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}