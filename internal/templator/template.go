package templator

const Template = `
package {{.PackageName}}

type Registrator struct {
	url []Path
}

type Path struct {
	Url string
	Method string
}

func NewRegistrator() *Registrator {
	return &Registrator{
		url: []Path{
			{{range $k, $url := .Urls}}
			  	{
					Url:"{{$url.Url}}"
					Method:"{{$url.Method}}"	
				},
			{{end}}
		},
	}
}

func (r *Registrator) GetRegistered() []Path {
	return r.url
}
`
