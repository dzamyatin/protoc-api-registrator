package atomWebsite

	type UrlRegistrator struct {
			url []Path
		}

	type Path struct {
			Url string
			Method string
		}

	func NewUrlRegistrator() *Registrator {
			return &UrlRegistrator{
					url: []Path{

						  	{
								Url:"/current",
								Method:"GET",
							},

						  	{
								Url:"/register"
								Method:"POST"
							},

						  	{
								Url:"/login"
								Method:"POST"
							},

						  	{
								Url:"/remember-password"
								Method:"POST"
							},

						  	{
								Url:"/change-password"
								Method:"POST"
							},

						  	{
								Url:"/confirm-email"
								Method:"POST"
							},

						  	{
								Url:"/send-email-confirmation"
								Method:"POST"
							},

						  	{
								Url:"/confirm-phone"
								Method:"POST"
							},

						  	{
								Url:"/send-phone-confirmation"
								Method:"POST"
							},

					},
				}
		}

	func (r *UrlRegistrator) GetRegistered() []Path {
			return r.url
		}
