package mantis

import (
	"log"
	"net/http"
	"os"
)

type Mantis struct {
	logger *log.Logger
	Router
}

func (m *Mantis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.logger.Println("Begin ServeHTTP")
	m.Handle(rw, r)
}

func newMantis() *Mantis {
	lg := log.New(os.Stdout, "[Mantis]", 0)
	rt := newRouter()
	return &Mantis{lg, rt}
}

type ClassicMantis struct {
	mts *Mantis
}

func (m *ClassicMantis) AddRouter(pattern string, h Handler, method ...string) {
	if len(method) > 0 {
		m.mts.AddRouter(pattern, h, method[0])
	} else {
		m.mts.AddRouter(pattern, h, "")
	}
}

func (m *ClassicMantis) Run() {
	port := ":8080"
	m.mts.logger.Println("Listen On Port ", port)
	m.mts.logger.Fatalln(http.ListenAndServe(port, m.mts))
}

func Classic() *ClassicMantis {
	m := newMantis()
	return &ClassicMantis{m}
}
