package steamcmd

import (
	"os"
	"os/exec"
)

var (
	pathToSteamCMD = func() string {
		getenv := os.Getenv("STEAM_CMD_EXE")
		if getenv != "" {
			return getenv
		}
		return "steamcmd"
	}()
)

type client struct {
	steamCmdPath string
	anon         bool
	username     string
	password     string // TODO: Is there a secure way to store this?
	debug        bool
}

func (s *client) getArgs() []string {
	args := []string{commandLogin}
	if s.anon {
		args = append(args, paramAnonymous)
	} else {
		// TODO: Log in with username and password
		panic("login with username and password not yet implemented")
		args = append(args, s.username, s.password)
	}
	return args
}

// TODO: Interactive stuff. Solve inline first.
//func (s *client) Interactive() error {
//	command := exec.Command(pathToSteamCMD, s.getArgs()...)
//
//	buf := bytes.NewBuffer([]byte{})
//	command.Stdout = io.MultiWriter(buf, os.Stdout)
//	scanner := bufio.NewScanner(buf)
//	scanner.Split(bufio.ScanWords)
//
//	err := command.Start()
//	if err != nil {
//		return err
//	}
//
//	for scanner.Scan() {
//		output := scanner.Text()
//		if strings.Contains(output, "Steam>") {
//			fmt.Println("ready for input")
//			break
//		}
//	}
//
//	return nil
//}

func (s *client) cmd(args ...string) *exec.Cmd {
	if !s.anon {
		// TODO: Log in with username and password
		panic("login with username and password not yet implemented")
	}

	command := exec.Command(s.steamCmdPath, args...)

	if s.debug {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}

	return command
}

// New uses anonymous by default unless environment variables STEAM_USERNAME and
// STEAM_PASSWORD are set, or WithUsernameAndPassword is passed in as an option.
func New(options ...loginOptions) *client {
	// TODO: By Default check env vars for username and password
	s := &client{
		steamCmdPath: pathToSteamCMD,
		anon:         true,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

type loginOptions func(*client)

// WithUsernameAndPassword sets the username and password used to interact with steamcmd. Without this option you
// are logged in anonymously.
func WithUsernameAndPassword(username, password string) loginOptions {
	return func(s *client) {
		s.anon = false
		s.username = username
		s.password = password
	}
}

// WithSteamCMD overrides the directory in which the library looks for a steamcmd executable
func WithSteamCMD(path string) loginOptions {
	return func(s *client) {
		s.steamCmdPath = path
	}
}

func EnableDebug() loginOptions {
	return func(c *client) {
		c.debug = true
	}
}
