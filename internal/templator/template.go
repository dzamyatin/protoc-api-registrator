package templator

const Template = `
package {{.PackageName}}

type Registrator struct {
	url []string
}

func NewRegistrator() *Registrator {
	return &Registrator{
		url: []string{
			{{range $k, $url := .Urls}}
			  	{{$url}}
			{{end}}
		},
	}
}

func (r *Registrator) GetRegistered() []string {
	return r.url
}
`
