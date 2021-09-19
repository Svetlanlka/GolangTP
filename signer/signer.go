package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

func ExecutePipeline(workers ...job) {
	var in chan interface{}
	wg := new(sync.WaitGroup)

	for _, work := range workers {
		out := make(chan interface{}, 1)

		wg.Add(1)
		go func(wg *sync.WaitGroup, work job, in chan interface{}, out chan interface{}) {
			defer wg.Done()
			work(in, out)
			close(out)
		}(wg, work, in, out)

		in = out
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	// start := time.Now()
	// end1 := time.Since(start)
	// fmt.Println("singlehash begin time ", end1)
	i := 0
	wgComplex := new(sync.WaitGroup)
	for dataRaw := range in {
		wgComplex.Add(1)
		go func(dataRaw interface{}, wgComplex *sync.WaitGroup, i int) {
			// start2 := time.Now()

			defer wgComplex.Done()
			data := fmt.Sprint(dataRaw)
			// fmt.Println(data)
			var data1, data2 string
			wg := new(sync.WaitGroup)
			dataMd5 := DataSignerMd5(data)
			// fmt.Println("md5(data) ", dataMd5)

			wg.Add(2)
			go func(data string, wg *sync.WaitGroup) {
				defer wg.Done()
				data1 = DataSignerCrc32(data)
				// fmt.Println("crc32(data) ", data1)
				// end3 := time.Since(start)
				// fmt.Println("singlehash time go", i, "|1:", end3)
			}(data, wg)
			go func(dataMd5 string, wg *sync.WaitGroup) {
				defer wg.Done()
				data2 = DataSignerCrc32(dataMd5)
				// fmt.Println("md5(crc32(data)) ", data2)
				// end4 := time.Since(start)
				// fmt.Println("singlehash time go", i, "|2:", end4)
			}(dataMd5, wg)
			wg.Wait()

			hash := (data1 + "~" + data2)
			out <- hash
			// fmt.Println("SingleHash ", hash)
			// end2 := time.Since(start2)
			// fmt.Println("singlehash time ", i, ":", end2)
		}(dataRaw, wgComplex, i)
		i++
		time.Sleep(time.Millisecond * 10)
	}
	wgComplex.Wait()
	// end := time.Since(start)
	// fmt.Println("Singlehash ", end)
}

func MultiHash(in, out chan interface{}) {
	// start := time.Now()
	// end1 := time.Since(start)
	// fmt.Println("multihash begin time ", end1)
	for dataRaw := range in {
		var hash sync.Map
		var result string = ""
		data, ok := dataRaw.(string)
		if !ok {
			panic("cant convert input data to string")
		}
		wg := new(sync.WaitGroup)

		for th := 0; th < 6; th++ {
			wg.Add(1)
			go func(th int, data string, hash *sync.Map, wg *sync.WaitGroup) {
				defer wg.Done()
				hash.Store(th, DataSignerCrc32(strconv.Itoa(th)+data))
			}(th, data, &hash, wg)
		}
		wg.Wait()

		for th := 0; th < 6; th++ {
			h, ok := hash.Load(th)
			if !ok {
				panic("load of value in MultiHash failed")
			}
			strHash := fmt.Sprint(h)
			// fmt.Println("res ", th, " ", strHash)
			result += strHash
		}

		out <- result
		// fmt.Println("MultiHash ", result)
	}
	// end := time.Since(start)
	// fmt.Println("MultiHash ", end)
}

func CombineResults(in, out chan interface{}) {
	// start := time.Now()
	var result string = ""
	var hash []string
	wg := new(sync.WaitGroup)

	for dataRaw := range in {
		wg.Add(1)
		go func(dataRaw interface{}, hash *[]string, wg *sync.WaitGroup) {
			defer wg.Done()
			data, ok := dataRaw.(string)
			if !ok {
				panic("cant convert input data to string")
			}
			*hash = append(*hash, data)
		}(dataRaw, &hash, wg)
	}
	wg.Wait()

	sort.Strings(hash)
	for i := range hash {
		result += hash[i]
		if i < len(hash)-1 {
			result += "_"
		}
	}
	out <- result
	// fmt.Println("CombineResults ", result)
	// end := time.Since(start)
	// fmt.Println("CombineResults ", end)
}
