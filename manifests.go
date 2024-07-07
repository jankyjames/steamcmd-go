package steamcmd

import (
	"encoding/json"
	"errors"
	"strings"
)

type RemoteManifest struct {
	Common struct {
		Name       string `json:"name"`
		Type       string `json:"type"`
		Parent     string `json:"parent"`
		Oslist     string `json:"oslist"`
		Osarch     string `json:"osarch"`
		Osextended string `json:"osextended"`
		Icon       string `json:"icon"`
		Clienticon string `json:"clienticon"`
		Clienttga  string `json:"clienttga"`
		Eulas      struct {
			Field1 struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Url     string `json:"url"`
				Version string `json:"version"`
			} `json:"0"`
		} `json:"eulas"`
		ReleaseState string `json:"ReleaseState"`
		Gameid       string `json:"gameid"`
	} `json:"common"`
	Config struct {
		Installdir string `json:"installdir"`
		Launch     struct {
			Field1 struct {
				Executable string `json:"executable"`
				Arguments  string `json:"arguments"`
				Config     struct {
					Oslist string `json:"oslist"`
				} `json:"config"`
				DescriptionLoc struct {
					English string `json:"english"`
				} `json:"description_loc"`
				Description string `json:"description"`
			} `json:"0"`
			Field2 struct {
				Executable string `json:"executable"`
				Arguments  string `json:"arguments"`
				Config     struct {
					Oslist string `json:"oslist"`
				} `json:"config"`
				DescriptionLoc struct {
					English string `json:"english"`
				} `json:"description_loc"`
				Description string `json:"description"`
			} `json:"1"`
		} `json:"launch"`
	} `json:"config"`
	Depots depots `json:"depots"`
}

func (rm *RemoteManifest) GetBuildIDForBranch(branchName string) string {
	return rm.Depots.Branches[branchName].BuildID
}

type depots struct {
	Depots         map[string]depot
	Branches       map[string]branch `json:"branches"`
	AdditionalData map[string]string
}

func (d *depots) UnmarshalJSON(i []byte) error {
	var src map[string]json.RawMessage

	err := json.Unmarshal(i, &src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(src["branches"], &d.Branches)
	if err != nil {
		return err
	}
	delete(src, "branches")

	d.Depots = map[string]depot{}
	d.AdditionalData = map[string]string{}
	for key, val := range src {
		var value string
		if err := json.Unmarshal(val, &value); err == nil {
			d.AdditionalData[key] = value
			continue
		}
		var newDepot depot
		err = json.Unmarshal(i, &newDepot)
		if err != nil {
			return err
		}
		d.Depots[key] = newDepot
	}

	return nil
}

type branch struct {
	BuildID          string `json:"buildid"`
	PasswordRequired string `json:"pwdrequired"`
	TimeUpdated      string `json:"timeupdated"`
}

type depot struct {
	Config        map[string]any      `json:"config"` // TODO: To be further defined
	Manifests     map[string]manifest `json:"manifests"`
	DepotFromApp  string              `json:"depotfromapp"`
	SharedInstall string              `json:"sharedinstall"`
	SystemDefined string              `json:"systemdefined"`
}

type manifest struct {
	GID      string `json:"gid"`
	Size     string `json:"size"`
	Download string `json:"download"`
}

var (
	ErrInvalidResponse = errors.New("invalid response from steamcmd")
)

func SteamResponseToJSON(i []byte) ([]byte, error) {
	payload := string(i)

	payloadBegin := strings.Index(payload, "{")
	payloadEnd := strings.LastIndex(payload, "}")
	if payloadBegin == -1 || payloadEnd == -1 {
		return nil, ErrInvalidResponse
	}
	payload = payload[payloadBegin : payloadEnd+1]

	payload = strings.ReplaceAll(payload, "\t", "")
	payload = strings.ReplaceAll(payload, `""`, `":"`)
	payload = strings.ReplaceAll(payload, "\"\n{", "\":{")
	payload = strings.ReplaceAll(payload, "\"\n", "\",\n")
	payload = strings.ReplaceAll(payload, "}\n\"", "},\n\"")
	jsonPayload := strings.ReplaceAll(payload, "\",\n}", "\"\n}")

	return []byte(jsonPayload), nil
}
