package auth

import "github.com/alicebob/miniredis"

func mockedIsExistRedis(key string){
	IsExistRedis = func(key string) bool {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		defer s.Close()

		redisKey := "6c3a65d23c5f26fc529f6c5ce01a6b31"

		s.Set(redisKey, "")

		if s.Exists(key) {
			return true
		}
		return false
	}
}
