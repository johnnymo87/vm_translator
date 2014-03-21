**Setup**
* Clone the repo with `go get github.com/johnnymo87/assembler`
* Compile it by running `go build` in the app root directory
* Run the tests with `ginkgo -r`
* Translate a file to binary with `./vm_translator -filename=data/StackArithmetic/SimpleAdd.vm
* SimpleAdd.asm will be in the same directory as the specified file
