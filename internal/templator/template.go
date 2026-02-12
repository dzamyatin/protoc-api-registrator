package templator

const Template = `
package {{.PackageName}}

type UrlRegistrator struct {
	url []Path
}

type Path struct {
	Url string
	Method string
}

func NewUrlRegistrator() *UrlRegistrator {
	return &UrlRegistrator{
		url: []Path{
			{{range $k, $url := .Urls}}
			  	{
					Url:"{{$url.Url}}",
					Method:"{{$url.Method}}",	
				},
			{{end}}
		},
	}
}

func (r *UrlRegistrator) GetRegistered() []Path {
	return r.url
}
`
