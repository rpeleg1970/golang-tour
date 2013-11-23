package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 *rot13Reader) Read(p []byte) (n int, err error) {
    n,err = r13.r.Read(p);
    for i:=0;i<n;i++ {
        switch {
        case p[i]>='a' && p[i]<='m':
            p[i]+=13
        case p[i]>='n' && p[i]<='z':
            p[i]-=13
        case p[i]>='A' && p[i]<='M':
            p[i]+=13
        case p[i]>='N' && p[i]<='Z':
            p[i]-=13            
        }
    }
    return n,err
}

func main() {
    s := strings.NewReader("Bar evat gb ehyr gurz nyy")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
