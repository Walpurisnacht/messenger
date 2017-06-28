package messenger

import (
	"encoding/json"
	"net/http"
	"bytes"
)

type PersistMenuSetting struct {
	Setting []PersistMenu `json:"persistent_menu"`
}

type PersistMenu struct {
	Locale string `json:"locale,omitempty"`
	ComposerInput bool `json:"composer_input_disabled,omitempty"`
	CallToActions []CallToActionsItem `json:"call_to_actions"`
}

// Create Persist Menu from setting
func (m *Messenger) NewPersistMenu(setting PersistMenuSetting) error {

	data, err := json.Marshal(setting)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messenger_profile", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = "access_token=" + m.token

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return checkFacebookError(resp.Body)
}