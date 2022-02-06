package singleton

/* Singleton interface */
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var (
	instance *singleton
	// once     sync.Once
)

/* GetInstance function return object */
func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{
			count: 100,
		}
	}
	return instance
}

/* diligent initialization */
func init() {
	instance = &singleton{
		count: 100,
	}
}

func GetInstanceGoroutines() Singleton {
	/* lazy initialization */
	// once.Do(func() {
	// 	time.Sleep(time.Second)

	// 	instance = &singleton{
	// 		count: 100,
	// 	}
	// })
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
