package k8sclient

// -- git clone
import (
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/rs/zerolog/log"
)

func Test_Func1(t *testing.T) {
	// Info("git clone https://github.com/go-git/go-git")

	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})

	if err != nil {
		log.Err(err).Msg("")
		return
	}
}
