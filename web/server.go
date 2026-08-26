package web

import (
	"html/template"
	"miaoxiu.example/domain"
	"miaoxiu.example/service"
	"net/http"
)

type Server struct {
	Catalog domain.Catalog
	Service *service.RegistrationService
}

func New(c domain.Catalog, s *service.RegistrationService) *Server {
	return &Server{Catalog: c, Service: s}
}
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", s.home)
	m.HandleFunc("/pattern/", s.pattern)
	m.HandleFunc("/register", s.register)
	return m
}
func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	const page = `<!doctype html><html><head><meta charset="utf-8"><title>苗绣工艺</title><style>body{font-family:sans-serif;margin:0;background:#fff8f0;color:#38251f}header{padding:20px;background:#8f2942;color:white}nav a{color:white;margin-right:18px}.grid{display:grid;grid-template-columns:repeat(4,1fr);gap:16px;padding:24px}.tile{padding:18px;background:white;border-radius:8px;transition:transform .2s}.tile:hover{transform:translateY(-5px)}table{margin:24px;border-collapse:collapse}td,th{padding:10px;border:1px solid #ddd}</style></head><body><header><h1>苗绣工艺专题站</h1><nav><a href="/">首页</a><a href="/#patterns">纹样⌄</a><a href="/#stitches">针法⌄</a><a href="/#artisans">传承人</a><a href="/register">体验报名</a></nav></header><main><h2 id="patterns">纹样图片墙</h2><div class="grid">{{range .Patterns}}<a class="tile" href="/pattern/{{.ID}}"><strong>{{.Name}}</strong><p>{{.Meaning}}</p></a>{{end}}</div><h2 id="stitches">针法</h2><div class="grid">{{range .Stitches}}<div class="tile"><strong>{{.Name}}</strong><p>{{.Description}}</p></div>{{end}}</div><h2>纹样寓意表</h2><table><tr><th>纹样</th><th>寓意</th><th>地域</th></tr>{{range .Patterns}}<tr><td>{{.Name}}</td><td>{{.Meaning}}</td><td>{{.Region}}</td></tr>{{end}}</table></main></body></html>`
	template.Must(template.New("p").Parse(page)).Execute(w, s.Catalog)
}
func (s *Server) pattern(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/pattern/"):]
	p, ok := s.Catalog.Pattern(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<html><body><h1>" + p.Name + "</h1><p>" + p.Meaning + "</p><p>地域：" + p.Region + "</p><a href='/'>返回</a></body></html>"))
}
func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		_, e := s.Service.Register(r.FormValue("name"), r.FormValue("phone"), r.FormValue("session"))
		if e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		http.Redirect(w, r, "/register?ok=1", 303)
		return
	}
	w.Write([]byte(`<html><body><h1>体验报名</h1><form method="post"><input name="name" placeholder="姓名"><input name="phone" placeholder="电话"><select name="session"><option>周六上午</option><option>周日下午</option></select><button>提交报名</button></form></body></html>`))
}
