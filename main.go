package main
/*
#cgo linux CFLAGS: -fplugin=./maliciousLib.so

void echo() {
    printf("yolo");
}
*/
import "C"
func main() {
    C.echo()
    return 0;
}

