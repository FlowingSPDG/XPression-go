# Go library for XPression(Ross Video)
## WARNING!
This library, project and codes are not associated or provided with Ross Video.  

Ross Video's "RossTalk" XPression commands library for Go.  
[Official Documents (help.rossvideo.com)](http://help.rossvideo.com/acuity-device/Topics/Protocol/External/XPN/RT-XPN-Comm.html#topic.RT-XPN-Comm)

### Author
Shugo "FlowingSPDG" Kawamura

### Examples
[CUI Example](https://github.com/FlowingSPDG/XPression-go/tree/main/examples/cui/main.go)  
Initialize : 
```go
// xp = XPression instance...
xp, err := XPression.New("localhost",7788)
if err != nil {
	panic(err)
}
```
You can use RossTalk commands such as ``CLRA``, ``UP``, ``DOWN`` and ``READ`` through ``XPression`` instance.  