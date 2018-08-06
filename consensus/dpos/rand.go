//reference http://xoshiro.di.unimi.it/

package dpos

type Random struct {
	s [4] uint64
}

func NewRandom(seed uint64) *Random {
	var s0 uint64 = seed + 0x9e3779b97f4a7c15
	var s1 uint64 = (s0 ^ (s0 >> 30)) * 0xbf58476d1ce4e5b9
	var s2 uint64 = (s1 ^ (s1 >> 27)) * 0x94d049bb133111eb
	var s3 uint64 = s2 ^ (s2 >> 31)
	return &Random{
		s: [4]uint64{s0, s1, s2, s3},
	}
}

func rotl(x uint64, k uint64) uint64 {
	return (x << k) | (x >> (64 - k))
}

func (r *Random) GenRandom() uint64 {
	var result uint64 = rotl(r.s[1] * 5, 7) * 9

	var t uint64 = r.s[1] << 17

	r.s[2] ^= r.s[0]
	r.s[3] ^= r.s[1]
	r.s[1] ^= r.s[2]
	r.s[0] ^= r.s[3]

	r.s[2] ^= t

	r.s[3] = rotl(r.s[3], 45)

	return result
}