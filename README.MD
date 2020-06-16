#### Mftestutil v1.0.0 - Testing utils

`func RemoveFileRequiredToNotExist(t *testing.T, afile string)`

Remove a file that should not exist prior to a test.

`func RemoveDirectoryRequiredToNotExist(t *testing.T, adirectory string)`

Remove a directory that should not exist before a test.

`func FileRequiredToNotExist(t *testing.T, afile string)`

Test that a file that should not exist does not.

`func DirectoryRequiredToNotExist(t *testing.T, adirectory string)`

Test that a directory  that should not exist does not.

`func RequiredDirectoryExists(t *testing.T, adirectory string)`

Test that a directory that should exist does.

`func RequiredFileExists(t *testing.T, afile string)`

Test that a file that should exist does.