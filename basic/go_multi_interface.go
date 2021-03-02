package basic

/*
嵌入式接口的实现和调用
未实现的接口可以传任意参数
*/

type Camera101Package struct {
}

type IDeviceBase interface {
	Heart(deviceId string) string
	Send(data []byte) error
	// DealPackage(data byte) (Camera101Package, error)
}

type IDevice struct {
	Addr string
}

type CamereServer struct {
	ResponseData []byte
	//*IDevice
	IDevice
}

func CallBack(device IDeviceBase) {
}

func (d *IDevice) Heart(deviceId string) string {
	return ""
}

func (cs *CamereServer) Send(data []byte) error {
	return nil
}

func Start() {
	CallBack(&CamereServer{})
}
