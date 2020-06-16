package mftestutil

import (
	"github.com/RedmonkeyDF/mffileutil"
	"os"
	"path/filepath"
	"testing"
)

func RemoveFileRequiredToNotExist(t *testing.T, afile string){

	thefile, errabs := filepath.Abs(afile)
	if errabs != nil {

		t.Fatalf("%s failed via DoRemoveFileRequiredToNotExist. Unexpected error in Abs call.  Error: \"%s\"", t.Name(), errabs)
	}
	fe, errfe := mffileutil.RegularfileExists(thefile)
	if errfe != nil {

		t.Fatalf("%s failed via DoRemoveFileRequiredToNotExist.  Unexpected error in RegularfileExists: \"%s\".", t.Name(), errfe)
	}

	if fe {

		errrm := os.Remove(thefile)
		if errrm != nil {

			t.Fatalf("%s failed via DoRemoveFileRequiredToNotExist.  Unexpected error in os.Remove: \"%s\".", t.Name(), errrm)
		}

		fe2, errfe2 := mffileutil.RegularfileExists(thefile)
		if errfe2 != nil {

			t.Fatalf("%s failed via DoRemoveFileRequiredToNotExist.  Unexpected error in RegularfileExists after os.Remove: \"%s\".", t.Name(), errfe)
		}

		if fe2 {
			t.Fatalf("%s failed via DoRemoveFileRequiredToNotExist.  File \"%s\" exists after os.Remove.", t.Name(), thefile)
		}
	}
}

func RemoveDirectoryRequiredToNotExist(t *testing.T, adirectory string){

	thedirectory, errabs := filepath.Abs(adirectory)
	if errabs != nil {

		t.Fatalf("%s failed via DoRemoveDirectoryRequiredToNotExist. Unexpected error in Abs call.  Error: \"%s\"", t.Name(), errabs)
	}
	de, errde := mffileutil.DirectoryExists(thedirectory)
	if errde != nil {

		t.Fatalf("%s failed via DoRemoveDirectoryRequiredToNotExist.  Unexpected error in DirectoryExists: \"%s\".", t.Name(), errde)
	}

	if de {

		errrm := mffileutil.RemoveDirectoryWithContents(thedirectory)
		if errrm != nil {

			t.Fatalf("%s failed via DoRemoveDirectoryRequiredToNotExist.  Unexpected error in RemoveDirectoryWithContents: \"%s\".", t.Name(), errrm)
		}

		de2, errde2 := mffileutil.DirectoryExists(thedirectory)
		if errde2 != nil {

			t.Fatalf("%s failed via DoRemoveDirectoryRequiredToNotExist.  Unexpected error in DirectoryExists after os.Remove: \"%s\".", t.Name(), errde2)
		}

		if de2 {
			t.Fatalf("%s failed via DoRemoveDirectoryRequiredToNotExist.  Directory \"%s\" exists after os.Remove.", t.Name(), thedirectory)
		}
	}
}

func FileRequiredToNotExist(t *testing.T, afile string){

	thefile, errabs := filepath.Abs(afile)
	if errabs != nil {
		t.Fatalf("%s failed via testFileRequiredToNotExist.  Unexpected error on Abs call.  Error: \"%s\".", t.Name(), errabs)
	}

	fe, errfe := mffileutil.RegularfileExists(thefile)
	if errfe != nil{
		t.Fatalf("%s failed via testFileRequiredToNotExist.  Unexpected error in RegularfileExists call.  Error: \"%s\".", t.Name(), errfe)
	}

	if fe {

		t.Fatalf("%s failed via testFileRequiredToNotExist.  Required file \"%s\" should not exist, but does.", t.Name(), thefile)
	}
}

func DirectoryRequiredToNotExist(t *testing.T, adirectory string){

	thedir, errabs := filepath.Abs(adirectory)
	if errabs != nil {
		t.Fatalf("%s failed via DoDirectoryRequiredToNotExist.  Unexpected error on Abs call.  Error: \"%s\".", t.Name(), errabs)
	}

	de, errde := mffileutil.DirectoryExists(thedir)
	if errde != nil{
		t.Fatalf("%s failed via DoDirectoryRequiredToNotExist.  Unexpected error in DirectoryExists call.  Error: \"%s\".", t.Name(), errde)
	}

	if de {

		t.Fatalf("%s failed via DoDirectoryRequiredToNotExist.  Required directory \"%s\" should not exist, but does.", t.Name(), thedir)
	}
}

func RequiredDirectoryExists(t *testing.T, adirectory string){

	thedir, errabs := filepath.Abs(adirectory)
	if errabs != nil {
		t.Fatalf("%s failed via testRequiredDirectoryExists.  Unexpected error on Abs call.  Error: \"%s\".", t.Name(), errabs)
	}

	de, errde := mffileutil.DirectoryExists(thedir)
	if errde != nil{
		t.Fatalf("%s failed via testRequiredDirectoryExists.  Unexpected error in DirectoryExists call.  Error: \"%s\".", t.Name(), errde)
	}

	if !de {

		t.Fatalf("%s failed via testRequiredDirectoryExists.  Required directory \"%s\" does not exist.", t.Name(), thedir)
	}
}

func RequiredFileExists(t *testing.T, afile string){

	thefile, errabs := filepath.Abs(afile)
	if errabs != nil {
		t.Fatalf("%s failed via testRequiredFileExists.  Unexpected error on Abs call.  Error: \"%s\".", t.Name(), errabs)
	}

	fe, errfe := mffileutil.RegularfileExists(thefile)
	if errfe != nil{
		t.Fatalf("%s failed via testRequiredFileExists.  Unexpected error in RegularfileExists call.  Error: \"%s\".", t.Name(), errfe)
	}

	if !fe {

		t.Fatalf("%s failed via testRequiredFileExists.  Required file \"%s\" does not exist.", t.Name(), thefile)
	}
}