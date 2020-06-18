package cert

import (
	"testing"
)

func TestValidCertDate(t *testing.T)  {
	c, err := New("Golang", "Bob", "2020-18-06")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. expected'GOLANG COURSE' got %v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T)  {
	_, err := New("", "Bob", "2020-18-06")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseToLong(t *testing.T)  {
	course := "azeazpoazie poazie poaziep oazi poaizepo aizpe oaizpeo aizp oeakzp ozkaepoazeo kazpeo kazpoe kazpoe akzp okazpoek azpoe k"
	_, err := New(course, "Bob", "2020-18-06")
	if err == nil {
		t.Errorf("Error should be returned on a too long cousrse (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	name := ""
	_, err := New("Golang", name, "2020-18-06")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestNameTooLong(t *testing.T)  {
	name := "fsdmlfk msdkf llfsdmlf ksdmlfk sdmlkf sdmlk fkfmsdl kfmlsdk fmlsdk flmdsjfklsjdfkjsdlk jfsf ldkjlsdkjf lsdkj lskdj lksdjflkdsjf lksj lksdjflksdjlkjfsldkfsjdfkjds"
	_, err := New("Golang", name, "2020-18-06")
	if err == nil {
		t.Errorf("Error name should  be inf to 30 char got len=%v", len(name))
	}
}