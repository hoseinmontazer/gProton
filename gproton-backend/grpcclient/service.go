package grpcclient

func (ps *ProtoSet) FindServiceAndMethod(svcName, methodName string) (*Service, *Method) {
	for _, s := range ps.Services {
		if s.Name == svcName {
			for _, m := range s.Methods {
				if m.Name == methodName {
					return &s, &m
				}
			}
		}
	}
	return nil, nil
}
