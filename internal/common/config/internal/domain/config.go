package domain

type Config struct {
	Database database
	Log      log
	App      app
	Tekton   tekton
	Sso      sso
}
