package loader

type Module struct {
	Name   string
	Start  func()
	Stop   func()
	Enable bool
}

var modules = make([]*Module, 0)

func Enable(name string, enable bool) {
	for _, m := range modules {
		if m.Name == name {
			m.Enable = enable
		}
	}
}

func Register(newModule *Module) {
	for _, m := range modules {
		if m.Name == newModule.Name {
			return
		}
	}
	modules = append([]*Module{newModule}, modules...)
}

func Run() {
	for _, m := range modules {
		go m.Start()
	}
}