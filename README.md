# Go nil pointer dereference

This repository demonstrates a common error in Go: dereferencing a nil pointer.  This leads to a runtime panic.

The `bug.go` file contains the erroneous code.  The `bugSolution.go` file shows the corrected version.

The error occurs because the pointer `y` is set to `nil`, and then the code attempts to access the value it points to using the `*` operator.  This is undefined behavior and will always cause a panic.

This example highlights the importance of checking for `nil` pointers before dereferencing them. Always use appropriate error handling (e.g. if checks) to avoid this common error.