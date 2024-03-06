package connectorconfig

import (
	"encoding/json"
	"strings"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type BlueSnapEnvironment string

const (
	BlueSnapEnvironmentSandbox    BlueSnapEnvironment = "sandbox"
	BlueSnapEnvironmentProduction BlueSnapEnvironment = "production"
)

type BlueSnapCredentials struct {
	StoreID     string              `json:"storeID" yaml:"storeID" validate:"required,gt=0"`
	Environment BlueSnapEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
	Username    *string             `json:"username" yaml:"username" validate:"required,gt=0"`
	Password    *string             `json:"password" yaml:"password" validate:"required,gt=0"`
	Descriptor  string              `json:"descriptor" yaml:"descriptor"`
}

func (c *BlueSnapCredentials) GetMID() string {
	return c.StoreID
}

func (c *BlueSnapCredentials) GetLibrary() Library {
	return LibraryBlueSnap
}

func (c *BlueSnapCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *BlueSnapCredentials) GetSecureFields() []*string {
	return []*string{c.Username, c.Password}
}

func (c *BlueSnapCredentials) Validate() error {
	return nil
}

func (c *BlueSnapCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *BlueSnapCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *BlueSnapCredentials) SupportsSca() bool {
	return false
}

func (c *BlueSnapCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

var blueSnapAllowedCountires = []string{"al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw", "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt", "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca", "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck", "cr", "hr", "cw", "cy", "cz", "dk", "dj", "dm", "do", "tl", "ec", "eg", "sv", "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "tf", "ga", "gm", "ge", "de", "gh", "gi", "gb", "gr", "gl", "gd", "gp", "gu", "gt", "gg", "gn", "gw", "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ie", "im", "il", "it", "ci", "jm", "jp", "je", "jo", "kz", "ke", "ki", "kw", "kg", "la", "lv", "ls", "lr", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my", "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn", "me", "ms", "ma", "mz", "na", "nr", "np", "nl", "nc", "nz", "ni", "ne", "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe", "ph", "pn", "pl", "pf", "pt", "pr", "qa", "re", "ro", "ru", "rw", "gs", "sh", "kn", "lc", "mf", "pm", "st", "vc", "ws", "sm", "sa", "sn", "rs", "sc", "sl", "sg", "sx", "sk", "si", "sb", "so", "za", "kr", "es", "lk", "sr", "sj", "sz", "se", "ch", "tj", "tw", "tz", "th", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc", "tv", "ug", "ua", "ae", "gb", "us", "uy", "uz", "vu", "ve", "vn", "vg", "vi", "wf", "eh", "zm", "zw"}

func (c *BlueSnapCredentials) SupportsCountry(country string) bool {
	if country == "" {
		return true
	}
	for _, v := range blueSnapAllowedCountires {
		if strings.EqualFold(v, country) {
			return true
		}
	}
	return false
}

func (c *BlueSnapCredentials) CanPlanModeUse(mode environment.Mode) bool {

	if mode == environment.ModeSandbox && c.Environment == BlueSnapEnvironmentSandbox {
		return true
	}

	if mode == environment.ModeProduction && c.Environment == BlueSnapEnvironmentProduction {
		return true
	}

	return false
}

func (c *BlueSnapCredentials) IsRecoveryAgent() bool {
	return false
}
