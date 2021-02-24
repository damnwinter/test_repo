	fmt.Println("Worker pool")
	workerool()
}

func workerool() {

	goroutinesNum := 20
	ch := make(chan struct{}, goroutinesNum)

	// Some very big value
	allWorks := int32(10000000)

	var ind int32
	for ind = 0; ind < allWorks; {
		go func() {
			ch <- struct{}{}
			// some work for worker



			<- ch
		}()
	}
	fmt.Println(ind)

}
