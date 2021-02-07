package boot

type Bootstrapper interface {
    Bootstrap() error
}

type Bootstrap func() error
