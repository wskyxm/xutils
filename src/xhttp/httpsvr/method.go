package httpsvr

func (s *Server)POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	return s.router.POST(relativePath, handlers...)
}

func (s *Server)GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return s.router.GET(relativePath, handlers...)
}

func (s *Server)DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	return s.router.DELETE(relativePath, handlers...)
}

func (s *Server)PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	return s.router.PUT(relativePath, handlers...)
}


