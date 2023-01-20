package dispatcher

import (
	"github.com/noisyboy-9/colossus-of-rhodes/process"
)

type Dispatcher interface {
	Dispatch(p *process.Process)
}
