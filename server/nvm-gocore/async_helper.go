package nvmgocore

func MakeChannel[T any]() chan T {
	return make(chan T)
}

func GetFromChannel[T any](channel chan T) T {
	result := <-channel
	return result
}
