package steamcmd

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type appInfo struct {
	s     *Client
	appID string
}

func (s *Client) GetAppInfo(appID string) (*RemoteManifest, error) {
	return (&appInfo{
		s:     s,
		appID: appID,
	}).do()
}

func (i *appInfo) do() (*RemoteManifest, error) {
	args := i.getArgs()

	command := i.s.cmd(args...)

	buf := bytes.NewBuffer([]byte{})

	if i.s.debug {
		multiOut := io.MultiWriter(buf, os.Stdout)
		command.Stdout = multiOut
	} else {
		command.Stdout = buf
	}

	err := command.Run()
	if err != nil {
		return nil, err
	}

	steamJson := SteamResponseToJSON(buf.Bytes())
	var manifest RemoteManifest
	err = json.Unmarshal(steamJson, &manifest)
	if err != nil {
		return nil, err
	}

	return &manifest, nil
}

func (i *appInfo) getArgs() []string {
	var args []string

	args = append(args, i.s.getArgs()...)

	args = append(args, commandAppInfo, i.appID)

	args = append(args, commandQuit)

	return args
}
