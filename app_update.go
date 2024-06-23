package steamcmd

type appUpdate struct {
	s         *client
	appID     string
	validate  bool
	directory *string
}

func (s *client) AppUpdate(appID string) *appUpdate {
	return &appUpdate{
		s:     s,
		appID: appID,
	}
}

func (i *appUpdate) Validate() *appUpdate {
	i.validate = true
	return i
}

func (i *appUpdate) ForceInstallDirectory(dir string) *appUpdate {
	i.directory = &dir
	return i
}

func (i *appUpdate) Do() error {
	args := i.getArgs()

	command := i.s.cmd(args...)

	err := command.Run()
	if err != nil {
		return err
	}

	return nil
}

const (
	commandLogin           = "+login"
	commandAppUpdate       = "+app_update"
	commandAppInfo         = "+app_info_print"
	commandForceInstallDir = "+force_install_dir"
	commandQuit            = "+quit"

	paramAnonymous = "anonymous"
	paramValidate  = "validate"
)

func (i *appUpdate) getArgs() []string {
	var args []string

	// Need to set force install directory BEFORE logging in.
	if i.directory != nil {
		args = append(args, commandForceInstallDir, *i.directory)
	}

	args = append(args, i.s.getArgs()...)

	args = append(args, commandAppUpdate, i.appID)

	if i.validate {
		args = append(args, paramValidate)
	}

	args = append(args, commandQuit)

	return args
}
