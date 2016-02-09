package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"sync"
	//"encoding/binary"
	"github.com/emirpasic/gods/sets/treeset"
)

type P struct {
	X, Y, Z int
	Name    string
	q Q
	lock sync.RWMutex
	mymap map[int]Q
	myset *treeset.Set
}

type Q struct {
	X, Y int32
	Name string
}


func (p P) MarshalBinary() ([]byte, error) {

	var b bytes.Buffer
	fmt.Fprintln(&b, p.X, p.Y, p.Z)
	fmt.Fprintln(&b, p.Name)
	fmt.Fprintln(&b, p.q.X, p.q.Y, p.q.Name)
	fmt.Fprintln(&b, len(p.mymap))

	for key,value := range p.mymap{
		fmt.Fprintln(&b, key, value.X, value.Y, value.Name)
	}

	fmt.Fprintln(&b, p.myset.Size())
	for _, v := range p.myset.Values() {
		fmt.Fprintln(&b, fmt.Sprintf("%v",v))
	}
	
	/* Skipping Marshalling of lock - not needed */
	//binary.Write(&b, binary.BigEndian, p.lock)

	//fmt.Fprintln(&b, p.lock.w, p.lock.writerSem, p.lock.readerSem, p.lock.readerCount, p.lock.readerWait)

	return b.Bytes(), nil
}

func (p *P) UnmarshalBinary(data []byte) error {

	b := bytes.NewBuffer(data)
	
	_, err := fmt.Fscanln(b, &p.X, &p.Y, &p.Z)
	_, err = fmt.Fscanln(b, &p.Name)
	_, err = fmt.Fscanln(b, &p.q.X, &p.q.Y, &p.q.Name)
	
	var len int = 0
	_, err = fmt.Fscanln(b, &len)
	p.mymap = make(map[int]Q)
	p.myset = treeset.NewWithStringComparator()

	for i:=0; i<len; i++ {
		var key int
		var val Q
		_, err = fmt.Fscanln(b, &key, &val.X, &val.Y, &val.Name)
		p.mymap[key] = val
	}

	var setlen int = 0
	_, err = fmt.Fscanln(b, &setlen)
	p.myset = treeset.NewWithStringComparator()
	
	for j:=0; j< setlen; j++ {
		var val string
		_, err = fmt.Fscanln(b, &val)
		p.myset.Add(val)
	}
	
	/* Skipping UnMarshalling of lock - not needed */
	


	return err
}



func main() {
var filename string = "integerdata.gob"
p := P{X: 3, Y: 4, Z: 5, Name: "Pythagoras", mymap: make(map[int]Q), q: Q{10, 11, "Trad"}, myset: treeset.NewWithStringComparator()}
p.lock.Lock()
p.mymap[1] = Q{6, 7, "Quad"}
p.mymap[2] = Q{8, 9, "R"}
p.myset.Add("A")
p.myset.Add("B")

save(filename, &p)


var q P 

load(filename, &q)


fmt.Println(" save : " ,p)
fmt.Println(" save treeset : " ,p.myset.Values())
fmt.Println(" load : " ,q)
fmt.Println(" load treeset : ", q.myset.Values())

}

func save(filename string, p *P) {
	dataFile1, err := os.Create(filename)

         if err != nil {
                 fmt.Println(" Create", err)
                 os.Exit(1)
         }

	enc := gob.NewEncoder(dataFile1) // Will write to network.
	
	// Encode (send) the value.
	
	err = enc.Encode(p)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	
	
	dataFile1.Close()
	
}

func load(filename string, q *P) {
	dataFile2, err := os.Open(filename)

         if err != nil {
                 fmt.Println(" Open : ", err)
                 os.Exit(1)
         }
	dec := gob.NewDecoder(dataFile2) // Will read from network.
	
	// Decode (receive) the value.
	err = dec.Decode(q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	
	dataFile2.Close()
}






