package dispatcher

import (
	"github.com/noisyboy-9/colossus/process"
)

type Dispatcher interface {
	Dispatch(p *process.Process)
}
