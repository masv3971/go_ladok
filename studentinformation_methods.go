package ladok3

// GenderString translate from KonID to the equal string value
func (s *GetStudentReply) GenderString() string {
	switch s.KonID { //TODO(masv): Is this correct?
	case 1:
		return "female"
	case 2:
		return "male"
	default:
		return "n/a"
	}
}
