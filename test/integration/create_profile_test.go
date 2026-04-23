//go:build integration

package test

func (s *Suite) Test_CreateProfile() {
	id, err := s.profile.Create(ctx, "ZALUPACHOS", 22, "idinahuy@mail.com", "+73003002020")
	s.NoError(err)

	p, err := s.profile.GetProfile(ctx, id.String())
	s.NoError(err)

	s.Equal("ZALUPACHOS", p.Name)
	s.Equal(22, p.Age)
	s.Equal("idinahuy@mail.com", p.Email)
	s.Equal("+73003002020", p.Phone)

}
