package bootstrap

type MixtureServer interface {
	InitServer() error // 初始化
	StartServer() // 启动
	CloseServer() // 关闭
	CleanAllServer(s chan struct{}) // close by syscall signal
	WaitAllWorld()
}

